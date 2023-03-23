package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigofmeneses/cosmic-snap/schemas"
)

func ShowCardsHandler(ctx *gin.Context) {
	cards := []schemas.Card{}
	// Find Cards
	if err := db.Find(&cards).Error; err != nil {
		logger.Errorf("error list cards: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error list cards in database")
		return
	}
	sendSuccess(ctx, http.StatusOK, "list-cards", cards)
}
