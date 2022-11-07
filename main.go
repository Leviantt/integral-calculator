package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public/"))

	router.HandleFunc("/api/calculate-integral", HandleCalculateIntegral).Methods("POST")
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	http.ListenAndServe(":4000", router)
}
