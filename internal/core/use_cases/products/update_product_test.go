package products

import (
	"errors"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUpdateProductUseCase(t *testing.T) {
	//Arrange
	repo := &mocks.ProductRepositoryMock{}
	id := uuid.NewString()
	product := entities.RestoreProduct(
		id,
		"Product 1",
		entities.PRODUCT_CATEGORY_DESSERTS,
		40.2,
		"Description",
		"Image",
	)
	repo.On("FindProductByID", mock.Anything).Return(product, nil)
	repo.On("Update", mock.Anything).Return(nil)
	name := "New Name"
	category := entities.PRODUCT_CATEGORY_SANDWICH.String()
	price := float64(30.3)
	description := "Description"
	image := "Image"
	inputDTO := &dtos.ProductDTO{
		ID: id,
		Name: name,
		Category: category,
		Price: price,
		Description: description,
		Image: image,
	}
	uc := NewUpdateProductUseCase(repo)

	// Act
	outputDTO, err := uc.Execute(inputDTO)

	// Asset
	assert.Nil(t, err)
	assert.Equal(t, outputDTO.ID, id)
	assert.Equal(t, outputDTO.Name, name)
	assert.Equal(t, outputDTO.Category, category)
	assert.Equal(t, outputDTO.Price, price)
	assert.Equal(t, outputDTO.Description, description)
	assert.Equal(t, outputDTO.Image, image)
}

func TestNewUpdateProductUseCase_InvalidID(t *testing.T) {
	//Arrange
	repo := &mocks.ProductRepositoryMock{}
	id := uuid.NewString()
	repo.On("FindProductByID", mock.Anything).Return(&entities.Product{}, errors.New("Invalid"))
	repo.On("Update", mock.Anything).Return(nil)
	name := "New Name"
	category := entities.PRODUCT_CATEGORY_SANDWICH.String()
	price := float64(30.3)
	description := "Description"
	image := "Image"
	inputDTO := &dtos.ProductDTO{
		ID: id,
		Name: name,
		Category: category,
		Price: price,
		Description: description,
		Image: image,
	}
	uc := NewUpdateProductUseCase(repo)

	// Act
	_, err := uc.Execute(inputDTO)

	// Asset
	assert.NotNil(t, err)
}


func TestNewUpdateProductUseCase_InvalidCategory(t *testing.T) {
	//Arrange
	repo := &mocks.ProductRepositoryMock{}
	id := uuid.NewString()
	product := entities.RestoreProduct(
		id,
		"Product 1",
		entities.PRODUCT_CATEGORY_DESSERTS,
		40.2,
		"Description",
		"Image",
	)
	repo.On("FindProductByID", mock.Anything).Return(product, nil)
	repo.On("Update", mock.Anything).Return(errors.New("Invalid"))
	name := "New Name"
	category := ""
	price := float64(30.3)
	description := "Description"
	image := "Image"
	inputDTO := &dtos.ProductDTO{
		ID: id,
		Name: name,
		Category: category,
		Price: price,
		Description: description,
		Image: image,
	}
	uc := NewUpdateProductUseCase(repo)

	// Act
	_, err := uc.Execute(inputDTO)

	// Asset
	assert.NotNil(t, err)
}
