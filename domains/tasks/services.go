package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheus-rib/golang-crud/database"
)

func findAllTasks() []Task {
	var tasks []Task
	database.DB.Find(&tasks)

	return tasks
}

func findTask(context *gin.Context) (Task, error) {
	taskId := context.Param("id")
	var task Task

	err := database.DB.First(&task, taskId).Error

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func insertTask(context *gin.Context) (Task, error) {
	var input CreateTaskInput
	err := context.ShouldBindJSON(&input)

	if err != nil {
		return Task{}, err
	}

	newTask := Task{
		Name:      input.Name,
		Completed: false,
	}

	err = database.DB.Create(&newTask).Error

	if err != nil {
		return Task{}, err
	}

	return newTask, nil
}

func updateTaskName(context *gin.Context) (Task, error, int) {
	task, err := findTask(context)

	if err != nil {
		return task, err, http.StatusNotFound
	}

	var input UpdateTaskNameInput
	err = context.ShouldBindJSON(&input)

	if err != nil {
		return Task{}, err, http.StatusBadRequest
	}

	database.DB.Model(&task).Updates(input)
	return task, nil, http.StatusOK
}

func toggleCompletedStatus(context *gin.Context) (error, int) {
	task, err := findTask(context)

	if err != nil {
		return err, http.StatusNotFound
	}

	task.Completed = !task.Completed
	database.DB.Save(&task)

	return nil, http.StatusNoContent
}

func deleteTask(context *gin.Context) error {
	task, err := findTask(context)

	if err != nil {
		return err
	}

	database.DB.Delete(&task)
	return nil
}
