package app

import (
	"tech-challenge-fase-1/internal/core/repositories"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)

type APIApp struct {
	httpServer        httpserver.HTTPServer
	productRepository repositories.ProductRepositoryInterface
}

func NewAPIApp(
	httpServer httpserver.HTTPServer,
	productRepository repositories.ProductRepositoryInterface,
) *APIApp {
	app := &APIApp{}
	// HTTP SERVER
	app.httpServer = httpServer

	// REPOSITORIES AND SERVICES
	app.productRepository = productRepository

	app.configRoutes()
	return app
}

func (app *APIApp) configRoutes() {
	registerRouters(app)
}

func (app *APIApp) Run() {
	app.httpServer.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
