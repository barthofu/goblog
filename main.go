package main

import (
	"blog/libs"
	"blog/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	println("Hello, World!")

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init database
	libs.InitDatabase()

	// configure gin
	router := gin.Default()

	// add routes
	routes.SetupRoutes(router)

	// run server on port 8080
	log.Println("Server running on port 8080")
	router.Run(":8080")
}
