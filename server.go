package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Static("/", "public")

	e.GET("/caesarCipher", func(c echo.Context) error {
		plaintext := c.QueryParam("plaintext")
		shift, error := strconv.Atoi(c.QueryParam("shift"))
		if error != nil {
			fmt.Println("Caesar Cipher", error)
		}
		result := caesarEncrypt(plaintext, shift)
		return c.JSON(http.StatusOK, result)
	})

	e.Logger.Fatal(e.Start(":8000"))
}

// caesarEncrypt is alphabet agnostic.
// TODO: determine alphabet from input
func caesarEncrypt(plaintext string, shift int) string {
	runes := []rune(plaintext)
	for i := range runes {
		runes[i] += int32(shift)
	}
	return string(runes)
}
