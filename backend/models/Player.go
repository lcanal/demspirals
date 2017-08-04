package models

import (
	"encoding/json"
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
	Team             Team   `gorm:"ForeignKey:ID"`
	Stats            []Stat `json:"stats" gorm:"ForeignKey:ID;AssociationForeignKey:PID"`
}

//MapStats takes in set of objects, maps each to player's Stats property. Flatten stats object from api source.
func (p *Player) MapStats(stats []byte) {
	err := jsonparser.ObjectEach(
		stats,
		func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			//var statsset map[string]string
			var newStat Stat
			err := json.Unmarshal(value, &newStat)
			if err != nil {
				return err
			}

			newStat.Name = string(key)
			newStat.PID = p.ID
			p.Stats = append(p.Stats, newStat)
			return err
		},
		"stats",
	)

	if err != nil {
		log.Fatalf("Error mapping %s\n", err.Error())
		return
	}
}
