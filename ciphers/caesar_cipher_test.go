package ciphers

import "testing"

func TestCaesarEncrypt(t *testing.T) {
	cases := []struct {
		plaintext string
		alphabet  string
		shift     int
		want      string
	}{
		{"abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz", -1, "zabcdefghijklmnopqrstuvwxy"},
	}
	for _, c := range cases {
		got := CaesarEncrypt(c.plaintext, c.alphabet, c.shift)
		if got != c.want {
			t.Errorf("CaesarEncrypt(%q, %q) == %q, want %q", c.plaintext, c.shift, got, c.want)
		}
	}
}
