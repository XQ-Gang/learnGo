package main

import (
	"runtime"
	"sync"
)

func main() {
	var l = sync.RWMutex{}
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)
	// 协程A
	go func() {
		// 第一次获取读锁
		l.RLock()
		// 第一个defer
		defer l.RUnlock()
		c <- 1
		// 让协程B执行
		runtime.Gosched()
		// 第二次获取读锁
		l.RLock()
		// 第二个defer
		defer l.RUnlock()
		wg.Done()
	}()
	// 协程B
	go func() {
		<-c
		l.Lock()
		defer l.Unlock()
		wg.Done()
	}()

	wg.Wait()
}
