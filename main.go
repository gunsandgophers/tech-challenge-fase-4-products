package main

import (
	_ "tech-challenge-fase-1/docs"
	"tech-challenge-fase-1/internal/infra/app"
	"tech-challenge-fase-1/internal/infra/database"
	httpserver "tech-challenge-fase-1/internal/infra/http"
	"tech-challenge-fase-1/internal/infra/repositories"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	httpServer := httpserver.NewGinHTTPServerAdapter()
	connection := database.NewPGXConnectionAdapter()
	productRepository := repositories.NewProductRepositoryDB(connection)
	app := app.NewAPIApp(httpServer, productRepository)
	app.Run()
	defer connection.Close()
}
