package main

import (
	"io"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", Health).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
