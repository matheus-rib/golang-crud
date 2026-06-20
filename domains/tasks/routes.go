package tasks

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/tasks", listTasks)
	router.GET("/tasks/:id", showTask)
	router.POST("/tasks", createTask)
	router.PATCH("/tasks/:id", renameTask)
	router.PUT("/tasks/:id", toggleTaskStatus)
	router.DELETE("/tasks/:id", removeTask)
}
