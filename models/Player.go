package models

//Player is the type that holds all player information
type Player struct {
	ID       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Team     Team   `json:"team"`
	Stats    Stat   `json:"stats"`
}
