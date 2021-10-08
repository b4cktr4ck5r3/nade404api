package router

import (
	"github.com/b4cktr4ck5r3/nade404api/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/version", handler.GetVersion)
	api.Get("/players", handler.GetPlayers)
	api.Get("/players/:steam_id", handler.GetPlayerBySteamID)
}
