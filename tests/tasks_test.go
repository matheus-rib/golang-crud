//go:build integration

package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/matheus-rib/golang-crud/config"
	"github.com/matheus-rib/golang-crud/database"
	"github.com/matheus-rib/golang-crud/domains/tasks"
	tests "github.com/matheus-rib/golang-crud/tests/setup"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	tests.SetupDotEnvFile()

	database.DB = tests.ConnectDatabase()
	database.DB.AutoMigrate(&tasks.Task{})

	exitCode := m.Run()

	database.DB.Migrator().DropTable(&tasks.Task{})

	os.Exit(exitCode)
}

type DataTaskResponse struct {
	Data tasks.Task `json:"data"`
}

type DataTaskListResponse struct {
	Data []tasks.Task `json:"data"`
}

func Test_Tasks(t *testing.T) {
	router := config.SetupRouter()

	var TestSubjectTask tasks.Task

	t.Run("Create new task", func(t *testing.T) {
		createTaskInput := tasks.CreateTaskInput{
			Name: "New task",
		}

		taskInputJson, _ := json.Marshal(createTaskInput)

		req, _ := http.NewRequest("POST", "/tasks", strings.NewReader(string(taskInputJson)))
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		var response DataTaskResponse
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		TestSubjectTask = response.Data

		assert.NoError(t, err, "Response body should be valid JSON")
		assert.Equal(t, http.StatusCreated, resp.Code)
		assert.Equal(t, createTaskInput.Name, TestSubjectTask.Name)
	})

	t.Run("List tasks", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/tasks", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		var response DataTaskListResponse
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		taskListResponse := response.Data

		assert.NoError(t, err, "Response body should be valid JSON")
		assert.Equal(t, http.StatusOK, resp.Code)

		assert.Len(t, taskListResponse, 1, "Should return theo nly created task")
		assert.Equal(t, TestSubjectTask.ID, taskListResponse[0].ID)
		assert.Equal(t, TestSubjectTask.Name, taskListResponse[0].Name)
	})

	t.Run("Get task", func(t *testing.T) {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/tasks/%d", TestSubjectTask.ID), nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		var response DataTaskResponse
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		taskResponse := response.Data

		assert.NoError(t, err, "Response body should be valid JSON")
		assert.Equal(t, http.StatusOK, resp.Code)

		var task tasks.Task
		err = database.DB.First(&task, TestSubjectTask.ID).Error
		assert.NoError(t, err, "Should have find the task")

		assert.Equal(t, taskResponse.ID, task.ID)
		assert.Equal(t, taskResponse.Name, task.Name)
	})

	t.Run("Update task name", func(t *testing.T) {
		updateTaskInput := tasks.UpdateTaskNameInput{
			Name: "Renamed task",
		}

		taskInputJson, _ := json.Marshal(updateTaskInput)

		req, _ := http.NewRequest("PATCH", fmt.Sprintf("/tasks/%d", TestSubjectTask.ID), strings.NewReader(string(taskInputJson)))
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		var response DataTaskResponse
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		taskResponse := response.Data

		assert.NoError(t, err, "Response body should be valid JSON")
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, updateTaskInput.Name, taskResponse.Name)

		var task tasks.Task
		err = database.DB.First(&task, TestSubjectTask.ID).Error
		assert.NoError(t, err, "Should have find the task")

		assert.Equal(t, taskResponse.ID, task.ID)
		assert.Equal(t, taskResponse.Name, task.Name)
	})

	t.Run("Complete task", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/tasks/%d", TestSubjectTask.ID), nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNoContent, resp.Code)

		var task tasks.Task
		err := database.DB.First(&task, TestSubjectTask.ID).Error
		assert.NoError(t, err, "Should have find the task")

		assert.Equal(t, TestSubjectTask.ID, task.ID)
		assert.Equal(t, true, task.Completed)
	})

	t.Run("Delete task", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/tasks/%d", TestSubjectTask.ID), nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNoContent, resp.Code)

		var task tasks.Task
		err := database.DB.First(&task, TestSubjectTask.ID).Error
		assert.Error(t, err)
	})
}
