package models

//Player is the type that holds all player information
type Player struct {
	ID          string      `json:"id"`
	Slug        string      `json:"slug"`
	Name        string      `json:"name"`
	Position    string      `json:"position"`
	SeasonStats SeasonStats `json:"stats"`
	ShortP      string      `json:"shortpos,omitempty"`
	TeamID      string      `json:"teamid,omitempty"`
	TeamName    string      `json:"teamname,omitempty"`
	Games       string      `json:"games,omitempty"`
}
