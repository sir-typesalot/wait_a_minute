package endpoints

import (
    "net/http"
    "github.com/gin-gonic/gin"
)
func TestRoute(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "Welcome Here!")
}