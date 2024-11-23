package products

import (
	"tech-challenge-fase-1/internal/core/repositories"
)

type DeleteProductUseCase struct {
	productRepository repositories.ProductRepositoryInterface
}

func NewDeleteProductUseCase(productRepository repositories.ProductRepositoryInterface) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		productRepository: productRepository,
	}
}

func (dpc *DeleteProductUseCase) Execute(productID string) error {
	_, err := dpc.productRepository.FindProductByID(productID)
	if err != nil {
		return err
	}

	return dpc.productRepository.Delete(productID)
}
