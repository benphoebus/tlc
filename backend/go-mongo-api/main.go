package main

import (
	"fmt"
	"log"
	"net/http"
	"go-mongo-api/configs"
	"github.com/gorilla/mux"
)

func main() {
	// initalise mux https://www.gorillatoolkit.org
	router := mux.NewRouter()

	// run database
	configs.ConnectDB()
	
	// http.ListenAndServe function to run the application on port 6000
	log.Fatal(http.ListenAndServe(":6000", router))
	fmt.Println("Starting PORT at PORT:6000")
}
