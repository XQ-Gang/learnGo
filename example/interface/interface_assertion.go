package main

import (
	"fmt"
)

func getInteger() any {
	return 100
}

func getFloat() any {
	return 1.1
}

func main() {
	res1, ok := getInteger().(int64)
	fmt.Println(res1) // 0
	fmt.Println(ok)   // false

	res2, ok := getInteger().(int)
	fmt.Println(res2) // 100
	fmt.Println(ok)   // true

	res3, ok := getFloat().(float32)
	fmt.Println(res3) // 0
	fmt.Println(ok)   // false

	res4, ok := getFloat().(float64)
	fmt.Println(res4) // 1.1
	fmt.Println(ok)   // true
}
