package routes
import (
		"net/http"
		"encoding/json"
		"log"
		"github.com/lcanal/demspirals/loader"
		"github.com/lcanal/demspirals/models"
		"fmt"
)
//TopTen Returns top ten players
func TopTen(w http.ResponseWriter, r *http.Request) {
	db := loader.GormConnectDB()
	//Just grab any ten for now
	var players []models.Player
	db.Limit(10).Find(&players)
	b,err := json.Marshal(players)
	if err != nil {
		log.Printf("Error marshalling top ten players: %s",err.Error())
	}

	fmt.Fprintf(w,string(b))
}