package main

import (
	"github.com/beaunus/polyglottal.crypto-suite/ciphers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Static("/", "public")

	e.GET("/api/v1/caesar", ciphers.CaesarHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
