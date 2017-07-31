package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lcanal/demspirals/routes"
	"github.com/spf13/viper"
)

func main() {
	//Read in and set all settings
	configRead("settings")
	httpPort := viper.GetString("httpPort")
	clientFiles := viper.GetString("clientFiles")

	muxie := http.NewServeMux()
	muxie.HandleFunc("/api/hello", hello)
	//muxie.HandleFunc("/api/teams", routes.TeamRoster)
	muxie.HandleFunc("/api/playerstats", routes.PlayerStats)
	muxie.Handle("/", http.FileServer(http.Dir(clientFiles)))

	//fmt.Println("Loading all players....")
	//routes.LoadAllPlayers()

	log.Fatal(http.ListenAndServe(":"+httpPort, muxie))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"hello\":\"world\"}")
}

func configRead(configName string) (string, string) {
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")

	viper.SetDefault("clientFiles", "client/build")
	viper.SetDefault("httpPort", "8080")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s", err.Error())
	}

	apiBaseURL := viper.GetString("apiBaseURL")
	if len(apiBaseURL) <= 0 {
		log.Fatal("Error: No api base url set (\"apiBaseURL\")")
	}

	accessToken := viper.GetString("creds.accessToken")
	if len(accessToken) <= 0 {
		log.Fatal("Error: No access token set (\"accessToken\")")
	}

	return apiBaseURL, accessToken
}
