package main

import (
	"testing"
)

func byteInSlice(b byte, s []byte) bool {
	if b == '\n' {
		return true
	}

	for _, x := range s {
		if x == b {
			return true
		}
	}
	return false
}

func TestRandString(t *testing.T) {
	l := 100
	b := RandString(l)
	if len(b) != l {
		t.Errorf("Length doesn't match. Expected %d, got %d", l+1, len(b))
	}

	if b[l-1] != '\n' {
		t.Error("Generated string does not end with newline")
	}

	for i := 0; i < len(b); i++ {
		if !byteInSlice(b[i], []byte(base64chars)) {
			t.Errorf("Invalid character %X in generated string", b[i])
		}
	}
}

func BenchmarkRandString100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(100)
	}
}

func BenchmarkRandString4096(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(4096)
	}
}
