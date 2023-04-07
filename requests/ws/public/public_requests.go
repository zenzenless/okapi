package public

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
)

type (
	Instruments struct {
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType"`
	}
	Tickers struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
	}
	OpenInterest struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
	}
	Candlesticks struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// channel String 是 频道名
		Channel okapi.CandleStickWsBarSize `json:"channel"`
	}
	Trades struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
	}
	EstimatedDeliveryExercisePrice struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// uly String 否 标的指数
		Uly string `json:"uly,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okapi.InstrumentType `json:"instType,omitempty"`
	}
	MarkPrice struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
	}
	MarkPriceCandlesticks struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// channel String 是 频道名
		Channel okapi.CandleStickWsBarSize `json:"channel"`
	}
	PriceLimit struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
	}
	OrderBook struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// channel String 是 频道名
		Channel string `json:"channel"`
	}
	OPTIONSummary struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// uly String 否 标的指数
		Uly string `json:"uly"`
	}
	FundingRate struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
	}
	IndexCandlesticks struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
		// channel String 是 频道名
		Channel string `json:"channel"`
	}
	IndexTickers struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string `json:"instId"`
	}
)

func (m *Instruments) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Instruments ------┐", str)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- Instruments ------------┘", str)
	return str
}
func (m *Tickers) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Tickers ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n└----------- Tickers ------------┘", str)
	return str
}
func (m *OpenInterest) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ OpenInterest ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n└----------- OpenInterest ------------┘", str)
	return str
}
func (m *Candlesticks) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Candlesticks ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n频道名: %v", str, m.Channel)
	str = fmt.Sprintf("%s\r\n└----------- Candlesticks ------------┘", str)
	return str
}
func (m *Trades) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ Trades ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n└----------- Trades ------------┘", str)
	return str
}
func (m *EstimatedDeliveryExercisePrice) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ EstimatedDeliveryExercisePrice ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n产品类型: %v", str, m.InstType)
	str = fmt.Sprintf("%s\r\n└----------- EstimatedDeliveryExercisePrice ------------┘", str)
	return str
}
func (m *MarkPrice) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ MarkPrice ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n└----------- MarkPrice ------------┘", str)
	return str
}
func (m *MarkPriceCandlesticks) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ MarkPriceCandlesticks ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n频道名: %v", str, m.Channel)
	str = fmt.Sprintf("%s\r\n└----------- MarkPriceCandlesticks ------------┘", str)
	return str
}
func (m *PriceLimit) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ PriceLimit ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n└----------- PriceLimit ------------┘", str)
	return str
}
func (m *OrderBook) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ OrderBook ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n频道名: %v", str, m.Channel)
	str = fmt.Sprintf("%s\r\n└----------- OrderBook ------------┘", str)
	return str
}
func (m *OPTIONSummary) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ OPTIONSummary ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n标的指数: %v", str, m.Uly)
	str = fmt.Sprintf("%s\r\n└----------- OPTIONSummary ------------┘", str)
	return str
}
func (m *FundingRate) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ FundingRate ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n└----------- FundingRate ------------┘", str)
	return str
}
func (m *IndexCandlesticks) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ IndexCandlesticks ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n频道名: %v", str, m.Channel)
	str = fmt.Sprintf("%s\r\n└----------- IndexCandlesticks ------------┘", str)
	return str
}
func (m *IndexTickers) String() string {
	var str string
	str = fmt.Sprintf("\r\n%s┌------ IndexTickers ------┐", str)
	str = fmt.Sprintf("%s\r\n产品ID: %v", str, m.InstID)
	str = fmt.Sprintf("%s\r\n└----------- IndexTickers ------------┘", str)
	return str
}
