package mapper

import (
	"github.com/kong/go-kong/kong"
	"github.com/lffranca/gonga/domain"
	"time"
)

func RouteModelsToEntities(items []*kong.Route) (values []*domain.Route) {
	for _, item := range items {
		values = append(values, RouteModelToEntity(item))
	}

	return
}

func RouteModelToEntity(item *kong.Route) *domain.Route {
	if item == nil {
		return nil
	}

	var createdAt *time.Time
	if item.CreatedAt != nil {
		t := time.UnixMilli(int64(*item.CreatedAt))
		createdAt = &t
	}

	var updatedAt *time.Time
	if item.UpdatedAt != nil {
		t := time.UnixMilli(int64(*item.UpdatedAt))
		updatedAt = &t
	}

	return &domain.Route{
		ID:                      item.ID,
		Name:                    item.Name,
		Hosts:                   item.Hosts,
		Headers:                 item.Headers,
		Methods:                 item.Methods,
		Paths:                   item.Paths,
		PathHandling:            item.PathHandling,
		PreserveHost:            item.PreserveHost,
		Protocols:               item.Protocols,
		RegexPriority:           item.RegexPriority,
		Service:                 ServiceModelToEntity(item.Service),
		StripPath:               item.StripPath,
		CreatedAt:               createdAt,
		UpdatedAt:               updatedAt,
		SNIs:                    item.SNIs,
		Sources:                 CIDRPortModelsToEntities(item.Sources),
		Destinations:            CIDRPortModelsToEntities(item.Destinations),
		Tags:                    item.Tags,
		HTTPSRedirectStatusCode: item.HTTPSRedirectStatusCode,
		RequestBuffering:        item.RequestBuffering,
		ResponseBuffering:       item.ResponseBuffering,
	}
}
