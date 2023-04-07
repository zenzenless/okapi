package private

import (
	"fmt"

	"github.com/james-zhang-bing/okapi"
	"github.com/james-zhang-bing/okapi/api/rest"
	"github.com/james-zhang-bing/okapi/api/websocket/wsclient"
)

// 实盘API交易地址如下：

// REST：https://www.okx.com/
// WebSocket公共频道：wss://ws.okx.com:8443/ws/v5/public
// WebSocket私有频道：wss://ws.okx.com:8443/ws/v5/private
// AWS 地址如下：

// REST：https://aws.okx.com
// WebSocket公共频道：wss://wsaws.okx.com:8443/ws/v5/public
// WebSocket私有频道：wss://wsaws.okx.com:8443/ws/v5/private

type Client struct {
	*wsclient.Client
}

func NewPrivateClient(addr okapi.BaseURL, apiKeys ...*rest.APIKey) (*Client, error) {
	if len(apiKeys) == 0 {
		return nil, fmt.Errorf("should specified api key")
	}
	if !(addr == okapi.AwsPrivateWsURL || addr == okapi.PrivateWsURL || addr == okapi.DemoPrivateWsURL) {
		return nil, fmt.Errorf("this is private client ,should use private url, you specific url :%s", addr)
	}
	c, err := wsclient.NewClient(addr, privateArgToID, apiKeys...)
	return &Client{
		Client: c,
	}, err
}
