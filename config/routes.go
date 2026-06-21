package config

import (
	"github.com/gin-gonic/gin"
	healthCheck "github.com/matheus-rib/golang-crud/domains/healthCheck"
	"github.com/matheus-rib/golang-crud/domains/tasks"
)

func registerRoutes(router *gin.Engine) {
	api := router.Group("/")

	healthCheck.RegisterRoutes(api)
	tasks.RegisterRoutes(api)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	registerRoutes(router)

	return router
}
