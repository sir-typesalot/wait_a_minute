package main

import (
	"wait_a_minute/endpoints"
)

func configRoutes() {
	// Link endpoints to functions
	// Simple status check
	router.GET("/status", endpoints.GetStatus)
	// Category Endpoints
	router.GET("/category/getCategory", endpoints.CategoryByID)
	router.GET("/category/getAllCategories", endpoints.AllCategories)
	// Topic Endpoints
	router.POST("/topic/createTopic", endpoints.CreateTopic)
	router.GET("/topic/getAllTopics", endpoints.CreateTopic)
	// Pointer Endpoints
	router.GET("/pointer/getPointers", endpoints.GetPointers)
	router.POST("/pointer/createPointer", endpoints.CreatePointer)
	// Request Endpoints
	router.GET("/requests/getRequests", endpoints.GetRequests)
	router.POST("/requests/approveRequest", endpoints.ApproveReq)
}
