package main

import (
	"fmt"
)

func getNum() any {
	return 100
}

func main() {
	res, ok := getNum().(int64)
	fmt.Println(res) // 0
	fmt.Println(ok)  // false
}
