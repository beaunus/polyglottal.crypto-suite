package ciphers

import "testing"

func TestScytaleEncrypt(t *testing.T) {
	cases := []struct {
		plaintext string
		numSides  int
		want      string
	}{
		{"IamhurtverybadlyHELP", 4, "IryyatbHmvaEhedLurlP"},
	}
	for _, c := range cases {
		got := ScytaleEncrypt(c.plaintext, c.numSides)
		if got != c.want {
			t.Errorf("ScytaleEncrypt(%q, %q) == %q, want %q", c.plaintext, c.numSides, got, c.want)
		}
	}
}
