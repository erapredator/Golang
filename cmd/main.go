package main

import (
	"net/http"
	"Tsis01/pkg"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", pkg.HealthCheck).Methods("GET")
	router.HandleFunc("/players", pkg.Players).Methods("GET")
	router.HandleFunc("/players/{name}", pkg.PlayerByName).Methods("GET")

	http.Handle("/", router)

	http.ListenAndServe(":8080", router)

}
