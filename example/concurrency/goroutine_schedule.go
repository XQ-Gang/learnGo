package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n) // 9 0 1 2 3 4 5 6 7 8
			wg.Done()
		}(i)
	}
	wg.Wait()
}
