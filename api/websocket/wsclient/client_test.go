package wsclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/james-zhang-bing/okapi"
	"github.com/james-zhang-bing/okapi/api/rest"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gotest.tools/assert"
)

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
				Addr: okapi.BaseURL(tt.fields.Addr),
				conn: tt.fields.conn,
			}
			if err := c.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Client.Dial() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPingPong(t *testing.T) {
	c, err := NewClient(okapi.AwsPublicWsURL, argToID)

	assert.NilError(t, err)
	err = c.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
	assert.NilError(t, err)
	s := `{
		"op": "subscribe",
		"args": [{
			"channel": "candle1m",
			"instId": "BTC-USDT"
		}]
	}`
	err = c.conn.WriteMessage(websocket.TextMessage, []byte(s))
	assert.NilError(t, err)
	m, r, err := c.conn.NextReader()
	assert.NilError(t, err)
	fmt.Println(m)
	b, err := ioutil.ReadAll(r)
	assert.NilError(t, err)
	fmt.Println(string(b))
	fmt.Println(time.Now().UnixMilli())
	_, r, err = c.conn.NextReader()
	assert.NilError(t, err)
	b, err = ioutil.ReadAll(r)
	assert.NilError(t, err)
	fmt.Println(string(b))
	_, r, err = c.conn.NextReader()
	assert.NilError(t, err)
	b, err = ioutil.ReadAll(r)
	assert.NilError(t, err)
	fmt.Println(string(b))
	time.Sleep(3 * time.Minute)
	for i := 0; i < 10; i++ {
		fmt.Println(time.Now().UnixMilli())
		_, r, err := c.conn.NextReader()
		assert.NilError(t, err)
		b, err := ioutil.ReadAll(r)
		assert.NilError(t, err)
		fmt.Println(string(b))
	}
}

func TestWebSocket(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	go http.ListenAndServe("127.0.0.1:8080", nil)
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/", nil)
	assert.NilError(t, err)
	ch := make(chan []byte, 10)
	go func() {
		for {
			_, rr, err := conn.NextReader()
			assert.NilError(t, err)
			b, err := ioutil.ReadAll(rr)
			assert.NilError(t, err)
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

func TestLogin(t *testing.T) {
	c, err := NewClient(okapi.DemoPrivateWsURL, argToID, rest.GetApiKeyFromEnv())
	assert.NilError(t, err)
	fmt.Println(*c.apikey)
	err = c.login()
	assert.NilError(t, err)

}

func TestMashall(t *testing.T) {
	j := `{
		"event": "login",
		"code": "60022",
		"msg": "Bulk login partially succeeded",
		"data":[
			{"apiKey": "86126n98-57ce-40fb-b714-afc0b9787083"}
		]
	}`

	var model map[string]any
	err := json.Unmarshal([]byte(j), &model)
	assert.NilError(t, err)
	fmt.Printf("%+v\n", model)
	for k, v := range model {
		rt := reflect.TypeOf(v)
		fmt.Printf("%s : type:%s\n", k, rt.Kind())
	}
	nj, err := json.Marshal(model)
	assert.NilError(t, err)
	fmt.Println(string(nj))

}
