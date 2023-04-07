package subaccount

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
)

type (
	ViewList struct {
		// subAcct String 可选 子账户名称，type 为 1 ， 2  或  4 ：subAcct 为必填项
		SubAcct string `json:"subAcct,omitempty"`
		// enable String 否 子账户状态， true ：正常使用  false ：冻结
		Enable bool `json:"enable,omitempty"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After int64 `json:"after,omitempty,string"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before int64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit int64 `json:"limit,omitempty,string"`
	}
	CreateAPIKey struct {
		Pwd string `json:"pwd"`
		// subAcct String 可选 子账户名称，type 为 1 ， 2  或  4 ：subAcct 为必填项
		SubAcct string `json:"subAcct"`
		// label String 子账户备注
		Label      string `json:"label"`
		Passphrase string `json:"Passphrase"`
		// ip String 否 子账户APIKey绑定ip地址，多个ip用半角逗号隔开，最多支持20个ip。  如果填写该字段，那该字段会被重置如果ip传""，则表示解除IP绑定
		IP []string `json:"ip,omitempty"`
		// perm String 否 子账户APIKey权限  read_only ：只读 ； trade  ：交易 多个权限用半角逗号隔开。 如果填写该字段，则该字段会被重置
		Perm okapi.APIKeyAccess `json:"perm,omitempty"`
	}
	QueryAPIKey struct {
		// apiKey String 是 子账户API的公钥
		APIKey string `json:"apiKey"`
		// subAcct String 可选 子账户名称，type 为 1 ， 2  或  4 ：subAcct 为必填项
		SubAcct string `json:"subAcct"`
	}
	DeleteAPIKey struct {
		Pwd string `json:"pwd"`
		// apiKey String 是 子账户API的公钥
		APIKey string `json:"apiKey"`
		// subAcct String 可选 子账户名称，type 为 1 ， 2  或  4 ：subAcct 为必填项
		SubAcct string `json:"subAcct"`
	}
	GetBalance struct {
		// subAcct String 可选 子账户名称，type 为 1 ， 2  或  4 ：subAcct 为必填项
		SubAcct string `json:"subAcct"`
	}
	HistoryTransfer struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// subAcct String 可选 子账户名称，type 为 1 ， 2  或  4 ：subAcct 为必填项
		SubAcct string `json:"subAcct,omitempty"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After int64 `json:"after,omitempty,string"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before int64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit int64 `json:"limit,omitempty,string"`
		// type String 报价方类型（当前未生效，将返回 "" ）
		Type okapi.TransferType `json:"type,omitempty,string"`
	}
	ManageTransfers struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// fromSubAccount String 是 转出子账户的子账户名称
		FromSubAccount string `json:"fromSubAccount"`
		ToSubAccount   string `json:"tiSubAccount"`
		// amt String 是 划转数量
		Amt float64 `json:"amt,string"`
		// from String 是 转出账户 6 ：资金账户  18 ：交易账户 from和to不可相同
		From okapi.AccountType `json:"from,string"`
		// to String 是 转入账户 6 ：资金账户  18 ：交易账户
		To okapi.AccountType `json:"to,string"`
	}
)

func (m *ViewList) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ ViewList ------┐", str)
	str = fmt.Sprintf("%s\r\n子账户名称: %v", str, m.SubAcct)
	str = fmt.Sprintf("%s\r\n子账户状态: %v", str, m.Enable)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n└----------- ViewList ------------┘", str)
	return str
}
func (m *CreateAPIKey) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ CreateAPIKey ------┐", str)
	str = fmt.Sprintf("%s\r\n: %v", str, m.Pwd)
	str = fmt.Sprintf("%s\r\n子账户名称: %v", str, m.SubAcct)
	str = fmt.Sprintf("%s\r\n子账户备注: %v", str, m.Label)
	str = fmt.Sprintf("%s\r\n: %v", str, m.Passphrase)
	str = fmt.Sprintf("%s\r\n子账户APIKey绑定ip地址: %v", str, m.IP)
	str = fmt.Sprintf("%s\r\n子账户APIKey权限: %v", str, m.Perm)
	str = fmt.Sprintf("%s\r\n└----------- CreateAPIKey ------------┘", str)
	return str
}
func (m *QueryAPIKey) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ QueryAPIKey ------┐", str)
	str = fmt.Sprintf("%s\r\n子账户API的公钥: %v", str, m.APIKey)
	str = fmt.Sprintf("%s\r\n子账户名称: %v", str, m.SubAcct)
	str = fmt.Sprintf("%s\r\n└----------- QueryAPIKey ------------┘", str)
	return str
}
func (m *DeleteAPIKey) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ DeleteAPIKey ------┐", str)
	str = fmt.Sprintf("%s\r\n: %v", str, m.Pwd)
	str = fmt.Sprintf("%s\r\n子账户API的公钥: %v", str, m.APIKey)
	str = fmt.Sprintf("%s\r\n子账户名称: %v", str, m.SubAcct)
	str = fmt.Sprintf("%s\r\n└----------- DeleteAPIKey ------------┘", str)
	return str
}
func (m *GetBalance) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetBalance ------┐", str)
	str = fmt.Sprintf("%s\r\n子账户名称: %v", str, m.SubAcct)
	str = fmt.Sprintf("%s\r\n└----------- GetBalance ------------┘", str)
	return str
}
func (m *HistoryTransfer) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ HistoryTransfer ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n子账户名称: %v", str, m.SubAcct)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n报价方类型（当前未生效: %v", str, m.Type)
	str = fmt.Sprintf("%s\r\n└----------- HistoryTransfer ------------┘", str)
	return str
}
func (m *ManageTransfers) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ ManageTransfers ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n转出子账户的子账户名称: %v", str, m.FromSubAccount)
	str = fmt.Sprintf("%s\r\n: %v", str, m.ToSubAccount)
	str = fmt.Sprintf("%s\r\n划转数量: %v", str, m.Amt)
	str = fmt.Sprintf("%s\r\n转出账户: %v", str, m.From)
	str = fmt.Sprintf("%s\r\n转入账户: %v", str, m.To)
	str = fmt.Sprintf("%s\r\n└----------- ManageTransfers ------------┘", str)
	return str
}
