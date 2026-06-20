package config

import (
	"github.com/gin-gonic/gin"
	healthCheck "github.com/matheus-rib/golang-crud/domains/healthCheck"
	"github.com/matheus-rib/golang-crud/domains/tasks"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/")

	healthCheck.RegisterRoutes(api)
	tasks.RegisterRoutes(api)
}
