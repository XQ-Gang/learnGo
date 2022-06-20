package main

import "fmt"

func main() {
	var n int = 61
	var ei interface{} = n
	n = 62
	fmt.Println("data in box:", ei) // 61

	var m int = 51
	ei = &m
	m = 52
	p := ei.(*int)
	fmt.Println("data in box:", *p) // 52
}
