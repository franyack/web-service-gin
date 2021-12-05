package middleware

import (
	"example/web-service-gin/internal/infraestructure/delivery/webapi/utils"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/utils/apierrors"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c.Request)
		if err != nil {
			apiError := apierrors.NewUnauthorizedApiError(err.Error())
			c.AbortWithStatusJSON(apiError.Status(), apiError)
			return
		}
		c.Next()
	}
}
