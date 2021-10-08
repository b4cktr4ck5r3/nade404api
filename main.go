package main

import (
	"log"

	"github.com/b4cktr4ck5r3/nade404api/database"
	"github.com/b4cktr4ck5r3/nade404api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(":3000")
}
