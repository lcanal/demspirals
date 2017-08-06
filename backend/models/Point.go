package models

import "github.com/jinzhu/gorm"

//Point stores list of fantasy points. Assigned to a player.
type Point struct {
	gorm.Model
	PlayerID     string  `json:"pid"`
	Category     string  `json:"category"`
	Abbreviation string  `json:"shortname"`
	Name         string  `json:"name"`
	LeagueName   string  `json:"leaguename"` //The name the league uses on how it counts points (ie, RY10)
	StatNum      float64 `json:"stat"`
	Value        float64 `json:"points"`
}
