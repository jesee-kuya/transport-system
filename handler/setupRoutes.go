package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *Transport) SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(t.Middleware.RouteChecker())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return router
}
