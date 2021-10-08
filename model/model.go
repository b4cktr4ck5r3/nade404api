package model

type Player struct {
	Id               int      `json:"id"`
	SteamID          string   `json:"steam_id"`
	Name             string   `json:"name"`
	Kills            *int     `json:"kills"`
	Deaths           *int     `json:"deaths"`
	Ratio            *float32 `json:"ratio"`
	Headshots        *int     `json:"headshots"`
	HeadshotsPercent *int     `json:"headshots_percent"`
	Assists          *int     `json:"assists"`
	FlashAssists     *int     `json:"flash_assists"`
}

type Players struct {
	Players []Player `json:"players"`
}
