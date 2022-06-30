package endpoints

import (
	"fmt"
	"net/http"
	"net/url"
	categoryModel "wait_a_minute/backend/category"
	topicModel "wait_a_minute/backend/topic"

	"github.com/gin-gonic/gin"
)

func CategoryByName(c *gin.Context) {
	// Get category name from query params
	name := c.Query("name")

	categoryData, err := categoryModel.GetCategoryByName(name, true)
	topicData, err := topicModel.GetAllTopics(categoryData.Category_ID)
	categoryList, err:= categoryModel.GetAllCategories()

	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusOK, "404.html", gin.H{
			"ErrorMsg": err,
		})
	} else {
		c.HTML(http.StatusOK, "category.html", gin.H{
			"Category": categoryData,
			"Topics": topicData,
			"ParentList": categoryList,
			"ParentName": "Category",
			"ContentName": "Topic",
			"AddItemURL": "/topic/create",
		})
	}
}

func AllCategories(c *gin.Context) {
	// Get slice of all categories
	data, err := categoryModel.GetAllCategories()
	// Check if any errors and render 
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusOK, "404.html", gin.H{
			"ErrorMsg": err,
		})
	} else {
		c.HTML(http.StatusOK, "categories.html", gin.H{
			"Content": data,
			"ContentName": "Category",
			"AddItemURL": "/category/create",
		})
	}
}

func CreateCategory(c *gin.Context) {
	// Parse form data
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	tags := c.PostForm("tags")
	// Create new category and get status
	status := categoryModel.CreateCategory(title, desc, tags)

	if status != 200 {
		c.HTML(http.StatusOK, "404.html", gin.H{
			"ErrorMsg": "",
		})
	} else {
		// Redirect to previous URL
		fmt.Println("Category Created")
		location := url.URL{Path: "/categories"}
    	c.Redirect(http.StatusFound, location.RequestURI())
	}
}
