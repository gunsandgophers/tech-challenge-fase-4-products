package fixtures

import (
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/infra/app"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)


func NewAPIAppIntegrationTest(productRepository repositories.ProductRepositoryInterface) *app.APIApp {
	httpServer := httpserver.NewGinHTTPServerAdapter()
	return app.NewAPIApp(httpServer, productRepository)
}

