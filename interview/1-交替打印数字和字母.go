package main

import (
	"sync"
)

// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

func main() {
	number, letter := make(chan bool), make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		num := 1
		for {
			ok := <-number
			print(num)
			num++
			print(num)
			num++
			if ok {
				break
			}
			letter <- false
		}
	}()
	go func() {
		number <- false
		let := 'A'
		for {
			<-letter
			print(string(let))
			let++
			print(string(let))
			if let == 'Z' {
				number <- true
				break
			} else {
				let++
				number <- false
			}
		}
	}()
	wg.Wait()
}
