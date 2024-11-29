package mocks

import (
	"tech-challenge-fase-1/internal/core/entities"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (r *ProductRepositoryMock) Insert(product *entities.Product) error {
	args := r.Called(product)
	return args.Error(0)
}

func (r *ProductRepositoryMock) Update(product *entities.Product) error {
	args := r.Called(product)
	return args.Error(0)
}

func (r *ProductRepositoryMock) FindProductByID(ID string) (*entities.Product, error) {
	args := r.Called(ID)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (r *ProductRepositoryMock) Delete(ID string) error {
	args := r.Called(ID)
	return args.Error(0)
}

func (r *ProductRepositoryMock) FindProductByCategory(
	category entities.ProductCategory,
	page, size int,
) ([]*entities.Product, error) {
	args := r.Called(category, page, size)
	return args.Get(0).([]*entities.Product), args.Error(1)
}
