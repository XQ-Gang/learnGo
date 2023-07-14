package main

import (
	"fmt"
	"unsafe"
)

// StringToBytes supports since Go1.20+.
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// BytesToString supports since Go1.20+.
func BytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

func main() {
	s := "ab"
	fmt.Printf("%p\n", &s)
	b := StringToBytes(s)
	fmt.Printf("%p\n", &b)
	s2 := BytesToString(b)
	fmt.Printf("%p\n", &s2)
	fmt.Println(b)
	fmt.Println(string(b))
	fmt.Println(s2)
	fmt.Println([]byte(s2))

	// Output:
	// 0xc000014270
	// 0xc000010030
	// 0xc000014280
	// [97 98]
	// ab
	// ab
	// [97 98]
}
