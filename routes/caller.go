package routes

import (
	"io"
	"log"
	"net/http"
	"time"

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
		log.Printf("Error general error attempting REST call to (%s): %s\n", url, err)
	}

	if resp.StatusCode != 200 {
		log.Printf("Error with request: %s\nURL: %s\n", resp.Status, url)
		//Retry 5 times
		for index := 0; index < 5; index++ {
			log.Printf("Retrying %d time(s)", index+1)
			time.Sleep(3 * time.Second)
			resp, _ := client.Do(req)
			if resp.StatusCode == 200 {
				return resp.Body
			}
		}

	}
	return resp.Body
}
