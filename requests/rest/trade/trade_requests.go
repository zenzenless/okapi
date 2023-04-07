package trade

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
)

type (
	PlaceOrder struct {
		ID string `json:"-"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// clOrdId String 否 客户自定义订单ID  字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID string `json:"clOrdId,omitempty"`
		// tag String 否 订单标签  字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		Tag string `json:"tag,omitempty"`
		// reduceOnly Boolean 否 是否只减仓， true  或  false ，默认 false 仅适用于 币币杠杆 ，以及买卖模式下的 交割/永续 仅适用于 单币种保证金模式 和 跨币种保证金模式
		ReduceOnly bool `json:"reduceOnly,omitempty"`
		// sz String 是 委托数量
		Sz float64 `json:"sz,string"`
		// px String 可选  委托价格，仅适用于 limit 、 post_only 、 fok 、 ioc 类型的订单
		Px float64 `json:"px,omitempty,string"`
		// tdMode String 是 交易模式 保证金模式： isolated ：逐仓 ； cross ：全仓  非保证金模式： cash ：非保证金
		TdMode okapi.TradeMode `json:"tdMode"`
		// side String 是 订单方向   buy ：买，  sell ：卖
		Side okapi.OrderSide `json:"side"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okapi.PositionSide `json:"posSide,omitempty"`
		// ordType String 是 订单类型  market ：市价单 limit ：限价单  post_only ：只做maker单  fok ：全部成交或立即取消  ioc ：立即成交并取消剩余  optimal_limit_ioc ：市价委托立即成交并取消剩余（仅适用交割、永续） 逐仓自主划转保证金模式不支持ordType为iceberg、twap的策略委托单
		OrdType okapi.OrderType `json:"ordType"`
		// tgtCcy String 否 市价单委托数量的类型，仅适用于 币币 市价订单 base_ccy : 交易货币 ； quote_ccy ：计价货币 买单默认 quote_ccy ， 卖单默认 base_ccy 计划委托不支持使用tgtCcy参数
		TgtCcy okapi.QuantityType `json:"tgtCcy,omitempty"`
	}
	CancelOrder struct {
		ID string `json:"-"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID string `json:"ordId,omitempty"`
		// clOrdId String 否 客户自定义订单ID  字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID string `json:"clOrdId,omitempty"`
	}
	AmendOrder struct {
		ID string `json:"-"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID string `json:"ordId,omitempty"`
		// clOrdId String 否 客户自定义订单ID  字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID string `json:"clOrdId,omitempty"`
		// reqId String 否 用户自定义修改事件ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
		ReqID string `json:"reqId,omitempty"`
		// newSz String 可选 修改的新数量， newSz 和 newPx 不可同时为空。对于部分成交订单，该数量应包含已成交数量。
		NewSz int64 `json:"newSz,omitempty,string"`
		// newPx String 可选 修改的新价格
		NewPx float64 `json:"newPx,omitempty,string"`
		// cxlOnFail Boolean 否 false ：不自动撤单  true ：自动撤单当订单修改失败时，该订单是否需要自动撤销。默认为 false
		CxlOnFail bool `json:"cxlOnFail,omitempty"`
	}
	ClosePosition struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okapi.PositionSide `json:"posSide,omitempty"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okapi.MarginMode `json:"mgnMode"`
	}
	OrderDetails struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID string `json:"ordId,omitempty"`
		// clOrdId String 否 客户自定义订单ID  字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID string `json:"clOrdId,omitempty"`
	}
	OrderList struct {
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After float64 `json:"after,omitempty,string"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before float64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit float64 `json:"limit,omitempty,string"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType,omitempty"`
		// ordType String 是 订单类型  market ：市价单 limit ：限价单  post_only ：只做maker单  fok ：全部成交或立即取消  ioc ：立即成交并取消剩余  optimal_limit_ioc ：市价委托立即成交并取消剩余（仅适用交割、永续） 逐仓自主划转保证金模式不支持ordType为iceberg、twap的策略委托单
		OrdType okapi.OrderType `json:"ordType,omitempty"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State okapi.OrderState `json:"state,omitempty"`
	}
	TransactionDetails struct {
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID string `json:"ordId,omitempty"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After float64 `json:"after,omitempty,string"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before float64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit float64 `json:"limit,omitempty,string"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType,omitempty"`
	}
	PlaceAlgoOrder struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// tdMode String 是 交易模式 保证金模式： isolated ：逐仓 ； cross ：全仓  非保证金模式： cash ：非保证金
		TdMode okapi.TradeMode `json:"tdMode"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// side String 是 订单方向   buy ：买，  sell ：卖
		Side okapi.OrderSide `json:"side"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okapi.PositionSide `json:"posSide,omitempty"`
		// ordType String 是 订单类型  market ：市价单 limit ：限价单  post_only ：只做maker单  fok ：全部成交或立即取消  ioc ：立即成交并取消剩余  optimal_limit_ioc ：市价委托立即成交并取消剩余（仅适用交割、永续） 逐仓自主划转保证金模式不支持ordType为iceberg、twap的策略委托单
		OrdType okapi.AlgoOrderType `json:"ordType"`
		// sz String 是 委托数量
		Sz okapi.JSONFloat64 `json:"sz,string"`
		// reduceOnly Boolean 否 是否只减仓， true  或  false ，默认 false 仅适用于 币币杠杆 ，以及买卖模式下的 交割/永续 仅适用于 单币种保证金模式 和 跨币种保证金模式
		ReduceOnly bool `json:"reduceOnly,omitempty"`
		// tgtCcy String 否 市价单委托数量的类型，仅适用于 币币 市价订单 base_ccy : 交易货币 ； quote_ccy ：计价货币 买单默认 quote_ccy ， 卖单默认 base_ccy 计划委托不支持使用tgtCcy参数
		TgtCcy okapi.QuantityType `json:"tgtCcy,omitempty"`
		Tag    string             `json:"Tag,omitempty"`
		*StopOrder
		*TriggerOrder
		*IcebergOrder
		*TWAPOrder
		*MoveStopOrder
	}
	StopOrder struct {
		// tpTriggerPx String 止盈触发价
		TpTriggerPx float64 `json:"tpTriggerPx,string,omitempty"`
		// tpOrdPx String 止盈委托价
		TpOrdPx float64 `json:"tpOrdPx,string,omitempty"`
		// slTriggerPx String 止损触发价
		SlTriggerPx float64 `json:"slTriggerPx,string,omitempty"`
		// slOrdPx String 止损委托价
		SlOrdPx float64 `json:"slOrdPx,string,omitempty"`
	}
	TriggerOrder struct {
		// triggerPx String 是 计划委托触发价格
		TriggerPx float64 `json:"triggerPx,string,omitempty"`
		// ordPx String 计划委托委托价格
		OrdPx float64 `json:"ordPx,string,omitempty"`
	}
	IcebergOrder struct {
		// pxVar String 可选 挂单价距离盘口的比例  pxVar 和 pxSpread 只能传入一个
		PxVar float64 `json:"pxVar,string,omitempty"`
		// pxSpread String 可选 挂单价距离盘口的价距
		PxSpread float64 `json:"pxSpread,string,omitempty"`
		// szLimit String 是 单笔数量
		SzLimit int64 `json:"szLimit,string"`
		// pxLimit String 是 挂单限制价
		PxLimit float64 `json:"pxLimit,string"`
	}
	TWAPOrder struct {
		IcebergOrder
		// timeInterval String 是 下单间隔
		TimeInterval string `json:"timeInterval"`
	}
	MoveStopOrder struct {
		// callbackRatio String 可选 回调幅度的比例，如0.05 callbackRatio 和 callbackSpread 只能传入一个
		CallbackRatio float64 `json:"callbackRatio,string,omitempty"`
		// callbackSpread String 可选 回调幅度的价距
		CallbackSpread float64 `json:"callbackSpread,string,omitempty"`
		// activePx String 否 激活价格
		ActivePx float64 `json:"activePx,string,omitempty"`
	}
	CancelAlgoOrder struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		AlgoID string `json:"AlgoId"`
	}
	AlgoOrderList struct {
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType,omitempty"`
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After float64 `json:"after,omitempty,string"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before float64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit float64 `json:"limit,omitempty,string"`
		// ordType String 是 订单类型  market ：市价单 limit ：限价单  post_only ：只做maker单  fok ：全部成交或立即取消  ioc ：立即成交并取消剩余  optimal_limit_ioc ：市价委托立即成交并取消剩余（仅适用交割、永续） 逐仓自主划转保证金模式不支持ordType为iceberg、twap的策略委托单
		OrdType okapi.AlgoOrderType `json:"ordType,omitempty"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State okapi.OrderState `json:"state,omitempty"`
	}
)

func (m *PlaceOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ PlaceOrder ------┐", str)
	str = fmt.Sprintf("%s\r\n: %v", str, m.ID)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n客户自定义订单ID: %v", str, m.ClOrdID)
	str = fmt.Sprintf("%s\r\n订单标签: %v", str, m.Tag)
	str = fmt.Sprintf("%s\r\n是否只减仓: %v", str, m.ReduceOnly)
	str = fmt.Sprintf("%s\r\n委托数量: %v", str, m.Sz)
	str = fmt.Sprintf("%s\r\n委托价格: %v", str, m.Px)
	str = fmt.Sprintf("%s\r\n交易模式: %v", str, m.TdMode)
	str = fmt.Sprintf("%s\r\n订单方向: %v", str, m.Side)
	str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.PosSide)
	str = fmt.Sprintf("%s\r\n订单类型: %v", str, m.OrdType)
	str = fmt.Sprintf("%s\r\n市价单委托数量的类型: %v", str, m.TgtCcy)
	str = fmt.Sprintf("%s\r\n└----------- PlaceOrder ------------┘", str)
	return str
}
func (m *CancelOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ CancelOrder ------┐", str)
	str = fmt.Sprintf("%s\r\n: %v", str, m.ID)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n订单ID: %v", str, m.OrdID)
	str = fmt.Sprintf("%s\r\n客户自定义订单ID: %v", str, m.ClOrdID)
	str = fmt.Sprintf("%s\r\n└----------- CancelOrder ------------┘", str)
	return str
}
func (m *AmendOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ AmendOrder ------┐", str)
	str = fmt.Sprintf("%s\r\n: %v", str, m.ID)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n订单ID: %v", str, m.OrdID)
	str = fmt.Sprintf("%s\r\n客户自定义订单ID: %v", str, m.ClOrdID)
	str = fmt.Sprintf("%s\r\n用户自定义修改事件ID: %v", str, m.ReqID)
	str = fmt.Sprintf("%s\r\n修改的新数量: %v", str, m.NewSz)
	str = fmt.Sprintf("%s\r\n修改的新价格: %v", str, m.NewPx)
	str = fmt.Sprintf("%s\r\n：不自动撤单: %v", str, m.CxlOnFail)
	str = fmt.Sprintf("%s\r\n└----------- AmendOrder ------------┘", str)
	return str
}
func (m *ClosePosition) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ ClosePosition ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.PosSide)
	str = fmt.Sprintf("%s\r\n保证金模式: %v", str, m.MgnMode)
	str = fmt.Sprintf("%s\r\n└----------- ClosePosition ------------┘", str)
	return str
}
func (m *OrderDetails) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ OrderDetails ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n订单ID: %v", str, m.OrdID)
	str = fmt.Sprintf("%s\r\n客户自定义订单ID: %v", str, m.ClOrdID)
	str = fmt.Sprintf("%s\r\n└----------- OrderDetails ------------┘", str)
	return str
}
func (m *OrderList) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ OrderList ------┐", str)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n订单类型: %v", str, m.OrdType)
	str = fmt.Sprintf("%s\r\n订单状态: %v", str, m.State)
	str = fmt.Sprintf("%s\r\n└----------- OrderList ------------┘", str)
	return str
}
func (m *TransactionDetails) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ TransactionDetails ------┐", str)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n订单ID: %v", str, m.OrdID)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- TransactionDetails ------------┘", str)
	return str
}
func (m *PlaceAlgoOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ PlaceAlgoOrder ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n交易模式: %v", str, m.TdMode)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n订单方向: %v", str, m.Side)
	str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.PosSide)
	str = fmt.Sprintf("%s\r\n订单类型: %v", str, m.OrdType)
	str = fmt.Sprintf("%s\r\n委托数量: %v", str, m.Sz)
	str = fmt.Sprintf("%s\r\n是否只减仓: %v", str, m.ReduceOnly)
	str = fmt.Sprintf("%s\r\n市价单委托数量的类型: %v", str, m.TgtCcy)
	str = fmt.Sprintf("%s\r\n: %v", str, m.Tag)
	str = fmt.Sprintf("%s\r\n: %v", str, m.StopOrder)
	str = fmt.Sprintf("%s\r\n: %v", str, m.TriggerOrder)
	str = fmt.Sprintf("%s\r\n: %v", str, m.IcebergOrder)
	str = fmt.Sprintf("%s\r\n: %v", str, m.TWAPOrder)
	str = fmt.Sprintf("%s\r\n: %v", str, m.MoveStopOrder)
	str = fmt.Sprintf("%s\r\n└----------- PlaceAlgoOrder ------------┘", str)
	return str
}
func (m *StopOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ StopOrder ------┐", str)
	str = fmt.Sprintf("%s\r\n止盈触发价: %v", str, m.TpTriggerPx)
	str = fmt.Sprintf("%s\r\n止盈委托价: %v", str, m.TpOrdPx)
	str = fmt.Sprintf("%s\r\n止损触发价: %v", str, m.SlTriggerPx)
	str = fmt.Sprintf("%s\r\n止损委托价: %v", str, m.SlOrdPx)
	str = fmt.Sprintf("%s\r\n└----------- StopOrder ------------┘", str)
	return str
}
func (m *TriggerOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ TriggerOrder ------┐", str)
	str = fmt.Sprintf("%s\r\n计划委托触发价格: %v", str, m.TriggerPx)
	str = fmt.Sprintf("%s\r\n计划委托委托价格: %v", str, m.OrdPx)
	str = fmt.Sprintf("%s\r\n└----------- TriggerOrder ------------┘", str)
	return str
}
func (m *IcebergOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ IcebergOrder ------┐", str)
	str = fmt.Sprintf("%s\r\n挂单价距离盘口的比例: %v", str, m.PxVar)
	str = fmt.Sprintf("%s\r\n挂单价距离盘口的价距: %v", str, m.PxSpread)
	str = fmt.Sprintf("%s\r\n单笔数量: %v", str, m.SzLimit)
	str = fmt.Sprintf("%s\r\n挂单限制价: %v", str, m.PxLimit)
	str = fmt.Sprintf("%s\r\n└----------- IcebergOrder ------------┘", str)
	return str
}
func (m *TWAPOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ TWAPOrder ------┐", str)
	str = fmt.Sprintf("%s\r\n: %v", str, m.IcebergOrder)
	str = fmt.Sprintf("%s\r\n下单间隔: %v", str, m.TimeInterval)
	str = fmt.Sprintf("%s\r\n└----------- TWAPOrder ------------┘", str)
	return str
}
func (m *MoveStopOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ MoveStopOrder ------┐", str)
	str = fmt.Sprintf("%s\r\n回调幅度的比例: %v", str, m.CallbackRatio)
	str = fmt.Sprintf("%s\r\n回调幅度的价距: %v", str, m.CallbackSpread)
	str = fmt.Sprintf("%s\r\n激活价格: %v", str, m.ActivePx)
	str = fmt.Sprintf("%s\r\n└----------- MoveStopOrder ------------┘", str)
	return str
}
func (m *CancelAlgoOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ CancelAlgoOrder ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n: %v", str, m.AlgoID)
	str = fmt.Sprintf("%s\r\n└----------- CancelAlgoOrder ------------┘", str)
	return str
}
func (m *AlgoOrderList) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ AlgoOrderList ------┐", str)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n订单类型: %v", str, m.OrdType)
	str = fmt.Sprintf("%s\r\n订单状态: %v", str, m.State)
	str = fmt.Sprintf("%s\r\n└----------- AlgoOrderList ------------┘", str)
	return str
}
