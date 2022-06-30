package endpoints

import (
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

	title := c.PostForm("title")
	desc := c.PostForm("desc")
	tags := c.PostForm("tags")
	parent := c.PostForm("parentContent")
	status := pointerModel.CreateNewPointer(title, desc, tags, parent)
	if status == 200 {
		fmt.Println("Topic Created")
		c.IndentedJSON(http.StatusOK, "Topic Created")
	} else {
		c.IndentedJSON(http.StatusInternalServerError, "Error creating Topic")
	}
}
