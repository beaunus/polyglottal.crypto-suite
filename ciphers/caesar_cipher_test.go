package ciphers

import "testing"

func TestCaesarEncrypt(t *testing.T) {
	cases := []struct {
		plaintext string
		shift     int
		want      string
	}{
		{"abcdefghijklmnopqrstuvwxyz", -1, "`abcdefghijklmnopqrstuvwxy"},
	}
	for _, c := range cases {
		got := CaesarEncrypt(c.plaintext, c.shift)
		if got != c.want {
			t.Errorf("CaesarEncrypt(%q, %q) == %q, want %q", c.plaintext, c.shift, got, c.want)
		}
	}
}
