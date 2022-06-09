package main

import (
	"fmt"
	"time"
)

func demo1() {
	var m = [...]int{1, 2, 3, 4, 5}

	for i, v := range m {
		go func() {
			time.Sleep(time.Second * 1)
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 2)
}

func demo2() {
	var m = [...]int{1, 2, 3, 4, 5}

	for i, v := range m {
		go func(i, v int) {
			time.Sleep(time.Second * 1)
			fmt.Println(i, v)
		}(i, v)
	}

	time.Sleep(time.Second * 2)
}

func main() {
	demo1()
	// 4 5
	// 4 5
	// 4 5
	// 4 5
	// 4 5

	demo2() // 输出结果由 goroutine 调度决定
	// 3 4
	// 4 5
	// 0 1
	// 2 3
	// 1 2
}
