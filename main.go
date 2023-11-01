package main

import (
	"net/http"

	"github.com/AKaganKaya/image-playground/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve static files (optional)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/", handlers.UploadForm).Methods("GET")
	r.HandleFunc("/upload", handlers.UploadImage).Methods("POST")

	http.ListenAndServe(":8080", r)
}
