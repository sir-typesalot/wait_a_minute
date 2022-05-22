package endpoints

import (
	"fmt"
	"net/http"
	"wait_a_minute/models/categoryModel"

	"github.com/gin-gonic/gin"
)

func CategoryByID(c *gin.Context) {
	// data, err := categoryModel.GetAllCategories()
	name := c.Query("name")

	data, err := categoryModel.GetCategoryByID(name, true)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, data)
	} else {
		c.IndentedJSON(http.StatusOK, data)
	}
}

func AllCategories(c *gin.Context) {
	data, err := categoryModel.GetAllCategories()
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, data)
	} else {
		c.IndentedJSON(http.StatusOK, data)
	}
}
