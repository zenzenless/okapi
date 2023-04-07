package account

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
)

type (
	GetBalance struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy []string `json:"ccy,omitempty"`
	}
	GetPositions struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID []string `json:"instId,omitempty"`
		// posId String 否 持仓ID 支持多个 posId 查询（不超过20个），半角逗号分割
		PosID []string `json:"posId,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType,omitempty"`
	}
	GetAccountAndPositionRisk struct {
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType,omitempty"`
	}
	GetBills struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After int64 `json:"after,omitempty,string"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before int64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit int64 `json:"limit,omitempty,string"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType,omitempty"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okapi.MarginMode `json:"mgnMode,omitempty"`
		// ctType String 否 linear ： 正向合约 inverse ：反向合约 仅 交割/永续 有效
		CtType okapi.ContractType `json:"ctType,omitempty"`
		// type String 报价方类型（当前未生效，将返回 "" ）
		Type okapi.BillType `json:"type,omitempty,string"`
		// subType String 否 账单子类型 1 ：买入  2 ：卖出  3 ：开多4 ：开空  5 ：平多  6 ：平空  9 ：市场借币扣息  11 ：转入12 ：转出  14 ：尊享借币扣息  160 ：手动追加保证金  161 ：手动减少保证金162 ：自动追加保证金  114 ：自动换币买入  115 ：自动换币卖出  118 ：系统换币转入119 ：系统换币转出  100 ：强减平多  101 ：强减平空  102 ：强减买入103 ：强减卖出  104 ：强平平多  105 ：强平平空  106 ：强平买入107 ：强平卖出  110 ：强平换币转入  111 ：强平换币转出  125 ：自动减仓平多126 ：自动减仓平空  127 ：自动减仓买入  128 ：自动减仓卖出  131 ：对冲买入132 ：对冲卖出  170 ：到期行权  171 ：到期被行权  172 ：到期作废112 ：交割平多  113 ：交割平空  117 ：交割/期权穿仓补偿  173 ：资金费支出174 ：资金费收入  200 :系统转入  201 :手动转入  202 :系统转出203 :手动转出  204 : 大宗交易买  205 : 大宗交易卖  206 : 大宗交易开多207 : 大宗交易开空  208 : 大宗交易平多  209 : 大宗交易平空  224 : 还债转入225 : 还债转出
		SubType okapi.BillSubType `json:"subType,omitempty,string"`
	}
	SetPositionMode struct {
		PositionMode okapi.PositionType `json:"positionMode"`
	}
	SetLeverage struct {
		// lever String 杠杆倍数，0.01到125之间的数值，仅适用于  币币杠杆/交割/永续
		Lever int64 `json:"lever,string"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okapi.MarginMode `json:"mgnMode"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okapi.PositionSide `json:"posSide,omitempty"`
	}
	GetMaxBuySellAmount struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// px String 可选  委托价格，仅适用于 limit 、 post_only 、 fok 、 ioc 类型的订单
		Px float64 `json:"px,string,omitempty"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID []string `json:"instId"`
		// tdMode String 是 交易模式 保证金模式： isolated ：逐仓 ； cross ：全仓  非保证金模式： cash ：非保证金
		TdMode okapi.TradeMode `json:"tdMode"`
	}
	GetMaxAvailableTradeAmount struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// reduceOnly Boolean 否 是否只减仓， true  或  false ，默认 false 仅适用于 币币杠杆 ，以及买卖模式下的 交割/永续 仅适用于 单币种保证金模式 和 跨币种保证金模式
		ReduceOnly bool `json:"reduceOnly,omitempty"`
		// tdMode String 是 交易模式 保证金模式： isolated ：逐仓 ； cross ：全仓  非保证金模式： cash ：非保证金
		TdMode okapi.TradeMode `json:"tdMode"`
	}
	IncreaseDecreaseMargin struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// amt String 是 划转数量
		Amt float64 `json:"amt,string"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide    okapi.PositionSide `json:"posSide"`
		ActionType okapi.CountAction  `json:"actionType"`
	}
	GetLeverage struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID []string `json:"instId"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okapi.MarginMode `json:"mgnMode"`
	}
	GetMaxLoan struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// mgnCcy String 可选 保证金币种，如  BTC 币币杠杆单币种全仓情况下必须指定保证金币种
		MgnCcy string `json:"mgnCcy,omitempty"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okapi.MarginMode `json:"mgnMode"`
	}
	GetFeeRates struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// category String 订单种类  normal ：普通委托 twap ：TWAP自动换币adl ：ADL自动减仓   full_liquidation ：强制平仓 partial_liquidation ：强制减仓    delivery ：交割   ddh ：对冲减仓类型订单
		Category okapi.FeeCategory `json:"category,omitempty,string"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
	}
	GetInterestAccrued struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After int64 `json:"after,omitempty,string"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before int64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit int64 `json:"limit,omitempty,string"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okapi.MarginMode `json:"mgnMode,omitempty"`
	}
	SetGreeks struct {
		// greeksType String 当前希腊字母展示方式 PA ：币本位  BS ：美元本位
		GreeksType okapi.GreekType `json:"greeksType"`
	}
)

func (m *GetBalance) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetBalance ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n└----------- GetBalance ------------┘", str)
	return str
}
func (m *GetPositions) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetPositions ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n持仓ID: %v", str, m.PosID)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- GetPositions ------------┘", str)
	return str
}
func (m *GetAccountAndPositionRisk) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetAccountAndPositionRisk ------┐", str)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- GetAccountAndPositionRisk ------------┘", str)
	return str
}
func (m *GetBills) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetBills ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n保证金模式: %v", str, m.MgnMode)
	str = fmt.Sprintf("%s\r\n正向合约: %v", str, m.CtType)
	str = fmt.Sprintf("%s\r\n报价方类型（当前未生效: %v", str, m.Type)
	str = fmt.Sprintf("%s\r\n账单子类型: %v", str, m.SubType)
	str = fmt.Sprintf("%s\r\n└----------- GetBills ------------┘", str)
	return str
}
func (m *SetLeverage) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ SetLeverage ------┐", str)
	str = fmt.Sprintf("%s\r\n杠杆倍数: %v", str, m.Lever)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n保证金模式: %v", str, m.MgnMode)
	str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.PosSide)
	str = fmt.Sprintf("%s\r\n└----------- SetLeverage ------------┘", str)
	return str
}
func (m *GetMaxBuySellAmount) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetMaxBuySellAmount ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n委托价格: %v", str, m.Px)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n交易模式: %v", str, m.TdMode)
	str = fmt.Sprintf("%s\r\n└----------- GetMaxBuySellAmount ------------┘", str)
	return str
}
func (m *GetMaxAvailableTradeAmount) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetMaxAvailableTradeAmount ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n是否只减仓: %v", str, m.ReduceOnly)
	str = fmt.Sprintf("%s\r\n交易模式: %v", str, m.TdMode)
	str = fmt.Sprintf("%s\r\n└----------- GetMaxAvailableTradeAmount ------------┘", str)
	return str
}
func (m *IncreaseDecreaseMargin) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ IncreaseDecreaseMargin ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n划转数量: %v", str, m.Amt)
	str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.PosSide)
	str = fmt.Sprintf("%s\r\n: %v", str, m.ActionType)
	str = fmt.Sprintf("%s\r\n└----------- IncreaseDecreaseMargin ------------┘", str)
	return str
}
func (m *GetLeverage) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetLeverage ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n保证金模式: %v", str, m.MgnMode)
	str = fmt.Sprintf("%s\r\n└----------- GetLeverage ------------┘", str)
	return str
}
func (m *GetMaxLoan) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetMaxLoan ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.MgnCcy)
	str = fmt.Sprintf("%s\r\n保证金模式: %v", str, m.MgnMode)
	str = fmt.Sprintf("%s\r\n└----------- GetMaxLoan ------------┘", str)
	return str
}
func (m *GetFeeRates) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetFeeRates ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n订单种类: %v", str, m.Category)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- GetFeeRates ------------┘", str)
	return str
}
func (m *GetInterestAccrued) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetInterestAccrued ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n保证金模式: %v", str, m.MgnMode)
	str = fmt.Sprintf("%s\r\n└----------- GetInterestAccrued ------------┘", str)
	return str
}
func (m *SetGreeks) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ SetGreeks ------┐", str)
	str = fmt.Sprintf("%s\r\n当前希腊字母展示方式: %v", str, m.GreeksType)
	str = fmt.Sprintf("%s\r\n└----------- SetGreeks ------------┘", str)
	return str
}
func (m *SetPositionMode) String() string {
	var str string
	str = fmt.Sprintf("%s\r\n┌------ SetPositionMode ------┐", str)
	str = fmt.Sprintf("%s\r\nPositionMode:%v", str, m.PositionMode)
	str = fmt.Sprintf("%s\r\n└----------- SetPositionMode ------------┘", str)
	return str
}
