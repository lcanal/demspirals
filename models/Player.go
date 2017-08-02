package models

//Player is the type that holds all player information
type Player struct {
	ID       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Pos      string `json:"pos"`
	TeamID   string `json:"teamid"`
	Team     Team   `json:"team" gorm:"ForeignKey:TeamID;AssociationForeignKey:ID"`
	Stats    Stat   `json:"stats" gorm:"ForeignKey:ID;AssociationForeignKey:PID"`
}
