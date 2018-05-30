package ciphers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// ScytaleEncryptHandler handles requests for a Scytale Cipher
func ScytaleEncryptHandler(c echo.Context) error {
	plaintext := c.QueryParam("plaintext")
	numSides, error := strconv.Atoi(c.QueryParam("numSides"))
	if error != nil {
		fmt.Println("Scytale", error)
	}
	ciphertext := ScytaleEncrypt(plaintext, numSides)
	result := struct {
		Ciphertext string
	}{
		Ciphertext: ciphertext,
	}
	return c.JSON(http.StatusOK, result)
}

// ScytaleEncrypt is alphabet agnostic.
func ScytaleEncrypt(plaintext string, numSides int) string {
	return "Maybe"
}
