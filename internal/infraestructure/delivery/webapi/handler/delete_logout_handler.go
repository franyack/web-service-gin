package handler

import (
	"example/web-service-gin/internal/business/gateway"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/utils"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/utils/apierrors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostLogoutHandler(c *gin.Context) {
	tokenMetadata, err := utils.ExtractTokenMetadata(c.Request)
	if err != nil {
		apiError := apierrors.NewUnauthorizedApiError("unauthorized")
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}
	authRepository := gateway.NewAuthRepository()
	deleted, delErr := authRepository.DeleteAuth(tokenMetadata.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		apiError := apierrors.NewUnauthorizedApiError("unauthorized")
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}

	response := map[string]string{
		"message": "Successfully logged out",
	}

	c.JSON(http.StatusOK, response)
}
