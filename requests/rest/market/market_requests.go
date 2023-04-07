package market

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
)

type (
	GetTickers struct {
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
	}
	GetIndexTickers struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// quoteCcy String 计价货币币种，如  BTC-USDT 中的 USDT
		QuoteCcy string `json:"quoteCcy,omitempty"`
	}
	GetOrderBook struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// sz String 是 委托数量
		Sz int `json:"sz,omitempty,string"`
	}
	GetCandlesticks struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After int64 `json:"after,omitempty,string"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before int64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit int64 `json:"limit,omitempty,string"`
		// bar String 否 时间粒度，默认值 1m 如 [1m/3m/5m/15m/30m/1H/2H/4H]香港时间开盘价k线：[6H/12H/1D/2D/3D/1W/1M/3M/6M/1Y] UTC时间开盘价k线：[/6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/1Wutc/1Mutc/3Mutc/6Mutc/1Yutc]
		Bar okapi.BarSize `json:"bar,omitempty"`
	}
	GetTrades struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit int64 `json:"limit,omitempty,string"`
	}
	GetIndexComponents struct {
		// index String 是 指数，如  BTC-USDT
		Index string `json:"index"`
	}
)

func (m *GetTickers) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetTickers ------┐", str)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- GetTickers ------------┘", str)
	return str
}
func (m *GetIndexTickers) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetIndexTickers ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n计价货币币种: %v", str, m.QuoteCcy)
	str = fmt.Sprintf("%s\r\n└----------- GetIndexTickers ------------┘", str)
	return str
}
func (m *GetOrderBook) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetOrderBook ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n委托数量: %v", str, m.Sz)
	str = fmt.Sprintf("%s\r\n└----------- GetOrderBook ------------┘", str)
	return str
}
func (m *GetCandlesticks) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetCandlesticks ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n时间粒度: %v", str, m.Bar)
	str = fmt.Sprintf("%s\r\n└----------- GetCandlesticks ------------┘", str)
	return str
}
func (m *GetTrades) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetTrades ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n└----------- GetTrades ------------┘", str)
	return str
}
func (m *GetIndexComponents) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetIndexComponents ------┐", str)
	str = fmt.Sprintf("%s\r\n指数: %v", str, m.Index)
	str = fmt.Sprintf("%s\r\n└----------- GetIndexComponents ------------┘", str)
	return str
}
