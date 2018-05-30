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
	plaintextRunes := []rune(plaintext)
	resultLength := len(plaintextRunes)
	result := make([]rune, resultLength)

	fromIndex := 0
	i := 0
	for i < resultLength {
		toIndex := i
		for toIndex < resultLength && fromIndex < resultLength {
			result[toIndex] = plaintextRunes[fromIndex]
			toIndex += numSides
			fromIndex++
		}
		i++
	}

	return string(result)
}

// ScytaleDecrypt is alphabet agnostic.
func ScytaleDecrypt(ciphertext string, numSides int) string {
	ciphertextRunes := []rune(ciphertext)
	resultLength := len(ciphertextRunes)
	result := make([]rune, resultLength)

	fromIndex := 0
	i := 0
	for i < resultLength {
		toIndex := i
		for toIndex < resultLength && fromIndex < resultLength {
			result[fromIndex] = ciphertextRunes[toIndex]
			toIndex += numSides
			fromIndex++
		}
		i++
	}

	return string(result)
}
