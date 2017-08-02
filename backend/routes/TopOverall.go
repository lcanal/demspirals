package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lcanal/demspirals/backend/loader"
	"github.com/lcanal/demspirals/backend/models"
)

//TopOverall returnes a cached sorted or does a live sort of the top 10 players
func TopOverall(w http.ResponseWriter, r *http.Request) {
	var start int

	vars := mux.Vars(r)
	num, err := strconv.Atoi(vars["num"])
	if err != nil {
		num = 15 //Default return 15
	}

	jsonPlayers, found := loader.ReadFromCache("topoverall")
	if found {
		fmt.Fprintf(w, string(jsonPlayers.([]byte)))
		return
	}

	db := loader.GormConnectDB()
	var players []models.Player
	var sortedPlayers []models.Player

	//db.Find(&players)
	db.Preload("Team").Preload("Stats").Find(&players)
	sort.Sort(ByStats(players))

	for index := start; index < num; index++ {
		sortedPlayers = append(sortedPlayers, players[index])
	}

	b, err := json.Marshal(sortedPlayers)
	if err != nil {
		log.Printf("Error marshalling top players: %s", err.Error())
	}

	fmt.Fprintf(w, string(b))

	loader.WriteToCache("topoverall", b)
}

//ByStats is meant to be an interface to golang's sort function.
type ByStats []models.Player

func (a ByStats) Len() int { return len(a) }

func (a ByStats) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByStats) Less(i, j int) bool {
	sumI := a[i].Stats.Receptions + a[i].Stats.Rushattempts
	sumJ := a[j].Stats.Receptions + a[j].Stats.Rushattempts
	return sumI > sumJ
}
