package products

import (
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProductUseCase(t *testing.T) {
	//Arrange
	repo := &mocks.ProductRepositoryMock{}
	repo.On("Insert", mock.Anything).Return(nil)
	uc := NewCreateProductUseCase(repo)
	inputDTO := &dtos.ProductDTO{
		Name: "Product 1",
		Category: entities.PRODUCT_CATEGORY_DESSERTS.String(),
		Price: float64(40.2),
		Description: "Description ...",
		Image: "Image ...",
	}

	// Act
	outputDTO, err := uc.Execute(inputDTO)

	// Asset
	assert.Nil(t, err)
	assert.Empty(t, inputDTO.ID)
	assert.NotNil(t, outputDTO)
	assert.NotEmpty(t, outputDTO.ID)
	assert.Equal(t, inputDTO.Name, outputDTO.Name)
	assert.Equal(t, inputDTO.Category, outputDTO.Category)
	assert.Equal(t, inputDTO.Price, outputDTO.Price)
	assert.Equal(t, inputDTO.Description, outputDTO.Description)
	assert.Equal(t, inputDTO.Image, outputDTO.Image)
}

