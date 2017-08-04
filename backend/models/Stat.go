package models

//Stat  tracks the all of the passing stats.
type Stat struct {
	PlayerID     string `json:"pid"`
	Name         string `json:"statname"`
	Category     string `json:"@category"`
	Abbreviation string `json:"@abbreviation"`
	Value        string `json:"#text"`
}
