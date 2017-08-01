package models

//Player is the type that holds all player information
type Player struct {
	ID       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Pos      string `json:"pos"`
	Teamid   string `json:"teamid"`
	Team     Team   `json:"team" gorm:"ForeignKey:ID"`
	Stats    Stat   `json:"stats"`
}
