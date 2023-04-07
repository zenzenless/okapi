package funding

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
)

type (
	Currency struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// name String 币种中文名称，不显示则无对应名称
		Name string `json:"name"`
		// chain String 币种链信息 有的币种下有多个链，必须要做区分，如 USDT 下有 USDT-ERC20 ， USDT-TRC20 ， USDT-Omni 多个链
		Chain string `json:"chain"`
		// minWd String 币种单笔最小提币量
		MinWd string `json:"minWd"`
		// minFee String 最小提币手续费数量
		MinFee string `json:"minFee"`
		// maxFee String 最大提币手续费数量
		MaxFee string `json:"maxFee"`
		// canDep Boolean 是否可充值， false 表示不可链上充值， true 表示可以链上充值
		CanDep bool `json:"canDep"`
		// canWd Boolean 是否可提币， false 表示不可链上提币， true 表示可以链上提币
		CanWd bool `json:"canWd"`
		// canInternal Boolean 是否可内部转账， false 表示不可内部转账， true 表示可以内部转账
		CanInternal bool `json:"canInternal"`
	}
	Balance struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// bal String 余额
		Bal string `json:"bal"`
		// frozenBal String 冻结（不可用）
		FrozenBal string `json:"frozenBal"`
		// availBal String 可用余额
		AvailBal string `json:"availBal"`
	}
	Transfer struct {
		// transId String 划转 ID
		TransID string `json:"transId"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// amt String 是 划转数量
		Amt okapi.JSONFloat64 `json:"amt"`
		// from String 是 转出账户 6 ：资金账户  18 ：交易账户 from和to不可相同
		From okapi.AccountType `json:"from,string"`
		// to String 是 转入账户 6 ：资金账户  18 ：交易账户
		To okapi.AccountType `json:"to,string"`
	}
	Bill struct {
		// billId String 账单 ID
		BillID string `json:"billId"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// bal String 余额
		Bal okapi.JSONFloat64 `json:"bal"`
		// balChg String 账户层面的余额变动数量
		BalChg okapi.JSONFloat64 `json:"balChg"`
		// type String 报价方类型（当前未生效，将返回 "" ）
		Type okapi.BillType `json:"type,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	DepositAddress struct {
		// addr String 充值地址
		Addr string `json:"addr"`
		// tag String 否 订单标签  字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		Tag string `json:"tag,omitempty"`
		// memo String 部分币种充值需要 memo，若不需要则不返回此字段
		Memo string `json:"memo,omitempty"`
		// pmtId String 部分币种充值需要此字段（payment_id），若不需要则不返回此字段
		PmtID string `json:"pmtId,omitempty"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// chain String 币种链信息 有的币种下有多个链，必须要做区分，如 USDT 下有 USDT-ERC20 ， USDT-TRC20 ， USDT-Omni 多个链
		Chain string `json:"chain"`
		// ctAddr String 合约地址后6位
		CtAddr string `json:"ctAddr"`
		// selected Boolean 该地址是否为页面选中的地址
		Selected bool `json:"selected"`
		// to String 是 转入账户 6 ：资金账户  18 ：交易账户
		To okapi.AccountType `json:"to,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	DepositHistory struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// chain String 币种链信息 有的币种下有多个链，必须要做区分，如 USDT 下有 USDT-ERC20 ， USDT-TRC20 ， USDT-Omni 多个链
		Chain string `json:"chain"`
		// txId String 否 区块转账哈希记录
		TxID string `json:"txId"`
		// from String 是 转出账户 6 ：资金账户  18 ：交易账户 from和to不可相同
		From string `json:"from"`
		// to String 是 转入账户 6 ：资金账户  18 ：交易账户
		To string `json:"to"`
		// depId String 否 充值记录 ID
		DepId string `json:"depId"`
		// amt String 是 划转数量
		Amt okapi.JSONFloat64 `json:"amt"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State okapi.DepositState `json:"state,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	Withdrawal struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// chain String 币种链信息 有的币种下有多个链，必须要做区分，如 USDT 下有 USDT-ERC20 ， USDT-TRC20 ， USDT-Omni 多个链
		Chain string `json:"chain"`
		// wdId String 提币申请ID
		WdID okapi.JSONInt64 `json:"wdId"`
		// amt String 是 划转数量
		Amt okapi.JSONFloat64 `json:"amt"`
	}
	WithdrawalHistory struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// chain String 币种链信息 有的币种下有多个链，必须要做区分，如 USDT 下有 USDT-ERC20 ， USDT-TRC20 ， USDT-Omni 多个链
		Chain string `json:"chain"`
		// txId String 否 区块转账哈希记录
		TxID string `json:"txId"`
		// from String 是 转出账户 6 ：资金账户  18 ：交易账户 from和to不可相同
		From string `json:"from"`
		// to String 是 转入账户 6 ：资金账户  18 ：交易账户
		To string `json:"to"`
		// tag String 否 订单标签  字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		Tag string `json:"tag,omitempty"`
		// pmtId String 部分币种充值需要此字段（payment_id），若不需要则不返回此字段
		PmtID string `json:"pmtId,omitempty"`
		// memo String 部分币种充值需要 memo，若不需要则不返回此字段
		Memo string `json:"memo,omitempty"`
		// amt String 是 划转数量
		Amt okapi.JSONFloat64 `json:"amt"`
		// fee String 订单交易手续费，平台向用户收取的交易手续费，手续费扣除为 负数 。如：  -0.01
		Fee okapi.JSONInt64 `json:"fee"`
		// wdId String 提币申请ID
		WdID okapi.JSONInt64 `json:"wdId"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State okapi.WithdrawalState `json:"state,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	PiggyBank struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// amt String 是 划转数量
		Amt okapi.JSONFloat64 `json:"amt"`
		// side String 是 订单方向   buy ：买，  sell ：卖
		Side okapi.ActionType `json:"side,string"`
	}
	PiggyBankBalance struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// amt String 是 划转数量
		Amt okapi.JSONFloat64 `json:"amt"`
		// earnings String 币种持仓收益
		Earnings okapi.JSONFloat64 `json:"earnings"`
	}
)

func (m *Currency) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Currency ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Name); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种中文名称: %v", str, m.Name)
	}
	if s := fmt.Sprintf("%v", m.Chain); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种链信息: %v", str, m.Chain)
	}
	if s := fmt.Sprintf("%v", m.MinWd); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种单笔最小提币量: %v", str, m.MinWd)
	}
	if s := fmt.Sprintf("%v", m.MinFee); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最小提币手续费数量: %v", str, m.MinFee)
	}
	if s := fmt.Sprintf("%v", m.MaxFee); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最大提币手续费数量: %v", str, m.MaxFee)
	}
	if s := fmt.Sprintf("%v", m.CanDep); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n是否可充值: %v", str, m.CanDep)
	}
	if s := fmt.Sprintf("%v", m.CanWd); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n是否可提币: %v", str, m.CanWd)
	}
	if s := fmt.Sprintf("%v", m.CanInternal); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n是否可内部转账: %v", str, m.CanInternal)
	}
	str = fmt.Sprintf("%s\r\n└----------- Currency ------------┘", str)
	return str
}
func (m *Balance) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Balance ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Bal); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n余额: %v", str, m.Bal)
	}
	if s := fmt.Sprintf("%v", m.FrozenBal); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n冻结（不可用）: %v", str, m.FrozenBal)
	}
	if s := fmt.Sprintf("%v", m.AvailBal); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n可用余额: %v", str, m.AvailBal)
	}
	str = fmt.Sprintf("%s\r\n└----------- Balance ------------┘", str)
	return str
}
func (m *Transfer) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Transfer ------┐", str)
	if s := fmt.Sprintf("%v", m.TransID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n划转: %v", str, m.TransID)
	}
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Amt); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n划转数量: %v", str, m.Amt)
	}
	if s := fmt.Sprintf("%v", m.From); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n转出账户: %v", str, m.From)
	}
	if s := fmt.Sprintf("%v", m.To); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n转入账户: %v", str, m.To)
	}
	str = fmt.Sprintf("%s\r\n└----------- Transfer ------------┘", str)
	return str
}
func (m *Bill) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Bill ------┐", str)
	if s := fmt.Sprintf("%v", m.BillID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n账单: %v", str, m.BillID)
	}
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Bal); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n余额: %v", str, m.Bal)
	}
	if s := fmt.Sprintf("%v", m.BalChg); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n账户层面的余额变动数量: %v", str, m.BalChg)
	}
	if s := fmt.Sprintf("%v", m.Type); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n报价方类型（当前未生效: %v", str, m.Type)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- Bill ------------┘", str)
	return str
}
func (m *DepositAddress) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ DepositAddress ------┐", str)
	if s := fmt.Sprintf("%v", m.Addr); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n充值地址: %v", str, m.Addr)
	}
	if s := fmt.Sprintf("%v", m.Tag); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单标签: %v", str, m.Tag)
	}
	if s := fmt.Sprintf("%v", m.Memo); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n部分币种充值需要: %v", str, m.Memo)
	}
	if s := fmt.Sprintf("%v", m.PmtID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n部分币种充值需要此字段（payment_id）: %v", str, m.PmtID)
	}
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Chain); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种链信息: %v", str, m.Chain)
	}
	if s := fmt.Sprintf("%v", m.CtAddr); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n合约地址后6位: %v", str, m.CtAddr)
	}
	if s := fmt.Sprintf("%v", m.Selected); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n该地址是否为页面选中的地址: %v", str, m.Selected)
	}
	if s := fmt.Sprintf("%v", m.To); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n转入账户: %v", str, m.To)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- DepositAddress ------------┘", str)
	return str
}
func (m *DepositHistory) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ DepositHistory ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Chain); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种链信息: %v", str, m.Chain)
	}
	if s := fmt.Sprintf("%v", m.TxID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n区块转账哈希记录: %v", str, m.TxID)
	}
	if s := fmt.Sprintf("%v", m.From); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n转出账户: %v", str, m.From)
	}
	if s := fmt.Sprintf("%v", m.To); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n转入账户: %v", str, m.To)
	}
	if s := fmt.Sprintf("%v", m.DepId); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n充值记录: %v", str, m.DepId)
	}
	if s := fmt.Sprintf("%v", m.Amt); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n划转数量: %v", str, m.Amt)
	}
	if s := fmt.Sprintf("%v", m.State); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单状态: %v", str, m.State)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- DepositHistory ------------┘", str)
	return str
}
func (m *Withdrawal) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Withdrawal ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Chain); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种链信息: %v", str, m.Chain)
	}
	if s := fmt.Sprintf("%v", m.WdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n提币申请ID: %v", str, m.WdID)
	}
	if s := fmt.Sprintf("%v", m.Amt); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n划转数量: %v", str, m.Amt)
	}
	str = fmt.Sprintf("%s\r\n└----------- Withdrawal ------------┘", str)
	return str
}
func (m *WithdrawalHistory) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ WithdrawalHistory ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Chain); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种链信息: %v", str, m.Chain)
	}
	if s := fmt.Sprintf("%v", m.TxID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n区块转账哈希记录: %v", str, m.TxID)
	}
	if s := fmt.Sprintf("%v", m.From); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n转出账户: %v", str, m.From)
	}
	if s := fmt.Sprintf("%v", m.To); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n转入账户: %v", str, m.To)
	}
	if s := fmt.Sprintf("%v", m.Tag); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单标签: %v", str, m.Tag)
	}
	if s := fmt.Sprintf("%v", m.PmtID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n部分币种充值需要此字段（payment_id）: %v", str, m.PmtID)
	}
	if s := fmt.Sprintf("%v", m.Memo); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n部分币种充值需要: %v", str, m.Memo)
	}
	if s := fmt.Sprintf("%v", m.Amt); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n划转数量: %v", str, m.Amt)
	}
	if s := fmt.Sprintf("%v", m.Fee); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单交易手续费: %v", str, m.Fee)
	}
	if s := fmt.Sprintf("%v", m.WdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n提币申请ID: %v", str, m.WdID)
	}
	if s := fmt.Sprintf("%v", m.State); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单状态: %v", str, m.State)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- WithdrawalHistory ------------┘", str)
	return str
}
func (m *PiggyBank) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ PiggyBank ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Amt); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n划转数量: %v", str, m.Amt)
	}
	if s := fmt.Sprintf("%v", m.Side); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单方向: %v", str, m.Side)
	}
	str = fmt.Sprintf("%s\r\n└----------- PiggyBank ------------┘", str)
	return str
}
func (m *PiggyBankBalance) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ PiggyBankBalance ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Amt); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n划转数量: %v", str, m.Amt)
	}
	if s := fmt.Sprintf("%v", m.Earnings); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种持仓收益: %v", str, m.Earnings)
	}
	str = fmt.Sprintf("%s\r\n└----------- PiggyBankBalance ------------┘", str)
	return str
}
