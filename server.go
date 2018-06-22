package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/beaunus/pretty-sweet-crypto-suite/ciphers"
)

type CipherPageData struct {
	Ciphers []Cipher
}

type Cipher struct {
	Title         string
	ID            string
	CipherMethods []CipherMethod
}

type CipherMethod struct {
	Name       string
	Parameters []Parameter
}

type Parameter struct {
	DefaultValue string
	FriendlyName string
	Placeholder  string
	Type         string
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/caesar", ciphers.CaesarHandler).Methods("GET")
	router.HandleFunc("/api/v1/scytale", ciphers.ScytaleHandler).Methods("GET")
	tmpl := template.Must(template.ParseFiles("public/views/ciphers.html"))
	router.HandleFunc("/ciphers", func(w http.ResponseWriter, r *http.Request) {
		data := CipherPageData{
			Ciphers: []Cipher{
				{Title: "Caesar Cipher",
					ID: "caesar", CipherMethods: []CipherMethod{
						{Name: "Encrypt", Parameters: []Parameter{
							{
								FriendlyName: "Plaintext",
								Placeholder:  "Characters not in the alphabet will be ignored.",
								Type:         "textarea",
							},
							{
								DefaultValue: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
								FriendlyName: "Alphabet",
								Type:         "textarea",
							},
						},
						},
					},
				},
			},
		}
		tmpl.Execute(w, data)
	})

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(http.ListenAndServe("localhost:8000", router))
	} else {
		log.Fatal(http.ListenAndServe(":"+port, router))
	}

}
