package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	categoryModel "wait_a_minute/backend/category"
	topicModel "wait_a_minute/backend/topic"

	"github.com/gin-gonic/gin"
)

func CategoryByName(c *gin.Context) {
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
			"AddItemURL": "/topic/new",
		})
	}
}

func AllCategories(c *gin.Context) {
	data, err := categoryModel.GetAllCategories()
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusOK, "404.html", gin.H{
			"ErrorMsg": err,
		})
	} else {
		c.HTML(http.StatusOK, "categories.html", gin.H{
			"Content": data,
			"ContentName": "Category",
			"AddItemURL": "/category/new",
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
