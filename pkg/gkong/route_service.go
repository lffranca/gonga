package gkong

import (
	"context"
	"github.com/kong/go-kong/kong"
	"github.com/lffranca/gonga/domain"
	"github.com/lffranca/gonga/pkg/gkong/mapper"
)

type RouteService service

func (pkg *RouteService) List(ctx context.Context, gateway *domain.Gateway, size *int, offset *string, tags []*string, matchAllTags *bool) ([]*domain.Route, *domain.Option, error) {
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

	items, resultOption, err := kongClient.Routes.List(ctx, options)
	if err != nil {
		return nil, nil, err
	}

	return mapper.RouteModelsToEntities(items), mapper.OptionModelToEntity(resultOption), nil
}
