package inmemory

import (
	"context"

	"github.com/landistas/badulaque/pkg/entities"
)

type InfraStorage interface {
	Save(ctx context.Context, productID string, product interface{}) error
	Read(ctx context.Context, productID string) (interface{}, error)
}

type ProductStorageAdapter struct {
	infraStorage InfraStorage
}

func (adapter ProductStorageAdapter) Add(ctx context.Context, product entities.Product) (*entities.Product, error) {
	err := adapter.infraStorage.Save(ctx, product.ID, product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (adapter ProductStorageAdapter) Get(ctx context.Context, productID string) (*entities.Product, error) {
	return nil, nil
}
