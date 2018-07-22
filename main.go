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
	// Adapter Layer
	invoiceConsumer := rest.NewJSONConsumer(&domain.Invoice{})
	invoiceIdConsumer := rest.NewPathVariableConsumer("id")
	expandConsumer := rest.NewQueryVariableConsumer("expand")
	invoicePresenter := rest.NewHALInvoicePresenter()
	repository := database.NewMySQLRepository()

	// Testdata
	invoice := repository.CreateInvoice(&domain.Invoice{Year: 2018, Month: 12})
	repository.CreateBooking(domain.Booking{InvoiceId: invoice.Id, Hours: 2, Description: "Programmiert", Day: 12, ActivityId: 1, ProjectId: 1})

	// Usecase Layer
	createInvoice := usecase.NewCreateInvoice(invoiceConsumer, invoicePresenter, repository)
	getInvoice := usecase.NewGetInvoice(invoiceIdConsumer, expandConsumer, invoicePresenter, repository)

	// HTTP
	r := mux.NewRouter()
	r.HandleFunc("/invoice", rest.MakeCreateInvoiceHandler(createInvoice)).Methods("POST")
	r.HandleFunc("/invoice/{id:[0-9]+}", rest.MakeGetInvoiceHandler(getInvoice)).Methods("GET")
	fmt.Println("POST http://localhost:8080/invoice")
	fmt.Println("GET http://localhost:8080/invoice/{id}")
	http.ListenAndServe(":8080", cors.AllowAll().Handler(r))
}
