package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/lcanal/demspirals/models"
	"github.com/spf13/viper"
)

//Main SQL driver
import _ "github.com/go-sql-driver/mysql"

//LoadAllPlayers grabs full team roster
func LoadAllPlayers() {
	players := make(map[string]models.Player)
	const MAXPAGECOUNT = 5

	for currentPage := 1; currentPage < MAXPAGECOUNT; currentPage++ {
		apiBase := viper.GetString("apiBaseURL") + "/football/nfl/rosters" + "?per_page=40&page="
		apiPagedURL := apiBase + strconv.Itoa(currentPage)

		data, _ := ioutil.ReadAll(CallAPI(apiPagedURL))
		if len(data) <= 0 {
			break
		}

		fmt.Printf("Current page count %d, current data length %d\n", currentPage, len(data))

		jsonparser.ArrayEach(
			data,
			func(player []byte, dataType jsonparser.ValueType, offset int, err error) {
				playername, _ := jsonparser.GetString(player, "name")
				playerpos, _ := jsonparser.GetString(player, "position_name")
				playerslug, _ := jsonparser.GetString(player, "slug")
				playerid, _ := jsonparser.GetString(player, "id")

				newPlayer := models.Player{
					ID:       playerid,
					Slug:     playerslug,
					Name:     playername,
					Position: playerpos,
				}
				players[playerslug] = newPlayer
			},
			"players",
		)
	}

	body, _ := json.Marshal(players)
	err := ioutil.WriteFile("data/players.json", body, 0644)
	if err != nil {
		log.Printf("Error writing players.json: %s\n", err.Error())
	}
}

//PlayerStats print player stas
func PlayerStats(w http.ResponseWriter, r *http.Request) {
	stats := make(map[string]models.PlayerStats)
	const MAXPAGECOUNT = 5

	for currentPage := 1; currentPage < MAXPAGECOUNT; currentPage++ {
		apiBase := viper.GetString("apiBaseURL") + "/football/nfl/player_season_stats?interval_type=regularseason&season_id=nfl-2016-2017" + "&per_page=40&page="
		apiPagedURL := apiBase + strconv.Itoa(currentPage)

		data, _ := ioutil.ReadAll(CallAPI(apiPagedURL))
		if len(data) <= 0 {
			break
		}

		fmt.Printf("Current page count beeeeeee %d, current data length %d\n", currentPage, len(data))

		jsonparser.ArrayEach(
			data,
			func(player []byte, dataType jsonparser.ValueType, offset int, err error) {
				pID, _ := jsonparser.GetString(player, "player_id")
				receptions, _ := jsonparser.GetInt(player, "receptions")
				/*receptionyards, _ := jsonparser.GetInt(player, "receptionyards")
				receptionyards10p, _ := jsonparser.GetInt(player, "receptionyards10p")
				receptionyards20p, _ := jsonparser.GetInt(player, "receptionyards20p")
				receptionyards30p, _ := jsonparser.GetInt(player, "receptionyards30p")
				receptionyards40p, _ := jsonparser.GetInt(player, "receptionyards40p")
				receptionyards50p, _ := jsonparser.GetInt(player, "receptionyards50p")
				touchdownpasses, _ := jsonparser.GetInt(player, "touchdownpasses")*/

				newStat := models.PlayerStats{
					Receptions: receptions,
				}

				stats[pID] = newStat
			},
			"player_season_stats",
		)

	}

	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	pass := viper.GetString("db.pass")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, user)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to db: %s", err.Error())
	}

	insertStatStmt, err := db.Prepare("REPLACE INTO playerstats (pid,runs,passes,receptions) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatalf("Error preparing db statement: %s\n", err.Error())
	}

	for pid, stat := range stats {
		fmt.Printf("Inserting %s... \n", pid)
		_, err := insertStatStmt.Exec(pid, stat.Runs, stat.Passes, stat.Receptions)
		if err != nil {
			log.Fatalf("Error inserting stat %s: %s", pid, err.Error())
		}
	}

	insertStatStmt.Close()
}
