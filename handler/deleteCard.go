package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigofmeneses/cosmic-snap/schemas"
)

// @BasePath /api/v1

// @Summary Delete card
// @Description Delete a card
// @Tags Cards
// @Produce json
// @Param id query string true "Card identification"
// @Success 200 {object} DeleteCardResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /card [delete]
func DeleteCardHandler(ctx *gin.Context) {
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
	// Delete Card
	if err := db.Delete(&card, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting card with id: %s", id))
		return
	}
	sendSuccess(ctx, http.StatusOK, "delete-card", &card)
}
