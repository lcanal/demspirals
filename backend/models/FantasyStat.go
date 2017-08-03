package models

//FantasyStat Stores list of stats along with headers and how many points they're worth. All based on keys.
type FantasyStat struct {
	PID          string            `json:"pid" gorm:"primary_key"`
	LongHeaders  map[string]string `json:"longheaders"`
	ShortHeaders map[string]string `json:"shortheaders"`
	Stat         map[string]int    `json:"stat"`
	StatPts      map[string]int    `json:"statpts"`
}
