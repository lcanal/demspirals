package routes

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/buger/jsonparser"
	"github.com/spf13/viper"
)

//TeamRoster grabs full team roster
func TeamRoster(w http.ResponseWriter, r *http.Request) {
	var playerString string

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.stattleship.com/football/nfl/rosters", nil)
	if err != nil {
		log.Printf("Error pulling team rosters: %s\n", err)
	}

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

			playerString += fmt.Sprintf("%-50s - %-40s %-40s\n", playerslug, playername, playerpos)
		},
		"players",
	)

	fmt.Fprintf(w, playerString)
}
