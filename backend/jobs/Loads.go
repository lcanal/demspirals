package jobs

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/lcanal/demspirals/backend/loader"
	"github.com/lcanal/demspirals/backend/models"
	"github.com/lcanal/demspirals/backend/routes"
	"github.com/spf13/viper"
)

//LoadAllPlayers grabs full team roster
func LoadAllPlayers(MAXPAGECOUNT int) {
	players := make(map[string]models.Player)

	for currentPage := 1; currentPage < MAXPAGECOUNT; currentPage++ {
		apiBase := viper.GetString("apiBaseURL") + "/football/nfl/rosters" + "?per_page=40&page="
		apiPagedURL := apiBase + strconv.Itoa(currentPage)

		data, _ := ioutil.ReadAll(routes.CallAPI(apiPagedURL))
		_, _, _, err := jsonparser.Get(data, "players")
		if err != nil {
			fmt.Printf("Ended player load at page %d\n", currentPage)
			break
		}

		jsonparser.ArrayEach(
			data,
			func(player []byte, dataType jsonparser.ValueType, offset int, err error) {
				playername, _ := jsonparser.GetString(player, "name")
				playerposition, _ := jsonparser.GetString(player, "position_name")
				playerpos, _ := jsonparser.GetString(player, "position_abbreviation")
				playerslug, _ := jsonparser.GetString(player, "slug")
				playerid, _ := jsonparser.GetString(player, "id")
				playertid, _ := jsonparser.GetString(player, "team_id")

				newPlayer := models.Player{
					ID:       playerid,
					Slug:     playerslug,
					Name:     playername,
					Position: playerposition,
					Pos:      playerpos,
					TeamID:   playertid,
				}
				players[playerslug] = newPlayer
			},
			"players",
		)
	}

	//Save all records to DB once players have been obtained.
	db := loader.GormConnectDB()
	for _, player := range players {
		if db.Create(player).Error != nil {
			db.Save(player)
		}
	}
	log.Printf("Finished loading %d players", len(players))
}

//LoadAllTeams Loads team stats. Assumes a single page call.
func LoadAllTeams() {
	teams := make(map[string]models.Team)
	apiPagedURL := viper.GetString("apiBaseURL") + "/football/nfl/teams?per_page=40"
	data, _ := ioutil.ReadAll(routes.CallAPI(apiPagedURL))
	_, _, _, err := jsonparser.Get(data, "teams")
	if err != nil {
		fmt.Printf("Error... no teams loaded!\n")
		return
	}

	jsonparser.ArrayEach(
		data,
		func(team []byte, dataType jsonparser.ValueType, offset int, err error) {
			id, _ := jsonparser.GetString(team, "id")
			name, _ := jsonparser.GetString(team, "name")
			nickname, _ := jsonparser.GetString(team, "nickname")
			color, _ := jsonparser.GetString(team, "color")
			hashtag, _ := jsonparser.GetString(team, "hashtag")
			slug, _ := jsonparser.GetString(team, "slug")

			newTeam := models.Team{
				ID:       id,
				Slug:     slug,
				Name:     name,
				Nickname: nickname,
				Color:    color,
				Hashtag:  hashtag,
			}
			teams[slug] = newTeam
		},
		"teams",
	)

	//Save all records to DB once teams have been obtained.
	db := loader.GormConnectDB()
	for _, team := range teams {
		if db.Create(team).Error != nil {
			db.Save(team)
		}
	}
	log.Printf("Finished loading %d teams", len(teams))
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
