package service

import (
	"context"
	"github.com/lffranca/gonga/internal/entity"
)

type Data interface {
	List(ctx context.Context, offset, limit *int, search, sort *string) ([]entity.Service, error)
	Create(ctx context.Context, service *entity.Service) (*entity.Service, error)
}
