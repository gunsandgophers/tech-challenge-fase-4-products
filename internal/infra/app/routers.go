package app

import (
	"tech-challenge-fase-1/internal/infra/controllers"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)

// Registra as rotas dos controllers
func registerRouters(app *APIApp) {
	productController := controllers.NewProductController(app.productRepository)

	baseUrl := "/api/v1"
	app.httpServer.(httpserver.HTTPRoutes).SetBasePath(baseUrl)

	// products
	app.httpServer.(httpserver.HTTPRoutes).POST("/product", productController.CreateProduct)
	app.httpServer.(httpserver.HTTPRoutes).PUT("/product/:id", productController.UpdateProduct)
	app.httpServer.(httpserver.HTTPRoutes).DELETE("/product/:id", productController.DeleteProduct)
	app.httpServer.(httpserver.HTTPRoutes).GET("/product/:id", productController.GetProduct)
	app.httpServer.(httpserver.HTTPRoutes).GET("/product/category/:category", productController.ListProductsByCategory)
	app.httpServer.(httpserver.HTTPRoutes).SetSwagger("/swagger/*any")
}
