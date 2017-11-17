# linegen
Fast random and fixed line generator

## Examples

By default `linegen` generates random strings of length 10

```
$ linegen | head -n3
GYtWNErOe
I4bEWNO/+
1g/rx3RsR
```

You can set length with `-len` argument

```
$ linegen -len 5 | head -n3
qUJ7
ZGQL
TkJY
```
Note, that last character is always newline.

To print out fixed string (like `yes` command) pass it as positional argument

```
$ linegen awesomeness | head -n3
awesomeness
awesomeness
awesomeness
```

## Performance

`linegen` was written to make generation and output fast. 

Random number generation uses all bits of random data to select a char from base64 string. 
This will avoid extra calls to PRNG that is expensive

Output is buffered in 64K buffer before printing on stdout.

On Fedora 25, with i5 CPU, Go 1.9 I have a following numbers

Random string generation

```
$ linegen -len 4096 | pv > /dev/null
... [ 372MiB/s] ...
```

Fixed string printing
```
$ linegen awesomeness | pv > /dev/null
... [1.13GiB/s] ...
```

go bench for random string generation

```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/dzeban/linegen
BenchmarkRandString100-4    	 5000000	       294 ns/op
BenchmarkRandString4096-4   	  200000	      9804 ns/op
```

## Usage

```
Usage: linegen [opts] [str]

linegen will generate infinite stream of random strings of given length or fixed strings passed in "str" argument

Options:
  -len int
    	Length of random string to generate (default 10)
```
