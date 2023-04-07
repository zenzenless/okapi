package publicdata

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
)

type (
	Instrument struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// baseCcy String 交易货币币种，如  BTC-USDT 中的 BTC
		BaseCcy string `json:"baseCcy,omitempty"`
		// quoteCcy String 计价货币币种，如  BTC-USDT 中的 USDT
		QuoteCcy string `json:"quoteCcy,omitempty"`
		// settleCcy String 盈亏结算和保证金币种，如  BTC  仅适用于 交割/永续/期权
		SettleCcy string `json:"settleCcy,omitempty"`
		// ctValCcy String 合约面值计价币种，仅适用于 交割/永续/期权
		CtValCcy string `json:"ctValCcy,omitempty"`
		// ctVal String 合约面值 仅支持 FUTURES/SWAP
		CtVal okapi.JSONFloat64 `json:"ctVal,omitempty"`
		// ctMult String 合约乘数，仅适用于 交割/永续/期权
		CtMult okapi.JSONFloat64 `json:"ctMult,omitempty"`
		// stk String 行权价格，仅适用于 期权
		Stk okapi.JSONFloat64 `json:"stk,omitempty"`
		// tickSz String 下单价格精度，如  0.0001
		TickSz okapi.JSONFloat64 `json:"tickSz,omitempty"`
		// lotSz String 下单数量精度，如 BTC-USDT-SWAP： 1
		LotSz okapi.JSONFloat64 `json:"lotSz,omitempty"`
		// minSz String 最小下单数量
		MinSz okapi.JSONFloat64 `json:"minSz,omitempty"`
		// lever String 杠杆倍数，0.01到125之间的数值，仅适用于  币币杠杆/交割/永续
		Lever okapi.JSONFloat64 `json:"lever"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// category String 订单种类  normal ：普通委托 twap ：TWAP自动换币adl ：ADL自动减仓   full_liquidation ：强制平仓 partial_liquidation ：强制减仓    delivery ：交割   ddh ：对冲减仓类型订单
		Category okapi.FeeCategory `json:"category,string"`
		// optType String 期权类型， C 或 P  仅适用于 期权
		OptType okapi.OptionType `json:"optType,omitempty"`
		// listTime String 上线日期  Unix时间戳的毫秒数格式，如  1597026383085
		ListTime okapi.JSONTime `json:"listTime"`
		// expTime String 否 请求有效截止时间。Unix时间戳的毫秒数格式，如  1597026383085 无效的expTime
		ExpTime okapi.JSONTime `json:"expTime,omitempty"`
		// ctType String 否 linear ： 正向合约 inverse ：反向合约 仅 交割/永续 有效
		CtType okapi.ContractType `json:"ctType,omitempty"`
		// alias String 合约日期别名 this_week ：本周  next_week ：次周quarter ：季度 next_quarter ：次季度  仅适用于 交割
		Alias okapi.AliasType `json:"alias,omitempty"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State okapi.InstrumentState `json:"state"`
	}
	DeliveryExerciseHistory struct {
		// details Array 各个账户的资产估值
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// px String 可选  委托价格，仅适用于 limit 、 post_only 、 fok 、 ioc 类型的订单
		Px okapi.JSONFloat64 `json:"px"`
		// type String 报价方类型（当前未生效，将返回 "" ）
		Type okapi.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// oi String 持仓量（按 张 折算）
		Oi okapi.JSONFloat64 `json:"oi"`
		// oiCcy String 持仓量（按 币 折算）
		OiCcy okapi.JSONFloat64 `json:"oiCcy"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	FundingRate struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// fundingRate String 资金费率
		FundingRate     okapi.JSONFloat64 `json:"fundingRate"`
		NextFundingRate okapi.JSONFloat64 `json:"NextFundingRate"`
		// fundingTime String 资金费时间 ，Unix时间戳的毫秒数格式，如  1597026383085
		FundingTime okapi.JSONTime `json:"fundingTime"`
		// nextFundingTime String 下一期资金费时间 ，Unix时间戳的毫秒数格式，如  1622851200000
		NextFundingTime okapi.JSONTime `json:"nextFundingTime"`
	}
	FundingRateHistory struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// fundingRate String 资金费率
		FundingRate okapi.JSONFloat64 `json:"fundingRate"`

		//realizedRate	String	实际资金费率
		RealizedRate okapi.JSONFloat64 `json:"realizedRate"`
		// fundingTime String 资金费时间 ，Unix时间戳的毫秒数格式，如  1597026383085
		FundingTime okapi.JSONTime `json:"fundingTime"`
	}
	LimitPrice struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// buyLmt String 最高买价
		BuyLmt okapi.JSONFloat64 `json:"buyLmt"`
		// sellLmt String 最低卖价
		SellLmt okapi.JSONFloat64 `json:"sellLmt"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// settlePx String 预估交割、行权价格
		SettlePx okapi.JSONFloat64 `json:"settlePx"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	OptionMarketData struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// uly String 否 标的指数
		Uly string `json:"uly"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// delta String 期权价格对uly价格的敏感度
		Delta okapi.JSONFloat64 `json:"delta"`
		// gamma String delta对uly价格的敏感度
		Gamma okapi.JSONFloat64 `json:"gamma"`
		// vega String 权价格对隐含波动率的敏感度
		Vega okapi.JSONFloat64 `json:"vega"`
		// theta String 期权价格对剩余期限的敏感度
		Theta okapi.JSONFloat64 `json:"theta"`
		// deltaBS String 美金本位持仓仓位delta，仅适用于 期权
		DeltaBS okapi.JSONFloat64 `json:"deltaBS"`
		// gammaBS String 美金本位持仓仓位gamma，仅适用于 期权
		GammaBS okapi.JSONFloat64 `json:"gammaBS"`
		// vegaBS String 美金本位持仓仓位vega，仅适用于 期权
		VegaBS okapi.JSONFloat64 `json:"vegaBS"`
		// thetaBS String 美金本位持仓仓位theta，仅适用于 期权
		ThetaBS okapi.JSONFloat64 `json:"thetaBS"`
		// lever String 杠杆倍数，0.01到125之间的数值，仅适用于  币币杠杆/交割/永续
		Lever okapi.JSONFloat64 `json:"lever"`
		// markVol String 标记波动率
		MarkVol okapi.JSONFloat64 `json:"markVol"`
		// bidVol String bid波动率
		BidVol okapi.JSONFloat64 `json:"bidVol"`
		// askVol String ask波动率
		AskVol okapi.JSONFloat64 `json:"askVol"`
		// realVol String 已实现波动率（目前该字段暂未启用）
		RealVol okapi.JSONFloat64 `json:"realVol"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// amt String 是 划转数量
		Amt okapi.JSONFloat64 `json:"amt"`
		// discountLv String 否 折算率等级 1 :第一档  2 :第二档  3 :第三档4 :第四档  5 :第五档
		DiscountLv okapi.JSONInt64 `json:"discountLv"`
		// discountInfo Array 币种折算率详情
		DiscountInfo []*DiscountInfo `json:"discountInfo"`
	}
	DiscountInfo struct {
		// discountRate String 折算率
		DiscountRate okapi.JSONInt64 `json:"discountRate"`
		// maxAmt String 最多可调整的保证金数量
		MaxAmt okapi.JSONInt64 `json:"maxAmt"`
		// minAmt String 最小申购量
		MinAmt okapi.JSONInt64 `json:"minAmt"`
	}
	SystemTime struct {
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	LiquidationOrder struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// totalLoss String 当前 underlying 下，当期内的总穿仓亏损 以 USDT 为单位，每天16:00（UTC+8）清零
		TotalLoss okapi.JSONFloat64 `json:"totalLoss"`
		// details Array 各个账户的资产估值
		Details []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy,omitempty"`
		// side String 是 订单方向   buy ：买，  sell ：卖
		Side okapi.OrderSide `json:"side"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		OosSide okapi.PositionSide `json:"posSide"`
		// bkPx String 破产价格，与系统爆仓账号委托成交的价格。
		BkPx okapi.JSONFloat64 `json:"bkPx"`
		// sz String 是 委托数量
		Sz okapi.JSONFloat64 `json:"sz"`
		// bkLoss String 穿仓亏损数量
		BkLoss okapi.JSONFloat64 `json:"bkLoss"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	MarkPrice struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// markPx String 标记价格
		MarkPx okapi.JSONFloat64 `json:"markPx"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	PositionTier struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
		// tier String 否 查指定档位
		Tier okapi.JSONInt64 `json:"tier"`
		// minSz String 最小下单数量
		MinSz okapi.JSONFloat64 `json:"minSz"`
		// maxSz String 最大持仓量
		MaxSz okapi.JSONFloat64 `json:"maxSz"`
		// mmr String 美金层面维持保证金 适用于 跨币种保证金模式 和 组合保证金模式
		Mmr okapi.JSONFloat64 `json:"mmr"`
		// imr String 美金层面占用保证金 适用于 跨币种保证金模式 和 组合保证金模式
		Imr okapi.JSONFloat64 `json:"imr"`
		// optMgnFactor String 期权保证金系数 （仅适用于期权）
		OptMgnFactor okapi.JSONFloat64 `json:"optMgnFactor,omitempty"`
		// quoteMaxLoan String 计价货币 最大借币量（仅适用于杠杆，且 instId 参数生效时），例如 BTC-USDT 里的 USDT最大借币量
		QuoteMaxLoan okapi.JSONFloat64 `json:"quoteMaxLoan,omitempty"`
		// baseMaxLoan String 交易货币 最大借币量（仅适用于杠杆，且 instId 参数生效时），例如 BTC-USDT 里的 BTC最大借币量
		BaseMaxLoan okapi.JSONFloat64 `json:"baseMaxLoan,omitempty"`
		// maxLever String 最高可用杠杆倍数
		MaxLever okapi.JSONFloat64 `json:"maxLever"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okapi.JSONTime `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		// basic Array 基础利率和借币限额
		Basic []*InterestRateAndLoanBasic `json:"basic"`
		// vip Array 专业用户
		Vip []*InterestRateAndLoanUser `json:"vip"`
		// regular Array 普通用户
		Regular []*InterestRateAndLoanUser `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy string `json:"ccy"`
		// rate String 最新出借利率
		Rate okapi.JSONFloat64 `json:"rate"`
		// quota String 基础借币限额
		Quota okapi.JSONFloat64 `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		// level String 当前在平台上真实交易量的用户等级，例如  lv1
		Level string `json:"level"`
		// irDiscount String 利率的折扣率
		IrDiscount okapi.JSONFloat64 `json:"irDiscount"`
		// loanQuotaCoef String 借币限额系数
		LoanQuotaCoef int `json:"loanQuotaCoef,string"`
	}
	State struct {
		// title String 系统维护说明的标题
		Title string `json:"title"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State string `json:"state"`
		// href String 系统维护详情的超级链接,若无返回值，默认值为空，如： “”
		Href string `json:"href"`
		// serviceType String 服务类型，  0 ：WebSocket ;5 ：交易服务； 99 ：其他（如：停止部分产品交易）
		ServiceType string `json:"serviceType"`
		// system String 系统， unified ：交易账户
		System string `json:"system"`
		// scheDesc String 改期进度说明，如：由 2021-01-26T16:30:00.000Z 改期到 2021-01-28T16:30:00.000Z
		ScheDesc string `json:"scheDesc"`
		// begin String 否 筛选的开始时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
		Begin okapi.JSONTime `json:"begin"`
		// end String 否 筛选的结束时间戳，Unix 时间戳为毫秒数格式，如 1597027383085
		End okapi.JSONTime `json:"end"`
	}
	ConvertContractUnit struct {
		// type	String	否	转换类型1: 币转张，当张为小数时，会进一取整2: 张转币 默认为 1
		ToUnitType okapi.UnitType `json:"type,omitempty"`
		//instId	String	是	产品ID，仅适用于交割/永续/期权
		InstId string `json:"instId"`
		//sz	String	是	数量，币转张时，为币的数量，张转币时，为张的数量。张的数量，只能是正整数
		Sz okapi.JSONFloat64 `json:"sz,string"`
		//px	String	可选	委托价格 币本位合约的张币转换时必填；U本位合约，usdt 与张的转换时，必填；coin 与张的转换时，可不填；期权的张币转换时，可不填。
		Px okapi.JSONFloat64 `json:"px,string"`
		//unit	String	否	币的单位，coin: 币，usdt: usdt 仅适用于交割和永续合约的U本位合约
		Unit string `json:"unit,omitempty"`
	}
)

func (m *Instrument) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Instrument ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Uly); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	}
	if s := fmt.Sprintf("%v", m.BaseCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n交易货币币种: %v", str, m.BaseCcy)
	}
	if s := fmt.Sprintf("%v", m.QuoteCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n计价货币币种: %v", str, m.QuoteCcy)
	}
	if s := fmt.Sprintf("%v", m.SettleCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n盈亏结算和保证金币种: %v", str, m.SettleCcy)
	}
	if s := fmt.Sprintf("%v", m.CtValCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n合约面值计价币种: %v", str, m.CtValCcy)
	}
	if s := fmt.Sprintf("%v", m.CtVal); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n合约面值: %v", str, m.CtVal)
	}
	if s := fmt.Sprintf("%v", m.CtMult); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n合约乘数: %v", str, m.CtMult)
	}
	if s := fmt.Sprintf("%v", m.Stk); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n行权价格: %v", str, m.Stk)
	}
	if s := fmt.Sprintf("%v", m.TickSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n下单价格精度: %v", str, m.TickSz)
	}
	if s := fmt.Sprintf("%v", m.LotSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n下单数量精度: %v", str, m.LotSz)
	}
	if s := fmt.Sprintf("%v", m.MinSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最小下单数量: %v", str, m.MinSz)
	}
	if s := fmt.Sprintf("%v", m.Lever); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n杠杆倍数: %v", str, m.Lever)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.Category); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单种类: %v", str, m.Category)
	}
	if s := fmt.Sprintf("%v", m.OptType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n期权类型: %v", str, m.OptType)
	}
	if s := fmt.Sprintf("%v", m.ListTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n上线日期: %v", str, m.ListTime)
	}
	if s := fmt.Sprintf("%v", m.ExpTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n请求有效截止时间。Unix时间戳的毫秒数格式: %v", str, m.ExpTime)
	}
	if s := fmt.Sprintf("%v", m.CtType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n正向合约: %v", str, m.CtType)
	}
	if s := fmt.Sprintf("%v", m.Alias); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n合约日期别名: %v", str, m.Alias)
	}
	if s := fmt.Sprintf("%v", m.State); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单状态: %v", str, m.State)
	}
	str = fmt.Sprintf("%s\r\n└----------- Instrument ------------┘", str)
	return str
}
func (m *DeliveryExerciseHistory) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ DeliveryExerciseHistory ------┐", str)
	if s := fmt.Sprintf("%v", m.Details); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n各个账户的资产估值: %v", str, m.Details)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- DeliveryExerciseHistory ------------┘", str)
	return str
}
func (m *DeliveryExerciseHistoryDetails) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ DeliveryExerciseHistoryDetails ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Px); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n委托价格: %v", str, m.Px)
	}
	if s := fmt.Sprintf("%v", m.Type); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n报价方类型（当前未生效: %v", str, m.Type)
	}
	str = fmt.Sprintf("%s\r\n└----------- DeliveryExerciseHistoryDetails ------------┘", str)
	return str
}
func (m *OpenInterest) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ OpenInterest ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Oi); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓量（按: %v", str, m.Oi)
	}
	if s := fmt.Sprintf("%v", m.OiCcy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓量（按: %v", str, m.OiCcy)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- OpenInterest ------------┘", str)
	return str
}
func (m *FundingRate) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ FundingRate ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.FundingRate); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n资金费率: %.3f%%", str, m.FundingRate*100)
	}
	if s := fmt.Sprintf("%v", m.NextFundingRate); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n: 下一期预测资金费率: %.3f%%", str, m.NextFundingRate*100)
	}
	if s := fmt.Sprintf("%v", m.FundingTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n资金费时间: %v", str, m.FundingTime)
	}
	if s := fmt.Sprintf("%v", m.NextFundingTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n下一期资金费时间: %v", str, m.NextFundingTime)
	}
	str = fmt.Sprintf("%s\r\n└----------- FundingRate ------------┘", str)
	return str
}
func (m *FundingRateHistory) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ FundingRateHistory ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.FundingRate); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n预计资金费率: %.3f%%", str, m.FundingRate*100)
	}
	if s := fmt.Sprintf("%v", m.RealizedRate); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n实际资金费率: %.3f%%", str, m.RealizedRate*100)
	}
	if s := fmt.Sprintf("%v", m.FundingTime); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n资金费时间: %v", str, m.FundingTime)
	}
	str = fmt.Sprintf("%s\r\n└----------- FundingRateHistory ------------┘", str)
	return str
}
func (m *LimitPrice) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ LimitPrice ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.BuyLmt); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最高买价: %v", str, m.BuyLmt)
	}
	if s := fmt.Sprintf("%v", m.SellLmt); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最低卖价: %v", str, m.SellLmt)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- LimitPrice ------------┘", str)
	return str
}
func (m *EstimatedDeliveryExercisePrice) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ EstimatedDeliveryExercisePrice ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.SettlePx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n预估交割、行权价格: %v", str, m.SettlePx)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- EstimatedDeliveryExercisePrice ------------┘", str)
	return str
}
func (m *OptionMarketData) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ OptionMarketData ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Uly); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.Delta); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n期权价格对uly价格的敏感度: %v", str, m.Delta)
	}
	if s := fmt.Sprintf("%v", m.Gamma); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\ndelta对uly价格的敏感度: %v", str, m.Gamma)
	}
	if s := fmt.Sprintf("%v", m.Vega); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n权价格对隐含波动率的敏感度: %v", str, m.Vega)
	}
	if s := fmt.Sprintf("%v", m.Theta); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n期权价格对剩余期限的敏感度: %v", str, m.Theta)
	}
	if s := fmt.Sprintf("%v", m.DeltaBS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金本位持仓仓位delta: %v", str, m.DeltaBS)
	}
	if s := fmt.Sprintf("%v", m.GammaBS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金本位持仓仓位gamma: %v", str, m.GammaBS)
	}
	if s := fmt.Sprintf("%v", m.VegaBS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金本位持仓仓位vega: %v", str, m.VegaBS)
	}
	if s := fmt.Sprintf("%v", m.ThetaBS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金本位持仓仓位theta: %v", str, m.ThetaBS)
	}
	if s := fmt.Sprintf("%v", m.Lever); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n杠杆倍数: %v", str, m.Lever)
	}
	if s := fmt.Sprintf("%v", m.MarkVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n标记波动率: %v", str, m.MarkVol)
	}
	if s := fmt.Sprintf("%v", m.BidVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nbid波动率: %v", str, m.BidVol)
	}
	if s := fmt.Sprintf("%v", m.AskVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\nask波动率: %v", str, m.AskVol)
	}
	if s := fmt.Sprintf("%v", m.RealVol); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n已实现波动率（目前该字段暂未启用）: %v", str, m.RealVol)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- OptionMarketData ------------┘", str)
	return str
}
func (m *GetDiscountRateAndInterestFreeQuota) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ GetDiscountRateAndInterestFreeQuota ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Amt); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n划转数量: %v", str, m.Amt)
	}
	if s := fmt.Sprintf("%v", m.DiscountLv); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n折算率等级: %v", str, m.DiscountLv)
	}
	if s := fmt.Sprintf("%v", m.DiscountInfo); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币种折算率详情: %v", str, m.DiscountInfo)
	}
	str = fmt.Sprintf("%s\r\n└----------- GetDiscountRateAndInterestFreeQuota ------------┘", str)
	return str
}
func (m *DiscountInfo) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ DiscountInfo ------┐", str)
	if s := fmt.Sprintf("%v", m.DiscountRate); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n折算率: %v", str, m.DiscountRate)
	}
	if s := fmt.Sprintf("%v", m.MaxAmt); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最多可调整的保证金数量: %v", str, m.MaxAmt)
	}
	if s := fmt.Sprintf("%v", m.MinAmt); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最小申购量: %v", str, m.MinAmt)
	}
	str = fmt.Sprintf("%s\r\n└----------- DiscountInfo ------------┘", str)
	return str
}
func (m *SystemTime) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ SystemTime ------┐", str)
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- SystemTime ------------┘", str)
	return str
}
func (m *LiquidationOrder) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ LiquidationOrder ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Uly); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.TotalLoss); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n当前: %v", str, m.TotalLoss)
	}
	if s := fmt.Sprintf("%v", m.Details); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n各个账户的资产估值: %v", str, m.Details)
	}
	str = fmt.Sprintf("%s\r\n└----------- LiquidationOrder ------------┘", str)
	return str
}
func (m *LiquidationOrderDetail) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ LiquidationOrderDetail ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Side); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单方向: %v", str, m.Side)
	}
	if s := fmt.Sprintf("%v", m.OosSide); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n持仓方向: %v", str, m.OosSide)
	}
	if s := fmt.Sprintf("%v", m.BkPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n破产价格: %v", str, m.BkPx)
	}
	if s := fmt.Sprintf("%v", m.Sz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n委托数量: %v", str, m.Sz)
	}
	if s := fmt.Sprintf("%v", m.BkLoss); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n穿仓亏损数量: %v", str, m.BkLoss)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- LiquidationOrderDetail ------------┘", str)
	return str
}
func (m *MarkPrice) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ MarkPrice ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.MarkPx); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n标记价格: %v", str, m.MarkPx)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- MarkPrice ------------┘", str)
	return str
}
func (m *PositionTier) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ PositionTier ------┐", str)
	if s := fmt.Sprintf("%v", m.InstID); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	}
	if s := fmt.Sprintf("%v", m.Uly); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	}
	if s := fmt.Sprintf("%v", m.InstType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	}
	if s := fmt.Sprintf("%v", m.Tier); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n查指定档位: %v", str, m.Tier)
	}
	if s := fmt.Sprintf("%v", m.MinSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最小下单数量: %v", str, m.MinSz)
	}
	if s := fmt.Sprintf("%v", m.MaxSz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最大持仓量: %v", str, m.MaxSz)
	}
	if s := fmt.Sprintf("%v", m.Mmr); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面维持保证金: %v", str, m.Mmr)
	}
	if s := fmt.Sprintf("%v", m.Imr); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n美金层面占用保证金: %v", str, m.Imr)
	}
	if s := fmt.Sprintf("%v", m.OptMgnFactor); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n期权保证金系数: %v", str, m.OptMgnFactor)
	}
	if s := fmt.Sprintf("%v", m.QuoteMaxLoan); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n计价货币: %v", str, m.QuoteMaxLoan)
	}
	if s := fmt.Sprintf("%v", m.BaseMaxLoan); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n交易货币: %v", str, m.BaseMaxLoan)
	}
	if s := fmt.Sprintf("%v", m.MaxLever); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最高可用杠杆倍数: %v", str, m.MaxLever)
	}
	if s := fmt.Sprintf("%v", m.TS); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n成交明细产生时间: %v", str, m.TS)
	}
	str = fmt.Sprintf("%s\r\n└----------- PositionTier ------------┘", str)
	return str
}
func (m *InterestRateAndLoanQuota) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ InterestRateAndLoanQuota ------┐", str)
	if s := fmt.Sprintf("%v", m.Basic); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n基础利率和借币限额: %v", str, m.Basic)
	}
	if s := fmt.Sprintf("%v", m.Vip); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n专业用户: %v", str, m.Vip)
	}
	if s := fmt.Sprintf("%v", m.Regular); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n普通用户: %v", str, m.Regular)
	}
	str = fmt.Sprintf("%s\r\n└----------- InterestRateAndLoanQuota ------------┘", str)
	return str
}
func (m *InterestRateAndLoanBasic) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ InterestRateAndLoanBasic ------┐", str)
	if s := fmt.Sprintf("%v", m.Ccy); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n保证金币种: %v", str, m.Ccy)
	}
	if s := fmt.Sprintf("%v", m.Rate); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n最新出借利率: %v", str, m.Rate)
	}
	if s := fmt.Sprintf("%v", m.Quota); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n基础借币限额: %v", str, m.Quota)
	}
	str = fmt.Sprintf("%s\r\n└----------- InterestRateAndLoanBasic ------------┘", str)
	return str
}
func (m *InterestRateAndLoanUser) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ InterestRateAndLoanUser ------┐", str)
	if s := fmt.Sprintf("%v", m.Level); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n当前在平台上真实交易量的用户等级: %v", str, m.Level)
	}
	if s := fmt.Sprintf("%v", m.IrDiscount); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n利率的折扣率: %v", str, m.IrDiscount)
	}
	if s := fmt.Sprintf("%v", m.LoanQuotaCoef); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n借币限额系数: %v", str, m.LoanQuotaCoef)
	}
	str = fmt.Sprintf("%s\r\n└----------- InterestRateAndLoanUser ------------┘", str)
	return str
}
func (m *State) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ State ------┐", str)
	if s := fmt.Sprintf("%v", m.Title); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n系统维护说明的标题: %v", str, m.Title)
	}
	if s := fmt.Sprintf("%v", m.State); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n订单状态: %v", str, m.State)
	}
	if s := fmt.Sprintf("%v", m.Href); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n系统维护详情的超级链接,若无返回值: %v", str, m.Href)
	}
	if s := fmt.Sprintf("%v", m.ServiceType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n服务类型: %v", str, m.ServiceType)
	}
	if s := fmt.Sprintf("%v", m.System); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n系统: %v", str, m.System)
	}
	if s := fmt.Sprintf("%v", m.ScheDesc); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n改期进度说明: %v", str, m.ScheDesc)
	}
	if s := fmt.Sprintf("%v", m.Begin); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n筛选的开始时间戳: %v", str, m.Begin)
	}
	if s := fmt.Sprintf("%v", m.End); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n筛选的结束时间戳: %v", str, m.End)
	}
	str = fmt.Sprintf("%s\r\n└----------- State ------------┘", str)
	return str
}
func (m *ConvertContractUnit) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ ConvertContractUnit ------┐", str)
	if s := fmt.Sprintf("%v", m.ToUnitType); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n转换类型1:: %v", str, m.ToUnitType)
	}
	if s := fmt.Sprintf("%v", m.InstId); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstId)
	}
	if s := fmt.Sprintf("%v", m.Sz); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n数量: %v", str, m.Sz)
	}
	if s := fmt.Sprintf("%v", m.Px); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n委托价格: %v", str, m.Px)
	}
	if s := fmt.Sprintf("%v", m.Unit); s != "" && s != "0" {
		str = fmt.Sprintf("%s\r\n币的单位: %v", str, m.Unit)
	}
	str = fmt.Sprintf("%s\r\n└----------- ConvertContractUnit ------------┘", str)
	return str
}
