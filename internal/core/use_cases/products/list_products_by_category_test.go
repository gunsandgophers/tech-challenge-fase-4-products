package products

import (
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewListProductsByCategoryUseCase(t *testing.T) {
	//Arrange
	repo := &mocks.ProductRepositoryMock{}
	var prods []*entities.Product
	prods = append(prods, entities.CreateProduct(
		"Product 1",
		entities.PRODUCT_CATEGORY_DESSERTS,
		40.2,
		"Description",
		"Image",
	))
	repo.On("FindProductByCategory", mock.Anything, mock.Anything, mock.Anything).Return(prods, nil)
	uc := NewListProductsByCategoryUseCase(repo)

	// Act
	outputDTOS, err := uc.Execute(entities.PRODUCT_CATEGORY_DESSERTS.String(), 1, 10)

	// Asset
	assert.Nil(t, err)
	assert.Len(t, outputDTOS, 1)
}


