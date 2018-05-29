package ciphers

import "testing"

func TestCaesarEncrypt(t *testing.T) {
	cases := []struct {
		in    string
		shift int
		want  string
	}{
		{"abcdefghijklmnopqrstuvwxyz", -1, "`abcdefghijklmnopqrstuvwxy"},
	}
	for _, c := range cases {
		got := CaesarEncrypt(c.in, c.shift)
		if got != c.want {
			t.Errorf("CaesarEncrypt(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
