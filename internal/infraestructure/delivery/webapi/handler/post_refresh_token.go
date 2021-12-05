package handler

import (
	"example/web-service-gin/internal/business/domain"
	"example/web-service-gin/internal/business/gateway"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/utils"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/utils/apierrors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func PostRefreshTokenHandler(c *gin.Context) {

	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		apiError := apierrors.NewApiError(err.Error(), "unprocessable-entity", http.StatusUnprocessableEntity, nil)
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}

	refreshToken := mapToken["refresh_token"]

	//verify the token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		apiError := apierrors.NewUnauthorizedApiError("Refresh token expired")
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		apiError := apierrors.NewUnauthorizedApiError("The token is invalid")
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			apiError := apierrors.NewApiError("error converting uuid from interface to string",
				"unprocessable-entity", http.StatusUnprocessableEntity, nil)
			c.AbortWithStatusJSON(apiError.Status(), apiError)
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			apiError := apierrors.NewApiError("Error occurred",
				"unprocessable-entity", http.StatusUnprocessableEntity, nil)
			c.AbortWithStatusJSON(apiError.Status(), apiError)
			return
		}
		//Delete the previous Refresh Token
		authRepository := gateway.NewAuthRepository()
		deleted, delErr := authRepository.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			apiError := apierrors.NewUnauthorizedApiError("unauthorized")
			c.AbortWithStatusJSON(apiError.Status(), apiError)
			return
		}
		//Create new pairs of refresh and access tokens
		newTokenDetails, createErr := utils.CreateToken(userId)
		if createErr != nil {
			apiError := apierrors.NewForbiddenApiError(createErr.Error())
			c.AbortWithStatusJSON(apiError.Status(), apiError)
			return
		}
		//save the tokens metadata to redis
		saveErr := authRepository.CreateAuth(userId, newTokenDetails)
		if saveErr != nil {
			apiError := apierrors.NewForbiddenApiError(saveErr.Error())
			c.AbortWithStatusJSON(apiError.Status(), apiError)
			return
		}
		tokens := &domain.Tokens{
			AccessToken:  &newTokenDetails.AccessToken,
			RefreshToken: &newTokenDetails.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		apiError := apierrors.NewUnauthorizedApiError("refresh expired")
		c.AbortWithStatusJSON(apiError.Status(), apiError)
	}
}
