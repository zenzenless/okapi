package main

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
	"github.com/james-zhang-bing/okapi/api/rest"
	pr "github.com/james-zhang-bing/okapi/api/websocket/private"
	"github.com/james-zhang-bing/okapi/api/websocket/public"
	"github.com/james-zhang-bing/okapi/requests/rest/market"
	"github.com/james-zhang-bing/okapi/requests/ws/private"
)

func main() {
	apikey := rest.NewAPIKey("your_key", "your_secret", "your_passphrase ")

	//for rest api
	restClient := rest.NewClient(apikey, okapi.AwsRestURL, okapi.AwsServer)

	res, err := restClient.Market.GetCandlesticks(market.GetCandlesticks{InstID: "BTC-USDT"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Candles[0])

	//for websocket public channel
	pubws, err := public.NewPublicClient(okapi.AwsPublicWsURL)
	if err != nil {
		panic(err)
	}
	candleSub := public.NewCandleChannel(string(okapi.Bar1m), "BTC-USDT-SWAP")
	err = pubws.Subscribe(candleSub)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 2; i++ {
		candle := <-candleSub.Received
		fmt.Println(candle)
	}

	//for websocket private channel

	priws, err := pr.NewPrivateClient(okapi.AwsPrivateWsURL, apikey)
	balanceCh, err := priws.SubscribeAccountChannel(&private.Account{})
	if err != nil {
		panic(err)
	}
	fmt.Println(<-balanceCh)
}
