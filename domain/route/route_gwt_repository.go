package route

import (
	"context"
	"github.com/lffranca/gonga/domain"
)

type RGatewayRepository interface {
	List(ctx context.Context, gateway *domain.Gateway, size *int, offset *string, tags []*string, matchAllTags *bool) ([]*domain.Route, *domain.Option, error)
}
