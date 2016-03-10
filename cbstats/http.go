package main

import (
	"fmt"
	"net/http"
)

const port = "8080"

// RunServer starts the web server on the specified port.
func RunServer() {
	fmt.Println("Listening on port", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	GetRoomsData()
	PopulateStats()
	RenderText(w)
}
