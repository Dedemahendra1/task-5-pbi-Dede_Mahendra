package main

import (
	"log"
	"net/http"

	"github.com/Dedemahendra1/go_tutorial/controller/authcontroller"
	"github.com/Dedemahendra1/go_tutorial/controller/photocontroller"
	"github.com/Dedemahendra1/go_tutorial/middlewares"
	"github.com/Dedemahendra1/go_tutorial/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/photo", photocontroller.Index).Methods("GET")
	api.HandleFunc("/photo/{id}", photocontroller.Show).Methods("GET")
	api.HandleFunc("/photo", photocontroller.Create).Methods("POST")
	api.HandleFunc("/photo/{id}", photocontroller.Update).Methods("PUT")
	api.HandleFunc("/photo", photocontroller.Delete).Methods("DELETE")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
