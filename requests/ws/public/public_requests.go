package public

import "github.com/james-zhang-bing/okapi"

type (
	Instruments struct {
		InstType okapi.InstrumentType `json:"instType"`
	}
	Tickers struct {
		InstID string `json:"instId"`
	}
	OpenInterest struct {
		InstID string `json:"instId"`
	}
	Candlesticks struct {
		InstID  string                     `json:"instId"`
		Channel okapi.CandleStickWsBarSize `json:"channel"`
	}
	Trades struct {
		InstID string `json:"instId"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string               `json:"instId"`
		Uly      string               `json:"uly,omitempty"`
		InstType okapi.InstrumentType `json:"instType,omitempty"`
	}
	MarkPrice struct {
		InstID string `json:"instId"`
	}
	MarkPriceCandlesticks struct {
		InstID  string                     `json:"instId"`
		Channel okapi.CandleStickWsBarSize `json:"channel"`
	}
	PriceLimit struct {
		InstID string `json:"instId"`
	}
	OrderBook struct {
		InstID  string `json:"instId"`
		Channel string `json:"channel"`
	}
	OPTIONSummary struct {
		InstID string `json:"instId"`
		Uly    string `json:"uly"`
	}
	FundingRate struct {
		InstID string `json:"instId"`
	}
	IndexCandlesticks struct {
		InstID  string `json:"instId"`
		Channel string `json:"channel"`
	}
	IndexTickers struct {
		InstID string `json:"instId"`
	}
)
