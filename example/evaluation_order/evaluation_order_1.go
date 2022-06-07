package main

import (
	"fmt"
)

var a, b, c = f() + v(), g(), sqr(u()) + v()

func f() int {
	fmt.Println("calling f")
	return c
}

func g() int {

	fmt.Println("calling g")
	return 1
}

func sqr(x int) int {
	fmt.Println("calling sqr")
	return x * x
}

func v() int {
	fmt.Println("calling v")
	return 1
}

func u() int {
	fmt.Println("calling u")
	return 2
}

func main() {
	// 1. 普通求值顺序 + 包级变量求值顺序
	fmt.Println(a, b, c)
	// calling g
	// calling u
	// calling sqr
	// calling v
	// calling f
	// calling v
	// 6 1 5
}
