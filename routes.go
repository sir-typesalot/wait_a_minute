package main

import (
	"wait_a_minute/endpoints"
)

func configRoutes() {
	// Link endpoints to functions
	router.GET("/status", endpoints.GetStatus)
	router.GET("/test", endpoints.TestModel)
}