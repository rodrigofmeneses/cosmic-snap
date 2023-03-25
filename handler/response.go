package handler

import (
	"fmt"
	"strings"

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

// Utils

func formatCardToResponse(card schemas.Card) schemas.CardResponse {
	tags := strings.Split(card.Tags, ",")

	cardResponse := schemas.CardResponse{
		ID:          card.ID,
		CreatedAt:   card.CreatedAt,
		UpdatedAt:   card.UpdatedAt,
		DeletedAt:   card.DeletedAt.Time,
		Name:        card.Name,
		Power:       card.Power,
		Cost:        card.Cost,
		Description: card.Description,
		Source:      card.Source,
		Image:       card.Image,
		Tags:        tags,
	}
	return cardResponse
}

// Docs

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
