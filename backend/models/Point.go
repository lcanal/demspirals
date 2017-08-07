package models

//Point stores list of fantasy points. Assigned to a player.
type Point struct {
	ID           uint    `json:"id" gorm:"primary_key"`
	PlayerID     string  `json:"pid"`
	Category     string  `json:"category"`
	Abbreviation string  `json:"shortname"`
	Name         string  `json:"name"`
	LeagueName   string  `json:"leaguename"` //The name the league uses on how it counts points (ie, RY10)
	StatID       int     `json:"statid"`     //Matching id
	StatNum      float64 `json:"stat"`
	Value        float64 `json:"points"`
}
