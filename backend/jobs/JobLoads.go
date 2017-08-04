package jobs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

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
	activePlayersEndPoint := apiBase + seasonKey + "/cumulative_player_stats.json?limit=25"

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

//CalculatePoints points for players
func CalculatePoints() {
	//All players
	pointValueFile := viper.New()
	pointValueFile.SetConfigName("pointvalues")
	pointValueFile.AddConfigPath(".")
	pointValueFile.AddConfigPath("config")
	pointModel := "espn."

	err := pointValueFile.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error pointvalues.json file: %s", err.Error())
	}

	db := loader.GormConnectDB()

	db.LogMode(true)
	var players []models.Player
	var points []models.Point

	log.Println("Reading stats from database...")
	db.Preload("Team").Preload("Stats").Find(&players)
	log.Printf("Saving stats to database...\n")
	db.DropTableIfExists(&models.Point{})
	db.CreateTable(&models.Point{})

	for _, player := range players {
		for _, stat := range player.Stats {
			var newStatFantasyPoints models.Point
			newStatFantasyPoints.Abbreviation = stat.Abbreviation
			newStatFantasyPoints.Category = stat.Category
			newStatFantasyPoints.PlayerID = stat.PlayerID
			newStatFantasyPoints.Name = stat.Name
			newStatFantasyPoints.StatNum, err = strconv.ParseFloat(stat.Value, 32)
			if err != nil {
				newStatFantasyPoints.StatNum = 0
			}

			switch stat.Name {
			case "RushYards":
				newStatFantasyPoints.Value = newStatFantasyPoints.StatNum * pointValueFile.GetFloat64(pointModel+"RY10")
			default:
				newStatFantasyPoints.Value = 0
			}

			if db.Create(&newStatFantasyPoints).Error != nil {
				db.Save(&newStatFantasyPoints)
			}
		}
	}

	log.Printf("Finished writing %d fantasy point rows\n", len(points))

}
