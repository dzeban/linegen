package main

import (
	"testing"
)

func TestRandString(t *testing.T) {
	l := 10
	b := RandString(l)
	if len(b) != l {
		t.Errorf("Length doesn't match. Expected %d, got %d", l+1, len(b))
	}

	if b[l-1] != '\n' {
		t.Error("Generated string does not end with newline")
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
