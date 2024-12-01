package httpserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type GinHTTPServerAdapter struct {
	Engine   *gin.Engine
	basePath string
}

func NewGinHTTPServerAdapter() *GinHTTPServerAdapter {
	httpServer := &GinHTTPServerAdapter{
		Engine: gin.Default(),
	}

	httpServer.Engine.SetTrustedProxies(nil)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"}
	httpServer.Engine.Use(cors.New(config))
	return httpServer
}

func (g *GinHTTPServerAdapter) SetBasePath(basePath string) {
	g.basePath = basePath
}

func (g *GinHTTPServerAdapter) Run(adds ...string) error {
	return g.Engine.Run(adds...)
}

func (g *GinHTTPServerAdapter) SetSwagger(path string) {
	g.Engine.GET(g.basePath+path, ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (g *GinHTTPServerAdapter) GET(path string, callback func(HTTPContext)) {
	g.Engine.GET(
		g.basePath+path,
		func(c *gin.Context) {
			callback(c)
		},
	)
}

func (g *GinHTTPServerAdapter) POST(path string, callback func(HTTPContext)) {
	g.Engine.POST(
		g.basePath+path,
		func(c *gin.Context) {
			callback(c)
		},
	)
}

func (g *GinHTTPServerAdapter) PUT(path string, callback func(HTTPContext)) {
	g.Engine.PUT(
		g.basePath+path,
		func(c *gin.Context) {
			callback(c)
		},
	)
}

func (g *GinHTTPServerAdapter) PATCH(path string, callback func(HTTPContext)) {
	g.Engine.PATCH(
		g.basePath+path,
		func(c *gin.Context) {
			callback(c)
		},
	)
}

func (g *GinHTTPServerAdapter) DELETE(path string, callback func(HTTPContext)) {
	g.Engine.DELETE(
		g.basePath+path,
		func(c *gin.Context) {
			callback(c)
		},
	)
}
