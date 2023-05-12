package controllers

import (
	"net/http"

	"bodotech.net/fizzbuzz/controllers/fizzbuzz"

	"github.com/gorilla/mux"
)

func HandleRequests() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	// TODO: maybe add auth method, for now this will be unsecure

	fizzBuzzCtrl := fizzbuzz.NewController()
	fizzBuzzSubRouter := router.PathPrefix("/" + fizzBuzzCtrl.GetName()).Subrouter()
	fizzBuzzCtrl.Register(fizzBuzzSubRouter)

	return router
}
