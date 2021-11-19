package gkong

import (
	"context"
	"errors"
	"github.com/kong/go-kong/kong"
	"github.com/lffranca/gonga/domain"
	"github.com/lffranca/gonga/pkg/gkong/mapper"
)

type ServiceService service

// Create creates an Service in Kong
func (pkg *ServiceService) Create(ctx context.Context, gateway *domain.Gateway, service *domain.Service) (*domain.Service, error) {
	kongClient, err := pkg.client.getClient(gateway)
	if err != nil {
		return nil, err
	}

	item := mapper.ServiceEntityToModel(service)
	if item == nil {
		return nil, errors.New("service param is required")
	}

	result, err := kongClient.Services.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return mapper.ServiceModelToEntity(result), nil
}

// Get fetches an Service in Kong.
func (pkg *ServiceService) Get(ctx context.Context, gateway *domain.Gateway, nameOrID *string) (*domain.Service, error) {
	kongClient, err := pkg.client.getClient(gateway)
	if err != nil {
		return nil, err
	}

	item, err := kongClient.Services.Get(ctx, nameOrID)
	if err != nil {
		return nil, err
	}

	return mapper.ServiceModelToEntity(item), nil
}

// GetForRoute fetches a Service associated with routeID in Kong.
func (pkg *ServiceService) GetForRoute(ctx context.Context, gateway *domain.Gateway, routeID *string) (*domain.Service, error) {
	kongClient, err := pkg.client.getClient(gateway)
	if err != nil {
		return nil, err
	}

	item, err := kongClient.Services.GetForRoute(ctx, routeID)
	if err != nil {
		return nil, err
	}

	return mapper.ServiceModelToEntity(item), nil
}

// Update updates an Service in Kong
func (pkg *ServiceService) Update(ctx context.Context, gateway *domain.Gateway, service *domain.Service) (*domain.Service, error) {
	kongClient, err := pkg.client.getClient(gateway)
	if err != nil {
		return nil, err
	}

	item := mapper.ServiceEntityToModel(service)
	if item == nil {
		return nil, errors.New("service param is required")
	}

	result, err := kongClient.Services.Update(ctx, item)
	if err != nil {
		return nil, err
	}

	return mapper.ServiceModelToEntity(result), nil
}

// Delete deletes an Service in Kong
func (pkg *ServiceService) Delete(ctx context.Context, gateway *domain.Gateway, nameOrID *string) error {
	kongClient, err := pkg.client.getClient(gateway)
	if err != nil {
		return err
	}

	return kongClient.Services.Delete(ctx, nameOrID)
}

// List fetches a list of Services in Kong.
func (pkg *ServiceService) List(ctx context.Context, gateway *domain.Gateway, size *int, offset *string, tags []*string, matchAllTags *bool) ([]*domain.Service, *domain.Option, error) {
	kongClient, err := pkg.client.getClient(gateway)
	if err != nil {
		return nil, nil, err
	}

	options := &kong.ListOpt{}

	if size != nil {
		options.Size = *size
	}

	if offset != nil {
		options.Offset = *offset
	}

	if tags != nil && len(tags) > 0 {
		options.Tags = tags
	}

	if matchAllTags != nil {
		options.MatchAllTags = *matchAllTags
	}

	items, resultOption, err := kongClient.Services.List(ctx, options)
	if err != nil {
		return nil, nil, err
	}

	return mapper.ServiceModelsToEntities(items), mapper.OptionModelToEntity(resultOption), nil
}

// ListAll fetches all Services in Kong.
func (pkg *ServiceService) ListAll(ctx context.Context, gateway *domain.Gateway) ([]*domain.Service, error) {
	kongClient, err := pkg.client.getClient(gateway)
	if err != nil {
		return nil, err
	}

	items, err := kongClient.Services.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.ServiceModelsToEntities(items), nil
}
