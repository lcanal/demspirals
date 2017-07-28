package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	//Read in and set all settings
	configRead("settings")
	httpPort := viper.GetString("httpPort")
	clientFiles := viper.GetString("clientFiles")

	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", hello)

	mux.Handle("/", http.FileServer(http.Dir(clientFiles)))
	log.Fatal(http.ListenAndServe(":"+httpPort, mux))
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
		panic(fmt.Errorf("Fata error config file: %s", err))
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
