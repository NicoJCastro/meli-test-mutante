package main

import (
	"log"
	"nicolascastro/go/isMutant/controllers"
	"nicolascastro/go/isMutant/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatalf("DATABASE_URL environment variable is not set")
	}

	err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	r := gin.Default()

	r.POST("/mutant/", controllers.CheckMutant)
	r.GET("/stats/", controllers.GetStats)

	port := ":8080"
	log.Printf("Server is running on port %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
