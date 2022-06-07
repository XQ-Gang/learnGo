package main

import (
	"fmt"
)

func main() {
	// 2. 赋值语句的求值
	n0, n1 := 1, 2
	n0, n1 = n0+n1, n0
	fmt.Println(n0, n1)
	// 3 1
}
