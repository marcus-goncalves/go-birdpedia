package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := routesBuilder()

	http.ListenAndServe(":8000", r)
}

func routesBuilder() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome).Methods("GET")

	// Cfg static files
	staticFileDir := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
