package main

import (
	"fmt"
	"time"
)

func allocSlice(min, high int) []int {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("slice a: len(%d), cap(%d), elements(%v)\n", len(a), cap(a), a)
	// slice a: len(10), cap(10), elements([1 2 3 4 5 6 7 8 9 10])

	go func() {
		time.Sleep(time.Second)
		fmt.Printf("slice a: len(%d), cap(%d), elements(%v)\n", len(a), cap(a), a)
		// slice a: len(10), cap(10), elements([1 2 3 4 5 6 7 18 19 10])
	}()

	return a[min:high]
}

func main() {
	// 1. 隐匿数据的暴露与切片数据篡改
	a1 := allocSlice(3, 7)
	fmt.Printf("slice a1: len(%d), cap(%d), elements(%v)\n", len(a1), cap(a1), a1)
	// slice a1: len(4), cap(7), elements([4 5 6 7])
	a2 := a1[:6]
	fmt.Printf("slice a2: len(%d), cap(%d), elements(%v)\n", len(a2), cap(a2), a2)
	// slice a2: len(6), cap(7), elements([4 5 6 7 8 9])
	a2[4] += 10
	a2[5] += 10
	time.Sleep(2 * time.Second)

	// 2. 新切片与原切片底层存储可能会“分家”
	var b = []int{1, 2, 3, 4}
	fmt.Printf("\nslice b: len(%d), cap(%d), elements(%v)\n", len(b), cap(b), b)
	// slice b: len(4), cap(4), elements([1 2 3 4])
	b1 := b[:2]
	fmt.Printf("slice b1: len(%d), cap(%d), elements(%v)\n", len(b1), cap(b1), b1)
	// slice b1: len(2), cap(4), elements([1 2])

	fmt.Println("\nappend 11 to b1:")
	b1 = append(b1, 11)
	fmt.Printf("slice b1: len(%d), cap(%d), elements(%v)\n", len(b1), cap(b1), b1)
	// slice b1: len(3), cap(4), elements([1 2 11])
	fmt.Printf("slice b: len(%d), cap(%d), elements(%v)\n", len(b), cap(b), b)
	// slice b: len(4), cap(4), elements([1 2 11 4])

	fmt.Println("\nappend 22 to b1:")
	b1 = append(b1, 22)
	fmt.Printf("slice b1: len(%d), cap(%d), elements(%v)\n", len(b1), cap(b1), b1)
	// slice b1: len(4), cap(4), elements([1 2 11 22])
	fmt.Printf("slice b: len(%d), cap(%d), elements(%v)\n", len(b), cap(b), b)
	// slice b: len(4), cap(4), elements([1 2 11 22])

	fmt.Println("\nappend 33 to b1:")
	b1 = append(b1, 33)
	fmt.Printf("slice b1: len(%d), cap(%d), elements(%v)\n", len(b1), cap(b1), b1)
	// slice b1: len(5), cap(8), elements([1 2 11 22 33])
	fmt.Printf("slice b: len(%d), cap(%d), elements(%v)\n", len(b), cap(b), b)
	// slice b: len(4), cap(4), elements([1 2 11 22])

	b1[0] *= 100
	fmt.Println("\nb1[0] multiply 100:")
	fmt.Printf("slice b1: len(%d), cap(%d), elements(%v)\n", len(b1), cap(b1), b1)
	// slice b1: len(5), cap(8), elements([100 2 11 22 33])
	fmt.Printf("slice b: len(%d), cap(%d), elements(%v)\n", len(b), cap(b), b)
	// slice b: len(4), cap(4), elements([1 2 11 22])
}
