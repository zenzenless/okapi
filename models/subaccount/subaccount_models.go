package subaccount

import "github.com/james-zhang-bing/okapi"

type (
	SubAccount struct {
		SubAcct string         `json:"subAcct,omitempty"`
		Label   string         `json:"label,omitempty"`
		Mobile  string         `json:"mobile,omitempty"`
		GAuth   bool           `json:"gAuth"`
		Enable  bool           `json:"enable"`
		TS      okapi.JSONTime `json:"ts"`
	}
	APIKey struct {
		SubAcct    string         `json:"subAcct,omitempty"`
		Label      string         `json:"label,omitempty"`
		APIKey     string         `json:"apiKey,omitempty"`
		SecretKey  string         `json:"secretKey,omitempty"`
		Passphrase string         `json:"Passphrase,omitempty"`
		Perm       string         `json:"perm,omitempty"`
		IP         string         `json:"ip,omitempty"`
		TS         okapi.JSONTime `json:"ts,omitempty"`
	}
	HistoryTransfer struct {
		SubAcct string          `json:"subAcct,omitempty"`
		Ccy     string          `json:"ccy,omitempty"`
		BillID  okapi.JSONInt64 `json:"billId,omitempty"`
		Type    okapi.BillType  `json:"type,omitempty,string"`
		TS      okapi.JSONTime  `json:"ts,omitempty"`
	}
	Transfer struct {
		TransID okapi.JSONInt64 `json:"transId"`
	}
)
