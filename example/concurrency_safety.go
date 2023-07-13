package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func concurrencySafetyDemo1() {
	const (
		one = "what the"
		two = "fuck"
	)
	fmt.Println([]byte(one))
	fmt.Println([]byte(two))

	var s string
	go func() {
		for {
			time.Sleep(1)
			ss := s
			fmt.Println(ss)
			if ss != one && ss != two {
				fmt.Printf("concurrent assignment error, ss=%s, []byte(ss)=%v", ss, []byte(ss))
				os.Exit(-1)
			}
		}
	}()
	for i := 1; true; i = 1 - i {
		if i == 0 {
			s = one
		} else {
			s = two
		}
		time.Sleep(1)
	}
}

func concurrencySafetyDemo2() {
	var s string

	for j := 0; j < 100000000; j++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			s = "ab"
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			s = "abc"
		}()
		wg.Wait()

		// 赋值异常判断
		if s != "ab" && s != "abc" {
			fmt.Printf("concurrent assignment error, j=%v s=%v", j, s)
			break
		}
	}
}

func main() {
	concurrencySafetyDemo1()
}
