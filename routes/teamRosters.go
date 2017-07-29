package routes

import (
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

	client := &http.Client{}
	players = make(map[string]models.Player)

	for currentPage := 1; currentPage < MAXPAGECOUNT; currentPage++ {
		apiBase := viper.GetString("apiBaseURL") + "/football/nfl/rosters" + "?per_page=40&page="
		apiPagedURL := apiBase + strconv.Itoa(currentPage)
		log.Printf("Grabbing URL: %s", apiPagedURL)

		req, _ := http.NewRequest("GET", apiPagedURL, nil)
		req.Header.Add("Authorization", "Token token="+viper.GetString("creds.accessToken"))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/vnd.stattleship.com; version=1")
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error pulling team rosters auth : %s\n", err)
		}

		if resp.StatusCode != 200 {
			log.Printf("Error with request: %s (%d)\n", resp.Status, resp.StatusCode)
		}

		data, _ := ioutil.ReadAll(resp.Body)
		//log.Printf("data: %s\n", body)

		jsonparser.ArrayEach(
			data,
			func(player []byte, dataType jsonparser.ValueType, offset int, err error) {
				playername, _ := jsonparser.GetString(player, "name")
				playerpos, _ := jsonparser.GetString(player, "position_name")
				playerslug, _ := jsonparser.GetString(player, "slug")

				newPlayer := models.Player{
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
	length := strconv.Itoa(len(players))
	fmt.Fprintf(w, length)
}
