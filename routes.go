package main

import (
	"wait_a_minute/endpoints"
)

func configRoutes() {
	// Link endpoints to functions
	// Simple status check
	router.GET("/status", endpoints.GetStatus)
	// Category Endpoints
	router.GET("/category/get", endpoints.CategoryByID)
	router.GET("/category/getAll", endpoints.AllCategories)
	// router.GET("/category/create", endpoints.AllCategories)

	// Topic Endpoints
	router.POST("/topic/create", endpoints.CreateTopic)
	router.GET("/topic/get", endpoints.GetTopics)
	// Pointer Endpoints
	router.GET("/pointer/get", endpoints.GetPointers)
	router.POST("/pointer/create", endpoints.CreatePointer)
	// Request Endpoints
	router.GET("/requests/get", endpoints.GetRequests)
	router.POST("/requests/approve", endpoints.ApproveReq)
}
