package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wait_a_minute/backend/pointer"

	"github.com/gin-gonic/gin"
)

type PointerArgs struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Tags string `json:"tags"`
	TopicID int `json:"topicID"`
}

func GetPointers(c *gin.Context) {

	topicID := c.Query("topicID")
	topicID_int, _ := strconv.Atoi(topicID)

	data, err := pointerModel.GetPointers(topicID_int)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, data)
	} else {
		c.IndentedJSON(http.StatusOK, data)
	}
}

func CreatePointer(c *gin.Context) {

	var vars PointerArgs
	decoder := json.NewDecoder(c.Request.Body)
	decoder.Decode(&vars)
	
	status := pointerModel.CreateNewPointer(vars.Name, vars.Desc, vars.Tags, vars.TopicID)
	if status == 200 {
		fmt.Println("Topic Created")
		c.IndentedJSON(http.StatusOK, "Topic Created")
	} else {
		c.IndentedJSON(http.StatusInternalServerError, "Error creating Topic")
	}
}
