package server

import (
	"context"
	"github.com/lffranca/gonga/domain"
)

type RouteRepository interface {
	List(ctx context.Context, gatewayID *string, size *int, offset *string, tags []*string, matchAllTags *bool) ([]*domain.Route, *domain.Option, error)
}
