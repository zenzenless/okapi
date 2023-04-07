package rest

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/james-zhang-bing/okapi"
	"github.com/james-zhang-bing/okapi/requests/rest/trade"
)

func TestTradeAPI(t *testing.T) {
	c := NewClient(&APIKey{}, okapi.RestURL, okapi.AwsServer)
	res, err := c.Trade.ClosePosition(trade.ClosePosition{})
	if err == nil {
		t.Fail()
	}
	t.Log(err)
	t.Log(res)
}

func TestSign(t *testing.T) {
	c := NewClient(GetApiKeyFromEnv(), okapi.DemoRestURL, okapi.DemoServer)
	ts, s := c.sign(http.MethodPost, "/api/v5/trade/order", `{"instId":"BTC-USDT-SWAP","ordType":"limit","px":"2000","side":"buy","sz":"0.1","tdMode":"isolated"}`)
	fmt.Println(ts)
	fmt.Println(s)
}
