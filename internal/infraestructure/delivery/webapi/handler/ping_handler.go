package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingHandler returns a successful pong answer to all HTTP requests.
func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
