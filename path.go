package okapi

var OkxApiURL = struct {
	OrdersHistory        string
	OrdersHistoryArchive string
	FillsHistory         string
	MarketCandles        string //最大返回1440条数据
	MarketHistoryCandles string
	AccountBalance string
	AccountPosition string
}{
	OrdersHistory:        "api/v5/trade/orders-historye",
	OrdersHistoryArchive: "api/v5/trade/orders-history-archive",
	FillsHistory:         "api/v5/trade/fills-history",
	MarketCandles:        "api/v5/market/candles",
	MarketHistoryCandles: "api/v5/market/history-candles",
	AccountBalance:       "api/v5/account/balance",
	AccountPosition:  "api/v5/account/positions",
}
