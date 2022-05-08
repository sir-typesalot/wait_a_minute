package main

import (
	"github.com/gin-gonic/gin"
)
// Declare router for server
var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()
	// Configure routes
	configRoutes()
	// Start serving the application
	router.Run()
}
