package rest

import (
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/james-zhang-bing/okapi"
	"github.com/james-zhang-bing/okapi/api"
)

var InstType = struct {
	SPOT    string
	FUTURES string
	MARGIN  string
	SWAP    string
	OPTION  string
}{
	SPOT:    "SPOT",
	FUTURES: "FUTURES",
	MARGIN:  "MARGIN",
	SWAP:    "SWAP",
	OPTION:  "OPTION",
}

func (c *ClientRest) GetOrdersOKX(instType string) ([]api.OkxOrder, error) {
	req, err := http.NewRequest("GET", c.BaseURL+okapi.OkxApiURL.OrdersHistoryArchive, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("instType", instType)
	req.URL.RawQuery = q.Encode()
	log.Debug(req.URL.String())
	c.AddSignature(req, "")

	body, err := c.v0Do(req)
	if err != nil {
		return nil, err
	}
	var order api.OkxOrderRes
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}
	return order.Data, nil
}

func (c *ClientRest) GetFillsOKX(instType string) ([]api.Fill, error) {
	req, err := http.NewRequest("GET", c.BaseURL+okapi.OkxApiURL.FillsHistory, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("instType", instType)
	req.URL.RawQuery = q.Encode()
	log.Debug(req.URL.String())
	c.AddSignature(req, "")

	body, err := c.v0Do(req)
	if err != nil {
		return nil, err
	}
	var fills api.FillsRes
	err = json.Unmarshal(body, &fills)
	if err != nil {
		return nil, err
	}
	return fills.Data, nil
}

func toNumber(str string) float64 {
	if str == "" {
		return 0
	}
	n, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return n
}
