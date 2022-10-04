package routes

import (
	"github.com/gorilla/mux"
	"go-mongo-api/controllers"
) 

func UserRoute(router *mux.Router) {
	// All routes related to users comes here
	router.HandleFunc("/user", controllers.CreateUser()).Methods("POST")
	router.HandleFunc("/user/{userId}", controllers.GetAUser()).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.EditAUser()).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteAUser()).Methods("DELETE")
	router.HandleFunc("/users", controllers.GetAllUser()).Methods("GET")
}