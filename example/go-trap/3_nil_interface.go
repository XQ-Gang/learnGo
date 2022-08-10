package main

import (
	"io"
)

type TxtReader struct{}

func (*TxtReader) Read(p []byte) (n int, err error) {
	// ... ...
	return 0, nil
}

func main() {
	var r *TxtReader = nil
	var i io.Reader = r
	println(i == nil) // false
	println(i)        // (0x10017cc88,0x0)
}
