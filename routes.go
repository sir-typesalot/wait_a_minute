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
	router.POST("/category/create", endpoints.CreateCategory)
	// Topic Endpoints
	router.POST("/topic/create", endpoints.CreateTopic)
	router.GET("/topics", endpoints.GetTopics)
	router.GET("/topic/get", endpoints.GetTopic)
	// Pointer Endpoints
	router.GET("/pointer/get", endpoints.GetPointers)
	router.POST("/pointer/create", endpoints.CreatePointer)

	// Admin Level Endpoints
	// Login Endpoints
	// TODO: NEED ENDPOINT HANDLERS FOR THESE 
	router.GET("/login", endpoints.GetPointers)
	router.POST("/login/auth", endpoints.CreatePointer)
	// Request Endpoints
	router.GET("/requests/get", endpoints.GetRequests)
	// router.GET("/requests/edit", endpoints.GetRequests)
	router.POST("/requests/approve", endpoints.ApproveReq)
}
