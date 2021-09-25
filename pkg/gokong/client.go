package gokong

import (
	"github.com/kong/go-kong/kong"
)

func NewClient(baseURL *string) (*kong.Client, error) {
	return kong.NewClient(baseURL, nil)
}
