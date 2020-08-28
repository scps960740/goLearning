package router

import (
	"net/http"

	"example.com/user/goWeb/router/middleware"
	"example.com/user/goWeb/handler/sd"
	"github.com/gin-gonic/gin"
)

// Load router
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	g.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world go")
	})

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
