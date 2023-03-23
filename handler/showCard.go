package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigofmeneses/cosmic-snap/schemas"
)

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
	sendSuccess(ctx, http.StatusOK, "show-card", &card)
}
