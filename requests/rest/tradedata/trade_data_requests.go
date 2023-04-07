package tradedata

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
)

type (
	GetTakerVolume struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Begin int64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		End int64 `json:"limit,omitempty,string"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType,string"`
		// period String 否 时间粒度，默认值 5m 。支持[5m/1H/1D]  5m 粒度最多只能查询两天之内的数据 1H 粒度最多只能查询30天之内的数据1D 粒度最多只能查询180天之内的数据
		Period okapi.BarSize `json:"period,string,omitempty"`
	}
	GetRatio struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Begin int64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		End int64 `json:"limit,omitempty,string"`
		// period String 否 时间粒度，默认值 5m 。支持[5m/1H/1D]  5m 粒度最多只能查询两天之内的数据 1H 粒度最多只能查询30天之内的数据1D 粒度最多只能查询180天之内的数据
		Period okapi.BarSize `json:"period,string,omitempty"`
	}
	GetOpenInterestAndVolumeStrike struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// expTime String 否 请求有效截止时间。Unix时间戳的毫秒数格式，如  1597026383085 无效的expTime
		ExpTime string `json:"expTime"`
		// period String 否 时间粒度，默认值 5m 。支持[5m/1H/1D]  5m 粒度最多只能查询两天之内的数据 1H 粒度最多只能查询30天之内的数据1D 粒度最多只能查询180天之内的数据
		Period okapi.BarSize `json:"period,string,omitempty"`
	}
)

func (m *GetTakerVolume) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetTakerVolume ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Begin)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.End)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n时间粒度: %v", str, m.Period)
	str = fmt.Sprintf("%s\r\n└----------- GetTakerVolume ------------┘", str)
	return str
}
func (m *GetRatio) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetRatio ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Begin)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.End)
	str = fmt.Sprintf("%s\r\n时间粒度: %v", str, m.Period)
	str = fmt.Sprintf("%s\r\n└----------- GetRatio ------------┘", str)
	return str
}
func (m *GetOpenInterestAndVolumeStrike) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetOpenInterestAndVolumeStrike ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n请求有效截止时间。Unix时间戳的毫秒数格式: %v", str, m.ExpTime)
	str = fmt.Sprintf("%s\r\n时间粒度: %v", str, m.Period)
	str = fmt.Sprintf("%s\r\n└----------- GetOpenInterestAndVolumeStrike ------------┘", str)
	return str
}
