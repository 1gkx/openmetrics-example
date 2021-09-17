package router

import (
	"github.com/1gkx/openmetrics/internal/yml"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data, err := yml.ParseFile("currencies.yaml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data.Unmarshal())
	})

	return r
}
