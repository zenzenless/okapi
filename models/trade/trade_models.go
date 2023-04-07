package trade

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
)

type (
	PlaceOrder struct {
		// clOrdId String 否 客户自定义订单ID  字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID string `json:"clOrdId"`
		// tag String 否 订单标签  字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		Tag string `json:"tag"`
		// sMsg String 事件执行失败时的msg
		SMsg string `json:"sMsg"`
		// sCode String 事件执行结果的code，0代表成功
		SCode okapi.JSONInt64 `json:"sCode"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID okapi.JSONFloat64 `json:"ordId"`
	}
	CancelOrder struct {
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID string `json:"ordId"`
		// clOrdId String 否 客户自定义订单ID  字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID string `json:"clOrdId"`
		// sMsg String 事件执行失败时的msg
		SMsg string `json:"sMsg"`
		// sCode String 事件执行结果的code，0代表成功
		SCode okapi.JSONFloat64 `json:"sCode"`
	}
	AmendOrder struct {
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID string `json:"ordId"`
		// clOrdId String 否 客户自定义订单ID  字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID string `json:"clOrdId"`
		// reqId String 否 用户自定义修改事件ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
		ReqID string `json:"reqId"`
		// sMsg String 事件执行失败时的msg
		SMsg string `json:"sMsg"`
		// sCode String 事件执行结果的code，0代表成功
		SCode okapi.JSONFloat64 `json:"sCode"`
	}
	ClosePosition struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okapi.PositionSide `json:"posSide"`
	}
	Order struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID string `json:"ordId"`
		// clOrdId String 否 客户自定义订单ID  字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID string `json:"clOrdId"`
		// tradeId String 最新成交ID
		TradeID string `json:"tradeId"`
		// tag String 否 订单标签  字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		Tag string `json:"tag"`
		// category String 订单种类  normal ：普通委托 twap ：TWAP自动换币adl ：ADL自动减仓   full_liquidation ：强制平仓 partial_liquidation ：强制减仓    delivery ：交割   ddh ：对冲减仓类型订单
		Category string `json:"category"`
		// feeCcy String 交易手续费币种
		FeeCcy string `json:"feeCcy"`
		// rebateCcy String 返佣金币种
		RebateCcy string `json:"rebateCcy"`
		// px String 可选  委托价格，仅适用于 limit 、 post_only 、 fok 、 ioc 类型的订单
		Px okapi.JSONFloat64 `json:"px"`
		// sz String 是 委托数量
		Sz okapi.JSONInt64 `json:"sz"`
		// pnl String 收益
		Pnl okapi.JSONFloat64 `json:"pnl"`
		// accFillSz String 累计成交数量
		AccFillSz okapi.JSONInt64 `json:"accFillSz"`
		// fillPx String 最新成交价格，如果成交数量为0，该字段也为 0
		FillPx okapi.JSONFloat64 `json:"fillPx"`
		// fillSz String 最新成交数量
		FillSz okapi.JSONInt64 `json:"fillSz"`
		// fillTime String 最新成交时间
		FillTime okapi.JSONFloat64 `json:"fillTime"`
		// avgPx String 成交均价，如果成交数量为0，该字段也为 0
		AvgPx okapi.JSONFloat64 `json:"avgPx"`
		// lever String 杠杆倍数，0.01到125之间的数值，仅适用于  币币杠杆/交割/永续
		Lever okapi.JSONFloat64 `json:"lever"`
		// tpTriggerPx String 止盈触发价
		TpTriggerPx okapi.JSONFloat64 `json:"tpTriggerPx"`
		// tpOrdPx String 止盈委托价
		TpOrdPx okapi.JSONFloat64 `json:"tpOrdPx"`
		// slTriggerPx String 止损触发价
		SlTriggerPx okapi.JSONFloat64 `json:"slTriggerPx"`
		// slOrdPx String 止损委托价
		SlOrdPx okapi.JSONFloat64 `json:"slOrdPx"`
		// fee String 订单交易手续费，平台向用户收取的交易手续费，手续费扣除为 负数 。如：  -0.01
		Fee okapi.JSONFloat64 `json:"fee"`
		// rebate String 返佣金额，平台向达到指定lv交易等级的用户支付的挂单奖励（返佣），如果没有返佣金，该字段为“”。手续费返佣为 正数 ，如： 0.01
		Rebate okapi.JSONFloat64 `json:"rebate"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State okapi.OrderState `json:"state"`
		// tdMode String 是 交易模式 保证金模式： isolated ：逐仓 ； cross ：全仓  非保证金模式： cash ：非保证金
		TdMode okapi.TradeMode `json:"tdMode"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okapi.PositionSide `json:"posSide"`
		// side String 是 订单方向   buy ：买，  sell ：卖
		Side okapi.OrderSide `json:"side"`
		// ordType String 是 订单类型  market ：市价单 limit ：限价单  post_only ：只做maker单  fok ：全部成交或立即取消  ioc ：立即成交并取消剩余  optimal_limit_ioc ：市价委托立即成交并取消剩余（仅适用交割、永续） 逐仓自主划转保证金模式不支持ordType为iceberg、twap的策略委托单
		OrdType okapi.OrderType `json:"ordType"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// tgtCcy String 否 市价单委托数量的类型，仅适用于 币币 市价订单 base_ccy : 交易货币 ； quote_ccy ：计价货币 买单默认 quote_ccy ， 卖单默认 base_ccy 计划委托不支持使用tgtCcy参数
		TgtCcy okapi.QuantityType `json:"tgtCcy"`
		// uTime String 订单状态更新时间，Unix时间戳的毫秒数格式，如： 1597026383085
		UTime okapi.JSONTime `json:"uTime"`
		// cTime String 订单创建时间，Unix时间戳的毫秒数格式， 如 ： 1597026383085
		CTime okapi.JSONTime `json:"cTime"`
	}
	TransactionDetail struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID string `json:"ordId"`
		// tradeId String 最新成交ID
		TradeID string `json:"tradeId"`
		// clOrdId String 否 客户自定义订单ID  字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID string `json:"clOrdId"`
		// billId String 账单 ID
		BillID string `json:"billId"`
		// tag String 否 订单标签  字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		Tag okapi.JSONFloat64 `json:"tag"`
		// fillPx String 最新成交价格，如果成交数量为0，该字段也为 0
		FillPx okapi.JSONFloat64 `json:"fillPx"`
		// fillSz String 最新成交数量
		FillSz okapi.JSONFloat64 `json:"fillSz"`
		// feeCcy String 交易手续费币种
		FeeCcy okapi.JSONFloat64 `json:"feeCcy"`
		// fee String 订单交易手续费，平台向用户收取的交易手续费，手续费扣除为 负数 。如：  -0.01
		Fee okapi.JSONFloat64 `json:"fee"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// side String 是 订单方向   buy ：买，  sell ：卖
		Side okapi.OrderSide `json:"side"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okapi.PositionSide `json:"posSide"`
		// execType String 流动性方向  T ：taker  M ：maker
		ExecType okapi.OrderFlowType `json:"execType"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	PlaceAlgoOrder struct {
		// algoId String 策略委托单ID
		AlgoID string `json:"algoId"`
		// sMsg String 事件执行失败时的msg
		SMsg string `json:"sMsg"`
		// sCode String 事件执行结果的code，0代表成功
		SCode okapi.JSONInt64 `json:"sCode"`
	}
	CancelAlgoOrder struct {
		// algoId String 策略委托单ID
		AlgoID string `json:"algoId"`
		// sMsg String 事件执行失败时的msg
		SMsg string `json:"sMsg"`
		// sCode String 事件执行结果的code，0代表成功
		SCode okapi.JSONInt64 `json:"sCode"`
	}
	AlgoOrder struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID string `json:"ordId"`
		// algoId String 策略委托单ID
		AlgoID string `json:"algoId"`
		// clOrdId String 否 客户自定义订单ID  字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID string `json:"clOrdId"`
		// tradeId String 最新成交ID
		TradeID string `json:"tradeId"`
		// tag String 否 订单标签  字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		Tag string `json:"tag"`
		// category String 订单种类  normal ：普通委托 twap ：TWAP自动换币adl ：ADL自动减仓   full_liquidation ：强制平仓 partial_liquidation ：强制减仓    delivery ：交割   ddh ：对冲减仓类型订单
		Category string `json:"category"`
		// feeCcy String 交易手续费币种
		FeeCcy string `json:"feeCcy"`
		// rebateCcy String 返佣金币种
		RebateCcy string `json:"rebateCcy"`
		// timeInterval String 是 下单间隔
		TimeInterval string `json:"timeInterval"`
		// px String 可选  委托价格，仅适用于 limit 、 post_only 、 fok 、 ioc 类型的订单
		Px okapi.JSONFloat64 `json:"px"`
		// pxVar String 可选 挂单价距离盘口的比例  pxVar 和 pxSpread 只能传入一个
		PxVar okapi.JSONFloat64 `json:"pxVar"`
		// pxSpread String 可选 挂单价距离盘口的价距
		PxSpread okapi.JSONFloat64 `json:"pxSpread"`
		// pxLimit String 是 挂单限制价
		PxLimit okapi.JSONFloat64 `json:"pxLimit"`
		// sz String 是 委托数量
		Sz okapi.JSONInt64 `json:"sz"`
		// szLimit String 是 单笔数量
		SzLimit okapi.JSONInt64 `json:"szLimit"`
		// actualSz String 实际委托量
		ActualSz okapi.JSONFloat64 `json:"actualSz"`
		// actualPx String 实际委托价
		ActualPx okapi.JSONFloat64 `json:"actualPx"`
		// pnl String 收益
		Pnl okapi.JSONFloat64 `json:"pnl"`
		// accFillSz String 累计成交数量
		AccFillSz okapi.JSONInt64 `json:"accFillSz"`
		// fillPx String 最新成交价格，如果成交数量为0，该字段也为 0
		FillPx okapi.JSONFloat64 `json:"fillPx"`
		// fillSz String 最新成交数量
		FillSz okapi.JSONInt64 `json:"fillSz"`
		// fillTime String 最新成交时间
		FillTime okapi.JSONFloat64 `json:"fillTime"`
		// avgPx String 成交均价，如果成交数量为0，该字段也为 0
		AvgPx okapi.JSONFloat64 `json:"avgPx"`
		// lever String 杠杆倍数，0.01到125之间的数值，仅适用于  币币杠杆/交割/永续
		Lever okapi.JSONFloat64 `json:"lever"`
		// tpTriggerPx String 止盈触发价
		TpTriggerPx okapi.JSONFloat64 `json:"tpTriggerPx"`
		// tpOrdPx String 止盈委托价
		TpOrdPx okapi.JSONFloat64 `json:"tpOrdPx"`
		// slTriggerPx String 止损触发价
		SlTriggerPx okapi.JSONFloat64 `json:"slTriggerPx"`
		// slOrdPx String 止损委托价
		SlOrdPx okapi.JSONFloat64 `json:"slOrdPx"`
		// ordPx String 计划委托委托价格
		OrdPx okapi.JSONFloat64 `json:"ordPx"`
		// fee String 订单交易手续费，平台向用户收取的交易手续费，手续费扣除为 负数 。如：  -0.01
		Fee okapi.JSONFloat64 `json:"fee"`
		// rebate String 返佣金额，平台向达到指定lv交易等级的用户支付的挂单奖励（返佣），如果没有返佣金，该字段为“”。手续费返佣为 正数 ，如： 0.01
		Rebate okapi.JSONFloat64 `json:"rebate"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State okapi.OrderState `json:"state"`
		// tdMode String 是 交易模式 保证金模式： isolated ：逐仓 ； cross ：全仓  非保证金模式： cash ：非保证金
		TdMode okapi.TradeMode `json:"tdMode"`
		// actualSide String 实际触发方向，  tp ：止盈  sl ： 止损
		ActualSide okapi.PositionSide `json:"actualSide"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okapi.PositionSide `json:"posSide"`
		// side String 是 订单方向   buy ：买，  sell ：卖
		Side okapi.OrderSide `json:"side"`
		// ordType String 是 订单类型  market ：市价单 limit ：限价单  post_only ：只做maker单  fok ：全部成交或立即取消  ioc ：立即成交并取消剩余  optimal_limit_ioc ：市价委托立即成交并取消剩余（仅适用交割、永续） 逐仓自主划转保证金模式不支持ordType为iceberg、twap的策略委托单
		OrdType okapi.AlgoOrderType `json:"ordType"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// tgtCcy String 否 市价单委托数量的类型，仅适用于 币币 市价订单 base_ccy : 交易货币 ； quote_ccy ：计价货币 买单默认 quote_ccy ， 卖单默认 base_ccy 计划委托不支持使用tgtCcy参数
		TgtCcy okapi.QuantityType `json:"tgtCcy"`
		// cTime String 订单创建时间，Unix时间戳的毫秒数格式， 如 ： 1597026383085
		CTime okapi.JSONTime `json:"cTime"`
		// triggerTime String 策略委托触发时间，Unix时间戳的毫秒数格式，如  1597026383085
		TriggerTime okapi.JSONTime `json:"triggerTime"`
	}
)

func (m *PlaceOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ PlaceOrder ------┐", str)
	if s := fmt.Sprintf("%v", m.ClOrdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n客户自定义订单ID: %v", str, m.ClOrdID)
	}
	if s := fmt.Sprintf("%v", m.Tag); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单标签: %v", str, m.Tag)
	}
	if s := fmt.Sprintf("%v", m.SMsg); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n事件执行失败时的msg: %v", str, m.SMsg)
	}
	if s := fmt.Sprintf("%v", m.SCode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n事件执行结果的code: %v", str, m.SCode)
	}
	if s := fmt.Sprintf("%v", m.OrdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单ID: %v", str, m.OrdID)
	}
	str = fmt.Sprintf("%s\r\n└----------- PlaceOrder ------------┘", str)
	return str
}
func (m *CancelOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ CancelOrder ------┐", str)
	if s := fmt.Sprintf("%v", m.OrdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单ID: %v", str, m.OrdID)
	}
	if s := fmt.Sprintf("%v", m.ClOrdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n客户自定义订单ID: %v", str, m.ClOrdID)
	}
	if s := fmt.Sprintf("%v", m.SMsg); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n事件执行失败时的msg: %v", str, m.SMsg)
	}
	if s := fmt.Sprintf("%v", m.SCode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n事件执行结果的code: %v", str, m.SCode)
	}
	str = fmt.Sprintf("%s\r\n└----------- CancelOrder ------------┘", str)
	return str
}
func (m *AmendOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ AmendOrder ------┐", str)
	if s := fmt.Sprintf("%v", m.OrdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单ID: %v", str, m.OrdID)
	}
	if s := fmt.Sprintf("%v", m.ClOrdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n客户自定义订单ID: %v", str, m.ClOrdID)
	}
	if s := fmt.Sprintf("%v", m.ReqID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n用户自定义修改事件ID: %v", str, m.ReqID)
	}
	if s := fmt.Sprintf("%v", m.SMsg); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n事件执行失败时的msg: %v", str, m.SMsg)
	}
	if s := fmt.Sprintf("%v", m.SCode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n事件执行结果的code: %v", str, m.SCode)
	}
	str = fmt.Sprintf("%s\r\n└----------- AmendOrder ------------┘", str)
	return str
}
func (m *ClosePosition) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ ClosePosition ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.PosSide); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.PosSide)
	}
	str = fmt.Sprintf("%s\r\n└----------- ClosePosition ------------┘", str)
	return str
}
func (m *Order) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Order ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.OrdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单ID: %v", str, m.OrdID)
	}
	if s := fmt.Sprintf("%v", m.ClOrdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n客户自定义订单ID: %v", str, m.ClOrdID)
	}
	if s := fmt.Sprintf("%v", m.TradeID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交ID: %v", str, m.TradeID)
	}
	if s := fmt.Sprintf("%v", m.Tag); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单标签: %v", str, m.Tag)
	}
	if s := fmt.Sprintf("%v", m.Category); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单种类: %v", str, m.Category)
	}
	if s := fmt.Sprintf("%v", m.FeeCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n交易手续费币种: %v", str, m.FeeCcy)
	}
	if s := fmt.Sprintf("%v", m.RebateCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n返佣金币种: %v", str, m.RebateCcy)
	}
	if s := fmt.Sprintf("%v", m.Px); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n委托价格: %v", str, m.Px)
	}
	if s := fmt.Sprintf("%v", m.Sz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n委托数量: %v", str, m.Sz)
	}
	if s := fmt.Sprintf("%v", m.Pnl); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n收益: %v", str, m.Pnl)
	}
	if s := fmt.Sprintf("%v", m.AccFillSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n累计成交数量: %v", str, m.AccFillSz)
	}
	if s := fmt.Sprintf("%v", m.FillPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交价格: %v", str, m.FillPx)
	}
	if s := fmt.Sprintf("%v", m.FillSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交数量: %v", str, m.FillSz)
	}
	if s := fmt.Sprintf("%v", m.FillTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交时间: %v", str, m.FillTime)
	}
	if s := fmt.Sprintf("%v", m.AvgPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交均价: %v", str, m.AvgPx)
	}
	if s := fmt.Sprintf("%v", m.Lever); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n杠杆倍数: %v", str, m.Lever)
	}
	if s := fmt.Sprintf("%v", m.TpTriggerPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n止盈触发价: %v", str, m.TpTriggerPx)
	}
	if s := fmt.Sprintf("%v", m.TpOrdPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n止盈委托价: %v", str, m.TpOrdPx)
	}
	if s := fmt.Sprintf("%v", m.SlTriggerPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n止损触发价: %v", str, m.SlTriggerPx)
	}
	if s := fmt.Sprintf("%v", m.SlOrdPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n止损委托价: %v", str, m.SlOrdPx)
	}
	if s := fmt.Sprintf("%v", m.Fee); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单交易手续费: %v", str, m.Fee)
	}
	if s := fmt.Sprintf("%v", m.Rebate); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n返佣金额: %v", str, m.Rebate)
	}
	if s := fmt.Sprintf("%v", m.State); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单状态: %v", str, m.State)
	}
	if s := fmt.Sprintf("%v", m.TdMode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n交易模式: %v", str, m.TdMode)
	}
	if s := fmt.Sprintf("%v", m.PosSide); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.PosSide)
	}
	if s := fmt.Sprintf("%v", m.Side); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单方向: %v", str, m.Side)
	}
	if s := fmt.Sprintf("%v", m.OrdType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单类型: %v", str, m.OrdType)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.TgtCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n市价单委托数量的类型: %v", str, m.TgtCcy)
	}
	if s := fmt.Sprintf("%v", m.UTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单状态更新时间: %v", str, m.UTime)
	}
	if s := fmt.Sprintf("%v", m.CTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单创建时间: %v", str, m.CTime)
	}
	str = fmt.Sprintf("%s\r\n└----------- Order ------------┘", str)
	return str
}
func (m *TransactionDetail) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ TransactionDetail ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.OrdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单ID: %v", str, m.OrdID)
	}
	if s := fmt.Sprintf("%v", m.TradeID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交ID: %v", str, m.TradeID)
	}
	if s := fmt.Sprintf("%v", m.ClOrdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n客户自定义订单ID: %v", str, m.ClOrdID)
	}
	if s := fmt.Sprintf("%v", m.BillID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n账单: %v", str, m.BillID)
	}
	if s := fmt.Sprintf("%v", m.Tag); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单标签: %v", str, m.Tag)
	}
	if s := fmt.Sprintf("%v", m.FillPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交价格: %v", str, m.FillPx)
	}
	if s := fmt.Sprintf("%v", m.FillSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交数量: %v", str, m.FillSz)
	}
	if s := fmt.Sprintf("%v", m.FeeCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n交易手续费币种: %v", str, m.FeeCcy)
	}
	if s := fmt.Sprintf("%v", m.Fee); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单交易手续费: %v", str, m.Fee)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.Side); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单方向: %v", str, m.Side)
	}
	if s := fmt.Sprintf("%v", m.PosSide); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.PosSide)
	}
	if s := fmt.Sprintf("%v", m.ExecType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n流动性方向: %v", str, m.ExecType)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- TransactionDetail ------------┘", str)
	return str
}
func (m *PlaceAlgoOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ PlaceAlgoOrder ------┐", str)
	if s := fmt.Sprintf("%v", m.AlgoID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n策略委托单ID: %v", str, m.AlgoID)
	}
	if s := fmt.Sprintf("%v", m.SMsg); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n事件执行失败时的msg: %v", str, m.SMsg)
	}
	if s := fmt.Sprintf("%v", m.SCode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n事件执行结果的code: %v", str, m.SCode)
	}
	str = fmt.Sprintf("%s\r\n└----------- PlaceAlgoOrder ------------┘", str)
	return str
}
func (m *CancelAlgoOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ CancelAlgoOrder ------┐", str)
	if s := fmt.Sprintf("%v", m.AlgoID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n策略委托单ID: %v", str, m.AlgoID)
	}
	if s := fmt.Sprintf("%v", m.SMsg); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n事件执行失败时的msg: %v", str, m.SMsg)
	}
	if s := fmt.Sprintf("%v", m.SCode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n事件执行结果的code: %v", str, m.SCode)
	}
	str = fmt.Sprintf("%s\r\n└----------- CancelAlgoOrder ------------┘", str)
	return str
}
func (m *AlgoOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ AlgoOrder ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.OrdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单ID: %v", str, m.OrdID)
	}
	if s := fmt.Sprintf("%v", m.AlgoID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n策略委托单ID: %v", str, m.AlgoID)
	}
	if s := fmt.Sprintf("%v", m.ClOrdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n客户自定义订单ID: %v", str, m.ClOrdID)
	}
	if s := fmt.Sprintf("%v", m.TradeID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交ID: %v", str, m.TradeID)
	}
	if s := fmt.Sprintf("%v", m.Tag); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单标签: %v", str, m.Tag)
	}
	if s := fmt.Sprintf("%v", m.Category); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单种类: %v", str, m.Category)
	}
	if s := fmt.Sprintf("%v", m.FeeCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n交易手续费币种: %v", str, m.FeeCcy)
	}
	if s := fmt.Sprintf("%v", m.RebateCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n返佣金币种: %v", str, m.RebateCcy)
	}
	if s := fmt.Sprintf("%v", m.TimeInterval); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n下单间隔: %v", str, m.TimeInterval)
	}
	if s := fmt.Sprintf("%v", m.Px); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n委托价格: %v", str, m.Px)
	}
	if s := fmt.Sprintf("%v", m.PxVar); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n挂单价距离盘口的比例: %v", str, m.PxVar)
	}
	if s := fmt.Sprintf("%v", m.PxSpread); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n挂单价距离盘口的价距: %v", str, m.PxSpread)
	}
	if s := fmt.Sprintf("%v", m.PxLimit); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n挂单限制价: %v", str, m.PxLimit)
	}
	if s := fmt.Sprintf("%v", m.Sz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n委托数量: %v", str, m.Sz)
	}
	if s := fmt.Sprintf("%v", m.SzLimit); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n单笔数量: %v", str, m.SzLimit)
	}
	if s := fmt.Sprintf("%v", m.ActualSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n实际委托量: %v", str, m.ActualSz)
	}
	if s := fmt.Sprintf("%v", m.ActualPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n实际委托价: %v", str, m.ActualPx)
	}
	if s := fmt.Sprintf("%v", m.Pnl); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n收益: %v", str, m.Pnl)
	}
	if s := fmt.Sprintf("%v", m.AccFillSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n累计成交数量: %v", str, m.AccFillSz)
	}
	if s := fmt.Sprintf("%v", m.FillPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交价格: %v", str, m.FillPx)
	}
	if s := fmt.Sprintf("%v", m.FillSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交数量: %v", str, m.FillSz)
	}
	if s := fmt.Sprintf("%v", m.FillTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交时间: %v", str, m.FillTime)
	}
	if s := fmt.Sprintf("%v", m.AvgPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交均价: %v", str, m.AvgPx)
	}
	if s := fmt.Sprintf("%v", m.Lever); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n杠杆倍数: %v", str, m.Lever)
	}
	if s := fmt.Sprintf("%v", m.TpTriggerPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n止盈触发价: %v", str, m.TpTriggerPx)
	}
	if s := fmt.Sprintf("%v", m.TpOrdPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n止盈委托价: %v", str, m.TpOrdPx)
	}
	if s := fmt.Sprintf("%v", m.SlTriggerPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n止损触发价: %v", str, m.SlTriggerPx)
	}
	if s := fmt.Sprintf("%v", m.SlOrdPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n止损委托价: %v", str, m.SlOrdPx)
	}
	if s := fmt.Sprintf("%v", m.OrdPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n计划委托委托价格: %v", str, m.OrdPx)
	}
	if s := fmt.Sprintf("%v", m.Fee); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单交易手续费: %v", str, m.Fee)
	}
	if s := fmt.Sprintf("%v", m.Rebate); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n返佣金额: %v", str, m.Rebate)
	}
	if s := fmt.Sprintf("%v", m.State); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单状态: %v", str, m.State)
	}
	if s := fmt.Sprintf("%v", m.TdMode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n交易模式: %v", str, m.TdMode)
	}
	if s := fmt.Sprintf("%v", m.ActualSide); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n实际触发方向: %v", str, m.ActualSide)
	}
	if s := fmt.Sprintf("%v", m.PosSide); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.PosSide)
	}
	if s := fmt.Sprintf("%v", m.Side); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单方向: %v", str, m.Side)
	}
	if s := fmt.Sprintf("%v", m.OrdType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单类型: %v", str, m.OrdType)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.TgtCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n市价单委托数量的类型: %v", str, m.TgtCcy)
	}
	if s := fmt.Sprintf("%v", m.CTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单创建时间: %v", str, m.CTime)
	}
	if s := fmt.Sprintf("%v", m.TriggerTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n策略委托触发时间: %v", str, m.TriggerTime)
	}
	str = fmt.Sprintf("%s\r\n└----------- AlgoOrder ------------┘", str)
	return str
}
