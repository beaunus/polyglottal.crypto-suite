package ciphers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

// CaesarHandler handles requests for a Caesar Cipher
func CaesarHandler(c echo.Context) error {
	plaintext := c.QueryParam("plaintext")
	alphabet := c.QueryParam("alphabet")
	shift, error := strconv.Atoi(c.QueryParam("shift"))
	if error != nil {
		fmt.Println("Caesar", error)
	}
	ciphertext := CaesarEncrypt(plaintext, alphabet, shift)
	result := struct {
		Ciphertext string
	}{
		Ciphertext: ciphertext,
	}
	return c.JSON(http.StatusOK, result)
}

// CaesarEncrypt depends on the alphabet.
func CaesarEncrypt(plaintext, alphabet string, shift int) string {
	for shift < 0 {
		shift += len(alphabet)
	}
	candidateRunes := []rune(plaintext)
	var runes []rune
	for _, rune := range candidateRunes {
		if strings.ContainsRune(alphabet, rune) {
			runes = append(runes, rune)
		}
	}
	alphabetRunes := []rune(alphabet)
	firstRuneInAlphabet := alphabetRunes[0]
	for i, rune := range runes {
		runeIndexInAlphabet := rune - firstRuneInAlphabet
		shiftedIndex := (runeIndexInAlphabet + int32(shift)) % int32(len(alphabet))
		runes[i] = alphabetRunes[shiftedIndex]
	}
	return string(runes)
}
