package rest

import (
	"fmt"
	"mintool/cointool/API/types"
	"net/http"
	"strings"
)

func (c *Client) GetAccountBalance(coinName ...string) (*types.BalanceMsg, error) {
	req, err := http.NewRequest("GET", c.BaseURL+OkxAPI.accountBalance, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	coinList := strings.Join(coinName, ",")
	q.Add("ccy", coinList)
	req.URL.RawQuery = q.Encode()
	log.Debug(req.URL.String())
	c.AddSignature(req, "")

	body, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	var balance types.BalanceMsg
	err = balance.Unmarshal(body)
	if err != nil {
		return nil, fmt.Errorf("unmarshal balance fail %s", err)
	}
	return &balance, nil
}

func (c *Client) GetAccountPositions(instType, instId, posID string) (*types.PositionMsg, error) {
	req, err := http.NewRequest("GET", c.BaseURL+OkxAPI.accountPosition, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()

	//if instType != "" {
		q.Add("instType", instType)
	//}
	//if instId != "" {
		q.Add("instId", instId)
	//}
	//if posID != "" {
		q.Add("posID", posID)
	//}

	req.URL.RawQuery = q.Encode()
	log.Debug(req.URL.String())
	c.AddSignature(req, "")

	body, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	var position types.PositionMsg
	err = position.Unmarshal(body)
	if err != nil {
		return nil, fmt.Errorf("unmarshal position fail %s", err)
	}
	return &position, nil
}
