package handler

import (
	"example/web-service-gin/internal/business/domain"
	"example/web-service-gin/internal/business/gateway"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/utils"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/utils/apierrors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostLogin(c *gin.Context) {
	user := domain.User{
		Id:       1,
		Username: "username",
		Password: "password",
	}

	var u *domain.User
	if err := c.ShouldBindJSON(&u); err != nil {
		apiError := apierrors.NewApiError("Invalid json provided",
			"unprocessable-entity",
			http.StatusUnprocessableEntity,
			nil)
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}
	//TODO: post_login_usecase-> Validate User
	if user.Username != u.Username || user.Password != u.Password {
		apiError := apierrors.NewApiError("Please provide valid login details",
			"unprocessable-entity",
			http.StatusUnprocessableEntity,
			nil)
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}
	tokenDetails, err := utils.CreateToken(u.Id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	authRepository := gateway.NewAuthRepository()
	if saveErr := authRepository.CreateAuth(user.Id, tokenDetails); saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	tokens := domain.Tokens{
		AccessToken:  &tokenDetails.AccessToken,
		RefreshToken: &tokenDetails.RefreshToken,
	}
	c.JSON(http.StatusOK, &tokens)
}
