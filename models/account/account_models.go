package account

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
)

type (
	Balance struct {
		// totalEq String 美金层面权益
		TotalEq okapi.JSONFloat64 `json:"totalEq"`
		// isoEq String 美金层面逐仓仓位权益 适用于 单币种保证金模式 和 跨币种保证金模式 和 组合保证金模式
		IsoEq okapi.JSONFloat64 `json:"isoEq"`
		// adjEq String 美金层面有效保证金 适用于 跨币种保证金模式 和 组合保证金模式
		AdjEq okapi.JSONFloat64 `json:"adjEq,omitempty"`
		// ordFroz String 美金层面全仓挂单占用保证金 适用于 跨币种保证金模式 和 组合保证金模式
		OrdFroz okapi.JSONFloat64 `json:"ordFroz,omitempty"`
		// imr String 美金层面占用保证金 适用于 跨币种保证金模式 和 组合保证金模式
		Imr okapi.JSONFloat64 `json:"imr,omitempty"`
		// mmr String 美金层面维持保证金 适用于 跨币种保证金模式 和 组合保证金模式
		Mmr okapi.JSONFloat64 `json:"mmr,omitempty"`
		// mgnRatio String 美金层面保证金率 适用于 跨币种保证金模式  和 组合保证金模式
		MgnRatio okapi.JSONFloat64 `json:"mgnRatio,omitempty"`
		// notionalUsd String 以美金价值为单位的持仓数量，即仓位美金价值 适用于 跨币种保证金模式 和 组合保证金模式
		NotionalUsd okapi.JSONFloat64 `json:"notionalUsd,omitempty"`
		// details Array 各个账户的资产估值
		Details []*BalanceDetails `json:"details,omitempty"`
		// uTime String 订单状态更新时间，Unix时间戳的毫秒数格式，如： 1597026383085
		UTime okapi.JSONTime `json:"uTime"`
	}
	BalanceDetails struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// eq String 币种总权益
		Eq okapi.JSONFloat64 `json:"eq"`
		// cashBal String 币种余额
		CashBal okapi.JSONFloat64 `json:"cashBal"`
		// isoEq String 美金层面逐仓仓位权益 适用于 单币种保证金模式 和 跨币种保证金模式 和 组合保证金模式
		IsoEq okapi.JSONFloat64 `json:"isoEq,omitempty"`
		// availEq String 可用保证金 适用于 单币种保证金模式 和 跨币种保证金模式 和 组合保证金模式
		AvailEq okapi.JSONFloat64 `json:"availEq,omitempty"`
		// disEq String 美金层面币种折算权益
		DisEq okapi.JSONFloat64 `json:"disEq"`
		// availBal String 可用余额
		AvailBal okapi.JSONFloat64 `json:"availBal"`
		// frozenBal String 冻结（不可用）
		FrozenBal okapi.JSONFloat64 `json:"frozenBal"`
		// ordFrozen String 挂单冻结数量
		OrdFrozen okapi.JSONFloat64 `json:"ordFrozen"`
		// liab String 币种负债额 适用于 跨币种保证金模式 和 组合保证金模式
		Liab okapi.JSONFloat64 `json:"liab,omitempty"`
		// upl String 未实现盈亏 适用于 单币种保证金模式 和 跨币种保证金模式 和 组合保证金模式
		Upl    okapi.JSONFloat64 `json:"upl,omitempty"`
		UplLib okapi.JSONFloat64 `json:"uplLib,omitempty"`
		// crossLiab String 币种全仓负债额 适用于 跨币种保证金模式 和 组合保证金模式
		CrossLiab okapi.JSONFloat64 `json:"crossLiab,omitempty"`
		// isoLiab String 币种逐仓负债额 适用于 跨币种保证金模式 和 组合保证金模式
		IsoLiab okapi.JSONFloat64 `json:"isoLiab,omitempty"`
		// mgnRatio String 美金层面保证金率 适用于 跨币种保证金模式  和 组合保证金模式
		MgnRatio okapi.JSONFloat64 `json:"mgnRatio,omitempty"`
		// interest String 计息 适用于 跨币种保证金模式 和 组合保证金模式
		Interest okapi.JSONFloat64 `json:"interest,omitempty"`
		// twap String 当前负债币种触发系统自动换币的风险 0、1、2、3、4、5其中之一，数字越大代表您的负债币种触发自动换币概率越高 适用于 跨币种保证金模式 和 组合保证金模式
		Twap okapi.JSONFloat64 `json:"twap,omitempty"`
		// maxLoan String 币种最大可借 适用于 跨币种保证金模式 和 组合保证金模式
		MaxLoan okapi.JSONFloat64 `json:"maxLoan,omitempty"`
		// eqUsd String 币种权益美金价值
		EqUsd okapi.JSONFloat64 `json:"eqUsd"`
		// notionalLever String 币种杠杆倍数 适用于 单币种保证金模式
		NotionalLever okapi.JSONFloat64 `json:"notionalLever,omitempty"`
		// stgyEq String 策略权益
		StgyEq okapi.JSONFloat64 `json:"stgyEq"`
		// isoUpl String 逐仓未实现盈亏 适用于 单币种保证金模式 和 跨币种保证金模式 和 组合保证金模式
		IsoUpl okapi.JSONFloat64 `json:"isoUpl,omitempty"`
		// uTime String 订单状态更新时间，Unix时间戳的毫秒数格式，如： 1597026383085
		UTime okapi.JSONTime `json:"uTime"`
	}
	Position struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// posCcy String 仓位资产币种，仅适用于 币币杠杆 仓位
		PosCcy string `json:"posCcy,omitempty"`
		// liabCcy String 负债币种，仅适用于 币币杠杆
		LiabCcy string `json:"liabCcy,omitempty"`
		// optVal String 期权市值，仅适用于 期权
		OptVal string `json:"optVal,omitempty"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// posId String 否 持仓ID 支持多个 posId 查询（不超过20个），半角逗号分割
		PosID string `json:"posId"`
		// tradeId String 最新成交ID
		TradeID string `json:"tradeId"`
		// pos String 持仓数量，逐仓自主划转模式下，转入保证金后会产生pos为 0 的仓位
		Pos okapi.JSONFloat64 `json:"pos"`
		// availPos String 可平仓数量，适用于币币杠杆 , 交割/永续 （开平仓模式）， 期权 （交易账户及保证金账户逐仓）。
		AvailPos okapi.JSONFloat64 `json:"availPos,omitempty"`
		// avgPx String 成交均价，如果成交数量为0，该字段也为 0
		AvgPx okapi.JSONFloat64 `json:"avgPx"`
		// upl String 未实现盈亏 适用于 单币种保证金模式 和 跨币种保证金模式 和 组合保证金模式
		Upl okapi.JSONFloat64 `json:"upl"`
		// uplRatio String 未实现收益率
		UplRatio okapi.JSONFloat64 `json:"uplRatio"`
		// lever String 杠杆倍数，0.01到125之间的数值，仅适用于  币币杠杆/交割/永续
		Lever okapi.JSONFloat64 `json:"lever"`
		// liqPx String 预估强平价  不适用于 期权
		LiqPx okapi.JSONFloat64 `json:"liqPx,omitempty"`
		// imr String 美金层面占用保证金 适用于 跨币种保证金模式 和 组合保证金模式
		Imr okapi.JSONFloat64 `json:"imr,omitempty"`
		// margin String 保证金余额，可增减，仅适用于 逐仓
		Margin okapi.JSONFloat64 `json:"margin,omitempty"`
		// mgnRatio String 美金层面保证金率 适用于 跨币种保证金模式  和 组合保证金模式
		MgnRatio okapi.JSONFloat64 `json:"mgnRatio"`
		// mmr String 美金层面维持保证金 适用于 跨币种保证金模式 和 组合保证金模式
		Mmr okapi.JSONFloat64 `json:"mmr"`
		// liab String 币种负债额 适用于 跨币种保证金模式 和 组合保证金模式
		Liab okapi.JSONFloat64 `json:"liab,omitempty"`
		// interest String 计息 适用于 跨币种保证金模式 和 组合保证金模式
		Interest okapi.JSONFloat64 `json:"interest"`
		// notionalUsd String 以美金价值为单位的持仓数量，即仓位美金价值 适用于 跨币种保证金模式 和 组合保证金模式
		NotionalUsd okapi.JSONFloat64 `json:"notionalUsd"`
		// adl String 信号区 分为5档，从1到5，数字越小代表adl强度越弱
		ADL okapi.JSONFloat64 `json:"adl"`
		// last String 最新成交价
		Last okapi.JSONFloat64 `json:"last"`
		// deltaBS String 美金本位持仓仓位delta，仅适用于 期权
		DeltaBS okapi.JSONFloat64 `json:"deltaBS"`
		// deltaPA String 币本位持仓仓位delta，仅适用于 期权
		DeltaPA okapi.JSONFloat64 `json:"deltaPA"`
		// gammaBS String 美金本位持仓仓位gamma，仅适用于 期权
		GammaBS okapi.JSONFloat64 `json:"gammaBS"`
		// gammaPA String 币本位持仓仓位gamma，仅适用于 期权
		GammaPA okapi.JSONFloat64 `json:"gammaPA"`
		// thetaBS String 美金本位持仓仓位theta，仅适用于 期权
		ThetaBS okapi.JSONFloat64 `json:"thetaBS"`
		// thetaPA String 币本位持仓仓位theta，仅适用于 期权
		ThetaPA okapi.JSONFloat64 `json:"thetaPA"`
		// vegaBS String 美金本位持仓仓位vega，仅适用于 期权
		VegaBS okapi.JSONFloat64 `json:"vegaBS"`
		// vegaPA String 币本位持仓仓位vega，仅适用于 期权
		VegaPA okapi.JSONFloat64 `json:"vegaPA"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okapi.PositionSide `json:"posSide"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okapi.MarginMode `json:"mgnMode"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// cTime String 订单创建时间，Unix时间戳的毫秒数格式， 如 ： 1597026383085
		CTime okapi.JSONTime `json:"cTime"`
		// uTime String 订单状态更新时间，Unix时间戳的毫秒数格式，如： 1597026383085
		UTime okapi.JSONTime `json:"uTime"`
		PTime okapi.JSONTime `json:"pTime,omitempty"`
	}
	BalanceAndPosition struct {
		// eventType String 事件类型，枚举值： snapshot ：首推快照  delivered ：交割      exercised ：行权  transferred ：划转 filled ：成交      liquidation ：强平  claw_back ：穿仓补偿   adl ：ADL自动减仓      funding_fee ：资金费  adjust_margin ：调整保证金      set_leverage ：设置杠杆  interest_deduction ：扣息
		EventType okapi.EventType `json:"eventType"`
		// pTime String 持仓信息的推送时间，Unix时间戳的毫秒数格式，如  1597026383085
		PTime okapi.JSONTime `json:"pTime"`
		// uTime String 订单状态更新时间，Unix时间戳的毫秒数格式，如： 1597026383085
		UTime okapi.JSONTime `json:"uTime"`
		// posData Array 持仓详细信息
		PosData []*Position `json:"posData"`
		// balData Array 币种资产信息
		BalData []*BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		// adjEq String 美金层面有效保证金 适用于 跨币种保证金模式 和 组合保证金模式
		AdjEq okapi.JSONFloat64 `json:"adjEq,omitempty"`
		// balData Array 币种资产信息
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		// posData Array 持仓详细信息
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// eq String 币种总权益
		Eq okapi.JSONFloat64 `json:"eq"`
		// disEq String 美金层面币种折算权益
		DisEq okapi.JSONFloat64 `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// posCcy String 仓位资产币种，仅适用于 币币杠杆 仓位
		PosCcy string `json:"posCcy,omitempty"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// notionalCcy String 以 币 为单位的持仓数量
		NotionalCcy okapi.JSONFloat64 `json:"notionalCcy"`
		// pos String 持仓数量，逐仓自主划转模式下，转入保证金后会产生pos为 0 的仓位
		Pos okapi.JSONFloat64 `json:"pos"`
		// notionalUsd String 以美金价值为单位的持仓数量，即仓位美金价值 适用于 跨币种保证金模式 和 组合保证金模式
		NotionalUsd okapi.JSONFloat64 `json:"notionalUsd"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okapi.PositionSide `json:"posSide"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okapi.MarginMode `json:"mgnMode"`
	}
	Bill struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// notes String 备注 仅适用于 资金划转 ，不是 资金划转 时，返回 ""
		Notes string `json:"notes"`
		// billId String 账单 ID
		BillID string `json:"billId"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID string `json:"ordId"`
		// balChg String 账户层面的余额变动数量
		BalChg okapi.JSONFloat64 `json:"balChg"`
		// posBalChg String 仓位层面的余额变动数量
		PosBalChg okapi.JSONFloat64 `json:"posBalChg"`
		// bal String 余额
		Bal okapi.JSONFloat64 `json:"bal"`
		// posBal String 仓位层面的余额数量
		PosBal okapi.JSONFloat64 `json:"posBal"`
		// sz String 是 委托数量
		Sz okapi.JSONFloat64 `json:"sz"`
		// pnl String 收益
		Pnl okapi.JSONFloat64 `json:"pnl"`
		// fee String 订单交易手续费，平台向用户收取的交易手续费，手续费扣除为 负数 。如：  -0.01
		Fee okapi.JSONFloat64 `json:"fee"`
		// from String 是 转出账户 6 ：资金账户  18 ：交易账户 from和to不可相同
		From okapi.AccountType `json:"from,string"`
		// to String 是 转入账户 6 ：资金账户  18 ：交易账户
		To okapi.AccountType `json:"to,string"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		MgnMode  okapi.MarginMode     `json:"MgnMode"`
		// type String 报价方类型（当前未生效，将返回 "" ）
		Type okapi.BillType `json:"type,string"`
		// subType String 否 账单子类型 1 ：买入  2 ：卖出  3 ：开多4 ：开空  5 ：平多  6 ：平空  9 ：市场借币扣息  11 ：转入12 ：转出  14 ：尊享借币扣息  160 ：手动追加保证金  161 ：手动减少保证金162 ：自动追加保证金  114 ：自动换币买入  115 ：自动换币卖出  118 ：系统换币转入119 ：系统换币转出  100 ：强减平多  101 ：强减平空  102 ：强减买入103 ：强减卖出  104 ：强平平多  105 ：强平平空  106 ：强平买入107 ：强平卖出  110 ：强平换币转入  111 ：强平换币转出  125 ：自动减仓平多126 ：自动减仓平空  127 ：自动减仓买入  128 ：自动减仓卖出  131 ：对冲买入132 ：对冲卖出  170 ：到期行权  171 ：到期被行权  172 ：到期作废112 ：交割平多  113 ：交割平空  117 ：交割/期权穿仓补偿  173 ：资金费支出174 ：资金费收入  200 :系统转入  201 :手动转入  202 :系统转出203 :手动转出  204 : 大宗交易买  205 : 大宗交易卖  206 : 大宗交易开多207 : 大宗交易开空  208 : 大宗交易平多  209 : 大宗交易平空  224 : 还债转入225 : 还债转出
		SubType okapi.BillSubType `json:"subType,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	Config struct {
		// level String 当前在平台上真实交易量的用户等级，例如  lv1
		Level string `json:"level"`
		// levelTmp String 特约用户的临时体验用户等级，例如  lv3
		LevelTmp string `json:"levelTmp"`
		// acctLv String 账户层级 1 ：简单交易模式， 2 ：单币种保证金模式， 3 ：跨币种保证金模式 ， 4 ：组合保证金模式
		AcctLv string `json:"acctLv"`
		// autoLoan Boolean 是否自动借币 true ：自动借币  false ：非自动借币
		AutoLoan bool `json:"autoLoan"`
		// uid String 账户ID，账户uid和app上的一致
		UID string `json:"uid"`
		// greeksType String 当前希腊字母展示方式 PA ：币本位  BS ：美元本位
		GreeksType okapi.GreekType `json:"greeksType"`
		// posMode String 持仓方式 long_short_mode ：双向持仓net_mode ：单向持仓 仅适用 交割/永续
		PosMode okapi.PositionType `json:"posMode"`
	}
	PositionMode struct {
		// posMode String 持仓方式 long_short_mode ：双向持仓net_mode ：单向持仓 仅适用 交割/永续
		PosMode okapi.PositionType `json:"posMode"`
	}
	Leverage struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// lever String 杠杆倍数，0.01到125之间的数值，仅适用于  币币杠杆/交割/永续
		Lever okapi.JSONInt64 `json:"lever"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okapi.MarginMode `json:"mgnMode"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okapi.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// maxBuy String 币币/币币杠杆 ：最大可买的交易币数量 单币种保证金模式下的全仓杠杆订单，为交易币数量 交割 / 永续 / 期权 ：最大可开多的合约张数
		MaxBuy okapi.JSONFloat64 `json:"maxBuy"`
		// maxSell String 币币/币币杠杆 ：最大可卖的计价币数量 单币种保证金模式下的全仓杠杆订单，为交易币数量 交割 / 永续 / 期权 ：最大可开空的合约张数
		MaxSell okapi.JSONFloat64 `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// availBuy String 最大买入可用数量
		AvailBuy okapi.JSONFloat64 `json:"availBuy"`
		// availSell String 最大卖出可用数量
		AvailSell okapi.JSONFloat64 `json:"availSell"`
	}
	MarginBalanceAmount struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// amt String 是 划转数量
		Amt okapi.JSONFloat64 `json:"amt"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okapi.PositionSide `json:"posSide,string"`
		// type String 报价方类型（当前未生效，将返回 "" ）
		Type okapi.CountAction `json:"type,string"`
	}
	Loan struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// mgnCcy String 可选 保证金币种，如  BTC 币币杠杆单币种全仓情况下必须指定保证金币种
		MgnCcy string `json:"mgnCcy"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// maxLoan String 币种最大可借 适用于 跨币种保证金模式 和 组合保证金模式
		MaxLoan okapi.JSONFloat64 `json:"maxLoan"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okapi.MarginMode `json:"mgnMode"`
		// side String 是 订单方向   buy ：买，  sell ：卖
		Side okapi.OrderSide `json:"side,string"`
	}
	Fee struct {
		// level String 当前在平台上真实交易量的用户等级，例如  lv1
		Level string `json:"level"`
		// taker String USDT&USDⓈ&Crypto 交易区的吃单手续费率，永续和交割合约时，为币本位合约费率
		Taker okapi.JSONFloat64 `json:"taker"`
		// maker String USDT&USDⓈ&Crypto 交易区挂单手续费率，永续和交割合约时，为币本位合约费率
		Maker okapi.JSONFloat64 `json:"maker"`
		// delivery String 交割手续费率
		Delivery okapi.JSONFloat64 `json:"delivery,omitempty"`
		// exercise String 行权手续费率
		Exercise okapi.JSONFloat64 `json:"exercise,omitempty"`
		// category String 订单种类  normal ：普通委托 twap ：TWAP自动换币adl ：ADL自动减仓   full_liquidation ：强制平仓 partial_liquidation ：强制减仓    delivery ：交割   ddh ：对冲减仓类型订单
		Category okapi.FeeCategory `json:"category,string"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	InterestAccrued struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// interest String 计息 适用于 跨币种保证金模式 和 组合保证金模式
		Interest okapi.JSONFloat64 `json:"interest"`
		// interestRate String 计息利率(小时)
		InterestRate okapi.JSONFloat64 `json:"interestRate"`
		// liab String 币种负债额 适用于 跨币种保证金模式 和 组合保证金模式
		Liab okapi.JSONFloat64 `json:"liab"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okapi.MarginMode `json:"mgnMode"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	InterestRate struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// interestRate String 计息利率(小时)
		InterestRate okapi.JSONFloat64 `json:"interestRate"`
	}
	Greek struct {
		// greeksType String 当前希腊字母展示方式 PA ：币本位  BS ：美元本位
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// maxWd String 币种单笔最大提币量
		MaxWd okapi.JSONFloat64 `json:"maxWd"`
	}
)

func (m *Balance) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Balance ------┐", str)
	if s := fmt.Sprintf("%v", m.TotalEq); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面权益: %v", str, m.TotalEq)
	}
	if s := fmt.Sprintf("%v", m.IsoEq); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面逐仓仓位权益: %v", str, m.IsoEq)
	}
	if s := fmt.Sprintf("%v", m.AdjEq); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面有效保证金: %v", str, m.AdjEq)
	}
	if s := fmt.Sprintf("%v", m.OrdFroz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面全仓挂单占用保证金: %v", str, m.OrdFroz)
	}
	if s := fmt.Sprintf("%v", m.Imr); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面占用保证金: %v", str, m.Imr)
	}
	if s := fmt.Sprintf("%v", m.Mmr); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面维持保证金: %v", str, m.Mmr)
	}
	if s := fmt.Sprintf("%v", m.MgnRatio); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面保证金率: %v", str, m.MgnRatio)
	}
	if s := fmt.Sprintf("%v", m.NotionalUsd); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n以美金价值为单位的持仓数量: %v", str, m.NotionalUsd)
	}
	if s := fmt.Sprintf("%v", m.Details); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n各个账户的资产估值: %v", str, m.Details)
	}
	if s := fmt.Sprintf("%v", m.UTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单状态更新时间: %v", str, m.UTime)
	}
	str = fmt.Sprintf("%s\r\n└----------- Balance ------------┘", str)
	return str
}
func (m *BalanceDetails) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ BalanceDetails ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Eq); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种总权益: %v", str, m.Eq)
	}
	if s := fmt.Sprintf("%v", m.CashBal); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种余额: %v", str, m.CashBal)
	}
	if s := fmt.Sprintf("%v", m.IsoEq); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面逐仓仓位权益: %v", str, m.IsoEq)
	}
	if s := fmt.Sprintf("%v", m.AvailEq); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n可用保证金: %v", str, m.AvailEq)
	}
	if s := fmt.Sprintf("%v", m.DisEq); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面币种折算权益: %v", str, m.DisEq)
	}
	if s := fmt.Sprintf("%v", m.AvailBal); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n可用余额: %v", str, m.AvailBal)
	}
	if s := fmt.Sprintf("%v", m.FrozenBal); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n冻结（不可用）: %v", str, m.FrozenBal)
	}
	if s := fmt.Sprintf("%v", m.OrdFrozen); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n挂单冻结数量: %v", str, m.OrdFrozen)
	}
	if s := fmt.Sprintf("%v", m.Liab); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种负债额: %v", str, m.Liab)
	}
	if s := fmt.Sprintf("%v", m.Upl); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n未实现盈亏: %v", str, m.Upl)
	}
	if s := fmt.Sprintf("%v", m.UplLib); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n: %v", str, m.UplLib)
	}
	if s := fmt.Sprintf("%v", m.CrossLiab); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种全仓负债额: %v", str, m.CrossLiab)
	}
	if s := fmt.Sprintf("%v", m.IsoLiab); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种逐仓负债额: %v", str, m.IsoLiab)
	}
	if s := fmt.Sprintf("%v", m.MgnRatio); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面保证金率: %v", str, m.MgnRatio)
	}
	if s := fmt.Sprintf("%v", m.Interest); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n计息: %v", str, m.Interest)
	}
	if s := fmt.Sprintf("%v", m.Twap); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n当前负债币种触发系统自动换币的风险: %v", str, m.Twap)
	}
	if s := fmt.Sprintf("%v", m.MaxLoan); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种最大可借: %v", str, m.MaxLoan)
	}
	if s := fmt.Sprintf("%v", m.EqUsd); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种权益美金价值: %v", str, m.EqUsd)
	}
	if s := fmt.Sprintf("%v", m.NotionalLever); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种杠杆倍数: %v", str, m.NotionalLever)
	}
	if s := fmt.Sprintf("%v", m.StgyEq); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n策略权益: %v", str, m.StgyEq)
	}
	if s := fmt.Sprintf("%v", m.IsoUpl); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n逐仓未实现盈亏: %v", str, m.IsoUpl)
	}
	if s := fmt.Sprintf("%v", m.UTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单状态更新时间: %v", str, m.UTime)
	}
	str = fmt.Sprintf("%s\r\n└----------- BalanceDetails ------------┘", str)
	return str
}
func (m *Position) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Position ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.PosCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n仓位资产币种: %v", str, m.PosCcy)
	}
	if s := fmt.Sprintf("%v", m.LiabCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n负债币种: %v", str, m.LiabCcy)
	}
	if s := fmt.Sprintf("%v", m.OptVal); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n期权市值: %v", str, m.OptVal)
	}
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.PosID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓ID: %v", str, m.PosID)
	}
	if s := fmt.Sprintf("%v", m.TradeID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交ID: %v", str, m.TradeID)
	}
	if s := fmt.Sprintf("%v", m.Pos); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓数量: %v", str, m.Pos)
	}
	if s := fmt.Sprintf("%v", m.AvailPos); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n可平仓数量: %v", str, m.AvailPos)
	}
	if s := fmt.Sprintf("%v", m.AvgPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交均价: %v", str, m.AvgPx)
	}
	if s := fmt.Sprintf("%v", m.Upl); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n未实现盈亏: %v", str, m.Upl)
	}
	if s := fmt.Sprintf("%v", m.UplRatio); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n未实现收益率: %v", str, m.UplRatio)
	}
	if s := fmt.Sprintf("%v", m.Lever); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n杠杆倍数: %v", str, m.Lever)
	}
	if s := fmt.Sprintf("%v", m.LiqPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n预估强平价: %v", str, m.LiqPx)
	}
	if s := fmt.Sprintf("%v", m.Imr); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面占用保证金: %v", str, m.Imr)
	}
	if s := fmt.Sprintf("%v", m.Margin); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金余额: %v", str, m.Margin)
	}
	if s := fmt.Sprintf("%v", m.MgnRatio); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面保证金率: %v", str, m.MgnRatio)
	}
	if s := fmt.Sprintf("%v", m.Mmr); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面维持保证金: %v", str, m.Mmr)
	}
	if s := fmt.Sprintf("%v", m.Liab); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种负债额: %v", str, m.Liab)
	}
	if s := fmt.Sprintf("%v", m.Interest); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n计息: %v", str, m.Interest)
	}
	if s := fmt.Sprintf("%v", m.NotionalUsd); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n以美金价值为单位的持仓数量: %v", str, m.NotionalUsd)
	}
	if s := fmt.Sprintf("%v", m.ADL); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n信号区: %v", str, m.ADL)
	}
	if s := fmt.Sprintf("%v", m.Last); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新成交价: %v", str, m.Last)
	}
	if s := fmt.Sprintf("%v", m.DeltaBS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金本位持仓仓位delta: %v", str, m.DeltaBS)
	}
	if s := fmt.Sprintf("%v", m.DeltaPA); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币本位持仓仓位delta: %v", str, m.DeltaPA)
	}
	if s := fmt.Sprintf("%v", m.GammaBS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金本位持仓仓位gamma: %v", str, m.GammaBS)
	}
	if s := fmt.Sprintf("%v", m.GammaPA); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币本位持仓仓位gamma: %v", str, m.GammaPA)
	}
	if s := fmt.Sprintf("%v", m.ThetaBS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金本位持仓仓位theta: %v", str, m.ThetaBS)
	}
	if s := fmt.Sprintf("%v", m.ThetaPA); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币本位持仓仓位theta: %v", str, m.ThetaPA)
	}
	if s := fmt.Sprintf("%v", m.VegaBS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金本位持仓仓位vega: %v", str, m.VegaBS)
	}
	if s := fmt.Sprintf("%v", m.VegaPA); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币本位持仓仓位vega: %v", str, m.VegaPA)
	}
	if s := fmt.Sprintf("%v", m.PosSide); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.PosSide)
	}
	if s := fmt.Sprintf("%v", m.MgnMode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金模式: %v", str, m.MgnMode)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.CTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单创建时间: %v", str, m.CTime)
	}
	if s := fmt.Sprintf("%v", m.UTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单状态更新时间: %v", str, m.UTime)
	}
	if s := fmt.Sprintf("%v", m.PTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓信息的推送时间: %v", str, m.PTime)
	}
	str = fmt.Sprintf("%s\r\n└----------- Position ------------┘", str)
	return str
}
func (m *BalanceAndPosition) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ BalanceAndPosition ------┐", str)
	if s := fmt.Sprintf("%v", m.EventType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n事件类型: %v", str, m.EventType)
	}
	if s := fmt.Sprintf("%v", m.PTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓信息的推送时间: %v", str, m.PTime)
	}
	if s := fmt.Sprintf("%v", m.UTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单状态更新时间: %v", str, m.UTime)
	}
	if s := fmt.Sprintf("%v", m.PosData); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓详细信息: %v", str, m.PosData)
	}
	if s := fmt.Sprintf("%v", m.BalData); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种资产信息: %v", str, m.BalData)
	}
	str = fmt.Sprintf("%s\r\n└----------- BalanceAndPosition ------------┘", str)
	return str
}
func (m *PositionAndAccountRisk) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ PositionAndAccountRisk ------┐", str)
	if s := fmt.Sprintf("%v", m.AdjEq); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面有效保证金: %v", str, m.AdjEq)
	}
	if s := fmt.Sprintf("%v", m.BalData); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种资产信息: %v", str, m.BalData)
	}
	if s := fmt.Sprintf("%v", m.PosData); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓详细信息: %v", str, m.PosData)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- PositionAndAccountRisk ------------┘", str)
	return str
}
func (m *PositionAndAccountRiskBalanceData) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ PositionAndAccountRiskBalanceData ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Eq); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种总权益: %v", str, m.Eq)
	}
	if s := fmt.Sprintf("%v", m.DisEq); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面币种折算权益: %v", str, m.DisEq)
	}
	str = fmt.Sprintf("%s\r\n└----------- PositionAndAccountRiskBalanceData ------------┘", str)
	return str
}
func (m *PositionAndAccountRiskPositionData) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ PositionAndAccountRiskPositionData ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.PosCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n仓位资产币种: %v", str, m.PosCcy)
	}
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.NotionalCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n以: %v", str, m.NotionalCcy)
	}
	if s := fmt.Sprintf("%v", m.Pos); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓数量: %v", str, m.Pos)
	}
	if s := fmt.Sprintf("%v", m.NotionalUsd); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n以美金价值为单位的持仓数量: %v", str, m.NotionalUsd)
	}
	if s := fmt.Sprintf("%v", m.PosSide); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.PosSide)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.MgnMode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金模式: %v", str, m.MgnMode)
	}
	str = fmt.Sprintf("%s\r\n└----------- PositionAndAccountRiskPositionData ------------┘", str)
	return str
}
func (m *Bill) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Bill ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Notes); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n备注: %v", str, m.Notes)
	}
	if s := fmt.Sprintf("%v", m.BillID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n账单: %v", str, m.BillID)
	}
	if s := fmt.Sprintf("%v", m.OrdID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单ID: %v", str, m.OrdID)
	}
	if s := fmt.Sprintf("%v", m.BalChg); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n账户层面的余额变动数量: %v", str, m.BalChg)
	}
	if s := fmt.Sprintf("%v", m.PosBalChg); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n仓位层面的余额变动数量: %v", str, m.PosBalChg)
	}
	if s := fmt.Sprintf("%v", m.Bal); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n余额: %v", str, m.Bal)
	}
	if s := fmt.Sprintf("%v", m.PosBal); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n仓位层面的余额数量: %v", str, m.PosBal)
	}
	if s := fmt.Sprintf("%v", m.Sz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n委托数量: %v", str, m.Sz)
	}
	if s := fmt.Sprintf("%v", m.Pnl); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n收益: %v", str, m.Pnl)
	}
	if s := fmt.Sprintf("%v", m.Fee); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单交易手续费: %v", str, m.Fee)
	}
	if s := fmt.Sprintf("%v", m.From); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n转出账户: %v", str, m.From)
	}
	if s := fmt.Sprintf("%v", m.To); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n转入账户: %v", str, m.To)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.MgnMode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n: %v", str, m.MgnMode)
	}
	if s := fmt.Sprintf("%v", m.Type); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n报价方类型（当前未生效: %v", str, m.Type)
	}
	if s := fmt.Sprintf("%v", m.SubType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n账单子类型: %v", str, m.SubType)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- Bill ------------┘", str)
	return str
}
func (m *Config) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Config ------┐", str)
	if s := fmt.Sprintf("%v", m.Level); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n当前在平台上真实交易量的用户等级: %v", str, m.Level)
	}
	if s := fmt.Sprintf("%v", m.LevelTmp); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n特约用户的临时体验用户等级: %v", str, m.LevelTmp)
	}
	if s := fmt.Sprintf("%v", m.AcctLv); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n账户层级: %v", str, m.AcctLv)
	}
	if s := fmt.Sprintf("%v", m.AutoLoan); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n是否自动借币: %v", str, m.AutoLoan)
	}
	if s := fmt.Sprintf("%v", m.UID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n账户ID: %v", str, m.UID)
	}
	if s := fmt.Sprintf("%v", m.GreeksType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n当前希腊字母展示方式: %v", str, m.GreeksType)
	}
	if s := fmt.Sprintf("%v", m.PosMode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓方式: %v", str, m.PosMode)
	}
	str = fmt.Sprintf("%s\r\n└----------- Config ------------┘", str)
	return str
}
func (m *PositionMode) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ PositionMode ------┐", str)
	if s := fmt.Sprintf("%v", m.PosMode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓方式: %v", str, m.PosMode)
	}
	str = fmt.Sprintf("%s\r\n└----------- PositionMode ------------┘", str)
	return str
}
func (m *Leverage) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Leverage ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Lever); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n杠杆倍数: %v", str, m.Lever)
	}
	if s := fmt.Sprintf("%v", m.MgnMode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金模式: %v", str, m.MgnMode)
	}
	if s := fmt.Sprintf("%v", m.PosSide); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.PosSide)
	}
	str = fmt.Sprintf("%s\r\n└----------- Leverage ------------┘", str)
	return str
}
func (m *MaxBuySellAmount) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ MaxBuySellAmount ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.MaxBuy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币币/币币杠杆: %v", str, m.MaxBuy)
	}
	if s := fmt.Sprintf("%v", m.MaxSell); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币币/币币杠杆: %v", str, m.MaxSell)
	}
	str = fmt.Sprintf("%s\r\n└----------- MaxBuySellAmount ------------┘", str)
	return str
}
func (m *MaxAvailableTradeAmount) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ MaxAvailableTradeAmount ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.AvailBuy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最大买入可用数量: %v", str, m.AvailBuy)
	}
	if s := fmt.Sprintf("%v", m.AvailSell); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最大卖出可用数量: %v", str, m.AvailSell)
	}
	str = fmt.Sprintf("%s\r\n└----------- MaxAvailableTradeAmount ------------┘", str)
	return str
}
func (m *MarginBalanceAmount) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ MarginBalanceAmount ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Amt); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n划转数量: %v", str, m.Amt)
	}
	if s := fmt.Sprintf("%v", m.PosSide); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.PosSide)
	}
	if s := fmt.Sprintf("%v", m.Type); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n报价方类型（当前未生效: %v", str, m.Type)
	}
	str = fmt.Sprintf("%s\r\n└----------- MarginBalanceAmount ------------┘", str)
	return str
}
func (m *Loan) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Loan ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.MgnCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.MgnCcy)
	}
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.MaxLoan); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种最大可借: %v", str, m.MaxLoan)
	}
	if s := fmt.Sprintf("%v", m.MgnMode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金模式: %v", str, m.MgnMode)
	}
	if s := fmt.Sprintf("%v", m.Side); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单方向: %v", str, m.Side)
	}
	str = fmt.Sprintf("%s\r\n└----------- Loan ------------┘", str)
	return str
}
func (m *Fee) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Fee ------┐", str)
	if s := fmt.Sprintf("%v", m.Level); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n当前在平台上真实交易量的用户等级: %v", str, m.Level)
	}
	if s := fmt.Sprintf("%v", m.Taker); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n交易区的吃单手续费率: %v", str, m.Taker)
	}
	if s := fmt.Sprintf("%v", m.Maker); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n交易区挂单手续费率: %v", str, m.Maker)
	}
	if s := fmt.Sprintf("%v", m.Delivery); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n交割手续费率: %v", str, m.Delivery)
	}
	if s := fmt.Sprintf("%v", m.Exercise); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n行权手续费率: %v", str, m.Exercise)
	}
	if s := fmt.Sprintf("%v", m.Category); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单种类: %v", str, m.Category)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- Fee ------------┘", str)
	return str
}
func (m *InterestAccrued) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ InterestAccrued ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Interest); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n计息: %v", str, m.Interest)
	}
	if s := fmt.Sprintf("%v", m.InterestRate); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n计息利率(小时): %v", str, m.InterestRate)
	}
	if s := fmt.Sprintf("%v", m.Liab); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种负债额: %v", str, m.Liab)
	}
	if s := fmt.Sprintf("%v", m.MgnMode); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金模式: %v", str, m.MgnMode)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- InterestAccrued ------------┘", str)
	return str
}
func (m *InterestRate) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ InterestRate ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.InterestRate); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n计息利率(小时): %v", str, m.InterestRate)
	}
	str = fmt.Sprintf("%s\r\n└----------- InterestRate ------------┘", str)
	return str
}
func (m *Greek) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Greek ------┐", str)
	if s := fmt.Sprintf("%v", m.GreeksType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n当前希腊字母展示方式: %v", str, m.GreeksType)
	}
	str = fmt.Sprintf("%s\r\n└----------- Greek ------------┘", str)
	return str
}
func (m *MaxWithdrawal) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ MaxWithdrawal ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.MaxWd); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种单笔最大提币量: %v", str, m.MaxWd)
	}
	str = fmt.Sprintf("%s\r\n└----------- MaxWithdrawal ------------┘", str)
	return str
}
