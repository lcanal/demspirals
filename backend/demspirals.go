package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lcanal/demspirals/backend/jobs"
	"github.com/lcanal/demspirals/backend/loader"
	"github.com/lcanal/demspirals/backend/models"
	"github.com/lcanal/demspirals/backend/routes"
	"github.com/spf13/viper"
)

func main() {
	//Read in and set all settings
	configRead("settings")
	httpPort := viper.GetString("httpPort")
	frontendFiles := viper.GetString("frontendFiles")

	muxie := mux.NewRouter()

	muxie.HandleFunc("/api/hello", hello)
	muxie.HandleFunc("/api/topoverall", routes.TopOverall)
	muxie.HandleFunc("/api/topoverall/{num}", routes.TopOverall)
	muxie.PathPrefix("/").Handler(http.FileServer(http.Dir("./" + frontendFiles)))
	http.Handle("/", muxie)

	//Check if we want to run loads at startup...
	doLoads := flag.Bool("doloads", false, "Run initial loads for loading teams, players, stats.")
	dropTables := flag.Bool("droptables", false, "Drop tables. Must be set along with doloads to run.")

	flag.Parse()
	if *doLoads {
		db := loader.GormConnectDB()

		if *dropTables {
			fmt.Println("Droping old tables ...")
			db.DropTableIfExists(&models.Stat{})
			db.DropTableIfExists(&models.Player{})
			db.DropTableIfExists(&models.Team{})
		}

		fmt.Println("Creating new tables...")
		//db.CreateTable(&models.Stat{})
		db.CreateTable(&models.Player{})
		//db.CreateTable(&models.Team{})

		fmt.Println("Loading all players....")
		go jobs.LoadAllPlayers()
		//fmt.Println("Loading all teams... ")
		//go jobs.LoadAllTeams()
		//fmt.Println("Loading all stats... ")
		//go jobs.LoadAllPlayerStats()
	}

	log.Printf("Starting server on :%s", httpPort)
	log.Printf("Loading frontend files from %s", frontendFiles)
	log.Printf("Using database %s:%s", viper.GetString("db.host"), viper.GetString("db.port"))

	log.Fatal(http.ListenAndServe(":"+httpPort, muxie))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"hello\":\"world\"}")
}

func configRead(configName string) (string, string) {
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")

	viper.SetDefault("frontendFiles", "frontend")
	viper.SetDefault("httpPort", "8080")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s", err.Error())
	}

	apiBaseURL := viper.GetString("apiBaseURL")
	if len(apiBaseURL) <= 0 {
		log.Fatal("Error: No api base url set (\"apiBaseURL\")")
	}

	accessToken := viper.GetString("creds.user")
	if len(accessToken) <= 0 {
		log.Fatal("Error: No access token set (\"creds.user\")")
	}

	dbHost := viper.GetString("db.host")
	if len(dbHost) <= 0 {
		log.Fatal("Error: No db listed in settings file.")
	}
	//Do a init connection
	loader.GormConnectDB()

	return apiBaseURL, accessToken
}
