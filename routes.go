package main

import (
	"wait_a_minute/backend/endpoints"
)

func configRoutes() {
	// Link endpoints to functions
	// Simple status check
	router.GET("/", endpoints.HomePage)
	router.GET("/status", endpoints.GetStatus)
	// Category Endpoints
	router.GET("/category/get", endpoints.CategoryByName)
	router.GET("/categories", endpoints.AllCategories)
	router.GET("/category/create", endpoints.CreateCategory)

	// Topic Endpoints
	router.POST("/topic/create", endpoints.CreateTopic)
	router.GET("/topic/get-all", endpoints.GetTopics)
	router.GET("/topic/get", endpoints.GetTopic)
	// NEED A GET 1 TOPIC ENDPOINT

	// Pointer Endpoints
	router.GET("/pointer/get", endpoints.GetPointers)
	router.POST("/pointer/create", endpoints.CreatePointer)
	// Request Endpoints
	router.GET("/requests/get", endpoints.GetRequests)
	router.POST("/requests/approve", endpoints.ApproveReq)
}
