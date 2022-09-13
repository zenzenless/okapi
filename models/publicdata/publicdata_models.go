package publicdata

import "github.com/james-zhang-bing/okapi"

type (
	Instrument struct {
		InstID    string                `json:"instId"`
		Uly       string                `json:"uly,omitempty"`
		BaseCcy   string                `json:"baseCcy,omitempty"`
		QuoteCcy  string                `json:"quoteCcy,omitempty"`
		SettleCcy string                `json:"settleCcy,omitempty"`
		CtValCcy  string                `json:"ctValCcy,omitempty"`
		CtVal     okapi.JSONFloat64     `json:"ctVal,omitempty"`
		CtMult    okapi.JSONFloat64     `json:"ctMult,omitempty"`
		Stk       okapi.JSONFloat64     `json:"stk,omitempty"`
		TickSz    okapi.JSONFloat64     `json:"tickSz,omitempty"`
		LotSz     okapi.JSONFloat64     `json:"lotSz,omitempty"`
		MinSz     okapi.JSONFloat64     `json:"minSz,omitempty"`
		Lever     okapi.JSONFloat64     `json:"lever"`
		InstType  okapi.InstrumentType  `json:"instType"`
		Category  okapi.FeeCategory     `json:"category,string"`
		OptType   okapi.OptionType      `json:"optType,omitempty"`
		ListTime  okapi.JSONTime        `json:"listTime"`
		ExpTime   okapi.JSONTime        `json:"expTime,omitempty"`
		CtType    okapi.ContractType    `json:"ctType,omitempty"`
		Alias     okapi.AliasType       `json:"alias,omitempty"`
		State     okapi.InstrumentState `json:"state"`
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		TS      okapi.JSONTime                    `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                     `json:"instId"`
		Px     okapi.JSONFloat64          `json:"px"`
		Type   okapi.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		InstID   string               `json:"instId"`
		Oi       okapi.JSONFloat64    `json:"oi"`
		OiCcy    okapi.JSONFloat64    `json:"oiCcy"`
		InstType okapi.InstrumentType `json:"instType"`
		TS       okapi.JSONTime       `json:"ts"`
	}
	FundingRate struct {
		InstID          string               `json:"instId"`
		InstType        okapi.InstrumentType `json:"instType"`
		FundingRate     okapi.JSONFloat64    `json:"fundingRate"`
		NextFundingRate okapi.JSONFloat64    `json:"NextFundingRate"`
		FundingTime     okapi.JSONTime       `json:"fundingTime"`
		NextFundingTime okapi.JSONTime       `json:"nextFundingTime"`
	}
	LimitPrice struct {
		InstID   string               `json:"instId"`
		InstType okapi.InstrumentType `json:"instType"`
		BuyLmt   okapi.JSONFloat64    `json:"buyLmt"`
		SellLmt  okapi.JSONFloat64    `json:"sellLmt"`
		TS       okapi.JSONTime       `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string               `json:"instId"`
		InstType okapi.InstrumentType `json:"instType"`
		SettlePx okapi.JSONFloat64    `json:"settlePx"`
		TS       okapi.JSONTime       `json:"ts"`
	}
	OptionMarketData struct {
		InstID   string               `json:"instId"`
		Uly      string               `json:"uly"`
		InstType okapi.InstrumentType `json:"instType"`
		Delta    okapi.JSONFloat64    `json:"delta"`
		Gamma    okapi.JSONFloat64    `json:"gamma"`
		Vega     okapi.JSONFloat64    `json:"vega"`
		Theta    okapi.JSONFloat64    `json:"theta"`
		DeltaBS  okapi.JSONFloat64    `json:"deltaBS"`
		GammaBS  okapi.JSONFloat64    `json:"gammaBS"`
		VegaBS   okapi.JSONFloat64    `json:"vegaBS"`
		ThetaBS  okapi.JSONFloat64    `json:"thetaBS"`
		Lever    okapi.JSONFloat64    `json:"lever"`
		MarkVol  okapi.JSONFloat64    `json:"markVol"`
		BidVol   okapi.JSONFloat64    `json:"bidVol"`
		AskVol   okapi.JSONFloat64    `json:"askVol"`
		RealVol  okapi.JSONFloat64    `json:"realVol"`
		TS       okapi.JSONTime       `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string            `json:"ccy"`
		Amt          okapi.JSONFloat64 `json:"amt"`
		DiscountLv   okapi.JSONInt64   `json:"discountLv"`
		DiscountInfo []*DiscountInfo   `json:"discountInfo"`
	}
	DiscountInfo struct {
		DiscountRate okapi.JSONInt64 `json:"discountRate"`
		MaxAmt       okapi.JSONInt64 `json:"maxAmt"`
		MinAmt       okapi.JSONInt64 `json:"minAmt"`
	}
	SystemTime struct {
		TS okapi.JSONTime `json:"ts"`
	}
	LiquidationOrder struct {
		InstID    string                    `json:"instId"`
		Uly       string                    `json:"uly,omitempty"`
		InstType  okapi.InstrumentType      `json:"instType"`
		TotalLoss okapi.JSONFloat64         `json:"totalLoss"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string             `json:"ccy,omitempty"`
		Side    okapi.OrderSide    `json:"side"`
		OosSide okapi.PositionSide `json:"posSide"`
		BkPx    okapi.JSONFloat64  `json:"bkPx"`
		Sz      okapi.JSONFloat64  `json:"sz"`
		BkLoss  okapi.JSONFloat64  `json:"bkLoss"`
		TS      okapi.JSONTime     `json:"ts"`
	}
	MarkPrice struct {
		InstID   string               `json:"instId"`
		InstType okapi.InstrumentType `json:"instType"`
		MarkPx   okapi.JSONFloat64    `json:"markPx"`
		TS       okapi.JSONTime       `json:"ts"`
	}
	PositionTier struct {
		InstID       string               `json:"instId"`
		Uly          string               `json:"uly,omitempty"`
		InstType     okapi.InstrumentType `json:"instType"`
		Tier         okapi.JSONInt64      `json:"tier"`
		MinSz        okapi.JSONFloat64    `json:"minSz"`
		MaxSz        okapi.JSONFloat64    `json:"maxSz"`
		Mmr          okapi.JSONFloat64    `json:"mmr"`
		Imr          okapi.JSONFloat64    `json:"imr"`
		OptMgnFactor okapi.JSONFloat64    `json:"optMgnFactor,omitempty"`
		QuoteMaxLoan okapi.JSONFloat64    `json:"quoteMaxLoan,omitempty"`
		BaseMaxLoan  okapi.JSONFloat64    `json:"baseMaxLoan,omitempty"`
		MaxLever     okapi.JSONFloat64    `json:"maxLever"`
		TS           okapi.JSONTime       `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string            `json:"ccy"`
		Rate  okapi.JSONFloat64 `json:"rate"`
		Quota okapi.JSONFloat64 `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		Level         string            `json:"level"`
		IrDiscount    okapi.JSONFloat64 `json:"irDiscount"`
		LoanQuotaCoef int               `json:"loanQuotaCoef,string"`
	}
	State struct {
		Title       string         `json:"title"`
		State       string         `json:"state"`
		Href        string         `json:"href"`
		ServiceType string         `json:"serviceType"`
		System      string         `json:"system"`
		ScheDesc    string         `json:"scheDesc"`
		Begin       okapi.JSONTime `json:"begin"`
		End         okapi.JSONTime `json:"end"`
	}
)
