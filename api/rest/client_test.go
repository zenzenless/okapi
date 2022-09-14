package rest

import (
	"net/http"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestClientDo(t *testing.T) {
	c := NewClient(&APIKey{})
	req, err := http.NewRequest("GET", "https://www.okx.com/api/v5/market/tickers?instType=SPOT", nil)
	assert.NoError(t, err)
	_, err = c.v0Do(req)
	assert.NoError(t, err)
}
