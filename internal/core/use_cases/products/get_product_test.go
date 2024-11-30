package products

import (
	"errors"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewGetProductUseCase(t *testing.T) {
	//Arrange
	repo := &mocks.ProductRepositoryMock{}
	prod := entities.CreateProduct(
		"Product 1",
		entities.PRODUCT_CATEGORY_DESSERTS,
		40.2,
		"Description",
		"Image",
	)
	repo.On("FindProductByID", prod.GetId()).Return(prod, nil)
	uc := NewGetProductUseCase(repo)

	// Act
	outputDTO, err := uc.Execute(prod.GetId())

	// Asset
	assert.Nil(t, err)
	assert.Equal(t, prod.GetId(), outputDTO.ID)
	assert.Equal(t, prod.GetName(), outputDTO.Name)
	assert.Equal(t, prod.GetCategory().String(), outputDTO.Category)
	assert.Equal(t, prod.GetPrice(), outputDTO.Price)
	assert.Equal(t, prod.GetDescription(), outputDTO.Description)
	assert.Equal(t, prod.GetImage(), outputDTO.Image)
}


func TestNewProductUseCase_InvalidID(t *testing.T) {
	//Arrange
	repo := &mocks.ProductRepositoryMock{}
	var prod *entities.Product
	repo.On("FindProductByID", mock.Anything).Return(prod, errors.New("Invalid"))
	uc := NewGetProductUseCase(repo)

	// Act
	_, err := uc.Execute("")

	// Asset
	assert.NotNil(t, err)
}

