package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/rwirdemann/restvoice/domain"
)

const baseUri = "http://localhost:8080/"
const contentType = "application/vnd.restvoice.v1.hal+json"

var client = &http.Client{}

func getActivities() []domain.Activity {
	req, _ := http.NewRequest("GET", baseUri+"activities", nil)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var activities []domain.Activity
	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &activities)
	} else {
		fmt.Printf("Got status code: %d\n", resp.StatusCode)
	}

	return activities
}

func main() {
	activities := getActivities()
	for _, a := range activities {
		fmt.Printf("%s\n", a.String())
	}
}
