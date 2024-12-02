package bdd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/infra/app"
	"tech-challenge-fase-1/internal/infra/controllers"
	"tech-challenge-fase-1/internal/tests/fixtures"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/mock"
)

type appCtxKey struct{}
type productCtxKey struct{}
type responseCtxKey struct{}

func newProductRequest(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, productCtxKey{}, controllers.ProductRequest{}), nil
}

func defineProductCategory(
	ctx context.Context,
	category string,
) (context.Context, error) {
	request := ctx.Value(productCtxKey{}).(controllers.ProductRequest)
	request.Category = category
	return context.WithValue(ctx, productCtxKey{}, request), nil
}

func defineProductName(
	ctx context.Context,
	name string,
) (context.Context, error) {
	request := ctx.Value(productCtxKey{}).(controllers.ProductRequest)
	request.Name = name
	return context.WithValue(ctx, productCtxKey{}, request), nil
}

func defineProductPrice(
	ctx context.Context,
	price string,
) (context.Context, error) {
	request := ctx.Value(productCtxKey{}).(controllers.ProductRequest)
	request.Price, _ = strconv.ParseFloat(price, 64)
	return context.WithValue(ctx, productCtxKey{}, request), nil
}

func defineProductDescription(
	ctx context.Context,
	description string,
) (context.Context, error) {
	request := ctx.Value(productCtxKey{}).(controllers.ProductRequest)
	request.Description = description
	// request.Image = image
	return context.WithValue(ctx, productCtxKey{}, request), nil
}

func defineProductImage(
	ctx context.Context,
	image string,
) (context.Context, error) {
	request := ctx.Value(productCtxKey{}).(controllers.ProductRequest)
	request.Image = image
	return context.WithValue(ctx, productCtxKey{}, request), nil
}

func sendProductRequest(
	ctx context.Context,
) (context.Context, error) {

	request := ctx.Value(productCtxKey{}).(controllers.ProductRequest)

	app := ctx.Value(appCtxKey{}).(*app.APIApp)
	w := httptest.NewRecorder()
	body, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/api/v1/product", bytes.NewReader(body))
	app.HTTPServer().ServeHTTP(w, req)
	return context.WithValue(ctx, responseCtxKey{}, w), nil
}

type ResponseCreateProduct struct {
	Data struct {
		Product dtos.ProductDTO `json:"product,omitempty"`
	} `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func theresShouldHaveANewProduct(ctx context.Context, name string, category string) error {
	response, _ := ctx.Value(responseCtxKey{}).(*httptest.ResponseRecorder)
	status := response.Result().StatusCode

	if status != http.StatusCreated {
		return errors.New("invalid status")
	}

	r := ResponseCreateProduct{}
	err := json.Unmarshal(response.Body.Bytes(), &r)
	if err != nil {
		return err
	}

	if r.Message != "operation: create-product successfull" {
		return errors.New("error on create-product")
	}

	request := ctx.Value(productCtxKey{}).(controllers.ProductRequest)
	if r.Data.Product.Name != request.Name {
		return errors.New("Name invalid")
	}
	if r.Data.Product.Category != request.Category {
		return errors.New("Category invalid")
	}

	return nil
}


func TestFeatures(t *testing.T) {
	productRepository := &mocks.ProductRepositoryMock{}
	productRepository.On("Insert", mock.Anything).Return(nil)
	app := fixtures.NewAPIAppIntegrationTest(productRepository)
  suite := godog.TestSuite{
    ScenarioInitializer: InitializeScenario,
    Options: &godog.Options{
      Format:   "pretty",
      Paths:    []string{"features"},
			DefaultContext: context.WithValue(context.Background(), appCtxKey{}, app),
      TestingT: t, // Testing instance that will run subtests.
    },
  }

  if suite.Run() != 0 {
    t.Fatal("non-zero status returned, failed to run feature tests")
  }
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Given(`^That I need to create a new product via API$`, newProductRequest)
	sc.Step(`^have the category as "([^"]*)"$`, defineProductCategory)
	sc.Step(`^name as "([^"]*)"$`, defineProductName)
	sc.Step(`^price as "([^"]*)"$`, defineProductPrice)
	sc.Step(`^description as "([^"]*)"$`, defineProductDescription)
	sc.Step(`^image as "([^"]*)"$`, defineProductImage)
	sc.When(`^I send the data$`, sendProductRequest)
	sc.Then(
		`^the product "([^"]*)" should be added to the list of products in the "([^"]*)" category$`,
		theresShouldHaveANewProduct,
	)
}

