package private

import (
	models "github.com/james-zhang-bing/okapi/models/trade"
	pr "github.com/james-zhang-bing/okapi/requests/ws/private"
)

type orderList struct {
	PrivateChannel[*pr.Order, models.Order]
}

// 获取订单信息，首次订阅不推送，只有当下单、撤单等事件触发时，推送数据
func (c *Client) SubscribeOrderListChannel(arg *pr.Order) (ch <-chan *models.Order, err error) {
	channel := "orders"
	a := orderList{}
	a.Init(channel, arg)
	err = c.Subscribe(&a)
	ch = a.ReceivedCh
	return
}

type algoOrderList struct {
	PrivateChannel[*pr.AlgoOrder, models.AlgoOrder]
}

// 只推送变化的订单 获取策略委托订单，首次订阅不推送，只有当下单、撤单等事件触发时，推送数据
func (c *Client) SubscribeAlgoOrderList(arg *pr.AlgoOrder) (ch <-chan *models.AlgoOrder, err error) {
	a := algoOrderList{}
	a.Init("orders-algo", arg)
	err = c.Subscribe(&a)
	ch = a.ReceivedCh
	return
}

type algAdvanceList struct {
	PrivateChannel[*pr.AlgoOrder, models.AlgoOrder]
}

func (c *Client) SubscribeAlgoAdvanceList(arg *pr.AlgoOrder) (ch <-chan *models.AlgoOrder, err error) {
	a := algAdvanceList{}
	a.Init("algo-advance", arg)
	err = c.Subscribe(&a)
	ch = a.ReceivedCh
	return
}
