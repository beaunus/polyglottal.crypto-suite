package ciphers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// ScytaleHandler handles requests for a Scytale Cipher
func ScytaleHandler(c echo.Context) error {
	plaintext := c.QueryParam("plaintext")
	ciphertext := c.QueryParam("ciphertext")
	numSides, error := strconv.Atoi(c.QueryParam("numSides"))
	if error != nil {
		fmt.Println("Scytale", error)
	}
	type result struct {
		Ciphertext string
		Plaintext  string
	}
	var resultCiphertext, resultPlaintext string
	if len(plaintext) > 0 {
		resultCiphertext = ScytaleEncrypt(plaintext, numSides)
	} else {
		resultPlaintext = ScytaleDecrypt(ciphertext, numSides)
	}
	return c.JSON(http.StatusOK, result{
		Ciphertext: resultCiphertext,
		Plaintext:  resultPlaintext,
	})
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
