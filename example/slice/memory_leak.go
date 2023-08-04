package main

import (
	"log"
	"runtime"
)

func main() {
	withSlice()
	withFull()
	withCopy()
}

func withSlice() {
	s := make([]int, 1_000_000)
	s1 := s[:5]
	runtime.GC()
	printMemStats() // Alloc = 7965 TotalAlloc = 7968 Sys = 20031 NumGC = 2
	runtime.KeepAlive(s1)
}

func withFull() {
	s := make([]int, 1_000_000)
	s2 := s[:5:5]
	runtime.GC()
	printMemStats() // Alloc = 7968 TotalAlloc = 15789 Sys = 28479 NumGC = 4
	runtime.KeepAlive(s2)
}

func withCopy() {
	s := make([]int, 1_000_000)
	s3 := make([]int, 5)
	copy(s3, s)
	runtime.GC()
	printMemStats() // Alloc = 154 TotalAlloc = 23609 Sys = 36671 NumGC = 6
	runtime.KeepAlive(s3)
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
