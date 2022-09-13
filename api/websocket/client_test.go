package websocket

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	client,err := NewClient()
	assert.NoError(t, err)
	go client.handleWsConn()

	// s := `{
	// 	"op": "subscribe",
	// 	"args": [{
	// 		"channel": "candle1m",
	// 		"instId": "BTC-USDT"
	// 	}]
	// }`
	ch := NewCandleChannel("1m", "BTC-USDT")
	err = client.Subscribe(ch)

	//err = client.conn.WriteMessage(websocket.TextMessage, []byte(s))
	assert.NoError(t, err)
	for i := 0; i < 10; i++ {
		ts := <-ch.Received
		t.Log(ts)
	}
}

func TestClient_Dial(t *testing.T) {
	type fields struct {
		Addr string
		conn *websocket.Conn
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "okx",
			fields: fields{
				Addr: "wss://wsaws.okx.com:8443/ws/v5/public",
			},
			wantErr: false,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Addr: tt.fields.Addr,
				conn: tt.fields.conn,
			}
			if err := c.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Client.Dial() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPingPong(t *testing.T) {
	c ,err:= NewClient()
	
	assert.NoError(t, err)
	err = c.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
	assert.NoError(t, err)
	s := `{
		"op": "subscribe",
		"args": [{
			"channel": "candle1m",
			"instId": "BTC-USDT"
		}]
	}`
	err = c.conn.WriteMessage(websocket.TextMessage, []byte(s))
	assert.NoError(t, err)
	m, r, err := c.conn.NextReader()
	assert.NoError(t, err)
	fmt.Println(m)
	b, err := ioutil.ReadAll(r)
	assert.NoError(t, err)
	fmt.Println(string(b))
	fmt.Println(time.Now().UnixMilli())
	_, r, err = c.conn.NextReader()
	assert.NoError(t, err)
	b, err = ioutil.ReadAll(r)
	assert.NoError(t, err)
	fmt.Println(string(b))
	_, r, err = c.conn.NextReader()
	assert.NoError(t, err)
	b, err = ioutil.ReadAll(r)
	assert.NoError(t, err)
	fmt.Println(string(b))
	time.Sleep(3 * time.Minute)
	for i := 0; i < 10; i++ {
		fmt.Println(time.Now().UnixMilli())
		_, r, err := c.conn.NextReader()
		assert.NoError(t, err)
		b, err := ioutil.ReadAll(r)
		assert.NoError(t, err)
		fmt.Println(string(b))
	}
}

func TestWebSocket(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	go http.ListenAndServe("127.0.0.1:8080", nil)
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/", nil)
	assert.NoError(t, err)
	ch := make(chan []byte, 10)
	go func() {
		for {
			_, rr, err := conn.NextReader()
			assert.NoError(t, err)
			b, err := ioutil.ReadAll(rr)
			assert.NoError(t, err)
			ch <- b
		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println("main ", i)
		time.Sleep(3 * time.Second)
		if len(ch) > 1 {

		}
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	i := 0
	for {
		time.Sleep(time.Second)
		i++
		if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintln(i))); err != nil {
			return
		}
		fmt.Println("HomeHandler ", i)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
}

func TestClientTimeOut(t *testing.T) {

}
