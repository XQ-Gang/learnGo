package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1. 在“复制品”上进行迭代
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)
	fmt.Println("a =", a) // a = [1 2 3 4 5]
	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}
		r = append(r, v)
	}
	fmt.Println("r =", r) // r = [1 2 3 4 5]
	fmt.Println("a =", a) // a = [1 2 3 4 5 6 7]

	// 2. 迭代变量是重用的
	var b = []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	for _, v := range b {
		wg.Add(1)
		go func() {
			time.Sleep(time.Second)
			fmt.Println(v) // 5 5 5 5 5
			wg.Done()
		}()
	}
	wg.Wait()
}
