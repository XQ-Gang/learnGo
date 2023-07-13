package main

import (
	"fmt"
	"time"
)

// 注意：for + select closed channel 可能会造成死循环
// select 中 break 无法跳出 for 循环
func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()
	for {
		select {
		case i, ok := <-ch:
			fmt.Println(i, ok)
		default:
			fmt.Println("default")
			break
		}
		time.Sleep(1 * time.Second)
	}
}
