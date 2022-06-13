package main

import "fmt"

func bar() {
	fmt.Println("raise a panic")
	panic(-1)
}

func foo1() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recovered from a panic")
		}
	}()
	bar()
}

func main() {
	foo1()
	fmt.Println("main exit normally")
	// raise a panic
	// recovered from a panic
	// main exit normally
}
