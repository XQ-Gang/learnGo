package main

import "fmt"

func main() {
	n0, n1 := 1, 2
	n0, n1 = n0+n1, n0
	fmt.Println(n0, n1) // 3 1
}
