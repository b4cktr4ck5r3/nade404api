package handler

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/b4cktr4ck5r3/nade404api/model"
	"github.com/gofiber/fiber/v2"
)

func HandleGet5ConfigLogs(c *fiber.Ctx) error {
	if strings.Contains(string(c.Body()), "get5_event") {
		var msg json.RawMessage
		event := model.Get5Event{
			Params: &msg,
		}

		if err := json.Unmarshal([]byte(ParseStringEventToJSON(string(c.Body()))), &event); err != nil {
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

		fmt.Println(PrettyPrint(event))

	}

	return c.Status(200).JSON(string(c.Body()))
}

func ParseStringEventToJSON(str string) string {
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

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
