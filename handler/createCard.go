package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rodrigofmeneses/cosmic-snap/schemas"
)

// @BasePath /api/v1

// @Summary Create card
// @Description Create a new card
// @Tags Cards
// @Accept json
// @Produce json
// @Param request body CreateCardRequest true "Request body"
// @Success 201 {object} CreateCardResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /card [post]
func CreateCardHandler(ctx *gin.Context) {
	request := CreateCardRequest{}

	ctx.BindJSON(&request)
	logger.Debugf("%v", request.Tags)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Data to DB

	toDbTags := strings.Join(request.Tags, ",")

	card := schemas.Card{
		Name:        request.Name,
		Cost:        *request.Cost,
		Power:       *request.Cost,
		Description: request.Description,
		Source:      request.Source,
		Image:       request.Image,
		Tags:        toDbTags,
	}

	if err := db.Create(&card).Error; err != nil {
		logger.Errorf("error creating card: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating card in database")
		return
	}

	// Data to Response

	tags := strings.Split(card.Tags, ",")

	cardResponse := schemas.CardResponse{
		Name:        card.Name,
		Power:       card.Cost,
		Cost:        card.Cost,
		Description: card.Description,
		Source:      card.Source,
		Image:       card.Image,
		Tags:        tags,
	}

	sendSuccess(ctx, http.StatusCreated, "create-card", cardResponse)
}
