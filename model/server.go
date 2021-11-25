package model

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
	Port      int         `json:"port"`
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
