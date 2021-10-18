package model

type Player struct {
	Id               int      `json:"id"`
	SteamID          string   `json:"steam_id"`
	Name             string   `json:"name"`
	Score            *int     `json:"score"`
	Rank             *int     `json:"rank"`
	Mvp              *int     `json:"mvp"`
	Kills            *int     `json:"kills"`
	Deaths           *int     `json:"deaths"`
	Ratio            *float32 `json:"ratio"`
	Headshots        *int     `json:"headshots"`
	HeadshotsPercent *int     `json:"headshots_percent"`
	Assists          *int     `json:"assists"`
	FlashAssists     *int     `json:"flash_assists"`
	NoScope          *int     `json:"no_scope"`
	ThruSmoke        *int     `json:"thru_smoke"`
	Blind            *int     `json:"blind"`
	Wallbang         *int     `json:"wallbang"`
}

type Players struct {
	Players []Player `json:"players"`
}
