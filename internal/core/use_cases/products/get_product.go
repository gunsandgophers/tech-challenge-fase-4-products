package products

import (
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/repositories"
)

type GetProductUseCase struct {
	productRepository repositories.ProductRepositoryInterface
}

func NewGetProductUseCase(
	productRepository repositories.ProductRepositoryInterface,
) *GetProductUseCase {
	return &GetProductUseCase{
		productRepository: productRepository,
	}
}

func (uc *GetProductUseCase) Execute(productID string) (dtos.ProductDTO, error) {
	product, err := uc.productRepository.FindProductByID(productID)
	if err != nil {
		return dtos.ProductDTO{}, err
	}
	return *dtos.NewProductDTOFromEntity(product), nil
}

