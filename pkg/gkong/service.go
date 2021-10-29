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
func (pkg *ServiceService) Create(ctx context.Context, service *domain.Service) (*domain.Service, error) {
	item := mapper.ServiceEntityToModel(service)
	if item == nil {
		return nil, errors.New("service param is required")
	}

	result, err := pkg.client.kong.Services.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return mapper.ServiceModelToEntity(result), nil
}

// Get fetches an Service in Kong.
func (pkg *ServiceService) Get(ctx context.Context, nameOrID *string) (*domain.Service, error) {
	item, err := pkg.client.kong.Services.Get(ctx, nameOrID)
	if err != nil {
		return nil, err
	}

	return mapper.ServiceModelToEntity(item), nil
}

// GetForRoute fetches a Service associated with routeID in Kong.
func (pkg *ServiceService) GetForRoute(ctx context.Context, routeID *string) (*domain.Service, error) {
	item, err := pkg.client.kong.Services.GetForRoute(ctx, routeID)
	if err != nil {
		return nil, err
	}

	return mapper.ServiceModelToEntity(item), nil
}

// Update updates an Service in Kong
func (pkg *ServiceService) Update(ctx context.Context, service *domain.Service) (*domain.Service, error) {
	item := mapper.ServiceEntityToModel(service)
	if item == nil {
		return nil, errors.New("service param is required")
	}

	result, err := pkg.client.kong.Services.Update(ctx, item)
	if err != nil {
		return nil, err
	}

	return mapper.ServiceModelToEntity(result), nil
}

// Delete deletes an Service in Kong
func (pkg *ServiceService) Delete(ctx context.Context, nameOrID *string) error {
	return pkg.client.kong.Services.Delete(ctx, nameOrID)
}

// List fetches a list of Services in Kong.
func (pkg *ServiceService) List(ctx context.Context, tags []*string, matchAllTags bool) ([]*domain.Service, error) {
	items, _, err := pkg.client.kong.Services.List(ctx, &kong.ListOpt{
		Tags:         tags,
		MatchAllTags: matchAllTags,
	})
	if err != nil {
		return nil, err
	}

	return mapper.ServiceModelsToEntities(items), nil
}

// ListAll fetches all Services in Kong.
func (pkg *ServiceService) ListAll(ctx context.Context) ([]*domain.Service, error) {
	items, err := pkg.client.kong.Services.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.ServiceModelsToEntities(items), nil
}
