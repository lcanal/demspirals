package jobs

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/buger/jsonparser"
	"github.com/lcanal/demspirals/backend/loader"
	"github.com/lcanal/demspirals/backend/models"
	"github.com/lcanal/demspirals/backend/routes"
	"github.com/spf13/viper"
)

//LoadAllPlayersAndTeams grabs full team roster
func LoadAllPlayersAndTeams() {
	players := make(map[string]models.Player)
	teams := make(map[string]models.Team)
	//stats := make(map[string][]models.Stat)

	seasonKey := "2016-regular"
	apiBase := viper.GetString("apiBaseURL") + "/v1.1/pull/nfl/"
	activePlayersEndPoint := apiBase + seasonKey + "/cumulative_player_stats.json?player=Tom-Brady-7549"

	//log.Printf("Endpoint URL: %s", activePlayersEndPoint)

	data, err := ioutil.ReadAll(routes.CallAPI(activePlayersEndPoint))
	if err != nil {
		log.Printf("Error loading url '%s': %s", activePlayersEndPoint, err.Error())
		return
	}

	//activePlayers, _, _, _ := jsonparser.Get(data, "activeplayers")
	_, errArray := jsonparser.ArrayEach(
		data,
		func(playerTeamTuple []byte, dataType jsonparser.ValueType, offset int, err error) {
			var newPlayer models.Player
			var newTeam models.Team

			//Individual Player
			player, _, _, _ := jsonparser.Get(playerTeamTuple, "player")
			errUn := json.Unmarshal(player, &newPlayer)
			if errUn != nil {
				log.Printf("Error converting json to player object: %s\n", errUn.Error())
				return
			}

			//Team player belongs to
			team, _, _, errGet := jsonparser.Get(playerTeamTuple, "team")
			if errGet == nil {
				errUn = json.Unmarshal(team, &newTeam)
				if errUn != nil {
					log.Printf("Error converting json to team object: %s\nObject: %s", errUn.Error(), string(team))
					return
				}
			} else {
				//No team, make empty
				newTeam = models.Team{
					ID:           "FA",
					Name:         "Free Agent",
					City:         "N/A",
					Abbreviation: "N/A",
				}
			}

			newPlayer.Team = newTeam
			newPlayer.TeamID = newTeam.ID
			newPlayer.MapStats(playerTeamTuple)

			players[newPlayer.ID] = newPlayer
			teams[newTeam.ID] = newTeam

			//I realize the above line will overwrite a team again with the same team. I'm cool with this.
		},
		"cumulativeplayerstats", "playerstatsentry",
	)

	if errArray != nil {
		log.Fatalf("Somethin'g wrong here... %s\n", errArray.Error())
	}
	//Load all teams and players at once, then save them one by one to the DB.
	//Note, one by one saving is due to ORM limitation.
	db := loader.GormConnectDB()
	//db.LogMode(true)
	for _, team := range teams {
		if db.Create(&team).Error != nil {
			db.Save(&team)
		}
	}
	log.Printf("Finished loading %d teams", len(teams))

	//This query both runs insert on player AND update team.
	for _, player := range players {
		if db.Create(&player).Error != nil {
			db.Save(&player)
		}

		//		var blankStat []models.Stat

		for _, stat := range player.Stats {
			db.Create(stat)
		}

	}
	log.Printf("Finished loading %d players", len(players))

}
