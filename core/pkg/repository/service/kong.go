package service

import (
	"context"
	"errors"
	"github.com/kong/go-kong/kong"
	"github.com/lffranca/gonga/core/entity"
	"github.com/lffranca/gonga/core/pkg/mapper"
)

func NewKongRepository(client *kong.Client) (*kongRepository, error) {
	if client == nil {
		return nil, errors.New("invalid params")
	}

	return &kongRepository{
		client: client,
	}, nil
}

type kongRepository struct {
	client *kong.Client
}

func (repo *kongRepository) List(ctx context.Context, offset, limit *int, search, sort *string) ([]entity.Service, error) {
	services, _, err := repo.client.Services.List(ctx, nil)
	if err != nil {
		return nil, err
	}

	return mapper.ServicesKongToEntity(services), nil
}

func (repo *kongRepository) Create(ctx context.Context, service *entity.Service) (*entity.Service, error) {
	return nil, nil
}
