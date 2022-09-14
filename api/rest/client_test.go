package rest

import (
	"net/http"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/james-zhang-bing/okapi"
	"github.com/james-zhang-bing/okapi/requests/rest/trade"
)

func TestClientDo(t *testing.T) {
	c := NewClient(&APIKey{}, okapi.RestURL, okapi.AwsServer)
	req, err := http.NewRequest("GET", "https://www.okx.com/api/v5/market/tickers?instType=SPOT", nil)
	assert.NoError(t, err)
	_, err = c.v0Do(req)
	assert.NoError(t, err)
}

func TestTradeAPI(t *testing.T) {
	c := NewClient(&APIKey{}, okapi.RestURL, okapi.AwsServer)
	res, err := c.Trade.ClosePosition(trade.ClosePosition{})
	if err == nil {
		t.Fail()
	}
	t.Log(err)
	t.Log(res)
}
