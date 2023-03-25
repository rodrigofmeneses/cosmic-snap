package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigofmeneses/cosmic-snap/schemas"
)

// @BasePath /api/v1

// @Summary Show card
// @Description Show a card
// @Tags Cards
// @Produce json
// @Param id query string true "Card identification"
// @Success 200 {object} ShowCardResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /card [get]
func ShowCardHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	// If not send id
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	card := schemas.Card{}
	// Find Card
	if err := db.First(&card, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("card with id: %s not found", id))
		return
	}
	// Response
	cardResponse := formatCardToResponse(card)
	sendSuccess(ctx, http.StatusOK, "show-card", &cardResponse)
}
