package public

import "github.com/james-zhang-bing/okapi"

type (
	GetInstruments struct {
		Uly      string               `json:"uly,omitempty"`
		InstID   string               `json:"instId,omitempty"`
		InstType okapi.InstrumentType `json:"instType"`
	}
	GetDeliveryExerciseHistory struct {
		Uly      string               `json:"uly"`
		After    int64                `json:"after,omitempty,string"`
		Before   int64                `json:"before,omitempty,string"`
		Limit    int64                `json:"limit,omitempty,string"`
		InstType okapi.InstrumentType `json:"instType"`
	}
	GetOpenInterest struct {
		Uly      string               `json:"uly,omitempty"`
		InstID   string               `json:"instId,omitempty"`
		InstType okapi.InstrumentType `json:"instType"`
	}
	GetFundingRate struct {
		InstID string `json:"instId"`
	}
	GetLimitPrice struct {
		InstID string `json:"instId"`
	}
	GetOptionMarketData struct {
		Uly     string `json:"uly"`
		ExpTime string `json:"expTime,omitempty"`
	}
	GetEstimatedDeliveryExercisePrice struct {
		Uly     string `json:"uly"`
		ExpTime string `json:"expTime,omitempty"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Uly        string  `json:"uly"`
		Ccy        string  `json:"ccy,omitempty"`
		DiscountLv float64 `json:"discountLv,string"`
	}
	GetLiquidationOrders struct {
		InstID   string               `json:"instId,omitempty"`
		Ccy      string               `json:"ccy,omitempty"`
		Uly      string               `json:"uly,omitempty"`
		After    int64                `json:"after,omitempty,string"`
		Before   int64                `json:"before,omitempty,string"`
		Limit    int64                `json:"limit,omitempty,string"`
		InstType okapi.InstrumentType `json:"instType"`
		MgnMode  okapi.MarginMode     `json:"mgnMode,omitempty"`
		Alias    okapi.AliasType      `json:"alias,omitempty"`
		State    okapi.OrderState     `json:"state,omitempty"`
	}
	GetMarkPrice struct {
		InstID   string               `json:"instId,omitempty"`
		Uly      string               `json:"uly,omitempty"`
		InstType okapi.InstrumentType `json:"instType"`
	}
	GetPositionTiers struct {
		InstID   string               `json:"instId,omitempty"`
		Uly      string               `json:"uly,omitempty"`
		InstType okapi.InstrumentType `json:"instType"`
		TdMode   okapi.TradeMode      `json:"tdMode"`
		Tier     okapi.JSONInt64      `json:"tier,omitempty"`
	}
	GetUnderlying struct {
		InstType okapi.InstrumentType `json:"instType"`
	}
	Status struct {
		State string `json:"state,omitempty"`
	}
)
