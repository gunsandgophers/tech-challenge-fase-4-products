package app

import (
	"tech-challenge-fase-1/internal/infra/controllers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Registra as rotas dos controllers
func registerRouters(app *APIApp) {
	productController := controllers.NewProductController(app.productRepository)

	baseUrl := "/api/v1"
	app.httpServer.SetBasePath(baseUrl)

	//products
	app.httpServer.POST("/product", productController.CreateProduct)
	app.httpServer.PUT("/product/:id", productController.UpdateProduct)
	app.httpServer.DELETE("/product/:id", productController.DeleteProduct)
	app.httpServer.GET("/product/:category", productController.ListProductsByCategory)

	app.httpServer.SetSwagger("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
