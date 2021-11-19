package service

import "errors"

type Options struct {
	GatewayRepository GatewayRepository
	ServiceRepository ServGatewayRepository
}

func (pkg *Options) validate() error {
	if pkg.GatewayRepository == nil {
		return errors.New("GatewayRepository is required param")
	}

	if pkg.ServiceRepository == nil {
		return errors.New("ServiceRepository is required param")
	}

	return nil
}
