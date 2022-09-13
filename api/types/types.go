package types

import "encoding/json"

type BalanceMsg struct {
	Code string  `json:"code"`
	Data []Datum `json:"data"`
	Msg  string  `json:"msg"`
}

func (b *BalanceMsg) Unmarshal(data []byte) error {
	return json.Unmarshal(data, b)

}

type Datum struct {

	// adjEq	String	美金层面有效保证金
	AdjEq string `json:"adjEq"`
	// details	Array	各币种资产详细信息
	Details []Detail `json:"details"`
	// imr	String	美金层面占用保证金
	Imr string `json:"imr"`
	// isoEq	String	美金层面逐仓仓位权益
	ISOEq string `json:"isoEq"`
	// mgnRatio	String	美金层面保证金率
	MgnRatio string `json:"mgnRatio"`
	// mmr	String	美金层面维持保证金
	Mmr string `json:"mmr"`
	// notionalUsd	String	以美金价值为单位的持仓数量，即仓位美金价值
	NotionalUsd string `json:"notionalUsd"`
	// ordFroz	String	美金层面全仓挂单占用保证金
	OrdFroz string `json:"ordFroz"`
	// totalEq	String	美金层面权益
	TotalEq string `json:"totalEq"`
	// uTime	String	账户信息的更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	UTime string `json:"uTime"`
}

type Detail struct {

	// > availBal	String	可用余额
	AvailBAL string `json:"availBal"`
	// > availEq	String	可用保证金
	AvailEq string `json:"availEq"`
	// > cashBal	String	币种余额
	CashBAL string `json:"cashBal"`
	// > ccy	String	币种
	Ccy string `json:"ccy"`
	// > crossLiab	String	币种全仓负债额
	CrossLiab string `json:"crossLiab"`
	// > disEq	String	美金层面币种折算权益
	DisEq string `json:"disEq"`
	// > eq	String	币种总权益
	Eq string `json:"eq"`
	// > eqUsd	String	币种权益美金价值
	EqUsd string `json:"eqUsd"`
	// > frozenBal	String	币种占用金额
	FrozenBAL string `json:"frozenBal"`
	// > interest	String	计息
	Interest string `json:"interest"`
	// isoEq	String	美金层面逐仓仓位权益
	ISOEq string `json:"isoEq"`
	// > isoLiab	String	币种逐仓负债额
	ISOLiab string `json:"isoLiab"`
	// > isoUpl	String	逐仓未实现盈亏
	ISOUpl string `json:"isoUpl"`
	// > liab	String	币种负债额
	Liab string `json:"liab"`
	// > maxLoan	String	币种最大可借
	MaxLoan string `json:"maxLoan"`
	// mgnRatio	String	美金层面保证金率
	MgnRatio string `json:"mgnRatio"`
	// > notionalLever	String	币种杠杆倍数
	NotionalLever string `json:"notionalLever"`
	// > ordFrozen	String	挂单冻结数量
	OrdFrozen string `json:"ordFrozen"`
	// > twap	String	当前负债币种触发系统自动换币的风险
	Twap string `json:"twap"`
	// uTime	String	账户信息的更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	UTime string `json:"uTime"`
	// > upl	String	未实现盈亏
	Upl string `json:"upl"`
	// > uplLiab	String	由于仓位未实现亏损导致的负债
	UplLiab string `json:"uplLiab"`
	// > stgyEq	String	策略权益
	StgyEq string `json:"stgyEq"`
	// > spotInUseAmt	String	现货对冲占用数量
	SpotInUseAmt string `json:"spotInUseAmt"`
}
type Balance struct {
	AdjEq       float64  `json:"adjEq"`
	Details     []Detail `json:"details"`
	Imr         float64  `json:"imr"`
	ISOEq       float64  `json:"isoEq"`
	MgnRatio    float64  `json:"mgnRatio"`
	Mmr         float64  `json:"mmr"`
	NotionalUsd float64  `json:"notionalUsd"`
	OrdFroz     float64  `json:"ordFroz"`
	TotalEq     float64  `json:"totalEq"`
	UTime       float64  `json:"uTime"`
}

type PositionMsg struct {
	Code string             `json:"code"`
	Msg  string             `json:"msg"`
	Data []PositionMsgDatum `json:"data"`
}

func (b *PositionMsg) Unmarshal(data []byte) error {
	return json.Unmarshal(data, b)

}

type PositionMsgDatum struct {

	//adl	String	信号区
	Adl string `json:"adl"`
	//availPos	String	可平仓数量，适用于 币币杠杆,交割/永续（开平仓模式），期权（交易账户及保证金账户逐仓）。
	AvailPos string `json:"availPos"`
	//avgPx	String	开仓平均价
	AvgPx string `json:"avgPx"`
	//cTime	String	持仓创建时间，Unix时间戳的毫秒数格式，如 1597026383085
	CTime string `json:"cTime"`
	//ccy	String	占用保证金的币种
	Ccy string `json:"ccy"`
	//deltaBS	String	美金本位持仓仓位delta，仅适用于期权
	DeltaBS string `json:"deltaBS"`
	//deltaPA	String	币本位持仓仓位delta，仅适用于期权
	DeltaPA string `json:"deltaPA"`
	//gammaBS	String	美金本位持仓仓位gamma，仅适用于期权
	GammaBS string `json:"gammaBS"`
	//gammaPA	String	币本位持仓仓位gamma，仅适用于期权
	GammaPA string `json:"gammaPA"`
	//imr	String	初始保证金，仅适用于全仓
	Imr string `json:"imr"`
	//instId	String	产品ID，如 BTC-USD-180216
	InstID string `json:"instId"`
	//instType	String	产品类型
	InstType string `json:"instType"`
	//interest	String	利息，已经生成的未扣利息
	Interest string `json:"interest"`
	//last	String	最新成交价
	Last string `json:"last"`
	//usdPx	String	美金价格
	UsdPx string `json:"usdPx"`
	//lever	String	杠杆倍数，不适用于期权
	Lever string `json:"lever"`
	//liab	String	负债额，仅适用于币币杠杆
	Liab string `json:"liab"`
	//liabCcy	String	负债币种，仅适用于币币杠杆
	LiabCcy string `json:"liabCcy"`
	//liqPx	String	预估强平价
	LiqPx string `json:"liqPx"`
	//markPx	String	标记价格
	MarkPx string `json:"markPx"`
	//margin	String	保证金余额，可增减，仅适用于逐仓
	Margin string `json:"margin"`
	//mgnMode	String	保证金模式
	MgnMode string `json:"mgnMode"`
	//mgnRatio	String	保证金率
	MgnRatio string `json:"mgnRatio"`
	//mmr	String	维持保证金
	Mmr string `json:"mmr"`
	//notionalUsd	String	以美金价值为单位的持仓数量
	NotionalUsd string `json:"notionalUsd"`
	//optVal	String	期权市值，仅适用于期权
	OptVal string `json:"optVal"`

	PTime string `json:"pTime"`
	//posId	String	持仓ID
	Pos string `json:"pos"`
	//net：单向持仓（交割/永续/期权：pos为正代表多头，pos为负代表空头。币币杠杆：posCcy为交易货币时，代表多头；posCcy为计价货币时，代表空头。）
	PosCcy string `json:"posCcy"`
	//posId	String	持仓ID
	PosID string `json:"posId"`
	//posSide	String	持仓方向
	PosSide string `json:"posSide"`
	//spotInUseAmt	String	现货对冲占用数量
	SpotInUseAmt string `json:"spotInUseAmt"`
	//spotInUseCcy	String	现货对冲占用币种，如 BTC
	SpotInUseCcy string `json:"spotInUseCcy"`
	//thetaBS	String	美金本位持仓仓位theta，仅适用于期权
	ThetaBS string `json:"thetaBS"`
	//thetaPA	String	币本位持仓仓位theta，仅适用于期权
	ThetaPA string `json:"thetaPA"`
	//tradeId	String	最新成交ID
	TradeID string `json:"tradeId"`
	//quoteBal	String	计价币余额 ，适用于 币币杠杆（逐仓自主划转模式）
	QuoteBAL string `json:"quoteBal"`
	//baseBal	String	交易币余额，适用于 币币杠杆（逐仓自主划转模式）
	BaseBAL string `json:"baseBal"`
	//uTime	String	最近一次持仓更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	UTime string `json:"uTime"`
	//upl	String	未实现收益
	Upl string `json:"upl"`
	//uplRatio	String	未实现收益率
	UplRatio string `json:"uplRatio"`
	//vegaBS	String	美金本位持仓仓位vega，仅适用于期权
	VegaBS string `json:"vegaBS"`
	//vegaPA	String	币本位持仓仓位vega，仅适用于期权
	VegaPA string `json:"vegaPA"`
}

