package tradedata

import "github.com/james-zhang-bing/okapi"

type (
	GetTakerVolume struct {
		Ccy      string               `json:"ccy"`
		Begin    int64                `json:"before,omitempty,string"`
		End      int64                `json:"limit,omitempty,string"`
		InstType okapi.InstrumentType `json:"instType,string"`
		Period   okapi.BarSize        `json:"period,string,omitempty"`
	}
	GetRatio struct {
		Ccy    string        `json:"ccy"`
		Begin  int64         `json:"before,omitempty,string"`
		End    int64         `json:"limit,omitempty,string"`
		Period okapi.BarSize `json:"period,string,omitempty"`
	}
	GetOpenInterestAndVolumeStrike struct {
		Ccy     string        `json:"ccy"`
		ExpTime string        `json:"expTime"`
		Period  okapi.BarSize `json:"period,string,omitempty"`
	}
)
