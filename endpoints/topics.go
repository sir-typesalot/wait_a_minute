package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wait_a_minute/models/topicModel"

	"github.com/gin-gonic/gin"
)

type Vars struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Tags string `json:"tags"`
	CategoryName string `json:"categoryName"`
}

func CreateTopic(c *gin.Context) {
	// data, err := categoryModel.GetAllCategories()
	var vars Vars
	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(&vars)
	
	status := topicModel.CreateNewTopic(vars.Name, vars.Desc, vars.Tags, vars.CategoryName)
	if status == 200 {
		fmt.Println("Topic Created")
		c.IndentedJSON(http.StatusOK, "Topic Created")
	} else {
		c.IndentedJSON(http.StatusInternalServerError, "Error creating Topic")
	}
}
