package private

import "github.com/james-zhang-bing/okapi"

type (
	Account struct {
		Ccy string `json:"ccy,omitempty"`
	}
	Position struct {
		Uly      string               `json:"uly,omitempty"`
		InstID   string               `json:"instId,omitempty"`
		InstType okapi.InstrumentType `json:"instType"`
	}
	Order struct {
		Uly      string               `json:"uly,omitempty"`
		InstID   string               `json:"instId,omitempty"`
		InstType okapi.InstrumentType `json:"instType"`
	}
	AlgoOrder struct {
		Uly      string               `json:"uly,omitempty"`
		InstID   string               `json:"instId,omitempty"`
		InstType okapi.InstrumentType `json:"instType"`
	}
)
