package models

//SeasonStats tracks the stats we care about. Generally for fantasy points
type SeasonStats struct {
	Runs   string `json:"runs,omitempty"`
	Passes string `json:"passes,omitempty"`
}
