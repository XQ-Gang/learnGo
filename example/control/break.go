package main

import (
	"context"
	"fmt"
)

// 1. 错误的 break，只能跳出 switch
func listing1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)

		switch i {
		default:
		case 2:
			break
		}
	}
}

// 2. 正确的 break，可以跳出 for
func listing2() {
loop:
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)

		switch i {
		default:
		case 2:
			break loop
		}
	}
}

// 3. 错误的 break，只能跳出 select
func listing3(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ch:
			// Do something
		case <-ctx.Done():
			break
		}
	}
}

// 4. 正确的 break，可以跳出 for
func listing4(ctx context.Context, ch <-chan int) {
loop:
	for {
		select {
		case <-ch:
			// Do something
		case <-ctx.Done():
			break loop
		}
	}
}
