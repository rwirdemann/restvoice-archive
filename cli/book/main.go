package book

import (
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/rwirdemann/restvoice/rest"
	"github.com/rwirdemann/restvoice/domain"
)

const baseUri = "http://localhost:8080/"
const contentType = "application/vnd.restvoice.v1.hal+json"

var client = &http.Client{}

func createInvoice(customer int, year int, month int) []byte {
	i := domain.Invoice{CustomerId: customer, Year: year, Month: month}
	jsonStr, _ := json.Marshal(i)
	req, _ := http.NewRequest("POST", baseUri+"invoice", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

type Booking struct {
	Hours       float32 `json:"hours"`
	Description string  `json:"description"`
	Day         int     `json:"day"`
}

type Operations struct {
	Links map[string]rest.Link `json:"_links"`
}

var ops = Operations{Links: make(map[string]rest.Link)}

func main() {
	body := createInvoice(1, 2018, 8)
	json.Unmarshal(body, &ops)
	if _, ok := ops.Links["book"]; ok {
		book(4.5, "NRG-333 Benutzersynchronisation", 17)
	}
}

func book(hours float32, description string, date int) {
	b := Booking{Hours: hours, Description: description, Day: date}
	jsonStr, _ := json.Marshal(b)
	uri := baseUri + ops.Links["book"].Href
	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", contentType)
	resp, _ := client.Do(req);
	defer resp.Body.Close()
	fmt.Printf("Status: %s", resp.Status)
}
