package models

//Player is the type that holds all player information
type Player struct {
	ID               string
	LastName         string
	FirstName        string
	JerseyNumber     string
	Position         string
	Height           string
	Weight           string
	BirthDate        string
	Age              string
	BirthCity        string
	BirthCountry     string
	IsRookie         string
	officialImageSrc string
	TeamID           string
	Team             Team `gorm:"ForeignKey:ID"`
	//Stats    Stat   `json:"stats" gorm:"ForeignKey:ID;AssociationForeignKey:PID"`
}
