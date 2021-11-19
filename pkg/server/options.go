package server

import "errors"

type Options struct {
	TemplatePath      *string
	StaticJSPath      *string
	GatewayRepository GatewayRepository
	ServiceRepository ServiceRepository
	RouteRepository   RouteRepository
}

func (pkg *Options) validate() error {
	if pkg.TemplatePath == nil {
		return errors.New("template path param is required")
	}

	if pkg.StaticJSPath == nil {
		return errors.New("static js path param is required")
	}

	if pkg.GatewayRepository == nil {
		return errors.New("GatewayRepository param is required")
	}

	if pkg.ServiceRepository == nil {
		return errors.New("ServiceRepository param is required")
	}

	if pkg.RouteRepository == nil {
		return errors.New("RouteRepository param is required")
	}

	return nil
}
