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

type Get5Config struct {
	Matchid              string            `json:"matchid"`
	NumMaps              int64             `json:"num_maps"`
	PlayersPerTeam       int64             `json:"players_per_team"`
	MinPlayersToReady    int64             `json:"min_players_to_ready"`
	MinSpectatorsToReady int64             `json:"min_spectators_to_ready"`
	SkipVeto             bool              `json:"skip_veto"`
	SideType             string            `json:"side_type"`
	Maplist              []string          `json:"maplist"`
	Team1                Team              `json:"team1"`
	Team2                Team              `json:"team2"`
	Cvars                map[string]string `json:"cvars"`
}

type Team struct {
	Name    string   `json:"name"`
	Tag     string   `json:"tag"`
	Flag    string   `json:"flag"`
	Logo    string   `json:"logo"`
	Players []string `json:"players"`
}

type Get5ConfigEditPayload struct {
	SideType string         `json:"side_type"`
	Maplist  []string       `json:"maplist"`
	Team1    PlayersPayload `json:"team1"`
	Team2    PlayersPayload `json:"team2"`
	Cvars    CvarsPayload   `json:"cvars"`
}

type PlayersPayload struct {
	Players []string `json:"players"`
}

type CvarsPayload struct {
	MpMaxrounds         string `json:"mp_maxrounds"`
	MpOvertimeEnable    string `json:"mp_overtime_enable"`
	SvDamagePrintEnable string `json:"sv_damage_print_enable"`
}
