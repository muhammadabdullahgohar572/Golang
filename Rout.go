package main

import (
	"log"
	"github.com/gorilla/mux"
	
)

func rout() {
	r := mux.NewRouter()
	r.HandleFunc("/Empyoee",CreateEmployee)
	log.Fatal(":8080",r)
}