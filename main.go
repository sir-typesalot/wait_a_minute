package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Declare router for server
var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	router.Static("frontend/assets", "./assets")
	router.LoadHTMLGlob("frontend/templates/*.html")

	// Configure routes
	configRoutes()
	fmt.Println("Configured Routes...")
	// Start serving the application
	fmt.Println("Running on port 8080")
	router.Run(":8080")
}
