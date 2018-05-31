package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/beaunus/pretty-sweet-crypto-suite/ciphers"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/caesar", ciphers.CaesarHandler).Methods("GET")
	router.HandleFunc("/api/v1/scytale", ciphers.ScytaleHandler).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(http.ListenAndServe("localhost:8000", router))
	} else {
		log.Fatal(http.ListenAndServe(":"+port, router))
	}

}
