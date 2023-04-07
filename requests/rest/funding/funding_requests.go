package funding

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
)

type (
	GetBalance struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy []string `json:"ccy,omitempty"`
	}
	FundsTransfer struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// amt String 是 划转数量
		Amt float64 `json:"amt,string"`
		// subAcct String 可选 子账户名称，type 为 1 ， 2  或  4 ：subAcct 为必填项
		SubAcct string `json:"subAcct,omitempty"`
		InstID  string `json:"instID,omitempty"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		ToInstID string `json:"instId,omitempty"`
		// type String 报价方类型（当前未生效，将返回 "" ）
		Type okapi.TransferType `json:"type,omitempty,string"`
		// from String 是 转出账户 6 ：资金账户  18 ：交易账户 from和to不可相同
		From okapi.AccountType `json:"from,string"`
		// to String 是 转入账户 6 ：资金账户  18 ：交易账户
		To okapi.AccountType `json:"to,string"`
	}
	AssetBillsDetails struct {
		// type String 报价方类型（当前未生效，将返回 "" ）
		Type okapi.BillType `json:"type,string,omitempty"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After int64 `json:"after,string,omitempty"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before int64 `json:"before,string,omitempty"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit int64 `json:"limit,string,omitempty"`
	}
	GetDepositAddress struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
	}
	GetDepositHistory struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// txId String 否 区块转账哈希记录
		TxID string `json:"txId,omitempty"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After int64 `json:"after,omitempty,string"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before int64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit int64 `json:"limit,omitempty,string"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State okapi.DepositState `json:"state,omitempty,string"`
	}
	Withdrawal struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// chain String 币种链信息 有的币种下有多个链，必须要做区分，如 USDT 下有 USDT-ERC20 ， USDT-TRC20 ， USDT-Omni 多个链
		Chain string `json:"chain,omitempty"`
		// toAddr String 是 如果选择链上提币， toAddr 必须是认证过的数字货币地址。某些数字货币地址格式为 地址:标签 ，如ARDOR-7JF3-8F2E-QUWZ-CAN7F:123456 如果选择内部转账， toAddr 必须是接收方地址，可以是邮箱、手机或者账户名。
		ToAddr string `json:"toAddr"`
		Pwd    string `json:"pwd"`
		// amt String 是 划转数量
		Amt float64 `json:"amt,string"`
		// fee String 订单交易手续费，平台向用户收取的交易手续费，手续费扣除为 负数 。如：  -0.01
		Fee float64 `json:"fee,string"`
		// dest String 是 提币方式 3 ：内部转账  4 ：链上提币
		Dest okapi.WithdrawalDestination `json:"dest,string"`
	}
	GetWithdrawalHistory struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// txId String 否 区块转账哈希记录
		TxID string `json:"txId,omitempty"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After int64 `json:"after,omitempty,string"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before int64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit int64 `json:"limit,omitempty,string"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State okapi.WithdrawalState `json:"state,omitempty,string"`
	}
	PiggyBankPurchaseRedemption struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// txId String 否 区块转账哈希记录
		TxID string `json:"txId,omitempty"`
		// after String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的 ordId
		After int64 `json:"after,omitempty,string"`
		// before String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的 ordId 时间戳分页时，不支持使用before参数
		Before int64 `json:"before,omitempty,string"`
		// limit String 否 返回结果的数量，最大为100，默认100条
		Limit int64 `json:"limit,omitempty,string"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State okapi.WithdrawalState `json:"state,omitempty,string"`
	}
	GetPiggyBankBalance struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
	}
)

func (m *GetBalance) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetBalance ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n└----------- GetBalance ------------┘", str)
	return str
}
func (m *FundsTransfer) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ FundsTransfer ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n划转数量: %v", str, m.Amt)
	str = fmt.Sprintf("%s\r\n子账户名称: %v", str, m.SubAcct)
	str = fmt.Sprintf("%s\r\n: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.ToInstID)
	str = fmt.Sprintf("%s\r\n报价方类型（当前未生效: %v", str, m.Type)
	str = fmt.Sprintf("%s\r\n转出账户: %v", str, m.From)
	str = fmt.Sprintf("%s\r\n转入账户: %v", str, m.To)
	str = fmt.Sprintf("%s\r\n└----------- FundsTransfer ------------┘", str)
	return str
}
func (m *AssetBillsDetails) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ AssetBillsDetails ------┐", str)
	str = fmt.Sprintf("%s\r\n报价方类型（当前未生效: %v", str, m.Type)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n└----------- AssetBillsDetails ------------┘", str)
	return str
}
func (m *GetDepositAddress) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetDepositAddress ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n└----------- GetDepositAddress ------------┘", str)
	return str
}
func (m *GetDepositHistory) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetDepositHistory ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n区块转账哈希记录: %v", str, m.TxID)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n订单状态: %v", str, m.State)
	str = fmt.Sprintf("%s\r\n└----------- GetDepositHistory ------------┘", str)
	return str
}
func (m *Withdrawal) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Withdrawal ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n币种链信息: %v", str, m.Chain)
	str = fmt.Sprintf("%s\r\n如果选择链上提币: %v", str, m.ToAddr)
	str = fmt.Sprintf("%s\r\n: %v", str, m.Pwd)
	str = fmt.Sprintf("%s\r\n划转数量: %v", str, m.Amt)
	str = fmt.Sprintf("%s\r\n订单交易手续费: %v", str, m.Fee)
	str = fmt.Sprintf("%s\r\n提币方式: %v", str, m.Dest)
	str = fmt.Sprintf("%s\r\n└----------- Withdrawal ------------┘", str)
	return str
}
func (m *GetWithdrawalHistory) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetWithdrawalHistory ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n区块转账哈希记录: %v", str, m.TxID)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n订单状态: %v", str, m.State)
	str = fmt.Sprintf("%s\r\n└----------- GetWithdrawalHistory ------------┘", str)
	return str
}
func (m *PiggyBankPurchaseRedemption) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ PiggyBankPurchaseRedemption ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n区块转账哈希记录: %v", str, m.TxID)
	str = fmt.Sprintf("%s\r\n请求此ID之前（更旧的数据）的分页内容: %v", str, m.After)
	str = fmt.Sprintf("%s\r\n请求此ID之后（更新的数据）的分页内容: %v", str, m.Before)
	str = fmt.Sprintf("%s\r\n返回结果的数量: %v", str, m.Limit)
	str = fmt.Sprintf("%s\r\n订单状态: %v", str, m.State)
	str = fmt.Sprintf("%s\r\n└----------- PiggyBankPurchaseRedemption ------------┘", str)
	return str
}
func (m *GetPiggyBankBalance) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetPiggyBankBalance ------┐", str)
	str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	str = fmt.Sprintf("%s\r\n└----------- GetPiggyBankBalance ------------┘", str)
	return str
}
