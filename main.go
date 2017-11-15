package main

import (
	"os"
)

const bufsize = 65536

func printBuffered(n int) {
	buf := make([]byte, bufsize)
	i := 0
	for {
		copy(buf[i:i+n], RandString(n))
		i += n
		if i >= len(buf) {
			os.Stdout.Write(buf)
			i = 0
		}
	}
}

func main() {
	n := 4096
	printBuffered(n)
}
