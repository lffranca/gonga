package model

import (
	"github.com/lffranca/gonga/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GatewayEntityToModel(item *domain.Gateway) *Gateway {
	var id primitive.ObjectID
	if item.ID != nil {
		id, _ = primitive.ObjectIDFromHex(*item.ID)
	}

	return &Gateway{
		ID:   id,
		Name: item.Name,
		Host: item.Host,
	}
}

type Gateway struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name *string            `bson:"name,omitempty"`
	Host *string            `bson:"host,omitempty"`
}

func (item *Gateway) Entity() *domain.Gateway {
	var id *string
	if !item.ID.IsZero() {
		hex := item.ID.Hex()
		id = &hex
	}

	return &domain.Gateway{
		ID:   id,
		Name: item.Name,
		Host: item.Host,
	}
}

func GatewayModelsToEntities(items []*Gateway) (values []*domain.Gateway) {
	for _, item := range items {
		values = append(values, item.Entity())
	}

	return
}
