package main

import (
	"fmt"
	"time"
)

func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(time.Second)
	select {
	case <-conn:
		timer.Stop()
		return true
	case <-timer.C: // 超时
		println("WaitChannel timeout!")
		return false
	}
}

func AfterDemo() {
	fmt.Println(time.Now())
	<-time.After(time.Second)
	fmt.Println(time.Now())
}

func AfterFuncDemo() {
	fmt.Println(time.Now())
	time.AfterFunc(time.Second, func() {
		fmt.Println(time.Now())
	})
	time.Sleep(time.Second * 2)
}

func main() {
	AfterDemo()
	AfterFuncDemo()
}
