package route

import (
	"context"
	"github.com/lffranca/gonga/domain"
)

type GatewayRepository interface {
	Get(ctx context.Context, id *string) (*domain.Gateway, error)
}
