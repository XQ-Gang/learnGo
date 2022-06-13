package main

import "fmt"

var x = []int{1}

func foo4() {
	old := x[:]
	defer func() { x = old }()
	x = append(x, 2)
	fmt.Println("foo:", x)
}

func main() {
	foo4()
	fmt.Println("main:", x)
	// foo: [1 2]
	// main: [1]
}
