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
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Adapter Layer
	invoiceConsumer := rest.NewJSONConsumer(&domain.Invoice{})
	invoiceIdConsumer := rest.NewPathVariableConsumer("id")
	expandConsumer := rest.NewQueryVariableConsumer("expand")
	invoicePresenter := rest.NewHALInvoicePresenter()
	userConsumer := rest.NewUserConsumer()
	activitiesPresenter := rest.NewActivitiesPresenter()
	repository := database.NewMySQLRepository()

	// Testdata
	invoice := repository.CreateInvoice(&domain.Invoice{Year: 2018, Month: 12})
	booking := domain.Booking{InvoiceId: invoice.Id, Hours: 2, Description: "Tests refaktorisiert", Day: 12, ActivityId: 1, ProjectId: 1}
	repository.CreateBooking(booking)
	acitivty := domain.Activity{UserId: "1234", Name: "Programmierung"}
	repository.CreateActivity(acitivty)

	// Usecase Layer
	createInvoice := usecase.NewCreateInvoice(invoiceConsumer, invoicePresenter, repository)
	getInvoice := usecase.NewGetInvoice(invoiceIdConsumer, expandConsumer, invoicePresenter, repository)
	getActivities := usecase.NewGetActivities(userConsumer, activitiesPresenter, repository)

	// HTTP
	r := mux.NewRouter()
	r.HandleFunc("/invoice", rest.JwtAuth(rest.MakeCreateInvoiceHandler(createInvoice))).Methods("POST")
	r.HandleFunc("/invoice/{id:[0-9]+}", rest.MakeGetInvoiceHandler(getInvoice)).Methods("GET")
	r.HandleFunc("/activities", rest.MakeGetActivitiesHandler(getActivities)).Methods("GET")
	fmt.Println("POST http://localhost:8080/invoice")
	fmt.Println("GET http://localhost:8080/invoice/{id}")
	fmt.Println("GET http://localhost:8080/activities")
	http.ListenAndServe(":8080", cors.AllowAll().Handler(r))
}
