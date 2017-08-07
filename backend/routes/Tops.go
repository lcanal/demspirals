package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lcanal/demspirals/backend/loader"
)

//TopOverall returnes a cached sorted or does a live sort of the top 10 players
func TopOverall(w http.ResponseWriter, r *http.Request) {
	cacheKey := "topplayers"
	vars := mux.Vars(r)
	posFilter := vars["position"]

	switch posFilter {
	case "rb":
		cacheKey = "toprb"
		posFilter = "WHERE players.position IN ('RB')"
	case "qb":
		cacheKey = "topqb"
		posFilter = "WHERE players.position IN ('QB')"
	case "wr":
		cacheKey = "topwr"
		posFilter = "WHERE players.position IN ('WR')"
	default:
		posFilter = ""
	}

	//Caching for results
	jsonPlayers, found := loader.ReadFromCache(cacheKey)
	if found {
		fmt.Fprintf(w, string(jsonPlayers.([]byte)))
		return
	}

	db := loader.GormConnectDB()
	//db.LogMode(true)

	type result struct {
		ID                 string  `json:"id"`
		Age                string  `json:"age"`
		LastName           string  `json:"lastname"`
		FirstName          string  `json:"firstname"`
		JerseyNumber       string  `json:"jerseynum"`
		Position           string  `json:"position"`
		PicURL             string  `json:"picurl"`
		Height             string  `json:"height"`
		Weight             string  `json:"weight"`
		Rookie             bool    `json:"rookie"`
		NFLID              string  `json:"nflid"`
		TeamID             string  `json:"teamid"`
		TeamName           string  `json:"teamname"`
		TeamCity           string  `json:"teamcity"`
		TotalFantasyPoints float64 `json:"totalfantasypoints"`
	}

	var results []result

	topQuery := `
	SELECT 
		players.*,
        teams.name as team_name,
        teams.city as team_city,
		SUM(points.value) total_fantasy_points
	FROM
    	players
	LEFT JOIN 
    	points
	ON 
    	players.id = points.player_id
    LEFT JOIN
    	teams
    ON 
		players.team_id = teams.id
	`
	topQuery = topQuery + posFilter + `
	GROUP BY
    	players.id
	ORDER BY
    	total_fantasy_points
	DESC
	`

	//log.Printf("TopQuery : %s\n", topQuery)
	db.Raw(topQuery).Scan(&results)

	/*b, err := json.Marshal(results)
	if err != nil {
		log.Printf("Error marshalling top players: %s", err.Error())
	}*/

	data := make(map[string]interface{})
	data["playerdata"] = results
	data["datatype"] = cacheKey
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshalling top players: %s", err.Error())
		return
	}

	fmt.Fprintf(w, string(b))
	loader.WriteToCache(cacheKey, b, 6*time.Hour)
}

//PlayerInfo returns proper headers for overall players.
func PlayerInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid := vars["pid"]

	//Caching for results
	jsonPlayerInfo, found := loader.ReadFromCache(pid)
	if found {
		fmt.Fprintf(w, string(jsonPlayerInfo.([]byte)))
		return
	}

	type result struct {
		ID           string
		PlayerID     string
		Category     string
		Abbreviation string
		Name         string
		LeagueName   string
		StatID       string
		StatNum      float64
		Value        float64
	}

	var results []result

	playerQuery := `
	SELECT 
    	points.*
	FROM players 
	JOIN points
	ON players.id = points.player_id
	WHERE
		players.id IN ('` + pid + `')
    AND points.category = "Receiving"
	`

	db := loader.GormConnectDB()
	//db.LogMode(true)
	//log.Printf("playerQuery : %s\n", playerQuery)
	db.Raw(playerQuery).Scan(&results)
	b, err := json.Marshal(results)
	if err != nil {
		log.Printf("Error marshalling top players: %s", err.Error())
		return
	}
	fmt.Fprintf(w, string(b))
	db.Close()
	loader.WriteToCache(pid, b, 4*time.Hour)
}
