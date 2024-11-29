package dtos

import (
	"tech-challenge-fase-1/internal/core/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProductDTOFromEntity(t *testing.T) {
	//Arrange
	product := entities.CreateProduct(
		"Product 1",
		entities.PRODUCT_CATEGORY_SANDWICH,
		10,
		"Product 1 description",
		"image...",
	)
	//Act
	dto := NewProductDTOFromEntity(product)
	//Assert
	assert.NotNil(t, dto)
	assert.Equal(t, product.GetId(), dto.ID)
	assert.Equal(t, product.GetName(), dto.Name)
	assert.Equal(t, product.GetCategory().String(), dto.Category)
	assert.Equal(t, product.GetPrice(), dto.Price)
	assert.Equal(t, product.GetDescription(), dto.Description)
	assert.Equal(t, product.GetImage(), dto.Image)
}
