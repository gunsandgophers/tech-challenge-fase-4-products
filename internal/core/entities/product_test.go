package entities

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	// Arrange
	name := "Product 1"
	category := PRODUCT_CATEGORY_SANDWICH
	price := float64(10.4)
	description := "Some description"
	image := "Some image"
	// Act
	product := CreateProduct(name, category, price, description, image)
	// Assert
	assert.NotEmpty(t, product.GetId())
	assert.Equal(t, product.GetName(), name)
	assert.Equal(t, product.GetCategory(), category)
	assert.Equal(t, product.GetPrice(), price)
	assert.Equal(t, product.GetDescription(), description)
	assert.Equal(t, product.GetImage(), image)
}

func TestRestoreProduct(t *testing.T) {
	// Arrange
	id := uuid.NewString()
	name := "Product 1"
	category := PRODUCT_CATEGORY_SANDWICH
	price := float64(10.4)
	description := "Some description"
	image := "Some image"
	// Act
	product := RestoreProduct(id, name, category, price, description, image)
	// Assert
	assert.Equal(t, product.GetId(), id)
	assert.Equal(t, product.GetName(), name)
	assert.Equal(t, product.GetCategory(), category)
	assert.Equal(t, product.GetPrice(), price)
	assert.Equal(t, product.GetDescription(), description)
	assert.Equal(t, product.GetImage(), image)
}

func TestUpdateProduct(t *testing.T) {
	// Arrange
	name := "Product 1"
	category := PRODUCT_CATEGORY_SANDWICH
	price := float64(10.4)
	description := "Some description"
	image := "Some image"
	// Act
	product := CreateProduct(name, category, price, description, image)

	name = "Product 2"
	category = PRODUCT_CATEGORY_DESSERTS
	price = float64(50.7)
	description = "Some description updated"
	image = "Some image updated"
	product.SetName(name)
	product.SetCategory(category)
	product.SetPrice(price)
	product.SetDescription(description)
	product.SetImage(image)
	// Assert
	assert.Equal(t, product.GetName(), name)
	assert.Equal(t, product.GetCategory(), category)
	assert.Equal(t, product.GetPrice(), price)
	assert.Equal(t, product.GetDescription(), description)
	assert.Equal(t, product.GetImage(), image)
}

func TestProductCategory(t *testing.T) {
	categories := []ProductCategory{
		PRODUCT_CATEGORY_DESSERTS,
		PRODUCT_CATEGORY_SANDWICH,
		PRODUCT_CATEGORY_DRINKS,
		PRODUCT_CATEGORY_SIDEDISHES,
	}

	for _, c := range categories {
		assert.Equal(t, c.String(), string(c))
	}
}
