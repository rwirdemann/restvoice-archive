package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"time"
	"strings"
	"github.com/rwirdemann/restvoice/foundation"
	"fmt"
)

const contentType = "application/vnd.restvoice.v1.hal+json"

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

func MakeVersionHandler(githash string, buildstamp string) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Githash: %s Buildstamp: %s\n", githash, buildstamp)))
	}
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

func MakeCreateBookingHandler(usecase foundation.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		body, _ := ioutil.ReadAll(r.Body)
		w.Write(usecase.Run(body, r).([]byte))
	}
}

const layout = "Mon, _2 Jan 2006 15:04:05 GMT"

func MakeGetActivitiesHandler(usecase foundation.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Cache-Control", "public, max-age=0")
		response := usecase.Run(r.Header).(CacheableActivities)

		// Test if client wants a full refresh
		cacheControl := r.Header.Get("Cache-Control")
		if strings.Contains(cacheControl, "no-cache") {
			w.Header().Set("Last-Modified", response.LastModified.Format(layout))
			w.Write(response.Activities)
		}

		// Cache logic: return 304 if nothing has changed since "Last-Modified-Since"
		lastModifiedSince := r.Header.Get("Last-Modified-Since")
		if lastModifiedSince != "" {
			t, err := time.Parse(layout, lastModifiedSince)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if truncateToSeconds(t).Equal(truncateToSeconds(response.LastModified)) {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}

		w.Header().Set("Last-Modified", response.LastModified.Format(layout))
		w.Write(response.Activities)
	}
}

func truncateToSeconds(t time.Time) time.Time {
	return t.Truncate(time.Duration(time.Second))
}

