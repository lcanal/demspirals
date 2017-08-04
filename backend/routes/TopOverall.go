package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lcanal/demspirals/backend/loader"
	"github.com/lcanal/demspirals/backend/models"
)

//TopOverall returnes a cached sorted or does a live sort of the top 10 players
func TopOverall(w http.ResponseWriter, r *http.Request) {
	//var start int

	/*vars := mux.Vars(r)
	num, err := strconv.Atoi(vars["num"])
	if err != nil {
		num = 15 //Default return 15
	}

	//Elim caching for now
	/*jsonPlayers, found := loader.ReadFromCache("topoverall")
	if found {
		fmt.Fprintf(w, string(jsonPlayers.([]byte)))
		return
	}*/

	db := loader.GormConnectDB()
	db.LogMode(true)
	var players []models.Player
	var sortedPlayers []models.Player
	var stats []models.Stat

	db.Find(&stats)
	db.Preload("Team").Find(&players)
	//sort.Sort(ByStats(players))

	/*for index := start; index < num; index++ {
		sortedPlayers = append(sortedPlayers, players[index])
	}*/
	for _, player := range players {
		db.Model(&player).Related(&stats)
		//log.Fatalf("waaaa\n")
		if player.ID == "7549" {
			fmt.Printf("Here B Brady Bunch \n%v", player)
		}
	}

	b, err := json.Marshal(sortedPlayers)
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
