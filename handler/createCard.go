package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigofmeneses/cosmic-snap/schemas"
)

func CreateCardHandler(ctx *gin.Context) {
	request := CreateCardRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	card := schemas.Card{
		Name:        request.Name,
		Cost:        *request.Cost,
		Power:       *request.Cost,
		Description: request.Description,
		Source:      request.Source,
		Image:       request.Image,
	}

	if err := db.Create(&card).Error; err != nil {
		logger.Errorf("error creating card: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating card in database")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create-card", card)
}
