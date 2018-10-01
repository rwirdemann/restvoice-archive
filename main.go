package main

import (
	"github.com/rwirdemann/restvoice/database"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"github.com/rwirdemann/restvoice/domain"
	"github.com/rwirdemann/restvoice/rest"
	"github.com/rwirdemann/restvoice/usecase"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"context"
)

var (
	fileName   string
	githash    string
	buildstamp string
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Adapter Layer
	invoiceConsumer := rest.NewJSONConsumer(&domain.Invoice{})
	bookingConsumer := rest.NewJSONConsumer(&domain.Booking{})
	invoiceIdConsumer := rest.NewPathVariableConsumer("id")
	expandConsumer := rest.NewQueryVariableConsumer("expand")
	invoicePresenter := rest.NewHALInvoicePresenter()
	present := rest.NewJSONPresenter()
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
	updateInvoice := usecase.NewUpdateInvoice(invoiceConsumer, invoiceIdConsumer, repository)
	createBooking := usecase.NewCreateBooking(bookingConsumer, invoiceIdConsumer, present, repository)
	getInvoice := usecase.NewGetInvoice(invoiceIdConsumer, expandConsumer, invoicePresenter, repository)
	getActivities := usecase.NewGetActivities(userConsumer, activitiesPresenter, repository)

	// HTTP
	r := mux.NewRouter()
	r.HandleFunc("/version", rest.MakeVersionHandler(githash, buildstamp)).Methods("GET")
	r.HandleFunc("/invoice", rest.JwtAuth(rest.MakeCreateInvoiceHandler(createInvoice))).Methods("POST")
	r.HandleFunc("/booking/{id:[0-9]+}", rest.JwtAuth(rest.MakeCreateBookingHandler(createBooking))).Methods("POST")
	r.HandleFunc("/charge/{id:[0-9]+}", rest.JwtAuth(rest.MakeUpdateInvoiceHandler(updateInvoice))).Methods("PUT")
	r.HandleFunc("/invoice/{id:[0-9]+}", rest.MakeGetInvoiceHandler(getInvoice)).Methods("GET")
	r.HandleFunc("/activities", rest.MakeGetActivitiesHandler(getActivities)).Methods("GET")
	fmt.Println("POST http://localhost:8080/invoice")
	fmt.Println("GET http://localhost:8080/invoice/{id}")
	fmt.Println("POST http://localhost:8080/booking/1")
	fmt.Println("GET http://localhost:8080/activities")

	logger := log.New(os.Stdout, "", 0)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	server := &http.Server{Addr: ":8080", Handler: r}
	go func() {
		logger.Printf("Listening on http://0.0.0.0%s\n", ":8080")
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()
	<-stop

	logger.Println("\nShutting down the server...")
	server.Shutdown(context.Background())
	logger.Println("Server gracefully stopped")
}
