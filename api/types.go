package api

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

// ErrCode is the error code returned by the API
type ErrCode struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// Tickers is the response of the Restful API api/v5/market/tickers
type Tickers struct {
	Code string              `json:"code"`
	Msg  string              `json:"msg"`
	Data []map[string]string `json:"data"`
}

// Kline candle data
type Candle struct {
	Time   int64   `json:"time"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
}

// KlineResOKX is the response of the Restful API api/v5/market/candles
type KlineResOKX struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

// generate by copilot
const (
	BAR_1m  = "1m"
	BAR_3m  = "3m"
	BAR_5m  = "5m"
	BAR_15m = "15m"
	BAR_30m = "30m"
	BAR_1h  = "1H"
	BAR_2h  = "2H"
	BAR_4h  = "4H"
	BAR_6h  = "6H"
	BAR_12h = "12H"
	BAR_1d  = "1D"
	BAR_1w  = "1W"
	BAR_1M  = "1M"
)

var BarMap = map[string]time.Duration{
	BAR_1m:  time.Minute,
	BAR_3m:  time.Minute * 3,
	BAR_5m:  time.Minute * 5,
	BAR_15m: time.Minute * 15,
	BAR_30m: time.Minute * 30,
	BAR_1h:  time.Hour,
	BAR_2h:  time.Hour * 2,
	BAR_4h:  time.Hour * 4,
	BAR_6h:  time.Hour * 6,
	BAR_12h: time.Hour * 12,
	BAR_1d:  time.Hour * 24,
	BAR_1w:  time.Hour * 24 * 7,
	BAR_1M:  time.Hour * 24 * 30,
}

// Unmarshl KlineRes to techan.TimeSeries
func (k *KlineResOKX) ToTimeSeries(dur string) (*techan.TimeSeries, error) {
	bar, ok := BarMap[dur]
	if !ok {
		return nil, fmt.Errorf("invalid duration %s", dur)
	}
	series := techan.NewTimeSeries()
	for _, candle := range k.Data {
		ts, err := strconv.ParseInt(candle[0], 10, 64)
		if err != nil {
			return nil, err
		}

		period := techan.NewTimePeriod(time.UnixMilli(ts), bar)
		// fmt.Printf("%v\n", period)
		tCandle := techan.NewCandle(period)
		tCandle.OpenPrice = big.NewFromString(candle[1])

		tCandle.MaxPrice = big.NewFromString(candle[2])
		tCandle.MinPrice = big.NewFromString(candle[3])

		tCandle.ClosePrice = big.NewFromString(candle[4])
		tCandle.Volume = big.NewFromString(candle[5])
		series.Candles = append(series.Candles, tCandle)
	}
	sort.Slice(series.Candles, func(i, j int) bool {
		return series.Candles[i].Period.Start.Unix() < series.Candles[j].Period.Start.Unix()
	})
	return series, nil
}

// OkxOrderRes is the response of the Restful API /api/v5/trade/orders-history-archive
type OkxOrderRes struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data []OkxOrder `json:"data"`
}

//
type OkxOrder struct {
	//产品类型
	InstType string `json:"instType"`
	//产品ID
	InstID string `json:"instId"`
	// 保证金币种，仅适用于单币种保证金模式下的全仓币币杠杆订单
	Ccy string `json:"ccy"`
	// 订单ID
	OrdID string `json:"ordId"`
	// 客户自定义订单ID
	ClOrdID string `json:"clOrdId"`
	// 订单标签
	Tag string `json:"tag"`
	//px 委托价格
	Px string `json:"px"`
	// Sz 委托数量
	Sz string `json:"sz"`
	// OrdType  订单类型market：市价单 limit：限价单 post_only：只做maker单 fok：全部成交或立即取消 ioc：立即成交并取消剩余 optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续）
	OrdType string `json:"ordType"`
	// Side 订单方向
	Side string `json:"side"`
	// PosSide 持仓方向
	PosSide string `json:"posSide"`
	//TdMode  交易模式
	TdMode string `json:"tdMode"`
	// 累计成交数量
	AccFillSz string `json:"accFillSz"`
	// fillPx 最新成交价格
	FillPx string `json:"fillPx"`
	// TradeID 最新成交ID
	TradeID string `json:"tradeId"`
	// 最新成交数量
	FillSz string `json:"fillSz"`
	// 最新成交时间
	FillTime string `json:"fillTime"`
	// 订单来源 13:策略委托单触发后的生成的限价单
	Source string `json:"source"`
	// 订单状态
	State string `json:"state"`
	// 成交均价
	AvgPx string `json:"avgPx"`
	//杠杆倍数，0.01到125之间的数值，仅适用于 币币杠杆/交割/永续
	Lever string `json:"lever"`
	// 止盈触发价
	TpTriggerPx string `json:"tpTriggerPx"`
	//止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格
	TpTriggerPxType string `json:"tpTriggerPxType"`
	// 止盈委托价
	TpOrdPx string `json:"tpOrdPx"`
	//  止损触发价
	SlTriggerPx string `json:"slTriggerPx"`
	// 止损触发价类型 last：最新价格 index：指数价格 mark：标记价格
	SlTriggerPxType string `json:"slTriggerPxType"`
	//止损委托价
	SlOrdPx string `json:"slOrdPx"`
	//  交易手续费币种
	FeeCcy string `json:"feeCcy"`
	// 交易手续费
	Fee string `json:"fee"`
	//  返佣金币种
	RebateCcy string `json:"rebateCcy"`
	//返佣金额，平台向达到指定lv交易等级的用户支付的挂单奖励（返佣），如果没有返佣金，该字段为“”。手续费返佣为正数，如：0.01
	Rebate string `json:"rebate"`
	//  市价单委托数量的类型base_ccy: 交易货币 ；quote_ccy：计价货币
	TgtCcy string `json:"tgtCcy"`
	// 收益
	Pnl string `json:"pnl"`
	//订单种类    normal：普通委托 twap：TWAP自动换币 adl：ADL自动减仓 full_liquidation：强制平仓 partial_liquidation：强制减仓 delivery：交割 ddh：对冲减仓类型订单
	Category string `json:"category"`
	// 订单状态更新时间，Unix时间戳的毫秒数格式，如1597026383085
	UTime string `json:"uTime"`
	// 订单创建时间，Unix时间戳的毫秒数格式，如 1597026383085
	CTime string `json:"cTime"`
}

/*
参数名	类型	描述
instType	String	产品类型
instId	String	产品 ID
tradeId	String	最新成交 ID
ordId	String	订单 ID
clOrdId	String	用户自定义订单ID
billId	String	账单 ID
tag	String	订单标签
fillPx	String	最新成交价格
fillSz	String	最新成交数量
side	String	订单方向 buy：买 sell：卖
posSide	String	持仓方向 long：多 short：空 单向持仓模式返回 net
execType	String	流动性方向 T：taker M：maker
feeCcy	String	交易手续费币种或者返佣金币种
fee	String	手续费金额或者返佣金额 ，手续费扣除 为 ‘负数’，如 -0.01 ； 手续费返佣 为 ‘正数’，如 0.01
ts	String	成交明细产生时间，Unix 时间戳的毫秒数格式，如 1597026383085
*/

type FillsRes struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []Fill `json:"data"`
}

type Fill struct {
	InstType string `json:"instType"`
	InstID   string `json:"instId"`
	TradeID  string `json:"tradeId"`
	OrdID    string `json:"ordId"`
	ClOrdID  string `json:"clOrdId"`
	BillID   string `json:"billId"`
	Tag      string `json:"tag"`
	FillPx   string `json:"fillPx"`
	FillSz   string `json:"fillSz"`
	Side     string `json:"side"`
	PosSide  string `json:"posSide"`
	ExecType string `json:"execType"`
	FeeCcy   string `json:"feeCcy"`
	Fee      string `json:"fee"`
	Ts       string `json:"ts"`
}
