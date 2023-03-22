package router

import (
	"github.com/rodrigofmeneses/cosmic-snap/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/card", handler.ShowCardHandler)
		v1.POST("/card", handler.CreateCardHandler)
		v1.DELETE("/card", handler.DeleteCardHandler)
		v1.PUT("/card", handler.UpdateCardHandler)
		v1.GET("/cards", handler.ShowCardsHandler)
	}
}