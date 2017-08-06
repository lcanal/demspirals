package models

import (
	"encoding/json"
	"log"

	"github.com/buger/jsonparser"
)

//Player is the type that holds all player information
type Player struct {
	ID           string
	LastName     string
	FirstName    string
	JerseyNumber string
	Position     string
	PicURL       string
	TeamID       string
	Team         Team
	Stats        []Stat `json:"stats"`
}

//MapStats takes in set of objects, maps each to player's Stats property. Flatten stats object from api source.
func (p *Player) MapStats(playerData []byte) {
	err := jsonparser.ObjectEach(
		playerData,
		func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			var newStat Stat
			err := json.Unmarshal(value, &newStat)
			if err != nil {
				return err
			}

			newStat.Name = string(key)
			newStat.PlayerID = p.ID
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

//MapTeam takes in set of objects, maps each to player's Stats property. Flatten stats object from api source.
func (p *Player) MapTeam(playerData []byte) {
	var newTeam Team
	team, _, _, err := jsonparser.Get(playerData, "team")
	if err == nil {
		errUn := json.Unmarshal(team, &newTeam)
		if errUn != nil {
			log.Printf("Error converting json to team object: %s\nObject: %s", errUn.Error(), string(team))
			return
		}
	} else {
		//No team, make empty
		newTeam = Team{
			ID:           "FA",
			Name:         "Free Agent",
			City:         "N/A",
			Abbreviation: "N/A",
		}
	}

	p.Team = newTeam
	p.TeamID = newTeam.ID
}

//MapExtra maps additional info like portrait url. Takes in raw json call.
func (p *Player) MapExtra(extraInfo []byte) {
	jsonparser.ArrayEach(
		extraInfo,
		func(playerData []byte, dataType jsonparser.ValueType, offset int, err error) {
			player, _, _, err := jsonparser.Get(playerData, "player")
			if err != nil {
				log.Printf("Error finding extra player data for %s %s", p.FirstName, p.LastName)
				return
			}

			currPlayerID, _ := jsonparser.GetString(player, "ID")
			if currPlayerID == p.ID {
				p.PicURL, _ = jsonparser.GetString(player, "officialImageSrc")
				return
			}
		},
		"activeplayers", "playerentry",
	)
}
