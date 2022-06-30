package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	categoryModel "wait_a_minute/backend/category"
	pointerModel "wait_a_minute/backend/pointer"
	"wait_a_minute/backend/topic"

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
	categoryList, err:= categoryModel.GetAllCategories()

	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusOK, "404.html", gin.H{
			"ErrorMsg": err,
		})
	} else {
		c.HTML(http.StatusOK, "topics.html", gin.H{
			"Content": data,
			"ParentList": categoryList,
			"ParentName": "Category",
			"ContentName": "Topic",
			"AddItemURL": "/topic/new",
		})
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

func GetTopic(c *gin.Context) {
	topicID := c.Query("topicID")
	
	var topicID_int int
	if topicID != "" {
		topicID_int, _ = strconv.Atoi(topicID)
	}

	topicData, err := topicModel.GetOneTopic(topicID_int)
	pointerData, err := pointerModel.GetPointers(topicID_int)
	topicList, err := topicModel.GetAllTopics(0)

	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusOK, "404.html", gin.H{
			"ErrorMsg": err,
		})
	} else {
		c.HTML(http.StatusOK, "topic.html", gin.H{
			"Topic": topicData,
			"ParentList": topicList,
			"ParentName": "Topic",
			"Pointers": pointerData,
			"ContentName": "Pointer",
			"AddItemURL": "/pointer/new",
		})
	}
}
