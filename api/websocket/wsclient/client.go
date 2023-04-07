package wsclient

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
	"sync"
	"time"

	logs "github.com/ipfs/go-log/v2"

	"github.com/james-zhang-bing/okapi"
	"github.com/james-zhang-bing/okapi/api/rest"

	"github.com/gorilla/websocket"

	"github.com/james-zhang-bing/okapi/api/websocket/types"
)

var log = logs.NewLogger("websocket")

// 实盘API交易地址如下：

// REST：https://www.okx.com/
// WebSocket公共频道：wss://ws.okx.com:8443/ws/v5/public
// WebSocket私有频道：wss://ws.okx.com:8443/ws/v5/private
// AWS 地址如下：

// REST：https://aws.okx.com
// WebSocket公共频道：wss://wsaws.okx.com:8443/ws/v5/public
// WebSocket私有频道：wss://wsaws.okx.com:8443/ws/v5/private

type Client struct {
	Addr            okapi.BaseURL
	conn            *websocket.Conn
	Channels        map[string]ChannelSub
	cLk             sync.Mutex
	incomingErr     error
	incoming        chan io.Reader
	stop            chan struct{}
	pong            chan struct{}
	req             chan []byte
	sendLk          sync.Mutex
	receivedOpResCh chan *types.Response
	apikey          *rest.APIKey
	ArgToID         func(map[string]string) string
}

type param struct {
	key   string
	value string
}

func argToID(arg map[string]string) (id string) {
	var params []*param
	for i, v := range arg {
		params = append(params, &param{
			key:   i,
			value: v,
		})
	}
	sort.Slice(params, func(i, j int) bool {
		return params[i].key < params[j].key
	})
	var str []string
	for _, v := range params {
		str = append(str, v.value)
	}
	id = strings.Join(str, "_")
	return id
}

func NewClient(addr okapi.BaseURL, argFunc func(map[string]string) string, apiKeys ...*rest.APIKey) (*Client, error) {
	var apiKey *rest.APIKey
	if addr == okapi.AwsPrivateWsURL || addr == okapi.PrivateWsURL || addr == okapi.DemoPrivateWsURL {
		if len(apiKeys) == 0 {
			return nil, fmt.Errorf("should specified api key")
		}
		apiKey = apiKeys[0]
	}
	c := Client{
		Addr:            addr,
		Channels:        make(map[string]ChannelSub),
		stop:            make(chan struct{}),
		pong:            make(chan struct{}, 1),
		req:             make(chan []byte, 1),
		receivedOpResCh: make(chan *types.Response),
		ArgToID:         argFunc,
		apikey:          apiKey,
	}

	err := c.Connect()
	if err != nil {
		return nil, err
	}
	go c.handleWsConn(context.TODO())
	if len(apiKeys) != 0 {
		err := c.login()
		if err != nil {
			return nil, err
		}
	}
	return &c, nil
}

func (c *Client) Close() {
	close(c.stop)

}

func (c *Client) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(string(c.Addr), nil)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) reconnect() error {
	c.conn.Close()
	return c.Connect()
}

func (c *Client) nextMsg() {
	//如不读取出来，会越来越多
	msgType, r, err := c.conn.NextReader()
	if err != nil {
		c.incomingErr = err
		close(c.incoming)
		return
	}
	if msgType != websocket.BinaryMessage && msgType != websocket.TextMessage {
		c.incomingErr = fmt.Errorf("unsupported message type")
		close(c.incoming)
		return
	}
	c.incoming <- r
}

// func (c *Client) handleEvent(e *types.Event) error {
// 	if e.Event == "error" {
// 		//log.Errorf("event :%+v",*e)
// 		return fmt.Errorf("event Error:%+v", *e)
// 	}
// 	if e.Event == "subscribe" {
// 		c.receivedOpResCh <- e
// 		return nil
// 	}
// 	log.Info("event :", *e)
// 	return nil
// }

func (c *Client) handleReader(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("io error %s", err)
	}
	// ping pong
	if string(b) == "pong" {
		log.Debug("received pong")
		c.pong <- struct{}{}
		return nil
	}

	log.Debug("received msg:", string(b))
	res, err := types.UnmarshalResponse(b)
	if err != nil {
		return fmt.Errorf("try to unmarshal response error: %s", err)
	}
	if types.IsOpRes(res) {
		return c.HandleOpRes(res)
	}

	c.cLk.Lock()
	defer c.cLk.Unlock()
	chID := c.ArgToID(res.Arg)
	if ch, ok := c.Channels[chID]; ok {
		go ch.ReceiveData(res.Data)
		return nil
	}
	return fmt.Errorf("can't not find the channel of id %s arg %+v", chID, res.Arg)
}

func (c *Client) send(msg []byte) error {
	c.sendLk.Lock()
	defer c.sendLk.Unlock()
	return c.conn.WriteMessage(websocket.TextMessage, msg)
}

func (c *Client) handleWsConn(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			err := c.handling(ctx)
			if err != nil {
				log.Errorf("handleWsConn err %s", err)
			}
			err = c.reconnect()
			if err != nil {
				log.Errorf("reconnect err %s", err)
				time.Sleep(10 * time.Second)
				continue
			}
			// re subscribe
			go func() {
				if c.apikey != nil {
					err = c.login()
					if err != nil {
						log.Errorf("re login failed: %v", err)
						time.Sleep(30 * time.Second)
						return
					}
				}
				subs := make([]ChannelSub, 0)
				c.cLk.Lock()
				for _, sub := range c.Channels {
					subs = append(subs, sub)
				}
				c.cLk.Unlock()
				for _, sub := range subs {
					err = c.Subscribe(sub)
					if err != nil {
						log.Errorf("re subscribe failed: %v", err)
					}
				}
			}()
		}
	}
}

func (c *Client) handling(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	timeoutTimer := time.NewTimer(60 * time.Second)
	defer timeoutTimer.Stop()
	c.incoming = make(chan io.Reader)
	c.incomingErr = nil
	go func() {
		for {
			c.send([]byte("ping"))
			select {
			case <-c.stop:
				return
			case <-ctx.Done():
				return
			default:
				time.Sleep(15 * time.Second)
			}
		}
	}()
	go c.nextMsg()
	for {
		timeoutTimer.Reset(30 * time.Second)
		select {
		case <-timeoutTimer.C:
			//log.Warn("ping time out, try to reconnect")
			return fmt.Errorf("ping time out, try to reconnect")
		case reader, ok := <-c.incoming:
			err := c.incomingErr
			if err != nil {
				return err
			}
			if ok {
				//如果积累了很多，怎么处理
				err = c.handleReader(reader)
				if err != nil {

					log.Error(err)
				}
				go c.nextMsg()
				continue
			}
			return fmt.Errorf("c.incoming !ok")
		case <-c.pong:
		case <-c.stop:
			cmsg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
			if err := c.conn.WriteMessage(websocket.CloseMessage, cmsg); err != nil {
				log.Warn("failed to write close message: ", err)
			}
			if err := c.conn.Close(); err != nil {
				log.Warnw("websocket close error", "error", err)
			}
			return nil
		case <-ctx.Done():
			return fmt.Errorf("websocket have close: %s ", ctx.Err())
		}
	}
}

type ChannelSub interface {
	GetArg() map[string]string
	ReceiveData([]byte)
}

func (c *Client) Subscribe(cs types.ChannelSub) error {
	arg := cs.GetArg()
	sub := &types.Op{
		Op: "subscribe",
		Args: []map[string]string{
			arg,
		},
	}
	j, err := json.Marshal(sub)
	if err != nil {
		return err
	}

	id := c.ArgToID(arg)
	c.cLk.Lock()
	c.Channels[id] = cs
	c.cLk.Unlock()
	s, err := c.sendAndWait(context.TODO(), j)
	if err != nil {
		delete(c.Channels, id)
		return fmt.Errorf("send msg err %s", err)
	}
	if s != nil && s.Code != 0 {
		delete(c.Channels, id)
		return fmt.Errorf("subscribe failed res:%s error: %s", s, err)
	}
	subid := c.ArgToID(s.Arg)
	if subid != id {
		delete(c.Channels, id)
		return fmt.Errorf("subscribe error: the id not match %s != %s", subid, id)
	}

	return nil
}

func (c *Client) sendAndWait(ctx context.Context, msg []byte) (*types.Response, error) {
	c.sendLk.Lock()
	defer c.sendLk.Unlock()
	err := c.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	select {
	case res := <-c.receivedOpResCh:
		return res, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (c *Client) sign() (string, string) {
	method := "GET"
	path := "/users/self/verify"

	t := time.Now().Unix()
	ts := fmt.Sprint(t)
	s := ts + method + path
	p := []byte(s)
	h := hmac.New(sha256.New, c.apikey.SecretKey)
	h.Write(p)
	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *Client) login() error {

	ts, signature := c.sign()
	arg := make(map[string]string)
	arg["apiKey"] = c.apikey.ApiKey
	arg["passphrase"] = c.apikey.Passphrase
	arg["timestamp"] = ts
	arg["sign"] = signature
	logOp := &types.Op{
		Op: "login",
		Args: []map[string]string{
			arg,
		},
	}
	req, err := json.Marshal(logOp)
	if err != nil {
		return fmt.Errorf("login failed Marshal json failed %s", err)
	}
	res, err := c.sendAndWait(context.TODO(), req)
	if err != nil {
		return fmt.Errorf("send login request failed %s", req)
	}
	if res != nil && res.Code != 0 {
		return fmt.Errorf("login failed res:%s error: %s", res, err)
	}
	log.Info("login success")
	return nil
}

func (c *Client) HandleOpRes(res *types.Response) error {
	select {
	case c.receivedOpResCh <- res:
		return nil
	default:
		return fmt.Errorf("maybe the response is too late, no chan can push")
	}

}
