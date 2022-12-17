package main

import (
	"log"
	"os"

	"github.com/zhdlxh48/leader-board-server/database"
	"github.com/zhdlxh48/leader-board-server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error loading .env file")
	}

	database.Initialize()

	app := fiber.New()

	app.Get("/", routes.GetScore)
	app.Post("/", routes.SaveScore)

	app.Get("/rank", routes.GetRanks)
	app.Get("/rank/:user", routes.GetUserRank)

	app.Listen(getPort())
}
