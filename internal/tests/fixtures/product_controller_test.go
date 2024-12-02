package fixtures

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateProduct(t *testing.T) {

	repo := &mocks.ProductRepositoryMock{}

	app := NewAPIAppIntegrationTest(repo)

	product := entities.RestoreProduct(
		uuid.NewString(), "Um produto", entities.PRODUCT_CATEGORY_DRINKS,
		13.37, "Uma descrição", "url",
	)

	repo.On("FindProductByID", product.GetId()).Return(product, nil).Once()
	repo.On("Update", mock.Anything).Return(nil).Once()

	w := httptest.NewRecorder()

	request := map[string]interface{}{
		"name":        "Um nome",
		"category":    entities.PRODUCT_CATEGORY_DRINKS.String(),
		"price":       13.37,
		"description": "Uma descrição",
		"image":       "url",
	}

	boby, _ := json.Marshal(request)

	req, err := http.NewRequest("PUT", "/api/v1/product/"+product.GetId(), bytes.NewReader(boby))
	app.HTTPServer().ServeHTTP(w, req)

	var response map[string]interface{}

	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, w.Result().StatusCode, 200)

	assert.Equal(t, response["message"], "operation: update-product successfull")
}

func TestUpdateDelete(t *testing.T) {

	repo := &mocks.ProductRepositoryMock{}

	app := NewAPIAppIntegrationTest(repo)

	product := entities.RestoreProduct(
		uuid.NewString(), "Um produto", entities.PRODUCT_CATEGORY_DRINKS,
		13.37, "Uma descrição", "url",
	)

	repo.On("FindProductByID", product.GetId()).Return(product, nil).Once()
	repo.On("Delete", mock.Anything).Return(nil).Once()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", "/api/v1/product/"+product.GetId(), strings.NewReader(""))
	app.HTTPServer().ServeHTTP(w, req)

	var response map[string]interface{}

	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, w.Result().StatusCode, 200)

	assert.Equal(t, response["message"], "operation: delete-product successfull")
}
