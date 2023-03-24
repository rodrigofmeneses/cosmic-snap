package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigofmeneses/cosmic-snap/docs"
	"github.com/rodrigofmeneses/cosmic-snap/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/card", handler.ShowCardHandler)
		v1.POST("/card", handler.CreateCardHandler)
		v1.DELETE("/card", handler.DeleteCardHandler)
		v1.PUT("/card", handler.UpdateCardHandler)
		v1.GET("/cards", handler.ShowCardsHandler)
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
