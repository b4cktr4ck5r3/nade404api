package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/b4cktr4ck5r3/nade404api/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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
