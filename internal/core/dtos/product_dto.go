package dtos

import "tech-challenge-fase-1/internal/core/entities"

type ProductDTO struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Category    string  `json:"category,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
	Image       string  `json:"image,omitempty"`
}

func NewProductDTOFromEntity(product *entities.Product) *ProductDTO {
	return &ProductDTO{
		ID:          product.GetId(),
		Name:        product.GetName(),
		Category:    product.GetCategory().String(),
		Price:       product.GetPrice(),
		Description: product.GetDescription(),
		Image:       product.GetImage(),
	}
}
