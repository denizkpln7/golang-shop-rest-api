package main

import (
	"denizkpln7/rest-api/database"
	"denizkpln7/rest-api/routesModule"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	database.Connect()

	port := os.Getenv("PORT")
	app := fiber.New()
	var route routesModule.RouteModule
	route.Setup(app)
	app.Listen(":" + port)
}
