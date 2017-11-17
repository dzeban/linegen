package main

import (
	"flag"
	"os"
)

var Len int

func init() {
	flag.IntVar(&Len, "len", 10, "Length of string to generate")
	flag.Parse()

}

func main() {
	printBuffered(Len)
}

func printBuffered(n int) {
	const bufsize = 65536

	buf := make([]byte, bufsize)
	i := 0
	for {
		copy(buf[i:i+n], RandString(n))
		i += n
		if i >= len(buf)-n {
			os.Stdout.Write(buf)
			i = 0
		}
	}
}
