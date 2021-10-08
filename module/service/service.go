package service

import (
	"context"

	"github.com/lffranca/gonga/entity"
)

func NewService(data Data) (*service, error) {
	return &service{
		data: data,
	}, nil
}

type service struct {
	data Data
}

func (module *service) List(ctx context.Context, offset, limit *int, search, sort *string) ([]entity.Service, error) {
	return module.data.List(ctx, offset, limit, search, sort)
}

func (module *service) Create(ctx context.Context, item *entity.Service) (*entity.Service, error) {
	return module.data.Create(ctx, item)
}
