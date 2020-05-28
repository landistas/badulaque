package usecases_test

import (
	"context"
	"testing"

	"github.com/landistas/badulaque/pkg/entities"
	"github.com/landistas/badulaque/pkg/usecases"
	"github.com/landistas/badulaque/tests"
	"github.com/stretchr/testify/require"
)

func TestCreateProductHappyPath(t *testing.T) {
	product := entities.Product{
		ID: "id",
	}
	useCase := usecases.NewDefaultProductUseCase(
		tests.StubProductStorageAdapter{
			StubInfraStorage: &tests.StubInfraStorage{},
		},
	)

	returnedProduct, err := useCase.CreateProduct(context.TODO(), product)
	require.NoError(t, err)
	require.NotNil(t, returnedProduct)
	require.Equal(t, product.ID, returnedProduct.ID)
}
