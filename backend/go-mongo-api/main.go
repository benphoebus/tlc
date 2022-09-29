package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// initalise mux https://www.gorillatoolkit.org
	router := mux.NewRouter()

	// function that uses net/http package as parameters to route to / path and a handler function that sets the header type to a JSON and returns a JSON of Hello from Mux & mongoDB using the encoding/json package
	// HTTP "GET" method attached to this function
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"data": "Hello from mux & mongodb"})
	}).Methods("GET")
	
	// http.ListenAndServe function to run the application on port 6000
	log.Fatal(http.ListenAndServe(":6000", router))
	fmt.Println("Starting PORT at PORT:6000")
}
