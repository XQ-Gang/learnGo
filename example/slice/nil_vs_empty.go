package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

func main() {
	// 创建 nil 切片
	var s1 []int
	bs1, _ := json.Marshal(s1)
	fmt.Println(s1, string(bs1), unsafe.SliceData(s1)) // 输出：[] null <nil>
	var s2 = []int(nil)
	bs2, _ := json.Marshal(s2)
	fmt.Println(s2, string(bs2), unsafe.SliceData(s2)) // 输出：[] null <nil>
	var s3 = *new([]int)
	bs3, _ := json.Marshal(s3)
	fmt.Println(s3, string(bs3), unsafe.SliceData(s3)) // 输出：[] null <nil>

	// 创建空切片
	s4 := make([]int, 0)
	bs4, _ := json.Marshal(s4)
	fmt.Println(s4, string(bs4), unsafe.SliceData(s4)) // 输出：[] [] 0x11a40c0
	s5 := []int{}
	bs5, _ := json.Marshal(s5)
	fmt.Println(s5, string(bs5), unsafe.SliceData(s5)) // 输出：[] [] 0x11a40c0

	// 创建零切片
	s6 := make([]int, 2, 5)
	bs6, _ := json.Marshal(s6)
	fmt.Println(s6, string(bs6), unsafe.SliceData(s6)) // 输出：[0 0] [0,0] 0xc00001c0f0

	// full s expression: s[:0:0]
	s7 := s6[:0:0]
	bs7, _ := json.Marshal(s7)
	fmt.Println(s7, string(bs7), unsafe.SliceData(s7)) // 输出：[] [] 0xc00001c0f0
	// 由上可知，虽然底层是零长数组，但还是指向原切片的底层数组
}
