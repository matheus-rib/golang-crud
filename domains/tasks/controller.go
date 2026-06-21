package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func listTasks(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": findAllTasks()})
}

func showTask(context *gin.Context) {
	task, err := findTask(context)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": task})
}

func createTask(context *gin.Context) {
	task, err := insertTask(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": task})
}

func renameTask(context *gin.Context) {
	task, err, statusCode := updateTaskName(context)

	if err != nil {
		context.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	context.JSON(statusCode, gin.H{"data": task})
}

func toggleTaskStatus(context *gin.Context) {
	err, statusCode := toggleCompletedStatus(context)

	if err != nil {
		context.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	context.JSON(statusCode, nil)
}

func removeTask(context *gin.Context) {
	err := deleteTask(context)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
