package main

import (
	"fmt"
	"time"
)

func TickWhenFirstStart() {
	ticker := time.NewTicker(time.Second)
	for ; ; <-ticker.C {
		fmt.Println("run_time#", time.Now().Unix())
	}
}

func WrongTicker() {
	for {
		select {
		case <-time.Tick(time.Second):
			fmt.Println("Resource leak!")
		}
	}
}

func main() {
	TickWhenFirstStart()
}
