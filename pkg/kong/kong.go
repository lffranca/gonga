package kong

import (
	"encoding/json"
	"github.com/lffranca/gonga/pkg/kong/healthroutes"
	"github.com/lffranca/gonga/pkg/kong/informationroutes"
	"net/url"
)

func NewKong(urlAdmin string) (*kong, error) {
	u, errUrl := url.Parse(urlAdmin)
	if errUrl != nil {
		return nil, errUrl
	}

	return &kong{
		url: u,
	}, nil
}

type kong struct {
	url        *url.URL
	credential *credential
}

func (gateway *kong) InformationRoutes() (*informationroutes.InformationRoutes, error) {
	response, errResponse := request(gateway, informationroutes.URL, HTTPMethodGET, nil, nil, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item informationroutes.InformationRoutes
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) HealthRoutes() (*healthroutes.StatusRoutes, error) {
	response, errResponse := request(gateway, healthroutes.URL, HTTPMethodGET, nil, nil, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item healthroutes.StatusRoutes
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}
