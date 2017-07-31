package models

//PlayerStats tracks the stats we care about. Generally for fantasy points
type PlayerStats struct {
	PID        string `json:"pid"`
	Runs       int64  `json:"runs,omitempty"`
	Passes     int64  `json:"passes,omitempty"`
	Receptions int64  `json:"receptions,omitempty"`
}
