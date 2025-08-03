package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kunjbosamia/sgnp-ticket-booking/controllers"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/whatsapp/webhook", controllers.WhatsAppWebhookVerifyHandler)
		v1.POST("/whatsapp/webhook", controllers.WhatsAppWebhookHandler)
	}
}
