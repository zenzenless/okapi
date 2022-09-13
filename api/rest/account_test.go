package rest

import (
	"fmt"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestAccountBalance(t *testing.T) {
	env := fromEnv()
	c := NewClient(NewAPIKey(env[0], env[1], env[2]))
	balance, err := c.GetAccountBalance()
	assert.NoError(t, err)
	fmt.Printf("%+v\n",balance)
}

func TestAccountPosition(t *testing.T){
	env := fromEnv()
	c := NewClient(NewAPIKey(env[0], env[1], env[2]))
	position, err := c.GetAccountPositions("","ETH-USDT-SWAP","")
	assert.NoError(t, err)
	fmt.Printf("%+v\n",position)
}