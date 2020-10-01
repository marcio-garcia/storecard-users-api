package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

// Start - starts de application
func Start() {
	router = gin.Default()
	mapRoutes()

	fmt.Println("Starting the server")
	router.Run(":8080")
}
