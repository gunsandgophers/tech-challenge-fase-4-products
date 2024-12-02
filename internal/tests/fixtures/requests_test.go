package fixtures

import (
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/infra/controllers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateProduct(t *testing.T) {

	productRequest := controllers.ProductRequest{
		Name:        "Um nome",
		Category:    entities.PRODUCT_CATEGORY_DRINKS.String(),
		Price:       13.37,
		Description: "Uma descrição",
		Image:       "url",
	}

	err := productRequest.ValidateProduct()

	assert.Nil(t, err)

}

func TestValidateProductWithInvalidName(t *testing.T) {

	productRequest := controllers.ProductRequest{
		Category:    entities.PRODUCT_CATEGORY_DRINKS.String(),
		Price:       13.37,
		Description: "Uma descrição",
		Image:       "url",
	}

	err := productRequest.ValidateProduct()

	assert.Error(t, err)

}

func TestValidateProductWithInvalidCategory(t *testing.T) {

	productRequest := controllers.ProductRequest{
		Name:        "Um nome",
		Price:       13.37,
		Description: "Uma descrição",
		Image:       "url",
	}

	err := productRequest.ValidateProduct()

	assert.Error(t, err)

}

func TestValidateProductWithInvalidPrice(t *testing.T) {

	productRequest := controllers.ProductRequest{
		Name:        "Um nome",
		Category:    entities.PRODUCT_CATEGORY_DRINKS.String(),
		Description: "Uma descrição",
		Image:       "url",
	}

	err := productRequest.ValidateProduct()

	assert.Error(t, err)

}
