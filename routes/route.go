package router

import (
	"github.com/charan678/go-rest-example/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/v1/health", handlers.Health)
	router.POST("/v1/example/", handlers.Example)
	return router
}
