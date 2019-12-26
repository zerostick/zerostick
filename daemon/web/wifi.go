package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func wifi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	log.Println(vars["id"])

}
