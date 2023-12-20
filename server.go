package main

import (
	"groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {
	port := ":8040"
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/artist/", handlers.Artists)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	log.Printf("Connect to our website throught http://localhost%s", port)
	http.ListenAndServe(port, nil)
}
