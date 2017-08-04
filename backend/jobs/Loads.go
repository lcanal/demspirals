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

//LoadAllPlayerData grabs full team roster
func LoadAllPlayerData() {
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

	_, errArray := jsonparser.ArrayEach(
		data,
		func(playerData []byte, dataType jsonparser.ValueType, offset int, err error) {
			var newPlayer models.Player

			//Individual Player
			player, _, _, _ := jsonparser.Get(playerData, "player")
			errUn := json.Unmarshal(player, &newPlayer)
			if errUn != nil {
				log.Printf("Error converting json to player object: %s\n", errUn.Error())
				return
			}

			newPlayer.MapStats(playerData)
			newPlayer.MapTeam(playerData)

			players[newPlayer.ID] = newPlayer
			teams[newPlayer.Team.ID] = newPlayer.Team

			//I realize the above line will overwrite a team again with the same team. I'm cool with this.
		},
		"cumulativeplayerstats", "playerstatsentry",
	)

	if errArray != nil {
		log.Fatalf("Somethin'g wrong here... %s\n", errArray.Error())
	}

	//Note, one by one saving is due to ORM limitation.
	db := loader.GormConnectDB()
	//db.LogMode(true)
	for _, team := range teams {
		if db.Create(&team).Error != nil {
			db.Save(&team)
		}
	}
	log.Printf("Finished loading %d teams\n", len(teams))

	//This query both runs insert on player AND update team.
	for _, player := range players {
		if db.Create(&player).Error != nil {
			db.Save(&player)
		}
	}
	log.Printf("Finished loading %d players\n", len(players))

	log.Printf("Done with loads.\n")

}
