package main

import (
	"github.com/burdianov/go-admin/database"
	"github.com/burdianov/go-admin/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen("127.0.0.1:8000")
}
