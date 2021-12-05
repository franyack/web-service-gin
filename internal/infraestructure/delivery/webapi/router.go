package webapi

import (
	"example/web-service-gin/internal/infraestructure/delivery/webapi/handler"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/middleware"
	"github.com/gin-gonic/gin"
)

// Configure the HealthyChecks application's endpoints.
func registerHealthyCheckEndpoints(diagnostic *gin.RouterGroup) {
	diagnostic.GET("", handler.PingHandler)
}

// In this func you can register the application's endpoints.
func registerApplicationEndpoints(appEndpoints *gin.RouterGroup) {
	appEndpoints.POST("/login", handler.PostLogin)
	appEndpoints.POST("/logout", middleware.TokenAuthMiddleware(), handler.PostLogoutHandler)

	appEndpoints.POST("/token/refresh", handler.PostRefreshTokenHandler)

	appEndpoints.GET("/items", handler.GetItems)
	appEndpoints.GET("/items/:item_id", handler.GetItemById)
	appEndpoints.POST("/items", middleware.TokenAuthMiddleware(), handler.PostItem)
}
