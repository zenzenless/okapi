package subaccount

import "github.com/james-zhang-bing/okapi"

type (
	ViewList struct {
		SubAcct string `json:"subAcct,omitempty"`
		Enable  bool   `json:"enable,omitempty"`
		After   int64  `json:"after,omitempty,string"`
		Before  int64  `json:"before,omitempty,string"`
		Limit   int64  `json:"limit,omitempty,string"`
	}
	CreateAPIKey struct {
		Pwd        string             `json:"pwd"`
		SubAcct    string             `json:"subAcct"`
		Label      string             `json:"label"`
		Passphrase string             `json:"Passphrase"`
		IP         []string           `json:"ip,omitempty"`
		Perm       okapi.APIKeyAccess `json:"perm,omitempty"`
	}
	QueryAPIKey struct {
		APIKey  string `json:"apiKey"`
		SubAcct string `json:"subAcct"`
	}
	DeleteAPIKey struct {
		Pwd     string `json:"pwd"`
		APIKey  string `json:"apiKey"`
		SubAcct string `json:"subAcct"`
	}
	GetBalance struct {
		SubAcct string `json:"subAcct"`
	}
	HistoryTransfer struct {
		Ccy     string             `json:"ccy,omitempty"`
		SubAcct string             `json:"subAcct,omitempty"`
		After   int64              `json:"after,omitempty,string"`
		Before  int64              `json:"before,omitempty,string"`
		Limit   int64              `json:"limit,omitempty,string"`
		Type    okapi.TransferType `json:"type,omitempty,string"`
	}
	ManageTransfers struct {
		Ccy            string            `json:"ccy"`
		FromSubAccount string            `json:"fromSubAccount"`
		ToSubAccount   string            `json:"tiSubAccount"`
		Amt            float64           `json:"amt,string"`
		From           okapi.AccountType `json:"from,string"`
		To             okapi.AccountType `json:"to,string"`
	}
)
