package main

import (
	"github.com/gin-gonic/gin"
	"wait_a_minute/endpoints"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()
	router.GET("/test", endpoints.TestRoute)

	// Start serving the application
	router.Run()
}
