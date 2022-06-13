package main

import "fmt"

func foo61() {
	sl := []int{1, 2, 3}
	defer func(a []int) {
		fmt.Println(a)
	}(sl)

	sl = []int{3, 2, 1}
	_ = sl
}

func foo62() {
	sl := []int{1, 2, 3}
	defer func(p *[]int) {
		fmt.Println(*p)
	}(&sl)

	sl = []int{3, 2, 1}
	_ = sl
}

func main() {
	foo61() // [1 2 3]
	foo62() // [3 2 1]
}
