package ciphers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/beaunus/pretty-sweet-crypto-suite/util"
)

// CaesarHandler handles requests for a Caesar Cipher
func CaesarHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	plaintext := util.SanitizeParameter("plaintext", &values)
	alphabet := util.SanitizeParameter("alphabet", &values)
	shiftString := util.SanitizeParameter("shift", &values)
	shift, error := strconv.Atoi(shiftString)
	if error != nil {
		fmt.Println("Caesar", error)
	}
	ciphertext := CaesarEncrypt(plaintext, alphabet, shift)
	result := struct {
		Ciphertext string
	}{
		Ciphertext: ciphertext,
	}
	json.NewEncoder(w).Encode(&result)
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
