package endpoints

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	// Get current time
	now := time.Now().Format("15:04:05")
	// Create response
	resp := map[string]string {
		"status": "System Healthy",
		"code": "200",
		"time": now,
	}
    c.IndentedJSON(http.StatusOK, resp)
}