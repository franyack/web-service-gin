package webapi

import (
	"example/web-service-gin/internal/infraestructure/delivery"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

// New Creates a new WebAPI strategy.
func New() delivery.Strategy {
	return &webAPI{}
}

type webAPI struct{}

func (*webAPI) Start() {
	router := gin.Default()

	configureRoutes(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	configureLogFormatter()

	if err := router.Run(":" + port); err != nil {
		log.Errorf(fmt.Sprintf("error running server: %v", err))
	}

}

func configureLogFormatter() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

// Configures the application's endpoints middleware.
// Only for application's endpoints, for example, for "/ping" this middleware aren't apply.
func configureApplicationMiddleware(applicationGroup *gin.RouterGroup) {
	applicationGroup.Use(middleware.IoC())
}

func configureRoutes(router *gin.Engine) {
	healthyCheckGroup := router.Group("/ping")
	applicationGroup := router.Group("/")

	configureApplicationMiddleware(applicationGroup)

	registerHealthyCheckEndpoints(healthyCheckGroup)
	registerApplicationEndpoints(applicationGroup)
}
