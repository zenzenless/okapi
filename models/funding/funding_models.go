package funding

import "github.com/james-zhang-bing/okapi"

type (
	Currency struct {
		Ccy         string `json:"ccy"`
		Name        string `json:"name"`
		Chain       string `json:"chain"`
		MinWd       string `json:"minWd"`
		MinFee      string `json:"minFee"`
		MaxFee      string `json:"maxFee"`
		CanDep      bool   `json:"canDep"`
		CanWd       bool   `json:"canWd"`
		CanInternal bool   `json:"canInternal"`
	}
	Balance struct {
		Ccy       string `json:"ccy"`
		Bal       string `json:"bal"`
		FrozenBal string `json:"frozenBal"`
		AvailBal  string `json:"availBal"`
	}
	Transfer struct {
		TransID string           `json:"transId"`
		Ccy     string           `json:"ccy"`
		Amt     okapi.JSONFloat64 `json:"amt"`
		From    okapi.AccountType `json:"from,string"`
		To      okapi.AccountType `json:"to,string"`
	}
	Bill struct {
		BillID string           `json:"billId"`
		Ccy    string           `json:"ccy"`
		Bal    okapi.JSONFloat64 `json:"bal"`
		BalChg okapi.JSONFloat64 `json:"balChg"`
		Type   okapi.BillType    `json:"type,string"`
		TS     okapi.JSONTime    `json:"ts"`
	}
	DepositAddress struct {
		Addr     string           `json:"addr"`
		Tag      string           `json:"tag,omitempty"`
		Memo     string           `json:"memo,omitempty"`
		PmtID    string           `json:"pmtId,omitempty"`
		Ccy      string           `json:"ccy"`
		Chain    string           `json:"chain"`
		CtAddr   string           `json:"ctAddr"`
		Selected bool             `json:"selected"`
		To       okapi.AccountType `json:"to,string"`
		TS       okapi.JSONTime    `json:"ts"`
	}
	DepositHistory struct {
		Ccy   string            `json:"ccy"`
		Chain string            `json:"chain"`
		TxID  string            `json:"txId"`
		From  string            `json:"from"`
		To    string            `json:"to"`
		DepId string            `json:"depId"`
		Amt   okapi.JSONFloat64  `json:"amt"`
		State okapi.DepositState `json:"state,string"`
		TS    okapi.JSONTime     `json:"ts"`
	}
	Withdrawal struct {
		Ccy   string           `json:"ccy"`
		Chain string           `json:"chain"`
		WdID  okapi.JSONInt64   `json:"wdId"`
		Amt   okapi.JSONFloat64 `json:"amt"`
	}
	WithdrawalHistory struct {
		Ccy   string               `json:"ccy"`
		Chain string               `json:"chain"`
		TxID  string               `json:"txId"`
		From  string               `json:"from"`
		To    string               `json:"to"`
		Tag   string               `json:"tag,omitempty"`
		PmtID string               `json:"pmtId,omitempty"`
		Memo  string               `json:"memo,omitempty"`
		Amt   okapi.JSONFloat64     `json:"amt"`
		Fee   okapi.JSONInt64       `json:"fee"`
		WdID  okapi.JSONInt64       `json:"wdId"`
		State okapi.WithdrawalState `json:"state,string"`
		TS    okapi.JSONTime        `json:"ts"`
	}
	PiggyBank struct {
		Ccy  string           `json:"ccy"`
		Amt  okapi.JSONFloat64 `json:"amt"`
		Side okapi.ActionType  `json:"side,string"`
	}
	PiggyBankBalance struct {
		Ccy      string           `json:"ccy"`
		Amt      okapi.JSONFloat64 `json:"amt"`
		Earnings okapi.JSONFloat64 `json:"earnings"`
	}
)
