package endpoints

import (
	"fmt"
	"net/http"
	"net/url"
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

	// Get a list of all topics
	data, err := topicModel.GetAllTopics(catID_int)
	// Get a list of all the categories
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
			"AddItemURL": "/topic/create",
		})
	}
}

func CreateTopic(c *gin.Context) {

	// Parse results from the form submission
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	tags := c.PostForm("tags")
	parent := c.PostForm("parentContent")
	// Create new topic and get status
	status := topicModel.CreateNewTopic(title, desc, tags, parent)

	if status != 200 {
		c.HTML(http.StatusOK, "404.html", gin.H{
			"ErrorMsg": "",
		})
	} else {
		// Redirect back to previous URL
		fmt.Println("Category Created")
		location := url.URL{Path: "/topics"}
    	c.Redirect(http.StatusFound, location.RequestURI())
	}
}

func GetTopic(c *gin.Context) {
	// Get topic from query params
	topicID := c.Query("topicID")
	// Check to convert topic id
	var topicID_int int
	if topicID != "" {
		topicID_int, _ = strconv.Atoi(topicID)
	}
	// Get topic data and pointers slice for topic
	topicData, err := topicModel.GetOneTopic(topicID_int)
	pointerData, err := pointerModel.GetPointers(topicID_int)
	// Get slice of topics for creation of new pointer form
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
			"AddItemURL": "/pointer/create",
		})
	}
}
