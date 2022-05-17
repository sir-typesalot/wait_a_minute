package endpoints

import (
	"net/http"
	"wait_a_minute/models"
	"github.com/gin-gonic/gin"
)

func TestModel(c *gin.Context) {
	data := models.TestIt()
    c.IndentedJSON(http.StatusOK, data)
}