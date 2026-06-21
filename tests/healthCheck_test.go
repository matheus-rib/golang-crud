//go:build integration

package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matheus-rib/golang-crud/config"
	"github.com/stretchr/testify/assert"
)

type DataHealthCheckResponse struct {
	Data string `json:"data"`
}

func Test_HealthCheckRoute(t *testing.T) {
	router := config.SetupRouter()

	t.Run("Health check route", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/health-check", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		var response DataHealthCheckResponse
		err := json.Unmarshal(resp.Body.Bytes(), &response)

		assert.NoError(t, err, "Response body should be valid JSON")
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "alive", response.Data)
	})
}
