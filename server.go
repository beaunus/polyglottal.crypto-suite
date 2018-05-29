package main

import (
	"github.com/beaunus/polyglottal.crypto-suite/ciphers"
	"github.com/graph-gophers/graphql-go"
	"github.com/labstack/echo"
)

var schema *graphql.Schema

func main() {
	e := echo.New()

	e.Static("/", "public")

	e.GET("/caesarCipher", ciphers.CaesarHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
