package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rodrigofmeneses/cosmic-snap/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
	})
}

func sendSuccess(ctx *gin.Context, code int, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("operation from handler %s successful", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
}
type CreateCardResponse struct {
	Message string               `json:"message"`
	Data    schemas.CardResponse `json:"data"`
}
type DeleteCardResponse struct {
	Message string               `json:"message"`
	Data    schemas.CardResponse `json:"data"`
}
type UpdateCardResponse struct {
	Message string               `json:"message"`
	Data    schemas.CardResponse `json:"data"`
}
type ListCardsResponse struct {
	Message string                 `json:"message"`
	Data    []schemas.CardResponse `json:"data"`
}
type ShowCardResponse struct {
	Message string               `json:"message"`
	Data    schemas.CardResponse `json:"data"`
}
