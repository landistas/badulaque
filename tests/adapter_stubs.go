package tests

import (
	"context"

	inmemory "github.com/landistas/badulaque/pkg/adapters/outputs/storage"
	"github.com/landistas/badulaque/pkg/entities"
)

type StubProductStorageAdapter struct {
	StubInfraStorage inmemory.InfraStorage
}

func (stubProductStorageAdapter StubProductStorageAdapter) Add(ctx context.Context, product entities.Product) (*entities.Product, error) {
	return &product, nil
}
func (stubProductStorageAdapter StubProductStorageAdapter) Get(ctx context.Context, productID string) (*entities.Product, error) {
	v, err := stubProductStorageAdapter.StubInfraStorage.Read(ctx, productID)
	product, ok := v.(entities.Product)
	if ok {
		return &product, err
	}
	return nil, err
}
