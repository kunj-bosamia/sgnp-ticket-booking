package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kunjbosamia/sgnp-ticket-booking/internal/config"
	"github.com/kunjbosamia/sgnp-ticket-booking/internal/db"
	"github.com/kunjbosamia/sgnp-ticket-booking/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Falling back to environment variables.")
	}

	config.LoadConfig()
	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	// TODO: Add middleware and route registration
	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
