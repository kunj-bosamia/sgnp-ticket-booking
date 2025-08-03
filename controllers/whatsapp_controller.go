package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /api/v1/whatsapp/webhook
func WhatsAppWebhookVerifyHandler(c *gin.Context) {
	mode := c.Query("hub.mode")
	token := c.Query("hub.verify_token")
	challenge := c.Query("hub.challenge")

	verifyToken := os.Getenv("VERIFY_TOKEN")

	if mode == "subscribe" && token == verifyToken {
		fmt.Println("WEBHOOK VERIFIED")
		c.String(http.StatusOK, challenge)
		return
	}

	c.Status(http.StatusForbidden)
}

// POST /api/v1/whatsapp/webhook
func WhatsAppWebhookHandler(c *gin.Context) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("\n\nWebhook received %s\n", timestamp)

	var body map[string]interface{}
	if err := c.BindJSON(&body); err != nil {
		fmt.Println("Error parsing request body:", err)
		c.Status(http.StatusBadRequest)
		return
	}

	fmt.Println("Payload:\n", body)
	c.Status(http.StatusOK)
}
