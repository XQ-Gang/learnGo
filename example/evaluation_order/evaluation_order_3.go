package main

import (
	"fmt"
)

func Expr(n int) int {
	fmt.Println(n)
	return n
}

func main() {
	// 3. switch 语句中的表达式求值
	switch Expr(2) {
	case Expr(1), Expr(2), Expr(3):
		fmt.Println("enter into case1")
		fallthrough
	case Expr(4):
		fmt.Println("enter into case2")
	}
	// 2
	// 1
	// 2
	// enter into case1
	// enter into case2
}
