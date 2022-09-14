package rest

import (
	"fmt"

	"strconv"
	"testing"
	"time"

	"github.com/james-zhang-bing/okapi/api"
	"github.com/stretchr/testify/assert"
)

func TestGetTickers(t *testing.T) {
	c := NewClient(&APIKey{})
	tickers, err := c.GetTickers()
	assert.NoError(t, err)
	assert.NotNil(t, tickers)
	assert.NotEqual(t, 0, len(tickers.Data))
}

func TestGetKlineOKX(t *testing.T) {
	c := NewClient(&APIKey{})
	startDate := time.Unix(1618150773, 0)
	start := startDate.UnixMilli()
	startStr := strconv.FormatInt(start, 10)
	kline, err := c.GetKlineOKX("BTC-USDT", startStr, "", "1m", "2")
	assert.NoError(t, err)
	assert.NotNil(t, kline)
	assert.NotEqual(t, 0, len(kline.Data))
	fmt.Println(kline)
	candles, err := kline.ToTimeSeries(api.BAR_1m)
	assert.NoError(t, err)
	fmt.Println(candles)
}

func TestClient_GetKlineSeries(t *testing.T) {
	c := NewClient(&APIKey{})
	startDate := time.Unix(1618150773, 0)
	start := startDate.UnixMilli()
	endDate:=startDate.Add(24*time.Hour)
	end:=endDate.UnixMilli()
	got, err := c.GetKlineSeries("BTC-USDT", end, start, api.BAR_15m, "100")
	assert.NoError(t, err)
	fmt.Println(got.Candles[0].Period)

}

func TestTime(t *testing.T) {
	fmt.Println(time.UnixMilli(1597026383085))
}
