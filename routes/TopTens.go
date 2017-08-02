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

//TopTen returnes a cached sorted or does a live sort of the top 10 players
func TopTen(w http.ResponseWriter, r *http.Request) {

	jsonPlayers, found := loader.ReadFromCache("toptenplayers")
	if found {
		fmt.Fprintf(w, string(jsonPlayers.([]byte)))
		return
	}

	db := loader.GormConnectDB()
	var players []models.Player
	var sortedPlayers []models.Player

	//db.Find(&players)
	db.Preload("Team").Find(&players)
	db.Preload("Stats").Find(&players) //Slowness culprit
	sort.Sort(ByStats(players))

	for index := 0; index < 20; index++ {
		sortedPlayers = append(sortedPlayers, players[index])
	}

	b, err := json.Marshal(sortedPlayers)
	if err != nil {
		log.Printf("Error marshalling top ten players: %s", err.Error())
	}

	fmt.Fprintf(w, string(b))

	loader.WriteToCache("toptenplayers", b)
}

type ByStats []models.Player

func (a ByStats) Len() int { return len(a) }

func (a ByStats) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByStats) Less(i, j int) bool {
	sumI := a[i].Stats.Receptions + a[i].Stats.Rushattempts
	sumJ := a[j].Stats.Receptions + a[j].Stats.Rushattempts
	return sumI > sumJ
}
