package types

import "testing"

func TestGenCommentForPositionMsgDatum(t *testing.T) {
	com := `//instType	String	产品类型
	//mgnMode	String	保证金模式
	//cross：全仓
	//isolated：逐仓
	//posId	String	持仓ID
	//posSide	String	持仓方向
	//long：双向持仓多头
	//short：双向持仓空头
	//net：单向持仓（交割/永续/期权：pos为正代表多头，pos为负代表空头。币币杠杆：posCcy为交易货币时，代表多头；posCcy为计价货币时，代表空头。）
	//pos	String	持仓数量，逐仓自主划转模式下，转入保证金后会产生pos为0的仓位
	//baseBal	String	交易币余额，适用于 币币杠杆（逐仓自主划转模式）
	//quoteBal	String	计价币余额 ，适用于 币币杠杆（逐仓自主划转模式）
	//posCcy	String	仓位资产币种，仅适用于币币杠杆仓位
	//availPos	String	可平仓数量，适用于 币币杠杆,交割/永续（开平仓模式），期权（交易账户及保证金账户逐仓）。
	//avgPx	String	开仓平均价
	//upl	String	未实现收益
	//uplRatio	String	未实现收益率
	//instId	String	产品ID，如 BTC-USD-180216
	//lever	String	杠杆倍数，不适用于期权
	//liqPx	String	预估强平价
	//不适用于期权
	//markPx	String	标记价格
	//imr	String	初始保证金，仅适用于全仓
	//margin	String	保证金余额，可增减，仅适用于逐仓
	//mgnRatio	String	保证金率
	//mmr	String	维持保证金
	//liab	String	负债额，仅适用于币币杠杆
	//liabCcy	String	负债币种，仅适用于币币杠杆
	//interest	String	利息，已经生成的未扣利息
	//tradeId	String	最新成交ID
	//optVal	String	期权市值，仅适用于期权
	//notionalUsd	String	以美金价值为单位的持仓数量
	//adl	String	信号区
	//分为5档，从1到5，数字越小代表adl强度越弱
	//ccy	String	占用保证金的币种
	//last	String	最新成交价
	//usdPx	String	美金价格
	//deltaBS	String	美金本位持仓仓位delta，仅适用于期权
	//deltaPA	String	币本位持仓仓位delta，仅适用于期权
	//gammaBS	String	美金本位持仓仓位gamma，仅适用于期权
	//gammaPA	String	币本位持仓仓位gamma，仅适用于期权
	//thetaBS	String	美金本位持仓仓位theta，仅适用于期权
	//thetaPA	String	币本位持仓仓位theta，仅适用于期权
	//vegaBS	String	美金本位持仓仓位vega，仅适用于期权
	//vegaPA	String	币本位持仓仓位vega，仅适用于期权
	//cTime	String	持仓创建时间，Unix时间戳的毫秒数格式，如 1597026383085
	//uTime	String	最近一次持仓更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	//spotInUseAmt	String	现货对冲占用数量
	//适用于组合保证金模式
	//spotInUseCcy	String	现货对冲占用币种，如 BTC
	//适用于组合保证金模式`
	genTypeComment(PositionMsgDatum{}, com)
}

func TestGenCommentForDatum(t *testing.T) {
	com := `// uTime	String	账户信息的更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	// totalEq	String	美金层面权益
	// isoEq	String	美金层面逐仓仓位权益
	// 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
	// adjEq	String	美金层面有效保证金
	// 适用于跨币种保证金模式和组合保证金模式
	// ordFroz	String	美金层面全仓挂单占用保证金
	// 适用于跨币种保证金模式和组合保证金模式
	// imr	String	美金层面占用保证金
	// 适用于跨币种保证金模式和组合保证金模式
	// mmr	String	美金层面维持保证金
	// 适用于跨币种保证金模式和组合保证金模式
	// mgnRatio	String	美金层面保证金率
	// 适用于跨币种保证金模式 和组合保证金模式
	// notionalUsd	String	以美金价值为单位的持仓数量，即仓位美金价值
	// 适用于跨币种保证金模式和组合保证金模式
	// details	Array	各币种资产详细信息
	// > ccy	String	币种
	// > eq	String	币种总权益
	// > cashBal	String	币种余额
	// > uTime	String	币种余额信息的更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	// > isoEq	String	币种逐仓仓位权益
	// 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
	// > availEq	String	可用保证金
	// 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
	// > disEq	String	美金层面币种折算权益
	// > availBal	String	可用余额
	// 适用于简单交易模式
	// > frozenBal	String	币种占用金额
	// > ordFrozen	String	挂单冻结数量
	// > liab	String	币种负债额
	// 适用于跨币种保证金模式和组合保证金模式
	// > upl	String	未实现盈亏
	// 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
	// > uplLiab	String	由于仓位未实现亏损导致的负债
	// 适用于跨币种保证金模式和组合保证金模式
	// > crossLiab	String	币种全仓负债额
	// 适用于跨币种保证金模式和组合保证金模式
	// > isoLiab	String	币种逐仓负债额
	// 适用于跨币种保证金模式和组合保证金模式
	// > mgnRatio	String	保证金率
	// 适用于单币种保证金模式
	// > interest	String	计息
	// 适用于跨币种保证金模式和组合保证金模式
	// > twap	String	当前负债币种触发系统自动换币的风险
	// 0、1、2、3、4、5其中之一，数字越大代表您的负债币种触发自动换币概率越高
	// 适用于跨币种保证金模式和组合保证金模式
	// > maxLoan	String	币种最大可借
	// 适用于跨币种保证金模式和组合保证金模式
	// > eqUsd	String	币种权益美金价值
	// > notionalLever	String	币种杠杆倍数
	// 适用于单币种保证金模式
	// > stgyEq	String	策略权益
	// > isoUpl	String	逐仓未实现盈亏
	// 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
	// > spotInUseAmt	String	现货对冲占用数量
	// 适用于组合保证金模式
	`
	genTypeComment(Datum{}, com)
	genTypeComment(Detail{},com)
}
