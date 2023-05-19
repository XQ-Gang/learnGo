package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	for ; ; <-ticker.C {
		fmt.Println("run_time#", time.Now().Unix())
	}
}
