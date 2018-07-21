package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"github.com/rwirdemann/crudvoice.v2/foundation"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

func MakeGetInvoicesHandler(usecase foundation.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(usecase.Run(r).([]byte))
	}
}

func MakeGetInvoiceHandler(usecase foundation.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(usecase.Run(r).([]byte))
	}
}

func MakeCreateInvoiceHandler(usecase foundation.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/vnd.restvoice+json")
		body, _ := ioutil.ReadAll(r.Body)
		w.Write(usecase.Run(body).([]byte))
	}
}
