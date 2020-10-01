package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping - method for checking the service is running
func Ping(context *gin.Context) {
	context.String(http.StatusOK, "pong")
}
