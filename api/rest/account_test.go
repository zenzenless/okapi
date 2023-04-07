package rest

import (
	"testing"

	"github.com/james-zhang-bing/okapi"
	"github.com/james-zhang-bing/okapi/requests/rest/account"

	"gotest.tools/assert"
)

func TestAccountGetPosition(t *testing.T) {
	c := NewClient(GetApiKeyFromEnv(), okapi.DemoRestURL, okapi.DemoServer)
	p, err := c.Account.GetPositions(account.GetPositions{
		InstType: "SWAP",
	})
	assert.NilError(t, err)
	t.Logf("%+v", (p.Positions[0]))
}
