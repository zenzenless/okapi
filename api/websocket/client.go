package websocket

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("websocket")

func init() {
	logging.SetAllLoggers(logging.LevelDebug)
}

// 实盘API交易地址如下：

// REST：https://www.okx.com/
// WebSocket公共频道：wss://ws.okx.com:8443/ws/v5/public
// WebSocket私有频道：wss://ws.okx.com:8443/ws/v5/private
// AWS 地址如下：

// REST：https://aws.okx.com
// WebSocket公共频道：wss://wsaws.okx.com:8443/ws/v5/public
// WebSocket私有频道：wss://wsaws.okx.com:8443/ws/v5/private

type Client struct {
	Addr        string
	conn        *websocket.Conn
	Channels    map[string]ChannelSub
	cLk         sync.Mutex
	incomingErr error
	incoming    chan io.Reader
	stop        chan struct{}
	pong        chan struct{}
	req         chan []byte
	sendLk      sync.Mutex
	sub         chan *Event
}

type param struct {
	key   string
	value string
}

func argToChannelID(arg map[string]string) (id string) {
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

func NewClient() (*Client,error) {

	c := Client{
		Addr:     "wss://wsaws.okx.com:8443/ws/v5/public",
		Channels: make(map[string]ChannelSub),
		incoming: make(chan io.Reader),
		stop:     make(chan struct{}),
		pong:     make(chan struct{}, 1),
		req:      make(chan []byte, 1),
		sub:      make(chan *Event),
	}
	err:=c.Connect()
	if err != nil {
		return nil,err
	}
	go c.handleWsConn()
	return &c,nil
}

func (c *Client) Close() {
	close(c.stop)

}

func (c *Client) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(c.Addr, nil)
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

func (c *Client) handleEvent(e *Event) error {
	if e.Event == "error" {
		//log.Errorf("event :%+v",*e)
		return fmt.Errorf("event Error:%+v", *e)
	}
	if e.Event == "subscribe" {
		c.sub <- e
		return nil
	}
	log.Info("event :", *e)
	return nil
}

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

	log.Debug(string(b))
	var e Event
	err = e.Unmarshal(b)
	if err != nil {
		return fmt.Errorf("Unmarshal event error: %s", err)
	}
	if e.Event != "" {
		return c.handleEvent(&e)
		//TODO:fix event error
	}
	c.cLk.Lock()
	defer c.cLk.Unlock()
	chID := argToChannelID(e.Arg)
	if ch, ok := c.Channels[chID]; ok {
		go ch.ReciveData(b)
		return nil
	}
	return fmt.Errorf("can't not find the channel of id %s", chID)
}

func (c *Client) send(msg []byte) error {
	c.sendLk.Lock()
	defer c.sendLk.Unlock()
	return c.conn.WriteMessage(websocket.TextMessage, msg)
}

func (c *Client) handleWsConn() {
	timeoutTimer := time.NewTimer(60 * time.Second)
	defer timeoutTimer.Stop()
	go func() {
		for {
			c.send([]byte("ping"))
			select {
			case <-c.stop:
				return
			default:
				time.Sleep(30 * time.Second)
			}
		}
	}()
	go c.nextMsg()
	for {
		timeoutTimer.Reset(30 * time.Second)
		select {
		case reader, ok := <-c.incoming:
			err := c.incomingErr

			if ok {
				//如果积累了很多，怎么处理
				err = c.handleReader(reader)
				if err != nil {
					log.Error(err)
				}
				go c.nextMsg()
				continue
			}
			if err != nil {
				log.Error(err)
			}
			return
		case <-timeoutTimer.C:
			log.Warn("ping time out, try to reconnect")
			err := c.reconnect()
			if err != nil {
				log.Error(err)
				//FIXME: should proccess error friendly
				panic(err)
			}
		case <-c.pong:
		case <-c.stop:
			cmsg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
			if err := c.conn.WriteMessage(websocket.CloseMessage, cmsg); err != nil {
				log.Warn("failed to write close message: ", err)
			}
			if err := c.conn.Close(); err != nil {
				log.Warnw("websocket close error", "error", err)
			}
			return
		}
	}
}

type ChannelSub interface {
	GetArg() map[string]string
	ReciveData([]byte)
}

func (c *Client) Subscribe(cs ChannelSub) error {
	arg := cs.GetArg()
	sub := &Op{
		Op: "subscribe",
		Args: []map[string]string{
			arg,
		},
	}
	j, err := json.Marshal(sub)
	if err != nil {
		return err
	}
	err = c.send(j)
	if err != nil {
		return fmt.Errorf("fail to subsribe :%s", err)
	}

	id := argToChannelID(arg)

	timer := time.NewTimer(30 * time.Second)
	select {
	case s := <-c.sub:
		subid := argToChannelID(s.Arg)
		if subid != id {
			return fmt.Errorf("subsribe error: the id not match %s != %s", subid, id)
		}
	case <-timer.C:
		return fmt.Errorf("subscribe time out:%s ", id)
	}

	c.cLk.Lock()
	defer c.cLk.Unlock()
	c.Channels[id] = cs
	return nil
}
