package models

//Player is the type that holds all player information
type Player struct {
	ID       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Teamid   string `json:"team"`
}
