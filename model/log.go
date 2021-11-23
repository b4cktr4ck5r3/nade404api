package model

type Get5Event struct {
	Matchid string      `json:"matchid"`
	Params  interface{} `json:"params"`
	Event   string      `json:"event"`
}

type SeriesStart struct {
	Team1Name string `json:"team1_name"`
	Team2Name string `json:"team2_name"`
}

type SeriesEnd struct {
	Team1Score string `json:"team1_series_score"`
	Team2Score string `json:"team2_series_score"`
	Winner     string `json:"winner"`
}

type SeriesCancel struct {
	Team1Score string `json:"team1_series_score"`
	Team2Score string `json:"team2_series_score"`
}

type MapVeto struct {
	Team    string `json:"team"`
	MapName string `json:"map_name"`
}

type MapPick struct {
	MapName   string `json:"map_name"`
	Team      string `json:"team"`
	MapNumber int    `json:"map_number"`
}

type SidePicked struct {
	MapName   string `json:"map_name"`
	Team      string `json:"team"`
	MapNumber int    `json:"map_number"`
	Side      string `json:"side"`
}

type KnifeStart struct {
	MapName   string `json:"map_name"`
	MapNumber int    `json:"map_number"`
}

type KnifeWon struct {
	MapName      string `json:"map_name"`
	MapNumber    int    `json:"map_number"`
	Winner       string `json:"winner"`
	SelectedSide string `json:"selected_side"`
}

type GoingLive struct {
	MapName   string `json:"map_name"`
	MapNumber int    `json:"map_number"`
}

type RoundEnd struct {
	MapName    string `json:"map_name"`
	MapNumber  int    `json:"map_number"`
	WinnerSide string `json:"winner_side"`
	Winner     string `json:"winner"`
	Team1Score string `json:"team1_score"`
	Team2Score string `json:"team2_score"`
	Reason     string `json:"reason"`
}

type SideSwap struct {
	MapName    string `json:"map_name"`
	MapNumber  int    `json:"map_number"`
	Team1Side  string `json:"team1_side"`
	Team2Side  string `json:"team2_side"`
	Team1Score string `json:"team1_score"`
	Team2Score string `json:"team2_score"`
}

type MapEnd struct {
	MapName    string `json:"map_name"`
	MapNumber  int    `json:"map_number"`
	Winner     string `json:"winner"`
	Team1Score string `json:"team1_score"`
	Team2Score string `json:"team2_score"`
}

type PlayerDeath struct {
	MapName       string `json:"map_name"`
	MapNumber     int    `json:"map_number"`
	Attacker      string `json:"attacker"`
	Victim        string `json:"victim"`
	Headshot      int    `json:"headshot"`
	Weapon        string `json:"weapon"`
	Assister      string `json:"assister"`
	FlashAssister string `json:"flash_assister"`
}

type BombPlanted struct {
	MapName   string `json:"map_name"`
	MapNumber int    `json:"map_number"`
	Client    string `json:"client"`
	Site      int    `json:"site"`
}

type BombDefused struct {
	MapName   string `json:"map_name"`
	MapNumber int    `json:"map_number"`
	Client    string `json:"client"`
	Site      int    `json:"site"`
}

type BombExploded struct {
	MapName   string `json:"map_name"`
	MapNumber int    `json:"map_number"`
	Client    string `json:"client"`
	Site      int    `json:"site"`
}

type ClientSay struct {
	MapName   string `json:"map_name"`
	MapNumber string `json:"map_number"`
	Client    string `json:"client"`
	Message   string `json:"message"`
}

type PlayerConnect struct {
	MapName   string `json:"map_name"`
	MapNumber int    `json:"map_number"`
	Client    string `json:"client"`
	Ip        string `json:"ip"`
}

type PlayerDisconnect struct {
	MapName   string `json:"map_name"`
	MapNumber int    `json:"map_number"`
	Client    string `json:"client"`
}

type MatchConfigLoadFailed struct {
	Reason string `json:"reason"`
}

type BackupLoad struct {
	File string `json:"file"`
}

type TeamReady struct {
	Team  string `json:"team"`
	Stage string `json:"stage"`
}

type TeamUnready struct {
	Team string `json:"team"`
}
