package main

import (
	"log"
	"runtime"
)

func main() {
	withDelete()
	withSetNil()
	withPointer()
}

func withDelete() {
	m := make(map[int]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		m[i] = i
	}
	runtime.GC()
	printMemStats() // Alloc = 39355 TotalAlloc = 39444 Sys = 52991 NumGC = 2
	runtime.KeepAlive(m)

	for i := 0; i < 1_000_000; i++ {
		delete(m, i)
	}
	runtime.GC()
	printMemStats() // Alloc = 39358 TotalAlloc = 39447 Sys = 52991 NumGC = 3
	runtime.KeepAlive(m)
}

func withSetNil() {
	m := make(map[int]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		m[i] = i
	}
	runtime.GC()
	printMemStats() // Alloc = 39358 TotalAlloc = 78744 Sys = 95314 NumGC = 5
	runtime.KeepAlive(m)

	m = nil
	runtime.GC()
	printMemStats() // Alloc = 152 TotalAlloc = 78746 Sys = 95314 NumGC = 6
	runtime.KeepAlive(m)
}

func withPointer() {
	m := make(map[int]*int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		m[i] = &i
	}
	runtime.GC()
	printMemStats() // Alloc = 39331 TotalAlloc = 117930 Sys = 96082 NumGC = 8
	runtime.KeepAlive(m)

	for i := 0; i < 1_000_000; i++ {
		delete(m, i)
	}
	runtime.GC()
	printMemStats() // Alloc = 39333 TotalAlloc = 117931 Sys = 96082 NumGC = 9
	runtime.KeepAlive(m)
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// Alloc：当前堆上对象占用的内存大小;
	// TotalAlloc：堆上总共分配出的内存大小;
	// Sys：程序从操作系统总共申请的内存大小;
	// NumGC：垃圾回收运行的次数。
	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}
