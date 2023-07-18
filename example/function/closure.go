package main

import (
	"fmt"
)

func main() {
	var fs = [2]func(){}

	for i := 0; i < 2; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer_closure i = ", i) }()
		fs[i] = func() { fmt.Println("closure i = ", i) }
	}

	for _, f := range fs {
		f()
	}
	// closure i = 2
	// closure i = 2
	// defer_closure i = 2
	// defer i = 1
	// defer_closure i = 2
	// defer i = 0
}
