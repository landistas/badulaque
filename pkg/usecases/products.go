package usecases

import (
	"context"
	"errors"

	"github.com/landistas/badulaque/pkg/entities"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, product entities.Product) (*entities.Product, error)
	GetProduct(ctx context.Context, productID string) (*entities.Product, error)
}

type ProductStorageAdapter interface {
	Add(ctx context.Context, product entities.Product) (*entities.Product, error)
	Get(ctx context.Context, productID string) (*entities.Product, error)
}

type DefaultProductUseCase struct {
	productStorageAdapter ProductStorageAdapter
}

func NewDefaultProductUseCase(productStorageAdapter ProductStorageAdapter) DefaultProductUseCase {
	return DefaultProductUseCase{
		productStorageAdapter: productStorageAdapter,
	}
}

func (useCase DefaultProductUseCase) CreateProduct(ctx context.Context, product entities.Product) (*entities.Product, error) {
	// 1. Validate that it has an ID
	// 2. Check that it doesn't exists
	// 3. Save

	if product.ID == "" {
		return nil, errors.New("id can not be nil")
	}

	returnedProduct, err := useCase.productStorageAdapter.Get(ctx, product.ID)
	if err != nil {
		return nil, err
	}

	if returnedProduct != nil {
		return nil, errors.New("product already exists")
	}

	return useCase.productStorageAdapter.Add(ctx, product)
}

func (useCase DefaultProductUseCase) GetProduct(ctx context.Context, productID string) (*entities.Product, error) {
	return nil, nil
}
