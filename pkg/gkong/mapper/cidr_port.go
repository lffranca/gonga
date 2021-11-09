package mapper

import (
	"github.com/kong/go-kong/kong"
	"github.com/lffranca/gonga/domain"
)

func CIDRPortModelsToEntities(items []*kong.CIDRPort) (values []*domain.CIDRPort) {
	for _, item := range items {
		values = append(values, CIDRPortModelToEntity(item))
	}

	return
}

func CIDRPortModelToEntity(item *kong.CIDRPort) *domain.CIDRPort {
	if item == nil {
		return nil
	}

	return &domain.CIDRPort{
		IP:   item.IP,
		Port: item.Port,
	}
}
