package gmongo

import (
	"context"
	"github.com/lffranca/gonga/domain"
	"github.com/lffranca/gonga/pkg/gmongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GatewayService service

func (pkg *GatewayService) List(ctx context.Context) ([]*domain.Gateway, error) {
	client, db, err := pkg.Client.Database(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	cursor, err := db.Collection(collectionGateway).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var items []*model.Gateway
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return model.GatewayModelsToEntities(items), nil
}

func (pkg *GatewayService) Get(ctx context.Context, id *string) (*domain.Gateway, error) {
	client, db, err := pkg.Client.Database(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	search := bson.A{
		bson.D{{
			"_id",
			*id,
		}},
	}
	if idP, err := primitive.ObjectIDFromHex(*id); err == nil {
		search = append(search, bson.D{{
			"_id",
			idP,
		}})
	}

	var item *model.Gateway
	if err := db.Collection(collectionGateway).FindOne(ctx, bson.D{
		{
			"$or",
			search,
		},
	}).Decode(&item); err != nil {
		return nil, err
	}

	return item.Entity(), nil
}

func (pkg *GatewayService) Save(ctx context.Context, item *domain.Gateway) (*domain.Gateway, error) {
	client, db, err := pkg.Client.Database(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	insertResult, err := db.Collection(collectionGateway).InsertOne(ctx, model.GatewayEntityToModel(item))
	if err != nil {
		return nil, err
	}

	if id, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		h := id.Hex()
		item.ID = &h
	}

	return item, nil
}
