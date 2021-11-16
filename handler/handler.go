package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/uuid"

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
	rows, err := database.DB.Query("SELECT id, steam, name, score, FIND_IN_SET( score, ( SELECT GROUP_CONCAT( score ORDER BY score DESC ) FROM rankme ) ) AS rank, mvp, kills, deaths, ROUND((kills/deaths),2) as ratio, headshots, ROUND((headshots/kills) * 100, 0) as headshots_percent, assists, assist_flash, no_scope, thru_smoke, blind, wallbang FROM rankme")
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
		err := rows.Scan(&player.Id, &player.SteamID, &player.Name, &player.Score, &player.Rank, &player.Mvp, &player.Kills, &player.Deaths, &player.Ratio, &player.Headshots, &player.HeadshotsPercent, &player.Assists, &player.FlashAssists, &player.NoScope, &player.ThruSmoke, &player.Blind, &player.Wallbang)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result.Players = append(result.Players, player)
	}

	if len(result.Players) <= 0 {
		return c.Status(404).JSON(&fiber.Map{
			"succes":  false,
			"message": "No player found",
		})
	}

	response := &fiber.Map{
		"success": true,
		"players": result,
		"message": "All players returned successfully",
	}

	if err := c.JSON(response); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"succes":  false,
			"message": err,
		})
	}

	return c.Status(200).JSON(response)
}

func GetPlayerBySteamID(c *fiber.Ctx) error {
	rows, err := database.DB.Query(fmt.Sprintf("SELECT id, steam, name, score, FIND_IN_SET( score, ( SELECT GROUP_CONCAT( score ORDER BY score DESC ) FROM rankme ) ) AS rank, mvp, kills, deaths, ROUND((kills/deaths),2) as ratio, headshots, ROUND((headshots/kills) * 100, 0) as headshots_percent, assists, assist_flash, no_scope, thru_smoke, blind, wallbang FROM rankme WHERE steam = '%s'", c.Params("steam_id")))
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	defer rows.Close()
	if rows.Next() {
		player := model.Player{}
		err := rows.Scan(&player.Id, &player.SteamID, &player.Name, &player.Score, &player.Rank, &player.Mvp, &player.Kills, &player.Deaths, &player.Ratio, &player.Headshots, &player.HeadshotsPercent, &player.Assists, &player.FlashAssists, &player.NoScope, &player.ThruSmoke, &player.Blind, &player.Wallbang)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result := player
		response := &fiber.Map{
			"success": true,
			"players": result,
			"message": fmt.Sprintf("Player %s returned successfully", result.SteamID),
		}

		if err := c.JSON(response); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": err,
			})
		}
		return c.Status(200).JSON(response)

	}
	return c.Status(404).JSON(&fiber.Map{
		"success": false,
		"message": "Player not found",
	})
}

func GetTop10PlayersByKd(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, steam, name, score, FIND_IN_SET( score, ( SELECT GROUP_CONCAT( score ORDER BY score DESC ) FROM rankme ) ) AS rank, mvp, kills, deaths, ROUND((kills/deaths),2) as ratio, headshots, ROUND((headshots/kills) * 100, 0) as headshots_percent, assists, assist_flash, no_scope, thru_smoke, blind, wallbang FROM `rankme` WHERE kills > 750 ORDER BY ratio DESC LIMIT 10")
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
		err := rows.Scan(&player.Id, &player.SteamID, &player.Name, &player.Score, &player.Rank, &player.Mvp, &player.Kills, &player.Deaths, &player.Ratio, &player.Headshots, &player.HeadshotsPercent, &player.Assists, &player.FlashAssists, &player.NoScope, &player.ThruSmoke, &player.Blind, &player.Wallbang)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result.Players = append(result.Players, player)
	}

	if len(result.Players) <= 0 {
		return c.Status(404).JSON(&fiber.Map{
			"succes":  false,
			"message": "No player found for top 10 by kd",
		})
	}

	response := &fiber.Map{
		"success": true,
		"players": result,
		"message": "All players returned successfully",
	}

	if err := c.JSON(response); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"succes":  false,
			"message": err,
		})
	}

	return c.Status(200).JSON(response)
}

func GetTop10PlayersByHs(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, steam, name, score, FIND_IN_SET( score, ( SELECT GROUP_CONCAT( score ORDER BY score DESC ) FROM rankme ) ) AS rank, mvp, kills, deaths, ROUND((kills/deaths),2) as ratio, headshots, ROUND((headshots/kills) * 100, 0) as headshots_percent, assists, assist_flash, no_scope, thru_smoke, blind, wallbang FROM `rankme` WHERE kills > 750 ORDER BY headshots_percent DESC LIMIT 10")
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
		err := rows.Scan(&player.Id, &player.SteamID, &player.Name, &player.Score, &player.Rank, &player.Mvp, &player.Kills, &player.Deaths, &player.Ratio, &player.Headshots, &player.HeadshotsPercent, &player.Assists, &player.FlashAssists, &player.NoScope, &player.ThruSmoke, &player.Blind, &player.Wallbang)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result.Players = append(result.Players, player)
	}

	if len(result.Players) <= 0 {
		return c.Status(404).JSON(&fiber.Map{
			"succes":  false,
			"message": "No player found for top 10 by hs",
		})
	}

	response := &fiber.Map{
		"success": true,
		"players": result,
		"message": "All players returned successfully",
	}

	if err := c.JSON(response); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"succes":  false,
			"message": err,
		})
	}

	return c.Status(200).JSON(response)
}

func Get5Config(c *fiber.Ctx) error {
	jsonFile, err := os.Open("get5config/" + c.Params("config_id") + ".json")
	if err != nil {
		fmt.Println(err)
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "config not found",
		})
	}
	fmt.Println("Successfully Opened config ")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var matchConfig model.Get5Config
	json.Unmarshal(byteValue, &matchConfig)
	fmt.Println(matchConfig)
	return c.Status(200).JSON(matchConfig)
}

func CreateGet5Config(c *fiber.Ctx) error {
	var payload model.Get5ConfigEditPayload

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Bad body request",
		})
	}

	jsonFile, err := os.Open("base_get5config.json")
	if err != nil {
		fmt.Println(err)
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "cannot read base config",
		})
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var matchConfig model.Get5Config
	json.Unmarshal(byteValue, &matchConfig)

	matchConfig.Matchid = uuid.New().String()

	if payload.SideType != "" {
		matchConfig.SideType = payload.SideType
	}

	if len(payload.Maplist) > 0 {
		matchConfig.Maplist = payload.Maplist
	}

	if len(payload.Team1.Players) > 0 {
		matchConfig.Team1.Players = payload.Team1.Players
	}

	if len(payload.Team2.Players) > 0 {
		matchConfig.Team2.Players = payload.Team2.Players
	}

	if payload.Cvars.MpMaxrounds != "" {
		matchConfig.Cvars["mp_maxrounds"] = payload.Cvars.MpMaxrounds
	}

	if payload.Cvars.MpOvertimeEnable != "" {
		matchConfig.Cvars["mp_overtime_enable"] = payload.Cvars.MpOvertimeEnable
	}

	if payload.Cvars.SvDamagePrintEnable != "" {
		matchConfig.Cvars["sv_damage_print_enable"] = payload.Cvars.SvDamagePrintEnable
	}

	file, _ := json.MarshalIndent(matchConfig, "", " ")

	_ = ioutil.WriteFile("get5config/"+matchConfig.Matchid+".json", file, 0644)

	return c.Status(200).JSON(&fiber.Map{
		"success":  true,
		"match_id": matchConfig.Matchid,
		"content":  matchConfig,
	})
}
