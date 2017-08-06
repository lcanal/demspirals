package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lcanal/demspirals/backend/loader"
)

//TopOverall returnes a cached sorted or does a live sort of the top 10 players
func TopOverall(w http.ResponseWriter, r *http.Request) {
	//var start int

	/*vars := mux.Vars(r)
	num, err := strconv.Atoi(vars["num"])
	if err != nil {
		num = 15 //Default return 15
	}*/

	//Caching for results
	jsonPlayers, found := loader.ReadFromCache("topoverall")
	if found {
		fmt.Fprintf(w, string(jsonPlayers.([]byte)))
		return
	}

	db := loader.GormConnectDB()
	//db.LogMode(true)

	//var players []models.Player
	//var points []models.Point

	type result struct {
		ID                 string
		LastName           string
		FirstName          string
		Position           string
		TotalFantasyPoints float64
	}

	var results []result

	topQuery := `
	SELECT 
		players.id,
		players.last_name,
		players.first_name,
		players.position,
		SUM(points.value) total_fantasy_points
	FROM
    	players
	LEFT JOIN 
    	points
	ON 
    	players.id = points.player_id
	GROUP BY
    	players.id
	ORDER BY
    	total_fantasy_points
	DESC
	`

	db.Raw(topQuery).Scan(&results)

	//sort.Sort(ByStats(players))

	/*for index := start; index < num; index++ {
		sortedPlayers = append(sortedPlayers, players[index])
	}*/
	/*or _, player := range players {
		//db.Model(&player).Related(&stats)
		//log.Fatalf("waaaa\n")
		if player.ID == "7549" {
			fmt.Printf("Here B Brady Bunch \n%v", player)
		}
	}*/

	b, err := json.Marshal(results)
	if err != nil {
		log.Printf("Error marshalling top players: %s", err.Error())
	}

	fmt.Fprintf(w, string(b))

	loader.WriteToCache("topoverall", b)
}

//ByStats is meant to be an interface to golang's sort function.
/*type ByStats []models.Player

func (a ByStats) Len() int { return len(a) }

func (a ByStats) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByStats) Less(i, j int) bool {
	sumI := a[i].Stats.Receptions + a[i].Stats.Rushattempts
	sumJ := a[j].Stats.Receptions + a[j].Stats.Rushattempts
	return sumI > sumJ
}*/
