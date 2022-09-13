package account

import "github.com/james-zhang-bing/okapi"

type (
	Balance struct {
		TotalEq     okapi.JSONFloat64 `json:"totalEq"`
		IsoEq       okapi.JSONFloat64 `json:"isoEq"`
		AdjEq       okapi.JSONFloat64 `json:"adjEq,omitempty"`
		OrdFroz     okapi.JSONFloat64 `json:"ordFroz,omitempty"`
		Imr         okapi.JSONFloat64 `json:"imr,omitempty"`
		Mmr         okapi.JSONFloat64 `json:"mmr,omitempty"`
		MgnRatio    okapi.JSONFloat64 `json:"mgnRatio,omitempty"`
		NotionalUsd okapi.JSONFloat64 `json:"notionalUsd,omitempty"`
		Details     []*BalanceDetails `json:"details,omitempty"`
		UTime       okapi.JSONTime    `json:"uTime"`
	}
	BalanceDetails struct {
		Ccy           string            `json:"ccy"`
		Eq            okapi.JSONFloat64 `json:"eq"`
		CashBal       okapi.JSONFloat64 `json:"cashBal"`
		IsoEq         okapi.JSONFloat64 `json:"isoEq,omitempty"`
		AvailEq       okapi.JSONFloat64 `json:"availEq,omitempty"`
		DisEq         okapi.JSONFloat64 `json:"disEq"`
		AvailBal      okapi.JSONFloat64 `json:"availBal"`
		FrozenBal     okapi.JSONFloat64 `json:"frozenBal"`
		OrdFrozen     okapi.JSONFloat64 `json:"ordFrozen"`
		Liab          okapi.JSONFloat64 `json:"liab,omitempty"`
		Upl           okapi.JSONFloat64 `json:"upl,omitempty"`
		UplLib        okapi.JSONFloat64 `json:"uplLib,omitempty"`
		CrossLiab     okapi.JSONFloat64 `json:"crossLiab,omitempty"`
		IsoLiab       okapi.JSONFloat64 `json:"isoLiab,omitempty"`
		MgnRatio      okapi.JSONFloat64 `json:"mgnRatio,omitempty"`
		Interest      okapi.JSONFloat64 `json:"interest,omitempty"`
		Twap          okapi.JSONFloat64 `json:"twap,omitempty"`
		MaxLoan       okapi.JSONFloat64 `json:"maxLoan,omitempty"`
		EqUsd         okapi.JSONFloat64 `json:"eqUsd"`
		NotionalLever okapi.JSONFloat64 `json:"notionalLever,omitempty"`
		StgyEq        okapi.JSONFloat64 `json:"stgyEq"`
		IsoUpl        okapi.JSONFloat64 `json:"isoUpl,omitempty"`
		UTime         okapi.JSONTime    `json:"uTime"`
	}
	Position struct {
		InstID      string               `json:"instId"`
		PosCcy      string               `json:"posCcy,omitempty"`
		LiabCcy     string               `json:"liabCcy,omitempty"`
		OptVal      string               `json:"optVal,omitempty"`
		Ccy         string               `json:"ccy"`
		PosID       string               `json:"posId"`
		TradeID     string               `json:"tradeId"`
		Pos         okapi.JSONFloat64    `json:"pos"`
		AvailPos    okapi.JSONFloat64    `json:"availPos,omitempty"`
		AvgPx       okapi.JSONFloat64    `json:"avgPx"`
		Upl         okapi.JSONFloat64    `json:"upl"`
		UplRatio    okapi.JSONFloat64    `json:"uplRatio"`
		Lever       okapi.JSONFloat64    `json:"lever"`
		LiqPx       okapi.JSONFloat64    `json:"liqPx,omitempty"`
		Imr         okapi.JSONFloat64    `json:"imr,omitempty"`
		Margin      okapi.JSONFloat64    `json:"margin,omitempty"`
		MgnRatio    okapi.JSONFloat64    `json:"mgnRatio"`
		Mmr         okapi.JSONFloat64    `json:"mmr"`
		Liab        okapi.JSONFloat64    `json:"liab,omitempty"`
		Interest    okapi.JSONFloat64    `json:"interest"`
		NotionalUsd okapi.JSONFloat64    `json:"notionalUsd"`
		ADL         okapi.JSONFloat64    `json:"adl"`
		Last        okapi.JSONFloat64    `json:"last"`
		DeltaBS     okapi.JSONFloat64    `json:"deltaBS"`
		DeltaPA     okapi.JSONFloat64    `json:"deltaPA"`
		GammaBS     okapi.JSONFloat64    `json:"gammaBS"`
		GammaPA     okapi.JSONFloat64    `json:"gammaPA"`
		ThetaBS     okapi.JSONFloat64    `json:"thetaBS"`
		ThetaPA     okapi.JSONFloat64    `json:"thetaPA"`
		VegaBS      okapi.JSONFloat64    `json:"vegaBS"`
		VegaPA      okapi.JSONFloat64    `json:"vegaPA"`
		PosSide     okapi.PositionSide   `json:"posSide"`
		MgnMode     okapi.MarginMode     `json:"mgnMode"`
		InstType    okapi.InstrumentType `json:"instType"`
		CTime       okapi.JSONTime       `json:"cTime"`
		UTime       okapi.JSONTime       `json:"uTime"`
	}
	BalanceAndPosition struct {
		EventType okapi.EventType   `json:"eventType"`
		PTime     okapi.JSONTime    `json:"pTime"`
		UTime     okapi.JSONTime    `json:"uTime"`
		PosData   []*Position       `json:"posData"`
		BalData   []*BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		AdjEq   okapi.JSONFloat64                    `json:"adjEq,omitempty"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		TS      okapi.JSONTime                       `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string            `json:"ccy"`
		Eq    okapi.JSONFloat64 `json:"eq"`
		DisEq okapi.JSONFloat64 `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		InstID      string               `json:"instId"`
		PosCcy      string               `json:"posCcy,omitempty"`
		Ccy         string               `json:"ccy"`
		NotionalCcy okapi.JSONFloat64    `json:"notionalCcy"`
		Pos         okapi.JSONFloat64    `json:"pos"`
		NotionalUsd okapi.JSONFloat64    `json:"notionalUsd"`
		PosSide     okapi.PositionSide   `json:"posSide"`
		InstType    okapi.InstrumentType `json:"instType"`
		MgnMode     okapi.MarginMode     `json:"mgnMode"`
	}
	Bill struct {
		Ccy       string               `json:"ccy"`
		InstID    string               `json:"instId"`
		Notes     string               `json:"notes"`
		BillID    string               `json:"billId"`
		OrdID     string               `json:"ordId"`
		BalChg    okapi.JSONFloat64    `json:"balChg"`
		PosBalChg okapi.JSONFloat64    `json:"posBalChg"`
		Bal       okapi.JSONFloat64    `json:"bal"`
		PosBal    okapi.JSONFloat64    `json:"posBal"`
		Sz        okapi.JSONFloat64    `json:"sz"`
		Pnl       okapi.JSONFloat64    `json:"pnl"`
		Fee       okapi.JSONFloat64    `json:"fee"`
		From      okapi.AccountType    `json:"from,string"`
		To        okapi.AccountType    `json:"to,string"`
		InstType  okapi.InstrumentType `json:"instType"`
		MgnMode   okapi.MarginMode     `json:"MgnMode"`
		Type      okapi.BillType       `json:"type,string"`
		SubType   okapi.BillSubType    `json:"subType,string"`
		TS        okapi.JSONTime       `json:"ts"`
	}
	Config struct {
		Level      string             `json:"level"`
		LevelTmp   string             `json:"levelTmp"`
		AcctLv     string             `json:"acctLv"`
		AutoLoan   bool               `json:"autoLoan"`
		UID        string             `json:"uid"`
		GreeksType okapi.GreekType    `json:"greeksType"`
		PosMode    okapi.PositionType `json:"posMode"`
	}
	PositionMode struct {
		PosMode okapi.PositionType `json:"posMode"`
	}
	Leverage struct {
		InstID  string             `json:"instId"`
		Lever   okapi.JSONFloat64  `json:"lever"`
		MgnMode okapi.MarginMode   `json:"mgnMode"`
		PosSide okapi.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstID  string            `json:"instId"`
		Ccy     string            `json:"ccy"`
		MaxBuy  okapi.JSONFloat64 `json:"maxBuy"`
		MaxSell okapi.JSONFloat64 `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		InstID    string            `json:"instId"`
		AvailBuy  okapi.JSONFloat64 `json:"availBuy"`
		AvailSell okapi.JSONFloat64 `json:"availSell"`
	}
	MarginBalanceAmount struct {
		InstID  string             `json:"instId"`
		Amt     okapi.JSONFloat64  `json:"amt"`
		PosSide okapi.PositionSide `json:"posSide,string"`
		Type    okapi.CountAction  `json:"type,string"`
	}
	Loan struct {
		InstID  string            `json:"instId"`
		MgnCcy  string            `json:"mgnCcy"`
		Ccy     string            `json:"ccy"`
		MaxLoan okapi.JSONFloat64 `json:"maxLoan"`
		MgnMode okapi.MarginMode  `json:"mgnMode"`
		Side    okapi.OrderSide   `json:"side,string"`
	}
	Fee struct {
		Level    string               `json:"level"`
		Taker    okapi.JSONFloat64    `json:"taker"`
		Maker    okapi.JSONFloat64    `json:"maker"`
		Delivery okapi.JSONFloat64    `json:"delivery,omitempty"`
		Exercise okapi.JSONFloat64    `json:"exercise,omitempty"`
		Category okapi.FeeCategory    `json:"category,string"`
		InstType okapi.InstrumentType `json:"instType"`
		TS       okapi.JSONTime       `json:"ts"`
	}
	InterestAccrued struct {
		InstID       string            `json:"instId"`
		Ccy          string            `json:"ccy"`
		Interest     okapi.JSONFloat64 `json:"interest"`
		InterestRate okapi.JSONFloat64 `json:"interestRate"`
		Liab         okapi.JSONFloat64 `json:"liab"`
		MgnMode      okapi.MarginMode  `json:"mgnMode"`
		TS           okapi.JSONTime    `json:"ts"`
	}
	InterestRate struct {
		Ccy          string            `json:"ccy"`
		InterestRate okapi.JSONFloat64 `json:"interestRate"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string            `json:"ccy"`
		MaxWd okapi.JSONFloat64 `json:"maxWd"`
	}
)
