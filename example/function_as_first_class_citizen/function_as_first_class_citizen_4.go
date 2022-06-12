package main

import "fmt"

// 求阶乘函数 - 递归方法
func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

// 求阶乘函数 - CPS风格
func factorialCPS(n int, f func(int)) {
	if n == 1 {
		f(1) //基本情况
	} else {
		factorialCPS(n-1, func(y int) { f(n * y) })
	}
}

func main() {
	fmt.Printf("%d\n", factorial(5))                       // 120
	factorialCPS(5, func(y int) { fmt.Printf("%d\n", y) }) // 120
}
