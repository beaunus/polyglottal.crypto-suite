package ciphers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode"

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
	fmt.Println("numSides", numSides)
	noSpacePlaintextRunes := []rune(strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, plaintext))

	// Assemble the runes into their appropriate columns.

	resultLength := len(noSpacePlaintextRunes)
	result := make([]rune, resultLength)

	fromIndex := 0
	for fromIndex < resultLength {
		i := 0
		fmt.Println("i", i)
		for i < resultLength {
			toIndex := i
			for toIndex < resultLength && fromIndex < resultLength {
				fmt.Println("fromIndex", fromIndex)
				fmt.Println("toIndex", toIndex)
				result[toIndex] = noSpacePlaintextRunes[fromIndex]
				toIndex += numSides
				fromIndex++
			}
			i++
		}
	}

	return string(result)
}
