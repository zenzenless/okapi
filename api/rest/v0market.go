package rest

import (
	"encoding/json"
	"net/http"
	"sort"

	"sync"
	"time"

	logging "github.com/ipfs/go-log/v2"
	"github.com/james-zhang-bing/okapi"
	"github.com/james-zhang-bing/okapi/api"
	"github.com/james-zhang-bing/okapi/requests/rest/market"
	responses "github.com/james-zhang-bing/okapi/responses/market"
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

var log = logging.Logger("req")

func init() {
	logging.SetLogLevel("req", "INFO")
}

//get okex kline
// /api/v5/market/candles
func (c *ClientRest) GetKlineOKX(instId, after, before, bar, limit string) (*api.KlineResOKX, error) {
	req, err := http.NewRequest("GET", string(c.baseURL)+okapi.OkxApiURL.MarketHistoryCandles, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("instId", instId)
	q.Add("after", after)
	q.Add("before", before)
	q.Add("bar", bar)
	q.Add("limit", limit)

	req.URL.RawQuery = q.Encode()
	log.Debug(req.URL.String())
	body, err := c.v0Do(req)
	if err != nil {
		return nil, err
	}
	kline := &api.KlineResOKX{}
	err = json.Unmarshal(body, kline)
	if err != nil {
		return nil, err
	}
	return kline, nil
}

// api/v5/market/tickers
func (c *ClientRest) GetTickers() (*api.Tickers, error) {
	req, err := http.NewRequest("GET", string(c.baseURL)+"/api/v5/market/tickers?instType=SPOT", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.v0Do(req)
	if err != nil {
		return nil, err
	}
	tickers := &api.Tickers{}
	err = json.Unmarshal(body, tickers)
	if err != nil {
		return nil, err
	}
	return tickers, nil
}

// use GetKlineOKX and mashal to techan.TimeSeries
func (c *ClientRest) GetKlineSeries(instId string, endTs, startTs int64, bar string, limit int64) (*techan.TimeSeries, error) {

	log.Infof("get kline series for %s from %s -> %s", instId, time.UnixMilli(startTs), time.UnixMilli(endTs))
	// startTime := strconv.FormatInt(startTs-10, 10)
	// endTime := strconv.FormatInt(endTs-10, 10)
	if endTs == 0 || time.Now().UnixMilli()-endTs < 0 {
		endTs = time.Now().UnixMilli()
	}

	// kl, err := c.GetKlineOKX(instId, endTime, startTime, bar, limit)
	// if err != nil {
	// 	return nil, err
	// }
	candles, err := c.Market.GetCandlesticksHistory(market.GetCandlesticks{
		InstID: instId,
		After:  endTs - 10,
		Before: startTs - 10,
		Limit:  limit,
		Bar:    okapi.BarSize(bar),
	})
	if err != nil {
		return nil, err
	}
	return toTechanSeries(&candles, bar)
}

func toTechanSeries(response *responses.Candle, dur string) (*techan.TimeSeries, error) {
	bar := okapi.BarSize(dur).Duration()
	series := techan.NewTimeSeries()
	for _, candle := range response.Candles {
		ts := time.Time(candle.TS)

		period := techan.NewTimePeriod(ts, bar)
		// fmt.Printf("%v\n", period)
		tCandle := techan.NewCandle(period)
		tCandle.OpenPrice = big.NewDecimal(candle.O)

		tCandle.MaxPrice = big.NewDecimal(candle.H)
		tCandle.MinPrice = big.NewDecimal(candle.L)

		tCandle.ClosePrice = big.NewDecimal(candle.C)
		tCandle.Volume = big.NewDecimal(candle.Vol)
		series.Candles = append(series.Candles, tCandle)
	}
	sort.Slice(series.Candles, func(i, j int) bool {
		return series.Candles[i].Period.Start.Unix() < series.Candles[j].Period.Start.Unix()
	})
	return series, nil
}

func (c *ClientRest) ConcurrentFetch(instId string, end, start time.Time, barString string) (*techan.TimeSeries, error) {
	gNum := make(chan struct{}, 10)
	done := make(chan struct{})
	receive := make(chan *techan.TimeSeries, 10)
	bar := api.BarMap[barString]
	if end.After(time.Now()) {
		end = time.Now().Add(100 * bar)
	}
	frequency := 5
	go func() {
		var wg sync.WaitGroup
		t := time.Now()
		times := 0
		for from := start; end.After(from); from = from.Add(100 * bar) {
			times++
			if times > frequency {
				d := time.Second - time.Now().Sub(t)
				if d > 0 {
					time.Sleep(d)
				}
				times = 0
				t = time.Now()

			}
			wg.Add(1)
			gNum <- struct{}{}
			go func(f time.Time) {
				defer func() {
					<-gNum
					wg.Done()
				}()
				s, e := f.UnixMilli(), f.Add(100*bar).UnixMilli()
				timeSeries, err := c.GetKlineSeries(instId, e, s, barString, 100)
				if err != nil {
					log.Error(err)
					panic(err)
				}
				receive <- timeSeries

			}(from)

		}
		wg.Wait()
		close(done)
	}()
	series := techan.NewTimeSeries()
	for {
		select {
		case timeSeries := <-receive:
			series.Candles = append(series.Candles, timeSeries.Candles...)
		case <-done:
			return series, nil
		}
	}

}
