package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "alive"})
}
