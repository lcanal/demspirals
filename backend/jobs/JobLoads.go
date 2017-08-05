package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/buger/jsonparser"
	"github.com/lcanal/demspirals/backend/loader"
	"github.com/lcanal/demspirals/backend/models"
	"github.com/lcanal/demspirals/backend/routes"
	"github.com/spf13/viper"
)

//LoadAllPlayerData grabs full team roster
func LoadAllPlayerData(wg *sync.WaitGroup) {
	players := make(map[string]models.Player)
	teams := make(map[string]models.Team)
	//stats := make(map[string][]models.Stat)

	seasonKey := "2016-regular"
	apiBase := viper.GetString("apiBaseURL") + "/v1.1/pull/nfl/"
	activePlayersEndPoint := apiBase + seasonKey + "/cumulative_player_stats.json"

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

	/*//This query both runs insert on player AND update team.
	for _, player := range players {
		if db.Create(&player).Error != nil {
			db.Save(&player)
		}
	}
	log.Printf("Finished loading %d players\n", len(players))*/

	//Load into DB, add 2 more functions to wait for
	wg.Add(2)
	go loadPlayers(players, wg)
	go loadStats(players, wg)

	defer wg.Done()
}

//CalculatePoints points for players
func CalculatePoints(wg *sync.WaitGroup) {
	//Wait for relevant functions to finish
	wg.Wait()

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

	//db.LogMode(true)
	var players []models.Player
	var points []models.Point

	log.Println("Calculating fantasy points..")
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

			points = append(points, newStatFantasyPoints)

		}
	}

	go loadPoints(points)

}

func loadPlayers(players map[string]models.Player, wg *sync.WaitGroup) {
	//Create raw load strings
	stmt := "INSERT INTO players (id,last_name,first_name,jersey_number,position,team_id) VALUES (?,?,?,?,?,?)"
	rawdb := loader.ConnectDB()
	st, err := rawdb.Prepare(stmt)
	if err != nil {
		fmt.Printf("Error in preparing statement for saving players: %s\n", err.Error())
		return
	}
	count := 0
	for _, player := range players {
		_, err := st.Exec(player.ID, player.LastName, player.FirstName, player.JerseyNumber, player.Position, player.TeamID)
		if err != nil {
			log.Fatalf("Error executing statement for saving players %s:\n %s", stmt, err.Error())
			return
		}

		count = count + 1
	}

	log.Printf("Finished loading %d players into database.\n", len(players))
	st.Close()
	defer wg.Done()
}

func loadStats(players map[string]models.Player, wg *sync.WaitGroup) {
	valueStrings := make([]string, 0)
	valueArgs := make([]interface{}, 0)
	rawdb := loader.ConnectDB()
	var totalStatCount int

	for _, player := range players {
		for _, stat := range player.Stats {
			valueStrings = append(valueStrings, "(?,?,?,?,?,?)")
			valueArgs = append(valueArgs, nil)
			valueArgs = append(valueArgs, stat.PlayerID)
			valueArgs = append(valueArgs, stat.Name)
			valueArgs = append(valueArgs, stat.Category)
			valueArgs = append(valueArgs, stat.Abbreviation)
			valueArgs = append(valueArgs, stat.Value)
			totalStatCount = totalStatCount + 1
		}

		query := fmt.Sprintf("INSERT INTO stats (id,player_id,name,category,abbreviation,value) VALUES %s", strings.Join(valueStrings, ","))
		_, err := rawdb.Exec(query, valueArgs...)
		if err != nil {
			log.Fatalf("Error executing statement for loading stats: %s:\n %s", query, err.Error())
			return
		}
		valueStrings = make([]string, 0)
		valueArgs = make([]interface{}, 0)
	}
	log.Printf("Finished loading %d stats into database.\n", totalStatCount)
	defer wg.Done()
}

func loadPoints(points []models.Point) {
	valueStrings := make([]string, 0)
	valueArgs := make([]interface{}, 0)
	rawdb := loader.ConnectDB()
	var totalPointCount int

	for _, point := range points {
		//Chunk them 25 at a time.
		if totalPointCount%24 == 0 && totalPointCount != 0 {
			valueStrings = append(valueStrings, "(?,?,?,?,?,?,?,?,?,?)")
			valueArgs = append(valueArgs, nil)
			valueArgs = append(valueArgs, nil)
			valueArgs = append(valueArgs, nil)
			valueArgs = append(valueArgs, nil)
			valueArgs = append(valueArgs, point.PlayerID)
			valueArgs = append(valueArgs, point.Category)
			valueArgs = append(valueArgs, point.Abbreviation)
			valueArgs = append(valueArgs, point.Name)
			valueArgs = append(valueArgs, point.StatNum)
			valueArgs = append(valueArgs, point.Value)

			query := fmt.Sprintf("INSERT INTO points (id,created_at,updated_at,deleted_at,player_id,category,abbreviation,name,stat_num,value) VALUES %s", strings.Join(valueStrings, ","))
			_, err := rawdb.Exec(query, valueArgs...)
			if err != nil {
				log.Fatalf("Error executing statement for loading pointsz: %s:\n %s", query, err.Error())
				return
			}
			valueStrings = make([]string, 0)
			valueArgs = make([]interface{}, 0)
		} else {
			valueStrings = append(valueStrings, "(?,?,?,?,?,?,?,?,?,?)")
			valueArgs = append(valueArgs, nil)
			valueArgs = append(valueArgs, nil)
			valueArgs = append(valueArgs, nil)
			valueArgs = append(valueArgs, nil)
			valueArgs = append(valueArgs, point.PlayerID)
			valueArgs = append(valueArgs, point.Category)
			valueArgs = append(valueArgs, point.Abbreviation)
			valueArgs = append(valueArgs, point.Name)
			valueArgs = append(valueArgs, point.StatNum)
			valueArgs = append(valueArgs, point.Value)
		}
		totalPointCount = totalPointCount + 1
	}
	//Catch the last few
	query := fmt.Sprintf("INSERT INTO points (id,created_at,updated_at,deleted_at,player_id,category,abbreviation,name,stat_num,value) VALUES %s", strings.Join(valueStrings, ","))
	_, err := rawdb.Exec(query, valueArgs...)
	if err != nil {
		log.Fatalf("Error executing statement for loading points: %s:\n %s", query, err.Error())
		return
	}
	log.Printf("Finished writing %d fantasy points\n", totalPointCount)
}
