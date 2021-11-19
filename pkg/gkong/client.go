package gkong

import (
	"errors"
	"github.com/kong/go-kong/kong"
	"github.com/lffranca/gonga/domain"
)

func New() (*Client, error) {
	client := new(Client)
	client.common.client = client
	client.Service = (*ServiceService)(&client.common)
	client.Route = (*RouteService)(&client.common)

	return client, nil
}

type Client struct {
	common  service
	Service *ServiceService
	Route   *RouteService
}

func (pkg *Client) getClient(gateway *domain.Gateway) (*kong.Client, error) {
	if gateway == nil {
		return nil, errors.New("gateway param is required")
	}

	return kong.NewClient(gateway.Host, nil)
}

type service struct {
	client *Client
}
