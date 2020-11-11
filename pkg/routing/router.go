package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"go-swag-demo-api/pkg/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		ping := new(controllers.PingController)
		v1.GET("/ping", ping.Status)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
