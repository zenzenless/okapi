package public

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
	"github.com/james-zhang-bing/okapi/models/publicdata"
)

type (
	GetInstruments struct {
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
	}
	GetDeliveryExerciseHistory struct {
		// uly String 否 标的指数
		Uly string `json:"uly"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After int64 `json:"after,omitempty,string"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before int64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit int64 `json:"limit,omitempty,string"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
	}
	GetOpenInterest struct {
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
	}
	GetFundingRate struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
	}
	GetLimitPrice struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
	}
	GetOptionMarketData struct {
		// uly String 否 标的指数
		Uly string `json:"uly"`
		// expTime String 否 请求有效截止时间。Unix时间戳的毫秒数格式，如  1597026383085 无效的expTime
		ExpTime string `json:"expTime,omitempty"`
	}
	GetEstimatedDeliveryExercisePrice struct {
		// uly String 否 标的指数
		Uly string `json:"uly"`
		// expTime String 否 请求有效截止时间。Unix时间戳的毫秒数格式，如  1597026383085 无效的expTime
		ExpTime string `json:"expTime,omitempty"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		// uly String 否 标的指数
		Uly string `json:"uly"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// discountLv String 否 折算率等级 1 :第一档  2 :第二档  3 :第三档4 :第四档  5 :第五档
		DiscountLv float64 `json:"discountLv,string"`
	}
	GetLiquidationOrders struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After int64 `json:"after,omitempty,string"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before int64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit int64 `json:"limit,omitempty,string"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okapi.MarginMode `json:"mgnMode,omitempty"`
		// alias String 合约日期别名 this_week ：本周  next_week ：次周quarter ：季度 next_quarter ：次季度  仅适用于 交割
		Alias okapi.AliasType `json:"alias,omitempty"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State okapi.OrderState `json:"state,omitempty"`
	}
	GetMarkPrice struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
	}
	GetPositionTiers struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// tdMode String 是 交易模式 保证金模式： isolated ：逐仓 ； cross ：全仓  非保证金模式： cash ：非保证金
		TdMode okapi.TradeMode `json:"tdMode"`
		// tier String 否 查指定档位
		Tier okapi.JSONInt64 `json:"tier,omitempty"`
	}
	GetUnderlying struct {
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
	}
	Status struct {
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State string `json:"state,omitempty"`
	}

	ConvertContractUnit publicdata.ConvertContractUnit

	GetFundingRateHistory struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// after String 否 请求此时间戳之前（更旧的数据）的分页内容，传的值为对应接口的fundingTime
		After int64 `json:"after,omitempty,string"`
		// before String 否 请求此时间戳之后（更新的数据）的分页内容，传的值为对应接口的fundingTime
		Before int64 `json:"before,omitempty,string"`
		//limit String 分页返回的结果集数量，最大为100，不填默认返回100条
		Limit int64 `json:"limit,omitempty,string"`
	}
)

func (m *GetInstruments) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetInstruments ------┐", str)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- GetInstruments ------------┘", str)
	return str
}
func (m *GetDeliveryExerciseHistory) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetDeliveryExerciseHistory ------┐", str)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- GetDeliveryExerciseHistory ------------┘", str)
	return str
}
func (m *GetOpenInterest) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetOpenInterest ------┐", str)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- GetOpenInterest ------------┘", str)
	return str
}
func (m *GetFundingRate) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetFundingRate ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n└----------- GetFundingRate ------------┘", str)
	return str
}
func (m *GetLimitPrice) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetLimitPrice ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n└----------- GetLimitPrice ------------┘", str)
	return str
}
func (m *GetOptionMarketData) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetOptionMarketData ------┐", str)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n请求有效截止时间。Unix时间戳的毫秒数格式: %v", str, m.ExpTime)
	str = fmt.Sprintf("%s\r\n└----------- GetOptionMarketData ------------┘", str)
	return str
}
func (m *GetEstimatedDeliveryExercisePrice) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetEstimatedDeliveryExercisePrice ------┐", str)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n请求有效截止时间。Unix时间戳的毫秒数格式: %v", str, m.ExpTime)
	str = fmt.Sprintf("%s\r\n└----------- GetEstimatedDeliveryExercisePrice ------------┘", str)
	return str
}
func (m *GetDiscountRateAndInterestFreeQuota) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetDiscountRateAndInterestFreeQuota ------┐", str)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n折算率等级: %v", str, m.DiscountLv)
	str = fmt.Sprintf("%s\r\n└----------- GetDiscountRateAndInterestFreeQuota ------------┘", str)
	return str
}
func (m *GetLiquidationOrders) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetLiquidationOrders ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n保证金模式: %v", str, m.MgnMode)
	str = fmt.Sprintf("%s\r\n合约日期别名: %v", str, m.Alias)
	str = fmt.Sprintf("%s\r\n订单状态: %v", str, m.State)
	str = fmt.Sprintf("%s\r\n└----------- GetLiquidationOrders ------------┘", str)
	return str
}
func (m *GetMarkPrice) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetMarkPrice ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- GetMarkPrice ------------┘", str)
	return str
}
func (m *GetPositionTiers) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetPositionTiers ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n交易模式: %v", str, m.TdMode)
	str = fmt.Sprintf("%s\r\n查指定档位: %v", str, m.Tier)
	str = fmt.Sprintf("%s\r\n└----------- GetPositionTiers ------------┘", str)
	return str
}
func (m *GetUnderlying) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetUnderlying ------┐", str)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- GetUnderlying ------------┘", str)
	return str
}
func (m *Status) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Status ------┐", str)
	str = fmt.Sprintf("%s\r\n订单状态: %v", str, m.State)
	str = fmt.Sprintf("%s\r\n└----------- Status ------------┘", str)
	return str
}
