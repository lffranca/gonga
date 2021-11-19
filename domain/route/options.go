package route

import "errors"

type Options struct {
	GatewayRepository      GatewayRepository
	RouteGatewayRepository RGatewayRepository
}

func (pkg *Options) validate() error {
	if pkg.GatewayRepository == nil {
		return errors.New("GatewayRepository is required param")
	}

	if pkg.RouteGatewayRepository == nil {
		return errors.New("RouteGatewayRepository is required param")
	}

	return nil
}
