package account

import "github.com/james-zhang-bing/okapi"

type (
	GetBalance struct {
		Ccy []string `json:"ccy,omitempty"`
	}
	GetPositions struct {
		InstID   []string             `json:"instId,omitempty"`
		PosID    []string             `json:"posId,omitempty"`
		InstType okapi.InstrumentType `json:"instType,omitempty"`
	}
	GetAccountAndPositionRisk struct {
		InstType okapi.InstrumentType `json:"instType,omitempty"`
	}
	GetBills struct {
		Ccy      string               `json:"ccy,omitempty"`
		After    int64                `json:"after,omitempty,string"`
		Before   int64                `json:"before,omitempty,string"`
		Limit    int64                `json:"limit,omitempty,string"`
		InstType okapi.InstrumentType `json:"instType,omitempty"`
		MgnMode  okapi.MarginMode     `json:"mgnMode,omitempty"`
		CtType   okapi.ContractType   `json:"ctType,omitempty"`
		Type     okapi.BillType       `json:"type,omitempty,string"`
		SubType  okapi.BillSubType    `json:"subType,omitempty,string"`
	}
	SetPositionMode struct {
		PositionMode okapi.PositionType `json:"positionMode"`
	}
	SetLeverage struct {
		Lever   int64              `json:"lever,string"`
		InstID  string             `json:"instId,omitempty"`
		Ccy     string             `json:"ccy,omitempty"`
		MgnMode okapi.MarginMode   `json:"mgnMode"`
		PosSide okapi.PositionSide `json:"posSide,omitempty"`
	}
	GetMaxBuySellAmount struct {
		Ccy    string          `json:"ccy,omitempty"`
		Px     float64         `json:"px,string,omitempty"`
		InstID []string        `json:"instId"`
		TdMode okapi.TradeMode `json:"tdMode"`
	}
	GetMaxAvailableTradeAmount struct {
		Ccy        string          `json:"ccy,omitempty"`
		InstID     string          `json:"instId"`
		ReduceOnly bool            `json:"reduceOnly,omitempty"`
		TdMode     okapi.TradeMode `json:"tdMode"`
	}
	IncreaseDecreaseMargin struct {
		InstID     string             `json:"instId"`
		Amt        float64            `json:"amt,string"`
		PosSide    okapi.PositionSide `json:"posSide"`
		ActionType okapi.CountAction  `json:"actionType"`
	}
	GetLeverage struct {
		InstID  []string         `json:"instId"`
		MgnMode okapi.MarginMode `json:"mgnMode"`
	}
	GetMaxLoan struct {
		InstID  string           `json:"instId"`
		MgnCcy  string           `json:"mgnCcy,omitempty"`
		MgnMode okapi.MarginMode `json:"mgnMode"`
	}
	GetFeeRates struct {
		InstID   string               `json:"instId,omitempty"`
		Uly      string               `json:"uly,omitempty"`
		Category okapi.FeeCategory    `json:"category,omitempty,string"`
		InstType okapi.InstrumentType `json:"instType"`
	}
	GetInterestAccrued struct {
		InstID  string           `json:"instId,omitempty"`
		Ccy     string           `json:"ccy,omitempty"`
		After   int64            `json:"after,omitempty,string"`
		Before  int64            `json:"before,omitempty,string"`
		Limit   int64            `json:"limit,omitempty,string"`
		MgnMode okapi.MarginMode `json:"mgnMode,omitempty"`
	}
	SetGreeks struct {
		GreeksType okapi.GreekType `json:"greeksType"`
	}
)
