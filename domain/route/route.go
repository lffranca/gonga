package route

import (
	"context"
	"errors"
	"github.com/lffranca/gonga/domain"
)

func New(options *Options) (*Route, error) {
	if options == nil {
		return nil, errors.New("options is required param")
	}

	if err := options.validate(); err != nil {
		return nil, err
	}

	route := new(Route)
	route.gatewayRepository = options.GatewayRepository
	route.routeRepository = options.RouteGatewayRepository

	return route, nil
}

type Route struct {
	gatewayRepository GatewayRepository
	routeRepository   RGatewayRepository
}

func (mod *Route) List(ctx context.Context, gatewayID *string, size *int, offset *string, tags []*string, matchAllTags *bool) ([]*domain.Route, *domain.Option, error) {
	if gatewayID == nil {
		return nil, nil, errors.New("gatewayID is required param")
	}

	gateway, err := mod.gatewayRepository.Get(ctx, gatewayID)
	if err != nil {
		return nil, nil, err
	}

	return mod.routeRepository.List(ctx, gateway, size, offset, tags, matchAllTags)
}
