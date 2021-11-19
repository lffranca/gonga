package service

import (
	"context"
	"errors"
	"github.com/lffranca/gonga/domain"
)

func New(options *Options) (*Service, error) {
	if options == nil {
		return nil, errors.New("options is required param")
	}

	if err := options.validate(); err != nil {
		return nil, err
	}

	service := new(Service)
	service.gatewayRepository = options.GatewayRepository
	service.serviceRepository = options.ServiceRepository

	return service, nil
}

type Service struct {
	gatewayRepository GatewayRepository
	serviceRepository ServGatewayRepository
}

func (mod *Service) List(ctx context.Context, gatewayID *string, size *int, offset *string, tags []*string, matchAllTags *bool) ([]*domain.Service, *domain.Option, error) {
	if gatewayID == nil {
		return nil, nil, errors.New("gatewayID is required param")
	}

	gateway, err := mod.gatewayRepository.Get(ctx, gatewayID)
	if err != nil {
		return nil, nil, err
	}

	return mod.serviceRepository.List(ctx, gateway, size, offset, tags, matchAllTags)
}
