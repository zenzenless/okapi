package subaccount

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
)

type (
	SubAccount struct {
		// subAcct String 可选 子账户名称，type 为 1 ， 2  或  4 ：subAcct 为必填项
		SubAcct string `json:"subAcct,omitempty"`
		// label String 子账户备注
		Label string `json:"label,omitempty"`
		// mobile String 子账户绑定手机号
		Mobile string `json:"mobile,omitempty"`
		// gAuth Boolean 子账户是否开启的登录时的谷歌验证  true ：已开启  false ：未开启
		GAuth bool `json:"gAuth"`
		// enable String 否 子账户状态， true ：正常使用  false ：冻结
		Enable bool `json:"enable"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	APIKey struct {
		// subAcct String 可选 子账户名称，type 为 1 ， 2  或  4 ：subAcct 为必填项
		SubAcct string `json:"subAcct,omitempty"`
		// label String 子账户备注
		Label string `json:"label,omitempty"`
		// apiKey String 是 子账户API的公钥
		APIKey     string `json:"apiKey,omitempty"`
		SecretKey  string `json:"secretKey,omitempty"`
		Passphrase string `json:"Passphrase,omitempty"`
		// perm String 否 子账户APIKey权限  read_only ：只读 ； trade  ：交易 多个权限用半角逗号隔开。 如果填写该字段，则该字段会被重置
		Perm string `json:"perm,omitempty"`
		// ip String 否 子账户APIKey绑定ip地址，多个ip用半角逗号隔开，最多支持20个ip。  如果填写该字段，那该字段会被重置如果ip传""，则表示解除IP绑定
		IP string `json:"ip,omitempty"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts,omitempty"`
	}
	HistoryTransfer struct {
		// subAcct String 可选 子账户名称，type 为 1 ， 2  或  4 ：subAcct 为必填项
		SubAcct string `json:"subAcct,omitempty"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// billId String 账单 ID
		BillID okapi.JSONInt64 `json:"billId,omitempty"`
		// type String 报价方类型（当前未生效，将返回 "" ）
		Type okapi.BillType `json:"type,omitempty,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts,omitempty"`
	}
	Transfer struct {
		// transId String 划转 ID
		TransID okapi.JSONInt64 `json:"transId"`
	}
)

func (m *SubAccount) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ SubAccount ------┐", str)
	if s := fmt.Sprintf("%v", m.SubAcct); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n子账户名称: %v", str, m.SubAcct)
	}
	if s := fmt.Sprintf("%v", m.Label); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n子账户备注: %v", str, m.Label)
	}
	if s := fmt.Sprintf("%v", m.Mobile); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n子账户绑定手机号: %v", str, m.Mobile)
	}
	if s := fmt.Sprintf("%v", m.GAuth); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n子账户是否开启的登录时的谷歌验证: %v", str, m.GAuth)
	}
	if s := fmt.Sprintf("%v", m.Enable); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n子账户状态: %v", str, m.Enable)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- SubAccount ------------┘", str)
	return str
}
func (m *APIKey) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ APIKey ------┐", str)
	if s := fmt.Sprintf("%v", m.SubAcct); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n子账户名称: %v", str, m.SubAcct)
	}
	if s := fmt.Sprintf("%v", m.Label); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n子账户备注: %v", str, m.Label)
	}
	if s := fmt.Sprintf("%v", m.APIKey); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n子账户API的公钥: %v", str, m.APIKey)
	}
	if s := fmt.Sprintf("%v", m.SecretKey); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n: %v", str, m.SecretKey)
	}
	if s := fmt.Sprintf("%v", m.Passphrase); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n: %v", str, m.Passphrase)
	}
	if s := fmt.Sprintf("%v", m.Perm); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n子账户APIKey权限: %v", str, m.Perm)
	}
	if s := fmt.Sprintf("%v", m.IP); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n子账户APIKey绑定ip地址: %v", str, m.IP)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- APIKey ------------┘", str)
	return str
}
func (m *HistoryTransfer) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ HistoryTransfer ------┐", str)
	if s := fmt.Sprintf("%v", m.SubAcct); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n子账户名称: %v", str, m.SubAcct)
	}
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.BillID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n账单: %v", str, m.BillID)
	}
	if s := fmt.Sprintf("%v", m.Type); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n报价方类型（当前未生效: %v", str, m.Type)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- HistoryTransfer ------------┘", str)
	return str
}
func (m *Transfer) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Transfer ------┐", str)
	if s := fmt.Sprintf("%v", m.TransID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n划转: %v", str, m.TransID)
	}
	str = fmt.Sprintf("%s\r\n└----------- Transfer ------------┘", str)
	return str
}
