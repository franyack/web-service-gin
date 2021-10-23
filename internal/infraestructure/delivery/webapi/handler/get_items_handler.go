package handler

import (
	"example/web-service-gin/internal/business/usecase"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/utils/apierrors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetItems(c *gin.Context) {

	itemsUseCase := usecase.NewGetItemsUserCase()
	items, err := itemsUseCase.GetItems()
	if err != nil {
		log.Error("error obtaining items: ", err)
		apiError := err.(apierrors.ApiError)
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}
	c.JSON(http.StatusOK, items)
}
