package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/beaunus/pretty-sweet-crypto-suite/ciphers"
)

func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		w.Header().Set("Expires", "0")                                         // Proxies.
		next.ServeHTTP(w, r)
	})
}

func main() {

	router := mux.NewRouter()
	router.Use(noCache)
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
