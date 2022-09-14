package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/james-zhang-bing/okapi"
)

func TestRequestOrder(t *testing.T) {
	env := fromEnv()
	c := NewClient(NewAPIKey(env[0], env[1], env[2]), okapi.RestURL, okapi.AwsServer)
	req, err := http.NewRequest("GET", "https://www.okx.com/api/v5/trade/orders-history", nil)
	q := req.URL.Query()
	q.Add("instType", "SPOT")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	c.AddSignature(req, "")
	body1 := strings.NewReader("adasdasd")
	a1, err := ioutil.ReadAll(body1)
	assert.NoError(t, err)
	a2, err := ioutil.ReadAll(body1)
	assert.NoError(t, err)
	fmt.Println(a1, a2)
	assert.NoError(t, err)
	_, err = c.v0Do(req)
	assert.NoError(t, err)
}

func TestGetOrdersOKX(t *testing.T) {
	env := fromEnv()
	c := NewClient(NewAPIKey(env[0], env[1], env[2]), okapi.RestURL, okapi.AwsServer)
	orders, err := c.GetOrdersOKX(InstType.SPOT)
	assert.NoError(t, err)
	fmt.Printf("%+v", len(orders))
}
func TestGetFills(t *testing.T) {
	env := fromEnv()
	c := NewClient(NewAPIKey(env[0], env[1], env[2]), okapi.RestURL, okapi.AwsServer)
	fills, err := c.GetFillsOKX(InstType.SPOT)
	assert.NoError(t, err)
	fmt.Printf("%+v", len(fills))
	j, err := json.Marshal(fills)
	assert.NoError(t, err)
	ioutil.WriteFile("fills.json", j, 0644)
}
