package main

import (
	"math/rand"
	"time"
)

// Create new random source that is unsynchronized, thus much faster
var randSource = rand.NewSource(time.Now().UnixNano())

const (
	base64chars  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/"
	bitSliceSize = 6 // 6 bits of random gives random integer in [0..63] range to choose from base64chars
	bitSlices    = len(base64chars) / bitSliceSize
	bitMask      = (1 << bitSliceSize) - 1
)

func RandString(n int) []byte {
	b := make([]byte, n)

	r := randSource.Int63()
	for i := 0; i < len(b); i++ {
		for j := 0; j < bitSlices; j++ {
			idx := r & bitMask
			b[i] = base64chars[idx]

			r >>= bitSliceSize
			i++
			if i >= len(b) {
				break
			}
		}
	}
	b[n-1] = '\n'
	return b
}
