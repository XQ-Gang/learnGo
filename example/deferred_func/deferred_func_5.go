package main

import "fmt"

func foo51() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func foo52() {
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func foo53() {
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func main() {
	fmt.Println("foo1 result:")
	foo51() // 3 2 1 0
	fmt.Println("\nfoo2 result:")
	foo52() // 3 2 1 0
	fmt.Println("\nfoo3 result:")
	foo53() // 4 4 4 4
}
