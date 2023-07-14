package main

import (
	"fmt"
)

func defer1() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func defer2() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Print(i)
		}()
	}
}

func defer3() {
	a := 1
	defer fmt.Println(a)
	a++
}

func defer4() {
	a := 1
	defer func() {
		fmt.Println(a)
	}()
	a++
}

func main() {
	defer1()
	fmt.Println()
	defer2()
	fmt.Println()
	defer3()
	defer4()
}
