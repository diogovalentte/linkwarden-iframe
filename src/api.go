// Package api implements the API routes and groups
package api

import (
	"github.com/gin-gonic/gin"

	"github.com/diogovalentte/linkwarden-iframe/src/routes"
)

// SetupRouter sets up the API routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		routes.HealthCheckRoute(v1)
	}
	{
		routes.LinksRoutes(v1)
	}

	return router
}
