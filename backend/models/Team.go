package models

//Team type per player
type Team struct {
	ID           string `gorm:"primary_key"`
	Name         string
	City         string
	Abbreviation string
}
