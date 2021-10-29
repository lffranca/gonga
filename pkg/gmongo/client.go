package gmongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const (
	collectionGateway = "gateway"
)

func New(mongoURI *string) (*Client, error) {
	if mongoURI == nil {
		return nil, errors.New("mongoURI is required")
	}

	client := new(Client)
	client.common.Client = client
	client.uri = mongoURI
	client.Gateway = (*GatewayService)(&client.common)

	return client, nil
}

type service struct {
	Client *Client
}

type Client struct {
	common  service
	uri     *string
	Gateway *GatewayService
}

func (pkg *Client) Database(ctx context.Context) (*mongo.Client, *mongo.Database, error) {
	cs, err := connstring.ParseAndValidate(*pkg.uri)
	if err != nil {
		return nil, nil, err
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(*pkg.uri))
	if err != nil {
		return nil, nil, err
	}

	db := client.Database(cs.Database)

	return client, db, nil
}
