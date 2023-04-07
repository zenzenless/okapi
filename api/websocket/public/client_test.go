package public

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/james-zhang-bing/okapi"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gotest.tools/assert"
)

func TestClient(t *testing.T) {
	client, err := NewPublicClient(okapi.AwsPublicWsURL)
	assert.NilError(t, err)
	//go client.handleWsConn(context.Background())

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
	assert.NilError(t, err)
	for i := 0; i < 10; i++ {
		ts := <-ch.Received
		t.Log(ts)
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
