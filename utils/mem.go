package utils

import (
	"log"
	"runtime"
)

func PrintMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// Alloc：当前堆上对象占用的内存大小;
	// TotalAlloc：堆上总共分配出的内存大小;
	// Sys：程序从操作系统总共申请的内存大小;
	// NumGC：垃圾回收运行的次数。
	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}
