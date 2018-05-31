package ciphers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/beaunus/pretty-sweet-crypto-suite/util"
)

// ScytaleHandler handles requests for a Scytale Cipher
func ScytaleHandler(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()
	plaintext := util.SanitizeParameter("plaintext", &values)
	ciphertext := util.SanitizeParameter("ciphertext", &values)
	numSidesString := util.SanitizeParameter("numSides", &values)
	numSides, error := strconv.Atoi(numSidesString)
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
	json.NewEncoder(w).Encode(&result{
		Ciphertext: resultCiphertext,
		Plaintext:  resultPlaintext,
	})
}

// ScytaleEncrypt is alphabet agnostic.
func ScytaleEncrypt(plaintext string, numSides int) string {
	plaintextRunes := []rune(plaintext)
	resultLength := len(plaintextRunes)
	result := make([]rune, resultLength)
	doTheSwaps(&plaintextRunes, &result, resultLength, numSides, true)
	return string(result)
}

// ScytaleDecrypt is alphabet agnostic.
func ScytaleDecrypt(ciphertext string, numSides int) string {
	ciphertextRunes := []rune(ciphertext)
	resultLength := len(ciphertextRunes)
	result := make([]rune, resultLength)
	doTheSwaps(&ciphertextRunes, &result, resultLength, numSides, false)
	return string(result)
}

func doTheSwaps(source, target *[]rune, resultLength, numSides int, isEncode bool) {
	fromIndex := 0
	i := 0
	for i < resultLength {
		toIndex := i
		for toIndex < resultLength && fromIndex < resultLength {
			if isEncode {
				(*target)[toIndex] = (*source)[fromIndex]
			} else {
				(*target)[fromIndex] = (*source)[toIndex]
			}
			toIndex += numSides
			fromIndex++
		}
		i++
	}
}
