package main

import (
	"crud/servidor"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandFunc("/usuarios", servidor.CriaUsuario).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":5000", router))
}
