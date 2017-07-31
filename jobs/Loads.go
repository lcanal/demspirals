package jobs

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/lcanal/demspirals/loader"
	"github.com/lcanal/demspirals/models"
	"github.com/lcanal/demspirals/routes"
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
			fmt.Printf("Ended load at page %d", currentPage)
			break
		}

		//fmt.Printf("Loading players page %d", currentPage)
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

	db := loader.ConnectDB()

	insertPlayerStmt, err := db.Prepare("REPLACE INTO players (id,slug,name,position) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatalf("Error preparing db statement: %s\n", err.Error())
	}

	for slug, player := range players {
		log.Printf("Inserting Player %s... \n", slug)
		_, err := insertPlayerStmt.Exec(player.ID, player.Slug, player.Name, player.Position)
		if err != nil {
			log.Fatalf("Error inserting stat %s: %s", slug, err.Error())
		}
	}
	insertPlayerStmt.Close()
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

	fmt.Printf("Loading teams...\n")

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

	db := loader.ConnectDB()

	insertTeamsStmt, err := db.Prepare("REPLACE INTO teams (id,slug,name,nickname,color,hashtag) VALUES (?,?,?,?,?,?)")
	if err != nil {
		log.Fatalf("Error preparing db statement: %s\n", err.Error())
	}

	for slug, team := range teams {
		log.Printf("Inserting Team %s... \n", team.Slug)
		_, err := insertTeamsStmt.Exec(team.ID, team.Slug, team.Name, team.Nickname, team.Color, team.Hashtag)
		if err != nil {
			log.Fatalf("Error inserting team %s: %s", slug, err.Error())
		}
	}
	insertTeamsStmt.Close()
}

//LoadAllPlayerStats print player stas
func LoadAllPlayerStats(MAXPAGECOUNT int) {
	stats := make(map[string]models.PlayerStats)

	for currentPage := 1; currentPage < MAXPAGECOUNT; currentPage++ {
		apiBase := viper.GetString("apiBaseURL") + "/football/nfl/player_season_stats?interval_type=regularseason&season_id=nfl-2016-2017" + "&per_page=40&page="
		apiPagedURL := apiBase + strconv.Itoa(currentPage)

		data, _ := ioutil.ReadAll(routes.CallAPI(apiPagedURL))
		_, _, _, err := jsonparser.Get(data, "player_season_stats")
		if err != nil {
			fmt.Printf("Ended load at page %d", currentPage)
			break
		}

		//fmt.Printf("Loading player stats page %d\n", currentPage)

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

	db := loader.ConnectDB()

	insertStatStmt, err := db.Prepare("REPLACE INTO playerstats (pid,runs,passes,receptions) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatalf("Error preparing db statement: %s\n", err.Error())
	}

	for pid, stat := range stats {
		log.Printf("Inserting Stats %s... \n", pid)
		_, err := insertStatStmt.Exec(pid, stat.Runs, stat.Passes, stat.Receptions)
		if err != nil {
			log.Fatalf("Error inserting stat %s: %s", pid, err.Error())
		}
	}
	insertStatStmt.Close()
}
