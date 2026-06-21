//go:build integration

package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matheus-rib/golang-crud/config"
	"github.com/stretchr/testify/assert"
)

func Test_HealthCheckRoute(t *testing.T) {
	router := config.SetupRouter()

	t.Run("Health check route", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/health-check", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})
}
