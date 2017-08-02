package models

//Team type per player
type Team struct {
	ID       string `json:"id" gorm:"primary_key""`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Color    string `json:"color"`
	Hashtag  string `json:"hashtag"`
	Nickname string `json:"nickname"`
}
