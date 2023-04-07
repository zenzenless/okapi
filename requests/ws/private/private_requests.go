package private

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
)

type (
	Account struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
	}
	Position struct {
		// uly String 否 交易品种
		InstFamily string `json:"instFamily,omitempty"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
	}
	Order struct {
		// uly String 否 交易品种
		InstFamily string `json:"instFamily,omitempty"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
	}
	AlgoOrder struct {
		// uly String 否 交易品种
		InstFamily string `json:"instFamily,omitempty"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
	}
)

func (m *Account) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Account ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n└----------- Account ------------┘", str)
	return str
}
func (m *Position) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Position ------┐", str)
	str = fmt.Sprintf("%s\r\n交易品种: %v", str, m.InstFamily)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- Position ------------┘", str)
	return str
}
func (m *Order) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Order ------┐", str)
	str = fmt.Sprintf("%s\r\n交易品种: %v", str, m.InstFamily)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- Order ------------┘", str)
	return str
}
func (m *AlgoOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ AlgoOrder ------┐", str)
	str = fmt.Sprintf("%s\r\n交易品种: %v", str, m.InstFamily)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- AlgoOrder ------------┘", str)
	return str
}
