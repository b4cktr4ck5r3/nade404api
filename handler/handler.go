package handler

import (
	"github.com/b4cktr4ck5r3/nade404api/database"
	"github.com/b4cktr4ck5r3/nade404api/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func GetVersion(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"version": 1.0,
	})
}

func GetPlayers(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, steam, name, kills, deaths, (kills/deaths) as ratio, headshots, ROUND((headshots/kills) * 100, 0) as headshots_percent, assists, assist_flash FROM rankme")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	defer rows.Close()
	result := model.Players{}
	for rows.Next() {
		player := model.Player{}
		err := rows.Scan(&player.Id, &player.SteamID, &player.Name, &player.Kills, &player.Deaths, &player.Ratio, &player.Headshots, &player.HeadshotsPercent, &player.Assists, &player.FlashAssists)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result.Players = append(result.Players, player)
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"players": result,
		"message": "All players returned successfully",
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"succes":  false,
			"message": err,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"players": result,
		"message": "All players returned successfully",
	})
}
