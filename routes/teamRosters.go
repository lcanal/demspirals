package routes

import (
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

//TeamRoster grabs full team roster
func TeamRoster(w http.ResponseWriter, r *http.Request) {
	var players map[string]models.Player
	const MAXPAGECOUNT = 5
	players = make(map[string]models.Player)

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

	//Gotta return something...
	body, _ := json.Marshal(players)
	err := ioutil.WriteFile("data/players.json", body, 0644)
	if err != nil {
		log.Printf("Error writing players.json: %s\n", err.Error())
	}
	fmt.Fprintf(w, string(body))
}

//PlayerStats print player stas
func PlayerStats(w http.ResponseWriter, r *http.Request) {
	const MAXPAGECOUNT = 5

	for currentPage := 1; currentPage < MAXPAGECOUNT; currentPage++ {
		apiBase := viper.GetString("apiBaseURL") + "/football/nfl/player_season_stats?interval_type=regularseason&season_id=nfl-2016-2017" + "&per_page=40&page="
		apiPagedURL := apiBase + strconv.Itoa(currentPage)

		data, _ := ioutil.ReadAll(CallAPI(apiPagedURL))
		if len(data) <= 0 {
			break
		}

		jsonparser.ArrayEach(
			data,
			func(pStats []byte, dataType jsonparser.ValueType, offset int, err error) {

			},
			"player_season_stats",
		)

	}
}
