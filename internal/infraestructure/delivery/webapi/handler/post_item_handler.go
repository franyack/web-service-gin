package handler

import (
	"example/web-service-gin/internal/business/domain"
	"example/web-service-gin/internal/business/gateway"
	"example/web-service-gin/internal/business/usecase"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/utils"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/utils/apierrors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func PostItem(c *gin.Context) {
	tokenAuth, errTokenAuth := utils.ExtractTokenMetadata(c.Request)
	if errTokenAuth != nil {
		apiError := apierrors.NewUnauthorizedApiError("unauthorized")
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}
	authRepository := gateway.NewAuthRepository()
	userId, errFetchAuth := authRepository.FetchAuth(tokenAuth)
	if errFetchAuth != nil {
		apiError := apierrors.NewUnauthorizedApiError("unauthorized")
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}

	log.Info(fmt.Sprintf("the userID is: %v", userId))
	var item domain.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		log.Error("error on shouldBindJSON object on post_item_handler ", err)
		apiError := apierrors.NewBadRequestApiError(fmt.Sprintf("it was an error with the payload. Error: %v", err.Error()))
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}

	if err := item.CheckItem(); err != nil {
		log.Error("error checking the fields of item: ", err)
		apiError := apierrors.NewBadRequestApiError(err.Error())
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}

	addItemUseCase := usecase.NewAddItemUserCase()
	err := addItemUseCase.AddItem(&item)
	if err != nil {
		log.Error("error adding item: ", err)
		apiError := err.(apierrors.ApiError)
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}

	message := domain.Message{
		Message: "The item was added successfully",
		Item:    &item,
	}
	c.JSON(http.StatusOK, message)
}
