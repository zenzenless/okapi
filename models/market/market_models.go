package market

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/james-zhang-bing/okapi"
)

type (
	Ticker struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// last String 最新成交价
		Last okapi.JSONFloat64 `json:"last"`
		// lastSz String 最新成交的数量
		LastSz okapi.JSONFloat64 `json:"lastSz"`
		// askPx String 卖一价
		AskPx okapi.JSONFloat64 `json:"askPx"`
		// askSz String 卖一价的挂单数数量
		AskSz okapi.JSONFloat64 `json:"askSz"`
		// bidPx String 买一价
		BidPx okapi.JSONFloat64 `json:"bidPx"`
		// bidSz String 买一价的挂单数量
		BidSz okapi.JSONFloat64 `json:"bidSz"`
		// open24h String 24小时开盘价
		Open24h okapi.JSONFloat64 `json:"open24h"`
		// high24h String 24小时最高价
		High24h okapi.JSONFloat64 `json:"high24h"`
		// low24h String 24小时最低价
		Low24h okapi.JSONFloat64 `json:"low24h"`
		// volCcy24h String 24小时成交量，以 币 为单位 如果是 衍生品 合约，数值为交易货币的数量。 如果是 币币/币币杠杆 ，数值为计价货币的数量。
		VolCcy24h okapi.JSONFloat64 `json:"volCcy24h"`
		// vol24h String 24小时成交量，以 张 为单位 如果是 衍生品 合约，数值为合约的张数。 如果是 币币/币币杠杆 ，数值为交易货币的数量。
		Vol24h okapi.JSONFloat64 `json:"vol24h"`
		// sodUtc0 String UTC 0 时开盘价
		SodUtc0 okapi.JSONFloat64 `json:"sodUtc0"`
		// sodUtc8 String UTC+8 时开盘价
		SodUtc8 okapi.JSONFloat64 `json:"sodUtc8"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	IndexTicker struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// idxPx String 最新指数价格
		IdxPx okapi.JSONFloat64 `json:"idxPx"`
		// high24h String 24小时最高价
		High24h okapi.JSONFloat64 `json:"high24h"`
		// low24h String 24小时最低价
		Low24h okapi.JSONFloat64 `json:"low24h"`
		// open24h String 24小时开盘价
		Open24h okapi.JSONFloat64 `json:"open24h"`
		// sodUtc0 String UTC 0 时开盘价
		SodUtc0 okapi.JSONFloat64 `json:"sodUtc0"`
		// sodUtc8 String UTC+8 时开盘价
		SodUtc8 okapi.JSONFloat64 `json:"sodUtc8"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	OrderBook struct {
		// asks Array  卖方深度
		Asks []*OrderBookEntity `json:"asks"`
		// bids Array  买方深度
		Bids []*OrderBookEntity `json:"bids"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	OrderBookWs struct {
		// asks Array  卖方深度
		Asks []*OrderBookEntity `json:"asks"`
		// bids Array  买方深度
		Bids []*OrderBookEntity `json:"bids"`
		// checksum Integer 检验和 （下方注解）
		Checksum int `json:"checksum,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	OrderBookEntity struct {
		DepthPrice      float64
		Size            float64
		LiquidatedOrder int
		OrderNumbers    int
	}
	Candle struct {
		O           float64
		H           float64
		L           float64
		C           float64
		Vol         float64
		VolCcy      float64
		VolCcyQuote float64
		TS          okapi.JSONTime
		Confirm     int
	}
	IndexCandle struct {
		O  float64
		H  float64
		L  float64
		C  float64
		TS okapi.JSONTime
	}
	Trade struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// tradeId String 最新成交ID
		TradeID okapi.JSONFloat64 `json:"tradeId"`
		// px String 可选  委托价格，仅适用于 limit 、 post_only 、 fok 、 ioc 类型的订单
		Px okapi.JSONFloat64 `json:"px"`
		// sz String 是 委托数量
		Sz okapi.JSONFloat64 `json:"sz"`
		// side String 是 订单方向   buy ：买，  sell ：卖
		Side okapi.TradeSide `json:"side,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	TotalVolume24H struct {
		// volUsd String 24小时平台总成交量，以美元为单位
		VolUsd okapi.JSONFloat64 `json:"volUsd"`
		// volCny String 24小时平台总成交量，以人民币为单位
		VolCny okapi.JSONFloat64 `json:"volCny"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	IndexComponent struct {
		// index String 是 指数，如  BTC-USDT
		Index string `json:"index"`
		// last String 最新成交价
		Last okapi.JSONFloat64 `json:"last"`
		// components String 成分
		Components []*Component `json:"components"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	Component struct {
		// exch String 交易所名称
		Exch string `json:"exch"`
		// symbol String 采集的币对名称
		Symbol string `json:"symbol"`
		// symPx String 采集的币对价格
		SymPx okapi.JSONFloat64 `json:"symPx"`
		// wgt String 权重
		Wgt okapi.JSONFloat64 `json:"wgt"`
		// cnvPx String 换算成指数后的价格
		CnvPx okapi.JSONFloat64 `json:"cnvPx"`
	}
)

func (o *OrderBookEntity) UnmarshalJSON(buf []byte) error {
	var (
		dp, s, lo, on string
		err           error
	)
	tmp := []interface{}{&dp, &s, &lo, &on}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in OrderBookEntity: %d != %d", g, e)
	}
	o.DepthPrice, err = strconv.ParseFloat(dp, 64)
	if err != nil {
		return err
	}
	o.Size, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	o.LiquidatedOrder, err = strconv.Atoi(lo)
	if err != nil {
		return err
	}
	o.OrderNumbers, err = strconv.Atoi(on)
	if err != nil {
		return err
	}

	return nil
}

func (c *Candle) UnmarshalJSON(buf []byte) error {
	var (
		o, h, l, cl, vol, volCcy, ts, volCcyQuote, confirm string
		err                                                error
	)
	tmp := []interface{}{&ts, &o, &h, &l, &cl, &vol, &volCcy, &volCcyQuote, &confirm}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d buf:%s tmp:%+v", g, e, string(buf), tmp)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	c.O, err = strconv.ParseFloat(o, 64)
	if err != nil {
		return err
	}

	c.H, err = strconv.ParseFloat(h, 64)
	if err != nil {
		return err
	}

	c.L, err = strconv.ParseFloat(l, 64)
	if err != nil {
		return err
	}

	c.C, err = strconv.ParseFloat(cl, 64)
	if err != nil {
		return err
	}

	c.Vol, err = strconv.ParseFloat(vol, 64)
	if err != nil {
		return err
	}

	c.VolCcy, err = strconv.ParseFloat(volCcy, 64)
	if err != nil {
		return err
	}
	c.Confirm, err = strconv.Atoi(confirm)
	if err != nil {
		return err
	}

	return nil
}

func (c *IndexCandle) UnmarshalJSON(buf []byte) error {
	var (
		o, h, l, cl, ts string
		err             error
	)
	tmp := []interface{}{&ts, &o, &h, &l, &cl}
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

	c.O, err = strconv.ParseFloat(o, 64)
	if err != nil {
		return err
	}

	c.H, err = strconv.ParseFloat(h, 64)
	if err != nil {
		return err
	}

	c.L, err = strconv.ParseFloat(l, 64)
	if err != nil {
		return err
	}

	c.C, err = strconv.ParseFloat(cl, 64)
	if err != nil {
		return err
	}

	return nil
}
func (m *Ticker) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Ticker ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Last); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交价: %v", str, m.Last)
	}
	if s := fmt.Sprintf("%v", m.LastSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交的数量: %v", str, m.LastSz)
	}
	if s := fmt.Sprintf("%v", m.AskPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n卖一价: %v", str, m.AskPx)
	}
	if s := fmt.Sprintf("%v", m.AskSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n卖一价的挂单数数量: %v", str, m.AskSz)
	}
	if s := fmt.Sprintf("%v", m.BidPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n买一价: %v", str, m.BidPx)
	}
	if s := fmt.Sprintf("%v", m.BidSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n买一价的挂单数量: %v", str, m.BidSz)
	}
	if s := fmt.Sprintf("%v", m.Open24h); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n24小时开盘价: %v", str, m.Open24h)
	}
	if s := fmt.Sprintf("%v", m.High24h); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n24小时最高价: %v", str, m.High24h)
	}
	if s := fmt.Sprintf("%v", m.Low24h); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n24小时最低价: %v", str, m.Low24h)
	}
	if s := fmt.Sprintf("%v", m.VolCcy24h); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n24小时成交量: %v", str, m.VolCcy24h)
	}
	if s := fmt.Sprintf("%v", m.Vol24h); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n24小时成交量: %v", str, m.Vol24h)
	}
	if s := fmt.Sprintf("%v", m.SodUtc0); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n时开盘价: %v", str, m.SodUtc0)
	}
	if s := fmt.Sprintf("%v", m.SodUtc8); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n时开盘价: %v", str, m.SodUtc8)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- Ticker ------------┘", str)
	return str
}
func (m *IndexTicker) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ IndexTicker ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.IdxPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新指数价格: %v", str, m.IdxPx)
	}
	if s := fmt.Sprintf("%v", m.High24h); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n24小时最高价: %v", str, m.High24h)
	}
	if s := fmt.Sprintf("%v", m.Low24h); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n24小时最低价: %v", str, m.Low24h)
	}
	if s := fmt.Sprintf("%v", m.Open24h); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n24小时开盘价: %v", str, m.Open24h)
	}
	if s := fmt.Sprintf("%v", m.SodUtc0); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n时开盘价: %v", str, m.SodUtc0)
	}
	if s := fmt.Sprintf("%v", m.SodUtc8); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n时开盘价: %v", str, m.SodUtc8)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- IndexTicker ------------┘", str)
	return str
}
func (m *OrderBook) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ OrderBook ------┐", str)
	if s := fmt.Sprintf("%v", m.Asks); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n卖方深度: %v", str, m.Asks)
	}
	if s := fmt.Sprintf("%v", m.Bids); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n买方深度: %v", str, m.Bids)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- OrderBook ------------┘", str)
	return str
}
func (m *OrderBookWs) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ OrderBookWs ------┐", str)
	if s := fmt.Sprintf("%v", m.Asks); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n卖方深度: %v", str, m.Asks)
	}
	if s := fmt.Sprintf("%v", m.Bids); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n买方深度: %v", str, m.Bids)
	}
	if s := fmt.Sprintf("%v", m.Checksum); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n检验和: %v", str, m.Checksum)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- OrderBookWs ------------┘", str)
	return str
}
func (m *Trade) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Trade ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.TradeID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交ID: %v", str, m.TradeID)
	}
	if s := fmt.Sprintf("%v", m.Px); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n委托价格: %v", str, m.Px)
	}
	if s := fmt.Sprintf("%v", m.Sz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n委托数量: %v", str, m.Sz)
	}
	if s := fmt.Sprintf("%v", m.Side); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单方向: %v", str, m.Side)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- Trade ------------┘", str)
	return str
}
func (m *TotalVolume24H) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ TotalVolume24H ------┐", str)
	if s := fmt.Sprintf("%v", m.VolUsd); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n24小时平台总成交量: %v", str, m.VolUsd)
	}
	if s := fmt.Sprintf("%v", m.VolCny); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n24小时平台总成交量: %v", str, m.VolCny)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- TotalVolume24H ------------┘", str)
	return str
}
func (m *IndexComponent) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ IndexComponent ------┐", str)
	if s := fmt.Sprintf("%v", m.Index); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n指数: %v", str, m.Index)
	}
	if s := fmt.Sprintf("%v", m.Last); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交价: %v", str, m.Last)
	}
	if s := fmt.Sprintf("%v", m.Components); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成分: %v", str, m.Components)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- IndexComponent ------------┘", str)
	return str
}
func (m *Component) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Component ------┐", str)
	if s := fmt.Sprintf("%v", m.Exch); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n交易所名称: %v", str, m.Exch)
	}
	if s := fmt.Sprintf("%v", m.Symbol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n采集的币对名称: %v", str, m.Symbol)
	}
	if s := fmt.Sprintf("%v", m.SymPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n采集的币对价格: %v", str, m.SymPx)
	}
	if s := fmt.Sprintf("%v", m.Wgt); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n权重: %v", str, m.Wgt)
	}
	if s := fmt.Sprintf("%v", m.CnvPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n换算成指数后的价格: %v", str, m.CnvPx)
	}
	str = fmt.Sprintf("%s\r\n└----------- Component ------------┘", str)
	return str
}
func (m *OrderBookEntity) String() string {
	var str string
	str = fmt.Sprintf("%s\r\n┌------ OrderBookEntity ------┐", str)
	if s := fmt.Sprintf("%v", m.DepthPrice); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nDepthPrice:%v", str, m.DepthPrice)
	}
	if s := fmt.Sprintf("%v", m.Size); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nSize:%v", str, m.Size)
	}
	if s := fmt.Sprintf("%v", m.LiquidatedOrder); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nLiquidatedOrder:%v", str, m.LiquidatedOrder)
	}
	if s := fmt.Sprintf("%v", m.OrderNumbers); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nOrderNumbers:%v", str, m.OrderNumbers)
	}
	str = fmt.Sprintf("%s\r\n└----------- OrderBookEntity ------------┘", str)
	return str
}
func (m *Candle) String() string {
	var str string
	str = fmt.Sprintf("%s\r\n┌------ Candle ------┐", str)
	if s := fmt.Sprintf("%v", m.O); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nO:%v", str, m.O)
	}
	if s := fmt.Sprintf("%v", m.H); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nH:%v", str, m.H)
	}
	if s := fmt.Sprintf("%v", m.L); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nL:%v", str, m.L)
	}
	if s := fmt.Sprintf("%v", m.C); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nC:%v", str, m.C)
	}
	if s := fmt.Sprintf("%v", m.Vol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nVol:%v", str, m.Vol)
	}
	if s := fmt.Sprintf("%v", m.VolCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nVolCcy:%v", str, m.VolCcy)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nTS:%v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- Candle ------------┘", str)
	return str
}
func (m *IndexCandle) String() string {
	var str string
	str = fmt.Sprintf("%s\r\n┌------ IndexCandle ------┐", str)
	if s := fmt.Sprintf("%v", m.O); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nO:%v", str, m.O)
	}
	if s := fmt.Sprintf("%v", m.H); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nH:%v", str, m.H)
	}
	if s := fmt.Sprintf("%v", m.L); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nL:%v", str, m.L)
	}
	if s := fmt.Sprintf("%v", m.C); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nC:%v", str, m.C)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nTS:%v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- IndexCandle ------------┘", str)
	return str
}
