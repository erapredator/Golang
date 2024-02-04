package pkg

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is my Barcelona players API.\nDone by: Kaliuev Yerlan ⚽️🔵🔴")
}