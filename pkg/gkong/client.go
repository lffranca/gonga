package gkong

import (
	"github.com/kong/go-kong/kong"
	"net/http"
	"net/url"
)

func New(baseURL *string, cli *http.Client) (*Client, error) {
	clientKong, err := kong.NewClient(baseURL, cli)
	if err != nil {
		return nil, err
	}

	client := new(Client)
	client.common.client = client
	client.kong = clientKong
	client.Service = (*ServiceService)(&client.common)

	return client, nil
}

type service struct {
	client *Client
}

type Client struct {
	common  service
	url     *url.URL
	kong    *kong.Client
	Service *ServiceService
}
