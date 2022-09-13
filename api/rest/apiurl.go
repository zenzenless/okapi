package rest

var OkxAPI = struct {
	ordersHistory        string
	ordersHistoryArchive string
	fillsHistory         string
	marketCandles        string //最大返回1440条数据
	marketHistoryCandles string
	accountBalance       string
	accountPosition      string
}{
	ordersHistory:        "api/v5/trade/orders-historye",
	ordersHistoryArchive: "api/v5/trade/orders-history-archive",
	fillsHistory:         "api/v5/trade/fills-history",
	marketCandles:        "api/v5/market/candles",
	marketHistoryCandles: "api/v5/market/history-candles",
	accountBalance:       "api/v5/account/balance",
	accountPosition:  "api/v5/account/positions",
}
