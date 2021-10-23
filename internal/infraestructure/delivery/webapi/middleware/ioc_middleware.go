package middleware

import (
	"example/web-service-gin/internal/business/gateway"
	"example/web-service-gin/internal/infraestructure/repository"
	"github.com/gin-gonic/gin"
)

// Register all your useCases, repositories and services in this middleware.
// This is the first middleware. Show webapi.Start()
func IoC() gin.HandlerFunc {
	return func(c *gin.Context) {
		gateway.RegisterItemsRepository(repository.NewMySqlItemsRepository())
	}
}
