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
	cumPlayerSeasonStats := apiBase + seasonKey + "/cumulative_player_stats.json"
	activePlayerInfo := apiBase + seasonKey + "/active_players.json"

	data, err := ioutil.ReadAll(routes.CallAPI(cumPlayerSeasonStats))
	if err != nil {
		log.Printf("Error loading url '%s': %s", cumPlayerSeasonStats, err.Error())
		return
	}

	extraInfo, erre := ioutil.ReadAll(routes.CallAPI(activePlayerInfo))
	if erre != nil {
		log.Printf("Error loading url '%s': %s", activePlayerInfo, err.Error())
		return
	}

	//Parse through
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
			newPlayer.MapExtra(extraInfo)

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

	pointModel := "espn."

	db := loader.GormConnectDB()

	//db.LogMode(true)
	/////////////////////////////Maybe load player game logs here for game specific points //////////////////////////
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

			mapStatsToPoints(&newStatFantasyPoints, stat, pointModel)
			points = append(points, newStatFantasyPoints)

		}
	}

	go loadPoints(points)

}

func loadPlayers(players map[string]models.Player, wg *sync.WaitGroup) {
	//Create raw load strings
	stmt := "INSERT INTO players (id,last_name,first_name,jersey_number,position,pic_url,team_id) VALUES (?,?,?,?,?,?,?)"
	rawdb := loader.ConnectDB()
	st, err := rawdb.Prepare(stmt)
	if err != nil {
		fmt.Printf("Error in preparing statement for saving players: %s\n", err.Error())
		return
	}
	count := 0
	for _, player := range players {
		_, err := st.Exec(player.ID, player.LastName, player.FirstName, player.JerseyNumber, player.Position, player.PicURL, player.TeamID)
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
			valueStrings = append(valueStrings, "(?,?,?,?,?,?,?,?,?)")
			valueArgs = append(valueArgs, nil)
			valueArgs = append(valueArgs, point.PlayerID)
			valueArgs = append(valueArgs, point.Category)
			valueArgs = append(valueArgs, point.Abbreviation)
			valueArgs = append(valueArgs, point.Name)
			valueArgs = append(valueArgs, point.LeagueName)
			valueArgs = append(valueArgs, point.StatID)
			valueArgs = append(valueArgs, point.StatNum)
			valueArgs = append(valueArgs, point.Value)

			query := fmt.Sprintf("INSERT INTO points (id,player_id,category,abbreviation,name,league_name,stat_id,stat_num,value) VALUES %s", strings.Join(valueStrings, ","))
			_, err := rawdb.Exec(query, valueArgs...)
			if err != nil {
				log.Fatalf("Error executing statement for loading pointsz: %s:\n %s", query, err.Error())
				return
			}
			valueStrings = make([]string, 0)
			valueArgs = make([]interface{}, 0)
		} else {
			valueStrings = append(valueStrings, "(?,?,?,?,?,?,?,?,?)")
			valueArgs = append(valueArgs, nil)
			valueArgs = append(valueArgs, point.PlayerID)
			valueArgs = append(valueArgs, point.Category)
			valueArgs = append(valueArgs, point.Abbreviation)
			valueArgs = append(valueArgs, point.Name)
			valueArgs = append(valueArgs, point.LeagueName)
			valueArgs = append(valueArgs, point.StatID)
			valueArgs = append(valueArgs, point.StatNum)
			valueArgs = append(valueArgs, point.Value)
		}
		totalPointCount = totalPointCount + 1
	}
	//Catch the last few
	query := fmt.Sprintf("INSERT INTO points (id,player_id,category,abbreviation,name,league_name,stat_id,stat_num,value) VALUES %s", strings.Join(valueStrings, ","))
	_, err := rawdb.Exec(query, valueArgs...)
	if err != nil {
		log.Fatalf("Error executing statement for loading points: %s:\n %s", query, err.Error())
		return
	}
	log.Printf("Finished writing %d fantasy points\n", totalPointCount)
}

func mapStatsToPoints(fPoint *models.Point, stat models.Stat, pointModel string) {
	//Map stat points to fantasy points. Note, only takes in one point at a time.
	pointValueFile := viper.New()
	pointValueFile.SetConfigName("pointvalues")
	pointValueFile.AddConfigPath(".")
	pointValueFile.AddConfigPath("config")

	err := pointValueFile.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error pointvalues file: %s", err.Error())
	}

	fPoint.Abbreviation = stat.Abbreviation
	fPoint.Category = stat.Category
	fPoint.PlayerID = stat.PlayerID
	fPoint.Name = stat.Name
	fPoint.StatID = stat.ID

	fPoint.StatNum, err = strconv.ParseFloat(stat.Value, 32)
	if err != nil {
		fPoint.StatNum = 0
	}

	pointMaps := pointValueFile.GetStringMap("espn")
	for pointType, pointObject := range pointMaps {
		if strings.ToLower(pointType) == strings.ToLower(stat.Name) {
			//Convert pointobject to map , then to a float.
			pTuple := pointObject.(map[string]interface{})
			leagueValue := pTuple["leaguevalue"].(float64)
			leagueName := pTuple["leaguename"].(string)

			fPoint.Value = fPoint.StatNum * leagueValue
			fPoint.LeagueName = leagueName
			//log.Printf("Logging %f points for %s stat (num:%f) (LeagueName: %s) (League Value: %f)\n", fPoint.Value, fPoint.Name, fPoint.StatNum, leagueName, leagueValue)
			return
		}
	}

	//If out of loop, means no matching season stat point found
	fPoint.Value = 0
	fPoint.LeagueName = ""
}
