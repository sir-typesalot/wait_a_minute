package main

import (
	"wait_a_minute/endpoints"
)

func configRoutes() {
	// Link endpoints to functions
	router.GET("/status", endpoints.GetStatus)
	// Category Endpoints
	router.GET("/category/getCategory", endpoints.CategoryByID)
	router.GET("/category/getAllCategories", endpoints.AllCategories)
	// Topic Endpoints
	router.POST("/topic/createTopic", endpoints.CreateTopic)
	router.POST("/topic/getAllTopics", endpoints.CreateTopic)
}
