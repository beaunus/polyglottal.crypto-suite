package ciphers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// CaesarHandler handles requests for a Caesar Cipher
func CaesarHandler(c echo.Context) error {
	plaintext := c.QueryParam("plaintext")
	shift, error := strconv.Atoi(c.QueryParam("shift"))
	if error != nil {
		fmt.Println("Caesar Cipher", error)
	}
	ciphertext := caesarEncrypt(plaintext, shift)
	result := struct {
		Ciphertext string
	}{
		Ciphertext: ciphertext,
	}
	return c.JSON(http.StatusOK, result)
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
