package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kunjbosamia/sgnp-ticket-booking/internal/constants"
	"github.com/kunjbosamia/sgnp-ticket-booking/internal/services"
	"github.com/kunjbosamia/sgnp-ticket-booking/internal/utils"
)

func WhatsAppWebhookHandler(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.BindJSON(&payload); err != nil {
		utils.RespondWithError(c, constants.ErrInvalidPayload)
		return
	}

	err := services.HandleWhatsAppWebhook(c.Request.Context(), payload)
	if err != nil {
		utils.RespondWithError(c, constants.ErrInternal)
		return
	}

	c.JSON(200, gin.H{"message": "Webhook processed successfully"})
}
