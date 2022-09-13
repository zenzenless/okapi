package rest

import (
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

	"github.com/james-zhang-bing/okapi/api"
)

type APIKey struct {
	APIKey     string
	SecretKey  string
	Passphrase string
}

// Add request headers
func (c *Client) AddSignature(req *http.Request, body string) {
	req.Header.Add("OK-ACCESS-KEY", c.APIKey.APIKey)
	req.Header.Add("OK-ACCESS-TIMESTAMP", time.Now().UTC().Format("2006-01-02T15:04:05.000Z"))
	req.Header.Add("OK-ACCESS-PASSPHRASE", c.APIKey.Passphrase)
	req.Header.Add("OK-ACCESS-SIGN", c.sign(req, body))

}

//sign calculates the signature for a request
func (c *Client) sign(req *http.Request, body string) string {
	path := strings.Replace(req.URL.Path, c.BaseURL, "", -1)

	prehash := req.Header.Get("OK-ACCESS-TIMESTAMP") + req.Method + fmt.Sprintf("%s?%s", path, req.URL.RawQuery) + body
	h := hmac.New(sha256.New, []byte(c.APIKey.SecretKey))
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
		APIKey:     apiKey,
		SecretKey:  secretKey,
		Passphrase: passphrase,
	}
}

type Client struct {
	APIKey  *APIKey
	BaseURL string
}

// NewClient creates a new client
func NewClient(apiKey *APIKey) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://aws.okx.com/",
	}
}

// client do a request
func (c *Client) Do(req *http.Request) ([]byte, error) {
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
