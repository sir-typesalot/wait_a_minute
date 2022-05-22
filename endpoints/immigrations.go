package endpoints

import (
	"fmt"
	"net/http"
	"wait_a_minute/models/immigrationModel"

	"github.com/gin-gonic/gin"
)

func GetRequests(c *gin.Context) {

	data, err := immigrationModel.GetRequests("created")
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, data)
	} else {
		c.IndentedJSON(http.StatusOK, data)
	}
}
