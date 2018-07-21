package main

import (
	"github.com/rwirdemann/restvoice/database"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"github.com/rs/cors"
	"github.com/rwirdemann/restvoice/domain"
	"github.com/rwirdemann/restvoice/rest"
	"github.com/rwirdemann/restvoice/usecase"
)

func main() {
	invoiceConsumer := rest.NewJSONConsumer(&domain.Invoice{})
	invoicePresenter := rest.NewRVInvoicePresenter()

	repository := database.NewMySQLRepository()
	createInvoice := usecase.NewCreateInvoice(invoiceConsumer, invoicePresenter, repository)

	r := mux.NewRouter()
	r.HandleFunc("/invoice", rest.MakeCreateInvoiceHandler(createInvoice)).Methods("POST")

	fmt.Println("POST http://localhost:8080/invoice")

	http.ListenAndServe(":8080", cors.AllowAll().Handler(r))
}
