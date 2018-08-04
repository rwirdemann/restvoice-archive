package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
		"github.com/rwirdemann/restvoice/domain"
	"fmt"
	"flag"
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

func main() {
	customerId := flag.Int("customer", 1, "customer id")
	year := flag.Int("year", 2018, "year")
	month := flag.Int("month", 8, "month")
	flag.Parse()

	body := createInvoice(*customerId, *year, *month)
	var created domain.Invoice
	json.Unmarshal(body, &created)
	fmt.Println("Id: ", created.Id)
}
