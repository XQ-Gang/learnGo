package main

import "runtime"

var stat runtime.MemStats

func main() {
	runtime.ReadMemStats(&stat)
	println(stat.HeapSys)
}
