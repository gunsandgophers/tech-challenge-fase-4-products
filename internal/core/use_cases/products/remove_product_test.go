package products

import (
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewDeleteProductUseCase(t *testing.T) {
	//Arrange
	repo := &mocks.ProductRepositoryMock{}
	repo.On("FindProductByID", mock.Anything).Return(&entities.Product{}, nil)
	repo.On("Delete", mock.Anything).Return(nil)
	uc := NewDeleteProductUseCase(repo)

	// Act
	err := uc.Execute(uuid.NewString())

	// Asset
	assert.Nil(t, err)
}



