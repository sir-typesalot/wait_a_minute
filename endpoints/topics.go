package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wait_a_minute/models/topicModel"

	"github.com/gin-gonic/gin"
)

type TopicArgs struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Tags string `json:"tags"`
	CategoryName string `json:"categoryName"`
}
func GetTopics(c *gin.Context) {

	categoryID := c.Query("categoryID")
	
	var catID_int int
	if categoryID != "" {
		catID_int, _ = strconv.Atoi(categoryID)
	} else {
		catID_int = 0
	}

	data, err := topicModel.GetAllTopics(catID_int)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, data)
	} else {
		c.IndentedJSON(http.StatusOK, data)
	}
}

func CreateTopic(c *gin.Context) {

	var vars TopicArgs
	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(&vars)
	
	status := topicModel.CreateNewTopic(vars.Name, vars.Desc, vars.Tags, vars.CategoryName)
	// TODO : Get some better data as a response here
	if status == 200 {
		fmt.Println("Topic Created")
		c.IndentedJSON(http.StatusOK, "Topic Created")
	} else {
		c.IndentedJSON(http.StatusInternalServerError, "Error creating Topic")
	}
}
