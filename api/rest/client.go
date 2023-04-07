package rest

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/james-zhang-bing/okapi"

	requests "github.com/james-zhang-bing/okapi/requests/rest/public"
	responses "github.com/james-zhang-bing/okapi/responses/public_data"
)

type APIKey struct {
	ApiKey     string
	SecretKey  []byte
	Passphrase string
}

// Add request headers
func (c *ClientRest) AddSignature(req *http.Request, body string) {
	req.Header.Add("OK-ACCESS-KEY", c.APIKey.ApiKey)
	req.Header.Add("OK-ACCESS-TIMESTAMP", time.Now().UTC().Format("2006-01-02T15:04:05.000Z"))
	req.Header.Add("OK-ACCESS-PASSPHRASE", c.APIKey.Passphrase)
	req.Header.Add("OK-ACCESS-SIGN", c.v0sign(req, body))

}

// v0sign calculates the signature for a request
func (c *ClientRest) v0sign(req *http.Request, body string) string {
	path := strings.Replace(req.URL.Path, string(c.baseURL), "", -1)

	prehash := req.Header.Get("OK-ACCESS-TIMESTAMP") + req.Method + fmt.Sprintf("%s?%s", path, req.URL.RawQuery) + body
	h := hmac.New(sha256.New, []byte(c.APIKey.SecretKey))
	h.Write([]byte(prehash))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func GetApiKeyFromEnv() *APIKey {

	apikey := os.Getenv("APIKEY")
	secretkey := os.Getenv("SECRET_KEY")
	passphrase := os.Getenv("PASS_PHRASE")

	return NewAPIKey(apikey, secretkey, passphrase)
}

// NewAPIKey creates a new API key
func NewAPIKey(apiKey string, secretKey string, passphrase string) *APIKey {
	return &APIKey{
		ApiKey:     apiKey,
		SecretKey:  []byte(secretKey),
		Passphrase: passphrase,
	}
}

type ClientRest struct {
	Account    *Account
	SubAccount *SubAccount
	Trade      *Trade
	Funding    *Funding
	Market     *Market
	PublicData *PublicData
	TradeData  *TradeData
	*APIKey
	//baseURL     string
	destination okapi.Destination
	baseURL     okapi.BaseURL
	client      *http.Client
}

// NewClient creates a new client
func NewClient(apiKey *APIKey, baseURL okapi.BaseURL, destination okapi.Destination) *ClientRest {
	c := &ClientRest{
		APIKey:      apiKey,
		baseURL:     baseURL,
		destination: destination,
		client:      http.DefaultClient,
	}
	c.Account = NewAccount(c)
	c.SubAccount = NewSubAccount(c)
	c.Trade = NewTrade(c)
	c.Funding = NewFunding(c)
	c.Market = NewMarket(c)
	c.PublicData = NewPublicData(c)
	c.TradeData = NewTradeData(c)
	return c
}

// client do a request
func (c *ClientRest) Do(method, path string, private bool, params ...interface{}) (*http.Response, error) {
	u := fmt.Sprintf("%s%s", c.baseURL, path)
	var (
		r    *http.Request
		err  error
		j    []byte
		body string
	)

	if method == http.MethodGet {

		r, err = http.NewRequest(http.MethodGet, u, nil)
		if err != nil {
			return nil, err
		}

		if len(params) > 0 {
			paramsMap := okapi.MI2MS(params[0])
			// param,ok:= params[0].(map[string]string)
			// if !ok{
			// 	return nil, fmt.Errorf("params is not the mam[string]string type")
			// }
			q := r.URL.Query()
			for k, v := range paramsMap {
				q.Add(k, strings.ReplaceAll(v, "\"", ""))
			}
			r.URL.RawQuery = q.Encode()
			if len(paramsMap) > 0 {
				path += "?" + r.URL.RawQuery
			}
		}
	} else {
		j, err = json.Marshal(params[0])
		if err != nil {
			return nil, err
		}
		body = string(j)
		if body == "{}" || body == "[]" {
			body = ""
		}
		// fmt.Println(body)
		r, err = http.NewRequest(method, u, bytes.NewBuffer(j))
		if err != nil {
			return nil, err
		}
		r.Header.Add("Content-Type", "application/json")
	}
	if err != nil {
		return nil, err
	}
	if private {
		timestamp, sign := c.sign(method, path, body)
		r.Header.Add("OK-ACCESS-KEY", c.ApiKey)
		r.Header.Add("OK-ACCESS-PASSPHRASE", c.Passphrase)
		r.Header.Add("OK-ACCESS-SIGN", sign)
		r.Header.Add("OK-ACCESS-TIMESTAMP", timestamp)
	}
	if c.destination == okapi.DemoServer {
		r.Header.Add("x-simulated-trading", "1")
	}
	if os.Getenv("HTTP_DEBUG") != "" {
		fh, _ := httputil.DumpRequest(r, true)
		fmt.Printf("HTTP_DEBUG:\n%s\n", string(fh))
	}

	return c.client.Do(r)
}

// Status
// Get event status of system upgrade
//
// https://www.okex.com/docs-v5/en/#rest-api-status
func (c *ClientRest) Status(req requests.Status) (response responses.Status, err error) {
	p := "/api/v5/system/status"
	m := okapi.S2M(req)
	res, err := c.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

func (c *ClientRest) sign(method, path, body string) (string, string) {
	format := "2006-01-02T15:04:05.999Z07:00"
	t := time.Now().UTC().Format(format)
	ts := fmt.Sprint(t)
	s := ts + method + path + body
	p := []byte(s)
	h := hmac.New(sha256.New, c.SecretKey)
	h.Write(p)
	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}

type Res interface {
	ErrStatus() error
}

func Decode(d *json.Decoder, v Res) error {
	err := d.Decode(v)
	if err == nil {
		err = v.ErrStatus()
	}
	return err
}
