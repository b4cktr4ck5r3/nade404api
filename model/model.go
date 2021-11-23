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

type PteroServerList struct {
	Object string                 `json:"object"`
	Data   []PteroServerListDatum `json:"data"`
	Meta   Meta                   `json:"meta"`
}

type PteroServerListDatum struct {
	Object     string           `json:"object"`
	Attributes PurpleAttributes `json:"attributes"`
}

type PurpleAttributes struct {
	ServerOwner    bool          `json:"server_owner"`
	Identifier     string        `json:"identifier"`
	InternalID     int64         `json:"internal_id"`
	UUID           string        `json:"uuid"`
	Name           string        `json:"name"`
	Node           string        `json:"node"`
	SFTPDetails    SFTPDetails   `json:"sftp_details"`
	Description    string        `json:"description"`
	Limits         Limits        `json:"limits"`
	Invocation     string        `json:"invocation"`
	DockerImage    string        `json:"docker_image"`
	EggFeatures    []string      `json:"egg_features"`
	FeatureLimits  FeatureLimits `json:"feature_limits"`
	Status         interface{}   `json:"status"`
	IsSuspended    bool          `json:"is_suspended"`
	IsInstalling   bool          `json:"is_installing"`
	IsTransferring bool          `json:"is_transferring"`
	Relationships  Relationships `json:"relationships"`
}

type FeatureLimits struct {
	Databases   int64 `json:"databases"`
	Allocations int64 `json:"allocations"`
	Backups     int64 `json:"backups"`
}

type Limits struct {
	Memory      int64  `json:"memory"`
	Swap        int64  `json:"swap"`
	Disk        int64  `json:"disk"`
	Io          int64  `json:"io"`
	CPU         int64  `json:"cpu"`
	Threads     string `json:"threads"`
	OOMDisabled bool   `json:"oom_disabled"`
}

type Relationships struct {
	Allocations Allocations `json:"allocations"`
	Variables   Variables   `json:"variables"`
}

type Allocations struct {
	Object string             `json:"object"`
	Data   []AllocationsDatum `json:"data"`
}

type AllocationsDatum struct {
	Object     string           `json:"object"`
	Attributes FluffyAttributes `json:"attributes"`
}

type FluffyAttributes struct {
	ID        int64       `json:"id"`
	IP        string      `json:"ip"`
	IPAlias   string      `json:"ip_alias"`
	Port      int64       `json:"port"`
	Notes     interface{} `json:"notes"`
	IsDefault bool        `json:"is_default"`
}

type Variables struct {
	Object string           `json:"object"`
	Data   []VariablesDatum `json:"data"`
}

type VariablesDatum struct {
	Object     string              `json:"object"`
	Attributes TentacledAttributes `json:"attributes"`
}

type TentacledAttributes struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	EnvVariable  string `json:"env_variable"`
	DefaultValue string `json:"default_value"`
	ServerValue  string `json:"server_value"`
	IsEditable   bool   `json:"is_editable"`
	Rules        string `json:"rules"`
}

type SFTPDetails struct {
	IP   string `json:"ip"`
	Port int64  `json:"port"`
}

type Meta struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Total       int64 `json:"total"`
	Count       int64 `json:"count"`
	PerPage     int64 `json:"per_page"`
	CurrentPage int64 `json:"current_page"`
	TotalPages  int64 `json:"total_pages"`
	Links       Links `json:"links"`
}

type Links struct {
	Next string `json:"next"`
}
