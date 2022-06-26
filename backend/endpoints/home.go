package endpoints

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"content": "This is the home page",
	})
}
