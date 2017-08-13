package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/lcanal/demspirals/backend/jobs"
	"github.com/lcanal/demspirals/backend/loader"
	"github.com/lcanal/demspirals/backend/models"
	"github.com/lcanal/demspirals/backend/routes"
	cache "github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
)

//Global cache
var mainCache = cache.New(5*time.Minute, 10*time.Minute)

func main() {
	//Read in and set all settings
	configRead("settings")
	httpPort := viper.GetString("httpPort")
	frontendFiles := viper.GetString("frontendFiles")
	loader.MainCache = mainCache

	muxie := mux.NewRouter()

	fe := http.FileServer(http.Dir("./" + frontendFiles))

	muxie.HandleFunc("/api/topplayers", routes.TopOverall)
	muxie.HandleFunc("/api/topplayers/{position}", routes.TopOverall)
	muxie.HandleFunc("/api/player/{position}/{pid}", routes.PlayerInfo)
	muxie.PathPrefix("/").Handler(fe)

	//Flags //////////////////////////////////////////////
	doLoads := flag.Bool("doloads", false, "Run initial loads for loading teams, players, stats.")
	dropTables := flag.Bool("droptables", false, "Drop tables. Must be set along with doloads to run.")
	calcPoints := flag.Bool("calcpoints", false, "Calculate fantasy points for all players.")
	init := flag.Bool("init", false, "Initialize all of your data. Should be run when you want to completely redo and recalculate all of your data.")
	//////////////////////////////////////////////////////

	//Global waiting group to coordinate functions. Makes sure to only calc points when all
	//Relevant functions have finished.
	wg := new(sync.WaitGroup)

	flag.Parse()
	if *init {
		*doLoads = true
		*dropTables = true
		*calcPoints = true
	}
	if *doLoads {
		db := loader.GormConnectDB()

		if *dropTables {
			fmt.Println("Droping old tables ...")
			db.DropTableIfExists(&models.Stat{})
			db.DropTableIfExists(&models.Player{})
			db.DropTableIfExists(&models.Team{})
		}

		fmt.Println("Creating new tables...")
		db.CreateTable(&models.Stat{})
		db.CreateTable(&models.Player{})
		db.CreateTable(&models.Team{})

		fmt.Println("Loading all players and teams....")
		wg.Add(1) //Tell other groups waiting on functions that this one counts as a wait for
		go jobs.LoadAllPlayerData(wg)
	}

	if *calcPoints {
		go jobs.CalculatePoints(wg)
	}

	log.Printf("Starting server on :%s", httpPort)
	log.Printf("Loading frontend files from %s", frontendFiles)
	log.Printf("Using database %s:%s", viper.GetString("db.host"), viper.GetString("db.port"))

	log.Fatal(http.ListenAndServe(":"+httpPort, muxie))
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
