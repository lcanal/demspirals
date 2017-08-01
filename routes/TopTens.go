package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/lcanal/demspirals/loader"
	"github.com/lcanal/demspirals/models"
)

//TopTen Returns top ten players
func TopTen(w http.ResponseWriter, r *http.Request) {
	db := loader.GormConnectDB()
	//Just grab any ten for now
	var players []models.Player

	db.Limit(10).Find(&players)

	for slug, player := range players {
		var team models.Team
		var stat models.Stat
		db.First(&team, "id = ?", player.Teamid)
		db.First(&stat, "pid = ?", player.ID)
		players[slug].Team = team
		players[slug].Stats = stat
	}

	b, err := json.Marshal(players)
	if err != nil {
		log.Printf("Error marshalling top ten players: %s", err.Error())
	}

	fmt.Fprintf(w, string(b))
}

//TopTen2 test function
func TopTen2(w http.ResponseWriter, r *http.Request) {
	db := loader.GormConnectDB()
	//Just grab any ten for now
	var players []models.Player
	var sortedPlayers []models.Player

	db.Find(&players)

	for slug, player := range players {
		var team models.Team
		var stat models.Stat
		db.First(&team, "id = ?", player.Teamid)
		db.First(&stat, "pid = ?", player.ID)
		players[slug].Team = team
		players[slug].Stats = stat
	}

	sort.Sort(ByStats(players))

	for index := 0; index < 10; index++ {
		sortedPlayers = append(sortedPlayers, players[index])
	}

	b, err := json.Marshal(sortedPlayers)
	if err != nil {
		log.Printf("Error marshalling top ten players: %s", err.Error())
	}

	fmt.Fprintf(w, string(b))
}

type ByStats []models.Player

func (a ByStats) Len() int { return len(a) }

func (a ByStats) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByStats) Less(i, j int) bool {
	sumI := a[i].Stats.Receptions + a[i].Stats.Rushattempts
	sumJ := a[j].Stats.Receptions + a[j].Stats.Rushattempts
	return sumI > sumJ
}
