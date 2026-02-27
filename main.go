package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment variables")
	}

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Run(":" + port)

}
