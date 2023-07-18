package main

import (
	"fmt"
)

func modifySlice1(s []int) {
	s[0] = 1024
	s = append(s, 2048)
	s = append(s, 4096)
}

func modifySlice2(s []int) {
	s = append(s, 2048)
	s = append(s, 4096)
	s[0] = 1024
}

func main() {
	s1 := []int{0, 1, 2}
	modifySlice1(s1)
	fmt.Println(s1) // [1024 1 2]

	s2 := []int{0, 1, 2}
	modifySlice2(s2)
	fmt.Println(s2) // [0 1 2]
}
