package main

import (
	"fmt"
	"unsafe"
)

// StringToBytes supports since Go1.20+.
func StringToBytes(s string) []byte {
	if s == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// BytesToString supports since Go1.20+.
func BytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func main() {
	s := "ab"
	fmt.Printf("%p\n", &s) // 0xc000014270
	b := StringToBytes(s)
	fmt.Printf("%p\n", &b) // 0xc000010030
	s2 := BytesToString(b)
	fmt.Printf("%p\n", &s2) // 0xc000014280
	fmt.Println(b)          // [97 98]
	fmt.Println(string(b))  // ab
	fmt.Println(s2)         // ab
	fmt.Println([]byte(s2)) // [97 98]
}
