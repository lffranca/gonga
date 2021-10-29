package server

import (
	"context"
	"github.com/lffranca/gonga/domain"
)

type GatewayDatabase interface {
	Get(ctx context.Context, id *string) (*domain.Gateway, error)
	Save(ctx context.Context, item *domain.Gateway) (*domain.Gateway, error)
}
