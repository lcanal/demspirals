package models

import (
	"fmt"
	"log"

	"github.com/buger/jsonparser"
)

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

//MapStats takes in set of objects, maps each to player's Stats property
func (p *Player) MapStats(stats []byte) {
	fmt.Printf("Object: %s", string(stats))
	err := jsonparser.ObjectEach(
		stats,
		func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			fmt.Printf("Key Is: %s  --- Value is %s\n", string(key), string(value))
			return nil
		},
		"stats",
	)

	if err != nil {
		log.Fatalf("Error mapping %s\n", err.Error())
		return
	}
}
