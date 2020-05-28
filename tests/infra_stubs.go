package tests

import (
	"context"

	inmemory "github.com/landistas/badulaque/pkg/adapters/outputs/storage"
)

type StubInfraStorage struct {
	products map[string]interface{}
}

var _ inmemory.InfraStorage = (*StubInfraStorage)(nil)

func (stubInfraStorage *StubInfraStorage) Save(ctx context.Context, id string, v interface{}) error {
	if stubInfraStorage.products == nil {
		stubInfraStorage.products = make(map[string]interface{})
	}
	stubInfraStorage.products[id] = v
	return nil
}
func (stubInfraStorage StubInfraStorage) Read(ctx context.Context, id string) (interface{}, error) {
	product := stubInfraStorage.products[id]
	return product, nil
}
