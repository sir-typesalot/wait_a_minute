package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wait_a_minute/backend/category"

	"github.com/gin-gonic/gin"
)

func CategoryByName(c *gin.Context) {
	name := c.Query("name")

	data, err := categoryModel.GetCategoryByName(name, true)
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
		c.HTML(http.StatusOK, "categories.html", gin.H{
			"Content": data,
		})
	}
}

func CreateCategory(c *gin.Context) {

	var vars PointerArgs
	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(&vars)
	
	status := categoryModel.CreateCategory(vars.Name, vars.Desc, vars.Tags)
	if status == 200 {
		fmt.Println("Category Created")
		c.IndentedJSON(http.StatusOK, "Category Created")
	} else {
		c.IndentedJSON(http.StatusInternalServerError, "Error creating Category")
	}
}
