package main

import (
	"fmt"
	"sync"
	"time"
)

type Container struct {
	sync.Mutex // <-- Added a mutex
	counters   map[string]int
}

func (c Container) inc(name string) {
	c.Lock() // <-- Added locking of the mutex
	defer c.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{counters: map[string]int{"a": 0, "b": 0}}

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
	}

	go doIncrement("a", 100000)
	go doIncrement("a", 100000)
	// panic: fatal error: concurrent map writes

	// Wait a bit for the goroutines to finish
	time.Sleep(300 * time.Millisecond)
	fmt.Println(c.counters)
}
