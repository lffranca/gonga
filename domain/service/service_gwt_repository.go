package service

import (
	"context"
	"github.com/lffranca/gonga/domain"
)

type ServGatewayRepository interface {
	List(ctx context.Context, gateway *domain.Gateway, size *int, offset *string, tags []*string, matchAllTags *bool) ([]*domain.Service, *domain.Option, error)
}
