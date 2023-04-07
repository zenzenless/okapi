package private

import (
	models "github.com/james-zhang-bing/okapi/models/account"
	pr "github.com/james-zhang-bing/okapi/requests/ws/private"
)

type positionChannel struct {
	PrivateChannel[*pr.Position, models.Position]
}

func (c *Client) SubscribePositionChannel(arg *pr.Position) (ch <-chan *models.Position, err error) {
	channel := "positions"
	a := positionChannel{}
	a.Init(channel, arg)
	err = c.Subscribe(&a)
	ch = a.ReceivedCh
	return
}

type accountChannel struct {
	PrivateChannel[*pr.Account, models.Balance]
}

func (c *Client) SubscribeAccountChannel(arg *pr.Account) (ch <-chan *models.Balance, err error) {
	channel := "account"
	a := accountChannel{}
	a.Init(channel, arg)
	err = c.Subscribe(&a)
	ch = a.ReceivedCh
	return
}

type balanceAndPosition struct {
	PrivateChannel[*pr.Account, models.BalanceAndPosition]
}

func (c *Client) SubscribeBalanceAndPosition() (ch <-chan *models.BalanceAndPosition, err error) {
	a := balanceAndPosition{}
	a.Init("balance_and_position", nil)
	err = c.Subscribe(&a)
	ch = a.ReceivedCh
	return
}
