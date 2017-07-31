package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/lcanal/demspirals/models"

	"github.com/lcanal/demspirals/jobs"
	"github.com/lcanal/demspirals/loader"
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
	//muxie.HandleFunc("/api/playerstats", routes.PlayerStats)
	muxie.Handle("/", http.FileServer(http.Dir(clientFiles)))

	//Check if we want to run loads at startup...
	doLoads := flag.Bool("doloads", false, "Run initial loads for loading teams, players, stats.")
	flag.Parse()
	if *doLoads {
		fmt.Println("Creating tables ...")
		db := loader.GormConnectDB()
		db.CreateTable(&models.Stat{})
		db.CreateTable(&models.Player{})
		db.CreateTable(&models.Team{})
		fmt.Println("Loading all players....")
		go jobs.LoadAllPlayers(10)
		go jobs.LoadAllTeams()
		go jobs.LoadAllPlayerStats(10)
	}	
	log.Printf("Starting server on :%s", httpPort)
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
