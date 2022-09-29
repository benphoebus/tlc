package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"

	"go-mongo-api/configs"
	"go-mongo-api/routes"
)

func main() {
	// initalise mux https://www.gorillatoolkit.org
	router := mux.NewRouter()

	// run database
	configs.ConnectDB()

	// routes 
	routes.UserRoute(router)
	
	// http.ListenAndServe function to run the application on port 6000
	log.Fatal(http.ListenAndServe(":6000", router))
	fmt.Println("Starting PORT at PORT:6000")
}
