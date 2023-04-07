package api

import (
	"github.com/james-zhang-bing/okapi"
	"github.com/james-zhang-bing/okapi/api/rest"

	"github.com/james-zhang-bing/okapi/api/websocket/private"
	"github.com/james-zhang-bing/okapi/api/websocket/public"
)

func NewApiClient(apikey *rest.APIKey, baseUrl, publicWsAddr, privateWsAddr okapi.BaseURL, model okapi.Destination) (*ClientGroup, error) {
	var g ClientGroup
	api := rest.NewClient(apikey, okapi.AwsRestURL, model)
	g.Rest = api
	publicWs, err := public.NewPublicClient(publicWsAddr)
	if err != nil {
		return nil, err
	}
	g.Public = publicWs
	if privateWsAddr == "" {
		return &g, nil
	}
	privateWs, err := private.NewPrivateClient(privateWsAddr, apikey)
	if err != nil {
		return nil, err
	}
	g.Private = privateWs
	return &g, nil
}

type ClientGroup struct {
	Rest    *rest.ClientRest
	Private *private.Client
	Public  *public.Client
}
