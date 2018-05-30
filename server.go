package main

import (
	"os"

	"github.com/beaunus/polyglottal.crypto-suite/ciphers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Static("/", "public")

	e.GET("/api/v1/caesar", ciphers.CaesarHandler)

	port := os.Getenv("PORT")
	if port == "" {
		e.Logger.Fatal(e.Start("localhost:8000"))
	} else {

		e.Logger.Fatal(e.Start(":" + port))
	}

}
