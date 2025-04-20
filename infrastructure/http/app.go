package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/about", AboutHandler).Methods("GET")

	router.HandleFunc("/projects/html", ProjectsJSONHandler).Methods("GET")
	router.HandleFunc("/contact", ContactHandler).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
