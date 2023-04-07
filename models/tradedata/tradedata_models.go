package tradedata

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/james-zhang-bing/okapi"
)

type (
	SupportCoin struct {
		// contract Array 合约交易大数据接口功能支持的币种
		Contract []string `json:"contract"`
		// option Array 期权交易大数据接口功能支持的币种
		Option []string `json:"option"`
		// spot Array 现货交易大数据接口功能支持的币种
		Spot []string `json:"spot"`
	}
	TakerVolume struct {
		SellVol float64
		BuyVol  float64
		TS      okapi.JSONTime
	}
	Ratio struct {
		Ratio float64
		TS    okapi.JSONTime
	}
	InterestAndVolumeRatio struct {
		Oi  float64
		Vol float64
		TS  okapi.JSONTime
	}
	PutCallRatio struct {
		OiRatio  float64
		VolRatio float64
		TS       okapi.JSONTime
	}
	InterestAndVolumeExpiry struct {
		CallOI  float64
		PutOI   float64
		CallVol float64
		PutVol  float64
		ExpTime okapi.JSONTime
		TS      okapi.JSONTime
	}
	InterestAndVolumeStrike struct {
		Strike  float64
		CallOI  float64
		PutOI   float64
		CallVol float64
		PutVol  float64
		TS      okapi.JSONTime
	}
	TakerFlow struct {
		CallBuyVol   float64
		CallSellVol  float64
		PutBuyVol    float64
		PutSellVol   float64
		CallBlockVol float64
		PutBlockVol  float64
		TS           okapi.JSONTime
	}
)

func (c *TakerVolume) UnmarshalJSON(buf []byte) error {
	var (
		sellVol, buyVol, ts string
		err                 error
	)
	tmp := []interface{}{&ts, &sellVol, &buyVol}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	c.SellVol, err = strconv.ParseFloat(sellVol, 64)
	if err != nil {
		return err
	}

	c.BuyVol, err = strconv.ParseFloat(buyVol, 64)
	if err != nil {
		return err
	}

	return nil
}

func (c *Ratio) UnmarshalJSON(buf []byte) error {
	var (
		ratio, ts string
		err       error
	)
	tmp := []interface{}{&ts, &ratio}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	c.Ratio, err = strconv.ParseFloat(ratio, 64)
	if err != nil {
		return err
	}

	return nil
}

func (c *InterestAndVolumeRatio) UnmarshalJSON(buf []byte) error {
	var (
		oi, vol, ts string
		err         error
	)
	tmp := []interface{}{&ts, &oi, &vol}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	if oi != "" {
		c.Oi, err = strconv.ParseFloat(oi, 64)
		if err != nil {
			return err
		}
	}
	if vol != "" {
		c.Vol, err = strconv.ParseFloat(vol, 64)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *PutCallRatio) UnmarshalJSON(buf []byte) error {
	var (
		oi, vol, ts string
		err         error
	)
	tmp := []interface{}{&ts, &oi, &vol}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	if oi != "" {
		c.OiRatio, err = strconv.ParseFloat(oi, 64)
		if err != nil {
			return err
		}
	}
	if vol != "" {
		c.VolRatio, err = strconv.ParseFloat(vol, 64)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *InterestAndVolumeExpiry) UnmarshalJSON(buf []byte) error {
	var (
		callOI, putOI, callVol, putVol, expTime, ts string
		err                                         error
	)
	tmp := []interface{}{&ts, &expTime, &callOI, &putOI, &callVol, &putVol}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	exp, err := time.Parse("20060102", expTime)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.ExpTime) = exp

	if callOI != "" {
		c.CallOI, err = strconv.ParseFloat(callOI, 64)
		if err != nil {
			return err
		}
	}
	if putOI != "" {
		c.PutOI, err = strconv.ParseFloat(putOI, 64)
		if err != nil {
			return err
		}
	}
	if callVol != "" {
		c.CallVol, err = strconv.ParseFloat(callVol, 64)
		if err != nil {
			return err
		}
	}
	if putVol != "" {
		c.PutVol, err = strconv.ParseFloat(putVol, 64)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *InterestAndVolumeStrike) UnmarshalJSON(buf []byte) error {
	var (
		callOI, putOI, callVol, putVol, strike, ts string
		err                                        error
	)
	tmp := []interface{}{&ts, &strike, &callOI, &putOI, &callVol, &putVol}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	if callOI != "" {
		c.CallOI, err = strconv.ParseFloat(callOI, 64)
		if err != nil {
			return err
		}
	}
	if putOI != "" {
		c.PutOI, err = strconv.ParseFloat(putOI, 64)
		if err != nil {
			return err
		}
	}
	if callVol != "" {
		c.CallVol, err = strconv.ParseFloat(callVol, 64)
		if err != nil {
			return err
		}
	}
	if putVol != "" {
		c.PutVol, err = strconv.ParseFloat(putVol, 64)
		if err != nil {
			return err
		}
	}
	if strike != "" {
		c.Strike, err = strconv.ParseFloat(strike, 64)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *TakerFlow) UnmarshalJSON(buf []byte) error {
	var (
		ts, callBuyVol, callSellVol, putBuyVol, putSellVol, callBlockVol, putBlockVol string
		err                                                                           error
	)
	tmp := []interface{}{&ts, &callBuyVol, &callSellVol, &putBuyVol, &putSellVol, &callBlockVol, &putBlockVol}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	if callBuyVol != "" {
		c.CallBlockVol, err = strconv.ParseFloat(callBuyVol, 64)
		if err != nil {
			return err
		}
	}
	if callSellVol != "" {
		c.CallSellVol, err = strconv.ParseFloat(callSellVol, 64)
		if err != nil {
			return err
		}
	}
	if putBuyVol != "" {
		c.PutBuyVol, err = strconv.ParseFloat(putBuyVol, 64)
		if err != nil {
			return err
		}
	}
	if putSellVol != "" {
		c.PutSellVol, err = strconv.ParseFloat(putSellVol, 64)
		if err != nil {
			return err
		}
	}
	if callBlockVol != "" {
		c.CallBuyVol, err = strconv.ParseFloat(callBlockVol, 64)
		if err != nil {
			return err
		}
	}
	if putBlockVol != "" {
		c.PutBuyVol, err = strconv.ParseFloat(putBlockVol, 64)
		if err != nil {
			return err
		}
	}

	return nil
}
func (m *SupportCoin) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ SupportCoin ------┐", str)
	if s := fmt.Sprintf("%v", m.Contract); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n合约交易大数据接口功能支持的币种: %v", str, m.Contract)
	}
	if s := fmt.Sprintf("%v", m.Option); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n期权交易大数据接口功能支持的币种: %v", str, m.Option)
	}
	if s := fmt.Sprintf("%v", m.Spot); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n现货交易大数据接口功能支持的币种: %v", str, m.Spot)
	}
	str = fmt.Sprintf("%s\r\n└----------- SupportCoin ------------┘", str)
	return str
}
func (m *TakerVolume) String() string {
	var str string
	str = fmt.Sprintf("%s\r\n┌------ TakerVolume ------┐", str)
	if s := fmt.Sprintf("%v", m.SellVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nSellVol:%v", str, m.SellVol)
	}
	if s := fmt.Sprintf("%v", m.BuyVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nBuyVol:%v", str, m.BuyVol)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nTS:%v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- TakerVolume ------------┘", str)
	return str
}
func (m *Ratio) String() string {
	var str string
	str = fmt.Sprintf("%s\r\n┌------ Ratio ------┐", str)
	if s := fmt.Sprintf("%v", m.Ratio); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nRatio:%v", str, m.Ratio)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nTS:%v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- Ratio ------------┘", str)
	return str
}
func (m *InterestAndVolumeRatio) String() string {
	var str string
	str = fmt.Sprintf("%s\r\n┌------ InterestAndVolumeRatio ------┐", str)
	if s := fmt.Sprintf("%v", m.Oi); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nOi:%v", str, m.Oi)
	}
	if s := fmt.Sprintf("%v", m.Vol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nVol:%v", str, m.Vol)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nTS:%v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- InterestAndVolumeRatio ------------┘", str)
	return str
}
func (m *PutCallRatio) String() string {
	var str string
	str = fmt.Sprintf("%s\r\n┌------ PutCallRatio ------┐", str)
	if s := fmt.Sprintf("%v", m.OiRatio); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nOiRatio:%v", str, m.OiRatio)
	}
	if s := fmt.Sprintf("%v", m.VolRatio); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nVolRatio:%v", str, m.VolRatio)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nTS:%v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- PutCallRatio ------------┘", str)
	return str
}
func (m *InterestAndVolumeExpiry) String() string {
	var str string
	str = fmt.Sprintf("%s\r\n┌------ InterestAndVolumeExpiry ------┐", str)
	if s := fmt.Sprintf("%v", m.CallOI); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nCallOI:%v", str, m.CallOI)
	}
	if s := fmt.Sprintf("%v", m.PutOI); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nPutOI:%v", str, m.PutOI)
	}
	if s := fmt.Sprintf("%v", m.CallVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nCallVol:%v", str, m.CallVol)
	}
	if s := fmt.Sprintf("%v", m.PutVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nPutVol:%v", str, m.PutVol)
	}
	if s := fmt.Sprintf("%v", m.ExpTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nExpTime:%v", str, m.ExpTime)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nTS:%v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- InterestAndVolumeExpiry ------------┘", str)
	return str
}
func (m *InterestAndVolumeStrike) String() string {
	var str string
	str = fmt.Sprintf("%s\r\n┌------ InterestAndVolumeStrike ------┐", str)
	if s := fmt.Sprintf("%v", m.Strike); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nStrike:%v", str, m.Strike)
	}
	if s := fmt.Sprintf("%v", m.CallOI); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nCallOI:%v", str, m.CallOI)
	}
	if s := fmt.Sprintf("%v", m.PutOI); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nPutOI:%v", str, m.PutOI)
	}
	if s := fmt.Sprintf("%v", m.CallVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nCallVol:%v", str, m.CallVol)
	}
	if s := fmt.Sprintf("%v", m.PutVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nPutVol:%v", str, m.PutVol)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nTS:%v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- InterestAndVolumeStrike ------------┘", str)
	return str
}
func (m *TakerFlow) String() string {
	var str string
	str = fmt.Sprintf("%s\r\n┌------ TakerFlow ------┐", str)
	if s := fmt.Sprintf("%v", m.CallBuyVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nCallBuyVol:%v", str, m.CallBuyVol)
	}
	if s := fmt.Sprintf("%v", m.CallSellVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nCallSellVol:%v", str, m.CallSellVol)
	}
	if s := fmt.Sprintf("%v", m.PutBuyVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nPutBuyVol:%v", str, m.PutBuyVol)
	}
	if s := fmt.Sprintf("%v", m.PutSellVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nPutSellVol:%v", str, m.PutSellVol)
	}
	if s := fmt.Sprintf("%v", m.CallBlockVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nCallBlockVol:%v", str, m.CallBlockVol)
	}
	if s := fmt.Sprintf("%v", m.PutBlockVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nPutBlockVol:%v", str, m.PutBlockVol)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nTS:%v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- TakerFlow ------------┘", str)
	return str
}
