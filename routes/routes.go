package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kunjbosamia/sgnp-ticket-booking/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/whatsapp/webhook", controllers.WhatsAppWebhookHandler)
	}
}
