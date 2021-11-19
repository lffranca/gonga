package server

import (
	"context"
	"github.com/lffranca/gonga/domain"
)

type ServiceRepository interface {
	List(ctx context.Context, gatewayID *string, size *int, offset *string, tags []*string, matchAllTags *bool) ([]*domain.Service, *domain.Option, error)
}
