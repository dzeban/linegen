package main

import (
	"flag"
	"fmt"
	"os"
)

var Len int
var Str []byte

func usage() {
	fmt.Printf("Usage: %s [opts] [str]\n", os.Args[0])
	fmt.Println()
	fmt.Printf("%s will generate infinite stream of random strings ", os.Args[0])
	fmt.Printf("of given length or fixed strings passed in \"str\" argument\n")
	fmt.Println()
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func init() {
	flag.Usage = usage
	flag.IntVar(&Len, "len", 10, "Length of random string to generate")
	flag.Parse()
	if flag.NArg() == 1 {
		Str = []byte(flag.Arg(0) + "\n")
	}
}

func main() {
	var f func(int) []byte
	if len(Str) > 0 {
		f = func(n int) []byte {
			return Str
		}
	} else {
		f = RandString
	}
	printBuffered(Len, f)
}

func printBuffered(n int, genFunc func(int) []byte) {
	const bufsize = 65536

	buf := make([]byte, bufsize)
	i := 0
	for {
		copy(buf[i:i+n], genFunc(n))
		i += n
		if i >= len(buf)-n {
			os.Stdout.Write(buf)
			i = 0
		}
	}
}
