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

	seasonKey := "2017-regular"
	activePlayersEndPoint := viper.GetString("apiBaseURL") + "/v1.1/pull/nfl/" + seasonKey + "/active_players.json"

	data, err := ioutil.ReadAll(routes.CallAPI(activePlayersEndPoint))
	if err != nil {
		log.Printf("Error loading url '%s': %s", activePlayersEndPoint, err.Error())
		return
	}

	activePlayers, _, _, _ := jsonparser.Get(data, "activeplayers")

	jsonparser.ArrayEach(
		activePlayers,
		func(playerTeamTuple []byte, dataType jsonparser.ValueType, offset int, err error) {
			var newPlayer models.Player
			var newTeam models.Team
			player, _, _, _ := jsonparser.Get(playerTeamTuple, "player")
			errUn := json.Unmarshal(player, &newPlayer)
			if errUn != nil {
				log.Printf("Error converting json to player object: %s\n", errUn.Error())
				return
			}

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

			players[newPlayer.ID] = newPlayer
			teams[newTeam.ID] = newTeam

			//I realize the above line will overwrite a team again with the same team. I'm cool with this.
		},
		"playerentry",
	)

	//Load all teams and players at once, then save them one by one to the DB.
	//Note, one by one saving is due to ORM limitation.
	db := loader.GormConnectDB()
	db.LogMode(true)
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
	}
	log.Printf("Finished loading %d players", len(players))
}

//LoadAllPlayerStats print player stas
func LoadAllPlayerStats(MAXPAGECOUNT int) {
	stats := make(map[string]models.Stat)
	//fstats := make(map[string]models.FantasyStat)

	for currentPage := 1; currentPage < MAXPAGECOUNT; currentPage++ {
		//apiBase := viper.GetString("apiBaseURL") + "/football/nfl/player_season_stats?interval_type=regularseason&season_id=nfl-2016-2017" + "&per_page=40&page="
		//apiPagedURL := apiBase + strconv.Itoa(currentPage)

		/*data, _ := ioutil.ReadAll(routes.CallAPI(apiPagedURL))
		_, _, _, err := jsonparser.Get(data, "player_season_stats")
		if err != nil {
			fmt.Printf("Ended  stats load at page %d\n", currentPage)
			break
		}*/

		//NEW API SOURCE
	}

	//Save all records to DB once stats have been obtained.
	db := loader.GormConnectDB()
	for _, stat := range stats {
		if db.Create(stat).Error != nil {
			db.Save(stat)
		}
	}
	log.Printf("Finished loading %d stats", len(stats))
}
