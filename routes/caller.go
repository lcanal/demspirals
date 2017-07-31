package routes

import (
	"io"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

//CallAPI is the major function that does the api calls.
func CallAPI(url string) io.ReadCloser {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Token token="+viper.GetString("creds.accessToken"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/vnd.stattleship.com; version=1")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error pulling team rosters auth : %s\n", err)
	}

	if resp.StatusCode != 200 {
		log.Printf("Error with request: %s (%d)\n", resp.Status, resp.StatusCode)
	}
	return resp.Body
}
