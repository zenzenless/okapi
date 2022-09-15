package rest

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"os"
	"strings"
	"time"

	"github.com/james-zhang-bing/okapi"
	"github.com/james-zhang-bing/okapi/api"
	requests "github.com/james-zhang-bing/okapi/requests/rest/public"
	responses "github.com/james-zhang-bing/okapi/responses/public_data"
)

type APIKey struct {
	apiKey     string
	secretKey  []byte
	passphrase string
}

// Add request headers
func (c *ClientRest) AddSignature(req *http.Request, body string) {
	req.Header.Add("OK-ACCESS-KEY", c.APIKey.apiKey)
	req.Header.Add("OK-ACCESS-TIMESTAMP", time.Now().UTC().Format("2006-01-02T15:04:05.000Z"))
	req.Header.Add("OK-ACCESS-PASSPHRASE", c.APIKey.passphrase)
	req.Header.Add("OK-ACCESS-SIGN", c.v0sign(req, body))

}

//v0sign calculates the signature for a request
func (c *ClientRest) v0sign(req *http.Request, body string) string {
	path := strings.Replace(req.URL.Path, string(c.baseURL), "", -1)

	prehash := req.Header.Get("OK-ACCESS-TIMESTAMP") + req.Method + fmt.Sprintf("%s?%s", path, req.URL.RawQuery) + body
	h := hmac.New(sha256.New, []byte(c.APIKey.secretKey))
	h.Write([]byte(prehash))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func fromEnv() []string {

	apikey := os.Getenv("APIKEY")
	secretkey := os.Getenv("SECRET_KEY")
	passphrase := os.Getenv("PASS_PHRASE")
	return []string{apikey, secretkey, passphrase}
}

// NewAPIKey creates a new API key
func NewAPIKey(apiKey string, secretKey string, passphrase string) *APIKey {
	return &APIKey{
		apiKey:     apiKey,
		secretKey:  []byte(secretKey),
		passphrase: passphrase,
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
func (c *ClientRest) v0Do(req *http.Request) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do req error :%s", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var errcode api.ErrCode
	err = json.Unmarshal(body, &errcode)
	if err != nil {
		return nil, fmt.Errorf("unmarshal json error :%s body :%s", err, string(body))
	}
	if errcode.Code != api.OKXSuccessCode {
		return nil, fmt.Errorf("do request:%s error:%v", req.URL, errcode)
	}
	return body, nil
}

func (c *ClientRest) Do(method, path string, private bool, params ...map[string]string) (*http.Response, error) {
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
			q := r.URL.Query()
			for k, v := range params[0] {
				q.Add(k, strings.ReplaceAll(v, "\"", ""))
			}
			r.URL.RawQuery = q.Encode()
			if len(params[0]) > 0 {
				path += "?" + r.URL.RawQuery
			}
		}
	} else {
		j, err = json.Marshal(params[0])
		if err != nil {
			return nil, err
		}
		body = string(j)
		if body == "{}" {
			body = ""
		}
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
		r.Header.Add("OK-ACCESS-KEY", c.apiKey)
		r.Header.Add("OK-ACCESS-PASSPHRASE", c.passphrase)
		r.Header.Add("OK-ACCESS-SIGN", sign)
		r.Header.Add("OK-ACCESS-TIMESTAMP", timestamp)
	}
	if c.destination == okapi.DemoServer {
		r.Header.Add("x-simulated-trading", "1")
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
	err = d.Decode(&response)
	if err == nil {
		err = response.ErrStatus()
	}
	return
}

func (c *ClientRest) sign(method, path, body string) (string, string) {
	format := "2006-01-02T15:04:05.999Z07:00"
	t := time.Now().UTC().Format(format)
	ts := fmt.Sprint(t)
	s := ts + method + path + body
	p := []byte(s)
	h := hmac.New(sha256.New, c.secretKey)
	h.Write(p)
	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}
