package webapi

import (
	"example/web-service-gin/internal/infraestructure/delivery/webapi/handler"
	"github.com/gin-gonic/gin"
)

// Configure the HealthyChecks application's endpoints.
func registerHealthyCheckEndpoints(diagnostic *gin.RouterGroup) {
	diagnostic.GET("", handler.PingHandler)
}

// In this func you can register the application's endpoints.
func registerApplicationEndpoints(appEndpoints *gin.RouterGroup) {
	appEndpoints.GET("/ping", handler.PingHandler)
	appEndpoints.GET("/items", handler.GetItems)
	appEndpoints.GET("/items/:item_id", handler.GetItemById)
}