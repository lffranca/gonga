package presenter

import "github.com/lffranca/gonga/domain"

func GatewaysFromEntities(items []*domain.Gateway) (values []*GatewayBody) {
	for _, item := range items {
		values = append(values, GatewayFromEntity(item))
	}

	return
}

func GatewaysToEntities(items []*GatewayBody) (values []*domain.Gateway) {
	for _, item := range items {
		values = append(values, item.Entity())
	}

	return
}

func GatewayFromEntity(item *domain.Gateway) *GatewayBody {
	return &GatewayBody{
		ID:   item.ID,
		Name: item.Name,
		Host: item.Host,
	}
}

type GatewayBody struct {
	ID   *string `json:"id,omitempty" form:"id"`
	Name *string `json:"name,omitempty" form:"name" binding:"required"`
	Host *string `json:"host,omitempty" form:"host" binding:"required"`
}

func (item *GatewayBody) Entity() *domain.Gateway {
	return &domain.Gateway{
		ID:   item.ID,
		Name: item.Name,
		Host: item.Host,
	}
}
