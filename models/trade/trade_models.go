package trade

import "github.com/james-zhang-bing/okapi"

type (
	PlaceOrder struct {
		ClOrdID string            `json:"clOrdId"`
		Tag     string            `json:"tag"`
		SMsg    string            `json:"sMsg"`
		SCode   okapi.JSONInt64   `json:"sCode"`
		OrdID   okapi.JSONFloat64 `json:"ordId"`
	}
	CancelOrder struct {
		OrdID   string            `json:"ordId"`
		ClOrdID string            `json:"clOrdId"`
		SMsg    string            `json:"sMsg"`
		SCode   okapi.JSONFloat64 `json:"sCode"`
	}
	AmendOrder struct {
		OrdID   string            `json:"ordId"`
		ClOrdID string            `json:"clOrdId"`
		ReqID   string            `json:"reqId"`
		SMsg    string            `json:"sMsg"`
		SCode   okapi.JSONFloat64 `json:"sCode"`
	}
	ClosePosition struct {
		InstID  string             `json:"instId"`
		PosSide okapi.PositionSide `json:"posSide"`
	}
	Order struct {
		InstID      string               `json:"instId"`
		Ccy         string               `json:"ccy"`
		OrdID       string               `json:"ordId"`
		ClOrdID     string               `json:"clOrdId"`
		TradeID     string               `json:"tradeId"`
		Tag         string               `json:"tag"`
		Category    string               `json:"category"`
		FeeCcy      string               `json:"feeCcy"`
		RebateCcy   string               `json:"rebateCcy"`
		Px          okapi.JSONFloat64    `json:"px"`
		Sz          okapi.JSONInt64      `json:"sz"`
		Pnl         okapi.JSONFloat64    `json:"pnl"`
		AccFillSz   okapi.JSONInt64      `json:"accFillSz"`
		FillPx      okapi.JSONFloat64    `json:"fillPx"`
		FillSz      okapi.JSONInt64      `json:"fillSz"`
		FillTime    okapi.JSONFloat64    `json:"fillTime"`
		AvgPx       okapi.JSONFloat64    `json:"avgPx"`
		Lever       okapi.JSONFloat64    `json:"lever"`
		TpTriggerPx okapi.JSONFloat64    `json:"tpTriggerPx"`
		TpOrdPx     okapi.JSONFloat64    `json:"tpOrdPx"`
		SlTriggerPx okapi.JSONFloat64    `json:"slTriggerPx"`
		SlOrdPx     okapi.JSONFloat64    `json:"slOrdPx"`
		Fee         okapi.JSONFloat64    `json:"fee"`
		Rebate      okapi.JSONFloat64    `json:"rebate"`
		State       okapi.OrderState     `json:"state"`
		TdMode      okapi.TradeMode      `json:"tdMode"`
		PosSide     okapi.PositionSide   `json:"posSide"`
		Side        okapi.OrderSide      `json:"side"`
		OrdType     okapi.OrderType      `json:"ordType"`
		InstType    okapi.InstrumentType `json:"instType"`
		TgtCcy      okapi.QuantityType   `json:"tgtCcy"`
		UTime       okapi.JSONTime       `json:"uTime"`
		CTime       okapi.JSONTime       `json:"cTime"`
	}
	TransactionDetail struct {
		InstID   string               `json:"instId"`
		OrdID    string               `json:"ordId"`
		TradeID  string               `json:"tradeId"`
		ClOrdID  string               `json:"clOrdId"`
		BillID   string               `json:"billId"`
		Tag      okapi.JSONFloat64    `json:"tag"`
		FillPx   okapi.JSONFloat64    `json:"fillPx"`
		FillSz   okapi.JSONFloat64    `json:"fillSz"`
		FeeCcy   okapi.JSONFloat64    `json:"feeCcy"`
		Fee      okapi.JSONFloat64    `json:"fee"`
		InstType okapi.InstrumentType `json:"instType"`
		Side     okapi.OrderSide      `json:"side"`
		PosSide  okapi.PositionSide   `json:"posSide"`
		ExecType okapi.OrderFlowType  `json:"execType"`
		TS       okapi.JSONTime       `json:"ts"`
	}
	PlaceAlgoOrder struct {
		AlgoID string          `json:"algoId"`
		SMsg   string          `json:"sMsg"`
		SCode  okapi.JSONInt64 `json:"sCode"`
	}
	CancelAlgoOrder struct {
		AlgoID string          `json:"algoId"`
		SMsg   string          `json:"sMsg"`
		SCode  okapi.JSONInt64 `json:"sCode"`
	}
	AlgoOrder struct {
		InstID       string               `json:"instId"`
		Ccy          string               `json:"ccy"`
		OrdID        string               `json:"ordId"`
		AlgoID       string               `json:"algoId"`
		ClOrdID      string               `json:"clOrdId"`
		TradeID      string               `json:"tradeId"`
		Tag          string               `json:"tag"`
		Category     string               `json:"category"`
		FeeCcy       string               `json:"feeCcy"`
		RebateCcy    string               `json:"rebateCcy"`
		TimeInterval string               `json:"timeInterval"`
		Px           okapi.JSONFloat64    `json:"px"`
		PxVar        okapi.JSONFloat64    `json:"pxVar"`
		PxSpread     okapi.JSONFloat64    `json:"pxSpread"`
		PxLimit      okapi.JSONFloat64    `json:"pxLimit"`
		Sz           okapi.JSONInt64      `json:"sz"`
		SzLimit      okapi.JSONInt64      `json:"szLimit"`
		ActualSz     okapi.JSONFloat64    `json:"actualSz"`
		ActualPx     okapi.JSONFloat64    `json:"actualPx"`
		Pnl          okapi.JSONFloat64    `json:"pnl"`
		AccFillSz    okapi.JSONInt64      `json:"accFillSz"`
		FillPx       okapi.JSONFloat64    `json:"fillPx"`
		FillSz       okapi.JSONInt64      `json:"fillSz"`
		FillTime     okapi.JSONFloat64    `json:"fillTime"`
		AvgPx        okapi.JSONFloat64    `json:"avgPx"`
		Lever        okapi.JSONFloat64    `json:"lever"`
		TpTriggerPx  okapi.JSONFloat64    `json:"tpTriggerPx"`
		TpOrdPx      okapi.JSONFloat64    `json:"tpOrdPx"`
		SlTriggerPx  okapi.JSONFloat64    `json:"slTriggerPx"`
		SlOrdPx      okapi.JSONFloat64    `json:"slOrdPx"`
		OrdPx        okapi.JSONFloat64    `json:"ordPx"`
		Fee          okapi.JSONFloat64    `json:"fee"`
		Rebate       okapi.JSONFloat64    `json:"rebate"`
		State        okapi.OrderState     `json:"state"`
		TdMode       okapi.TradeMode      `json:"tdMode"`
		ActualSide   okapi.PositionSide   `json:"actualSide"`
		PosSide      okapi.PositionSide   `json:"posSide"`
		Side         okapi.OrderSide      `json:"side"`
		OrdType      okapi.AlgoOrderType  `json:"ordType"`
		InstType     okapi.InstrumentType `json:"instType"`
		TgtCcy       okapi.QuantityType   `json:"tgtCcy"`
		CTime        okapi.JSONTime       `json:"cTime"`
		TriggerTime  okapi.JSONTime       `json:"triggerTime"`
	}
)
