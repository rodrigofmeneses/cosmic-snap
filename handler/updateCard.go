package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigofmeneses/cosmic-snap/schemas"
)

// @BasePath /api/v1

// @Summary Update card
// @Description Update a card
// @Tags Cards
// @Accept json
// @Produce json
// @Param id query string true "Card identification"
// @Param request body UpdateCardRequest true "Request body"
// @Success 200 {object} UpdateCardResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /card [put]
func UpdateCardHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	request := UpdateCardRequest{}
	ctx.BindJSON(&request)

	// If not send id
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	// If body valid
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	card := schemas.Card{}
	// Find card
	if err := db.First(&card, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "card not found")
		return
	}
	// Update card
	if request.Name != "" {
		card.Name = request.Name
	}
	if request.Cost != nil {
		card.Cost = *request.Cost
	}
	if request.Power != nil {
		card.Power = *request.Power
	}
	if request.Description != "" {
		card.Description = request.Description
	}
	if request.Source != "" {
		card.Source = request.Source
	}
	if request.Image != "" {
		card.Image = request.Image
	}
	// Save card
	if err := db.Save(&card).Error; err != nil {
		logger.Errorf("error updating card: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating card")
		return
	}
	sendSuccess(ctx, http.StatusOK, "update-card", card)
}
