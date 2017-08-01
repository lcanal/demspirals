package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
