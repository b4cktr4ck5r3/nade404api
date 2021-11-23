package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

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

func HandleGet5ConfigLogs(c *fiber.Ctx) error {
	if strings.Contains(string(c.Body()), "get5_event") {
		var msg json.RawMessage
		event := model.Get5Event{
			Params: &msg,
		}

		if err := json.Unmarshal([]byte(toto(string(c.Body()))), &event); err != nil {
			fmt.Println(err)
		}

		switch event.Event {
		case "series_start":
			var p model.SeriesStart
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "series_end":
			var p model.SeriesEnd
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "series_cancel":
			var p model.SeriesCancel
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "map_veto":
			var p model.MapVeto
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "map_pick":
			var p model.MapPick
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "side_picked":
			var p model.SidePicked
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "knife_start":
			var p model.KnifeStart
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "knife_won":
			var p model.KnifeWon
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "going_live":
			var p model.GoingLive
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "round_end":
			var p model.RoundEnd
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "side_swap":
			var p model.SideSwap
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "map_end":
			var p model.MapEnd
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "player_death":
			var p model.PlayerDeath
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "bomb_planted":
			var p model.BombPlanted
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "bomb_defused":
			var p model.BombDefused
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "bomb_exploded":
			var p model.BombExploded
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "client_say":
			var p model.ClientSay
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "player_connect":
			var p model.PlayerConnect
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "player_disconnect":
			var p model.PlayerDisconnect
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "match_config_load_fail":
			var p model.MatchConfigLoadFailed
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "backup_loaded":
			var p model.BackupLoad
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "team_ready":
			var p model.TeamReady
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		case "team_unready":
			var p model.TeamUnready
			if err := json.Unmarshal(msg, &p); err != nil {
				fmt.Println(err)
			} else {
				event.Params = p
			}
		default:
			fmt.Println("unknown event type: " + event.Event)
		}

		fmt.Println(prettyPrint(event))

	}

	return c.Status(200).JSON(string(c.Body()))
}

func toto(str string) string {
	var sb strings.Builder
	braceLeftCount := 0
	braceRightCount := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '{' || braceLeftCount > 0 {
			sb.WriteRune(rune(str[i]))
			if str[i] == '{' {
				braceLeftCount++
			} else if str[i] == '}' {
				braceRightCount++
			}

			if len(sb.String()) > 0 && braceLeftCount == braceRightCount {
				break
			}
		}
	}

	return sb.String()
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func GetPteroServerList(c *fiber.Ctx) error {
	serverNotFound := true
	client := http.Client{}
	url := "https://p.ezstrat.com/api/client/"

	for serverNotFound {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": err,
			})
		}

		req.Header = http.Header{
			"Accept":        []string{"application/json"},
			"Authorization": []string{"Bearer odexNvfil7D21kXHc3UD9xRa5xFQOJ2PgVU74IwsZ0uV6OJK"},
		}

		res, err := client.Do(req)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": err,
			})
		}

		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			t := new(model.PteroServerList)
			if err := json.NewDecoder(res.Body).Decode(t); err != nil {
				return c.Status(500).JSON(&fiber.Map{
					"succes":  false,
					"message": err,
				})
			}

			server, err := FindServerInList(t)
			if err != nil {
				if t.Meta.Pagination.CurrentPage < t.Meta.Pagination.TotalPages {
					url = t.Meta.Pagination.Links.Next
				} else {
					serverNotFound = false
				}

			} else {
				serverNotFound = false
				return c.Status(500).JSON(&fiber.Map{
					"succes":      true,
					"message":     "Server found",
					"serverInfos": server,
				})
			}
		} else {
			return c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": "Request didn't work",
			})
		}

	}

	return c.Status(200).JSON("null")
}

func FindServerInList(t *model.PteroServerList) (model.PteroServerListDatum, error) {
	fmt.Println("CALL FINDSERVER")
	for i := 0; i < len(t.Data); i++ {
		data := t.Data[i]
		if len(data.Attributes.Relationships.Allocations.Data) > 0 {
			ip := data.Attributes.Relationships.Allocations.Data[0].Attributes.IP
			port := data.Attributes.Relationships.Allocations.Data[0].Attributes.Port
			if ip == "51.158.82.97" && port == 27015 {
				return t.Data[i], nil
			}
		}
	}

	return model.PteroServerListDatum{}, errors.New("Server not found")
}

func GetPteroServerDetails(c *fiber.Ctx) error {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://p.ezstrat.com/api/client/servers/"+c.Params("server_id"), nil)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"succes":  false,
			"message": err,
		})
	}

	req.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer odexNvfil7D21kXHc3UD9xRa5xFQOJ2PgVU74IwsZ0uV6OJK"},
	}

	res, err := client.Do(req)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"succes":  false,
			"message": err,
		})
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		t := new(model.PteroServerListDatum)
		if err := json.NewDecoder(res.Body).Decode(t); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": err,
			})
		}
		fmt.Println(t.Attributes.Relationships.Allocations)
	}
	fmt.Println(res)
	return c.Status(200).JSON("null")
}
