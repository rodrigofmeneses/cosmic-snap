package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigofmeneses/cosmic-snap/schemas"
)

// @BasePath /api/v1

// @Summary List cards
// @Description List all cards
// @Tags Cards
// @Produce json
// @Success 200 {object} ListCardsResponse
// @Failure 500 {object} ErrorResponse
// @Router /cards [get]
func ShowCardsHandler(ctx *gin.Context) {
	cards := []schemas.Card{}
	// Find Cards
	if err := db.Find(&cards).Error; err != nil {
		logger.Errorf("error list cards: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error list cards in database")
		return
	}
	// Response
	cardsResponse := []schemas.CardResponse{}
	for i := 0; i < len(cards); i++ {
		cardsResponse = append(cardsResponse, formatCardToResponse(cards[i]))
	}
	sendSuccess(ctx, http.StatusOK, "list-cards", cardsResponse)
}
