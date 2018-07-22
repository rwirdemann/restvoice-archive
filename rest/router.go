package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"github.com/rwirdemann/crudvoice.v2/foundation"
)

const contentType = "application/vnd.restvoice.v1.hal+json"

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

func MakeGetInvoiceHandler(usecase foundation.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Write(usecase.Run(r, r).([]byte))
	}
}

func MakeCreateInvoiceHandler(usecase foundation.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		body, _ := ioutil.ReadAll(r.Body)
		w.Write(usecase.Run(body).([]byte))
	}
}
