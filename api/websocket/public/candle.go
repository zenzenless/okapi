package public

import (
	"encoding/json"
	"strconv"
	"time"

	logs "github.com/ipfs/go-log/v2"

	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

var log = logs.NewLoggerWithLevel("ws_public", "DEBUG")

type ChannelSub interface {
	GetArg() map[string]string
	ReceiveData([]byte)
}

var _ ChannelSub = &CandleChannel{}

type CandleChannel struct {
	bar      time.Duration
	channel  string
	instId   string
	Received chan *techan.Candle
}

func NewCandleChannel(bar, instID string) *CandleChannel {
	dur, err := time.ParseDuration(bar)
	if err != nil {
		panic(err)
	}
	return &CandleChannel{
		channel:  "candle" + bar,
		instId:   instID,
		Received: make(chan *techan.Candle, 1024),
		bar:      dur,
	}
}

func (c *CandleChannel) GetArg() map[string]string {
	return map[string]string{
		"channel": c.channel,
		"instId":  c.instId,
	}
}

func (c *CandleChannel) ReceiveData(b []byte) {
	//必须保证每次拿到数据都是新鲜的
	var data [][]string
	err := json.Unmarshal(b, &data)
	if err != nil {
		log.Error(err)
		return
	}
	ts, err := c.ToSeries(data)
	if err != nil {
		log.Error(err)
		return
	}
	select {
	case c.Received <- ts:
	default:
		log.Warnf("the chan for receiving candles was full")
	}

}

func (c *CandleChannel) ToSeries(data [][]string) (*techan.Candle, error) {
	bar := c.bar
	candle := data[0]
	ts, err := strconv.ParseInt(candle[0], 10, 64)
	if err != nil {
		return nil, err
	}
	period := techan.NewTimePeriod(time.UnixMilli(ts), bar)
	tCandle := techan.NewCandle(period)

	tCandle.OpenPrice = big.NewFromString(candle[1])

	tCandle.MaxPrice = big.NewFromString(candle[2])
	tCandle.MinPrice = big.NewFromString(candle[3])

	tCandle.ClosePrice = big.NewFromString(candle[4])
	tCandle.Volume = big.NewFromString(candle[5])

	return tCandle, nil
}
