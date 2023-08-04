package main

import (
	"fmt"
	"sync"
	"unsafe"
)

// 必须用key来初始化结构体
type NoUnkeyedLiterals struct{}

// 不允许结构体比较
type DoNotCompare [0]func()

// 不允许结构体拷贝、值传递
type DoNotCopy [0]sync.Mutex

type User struct {
	// 必须用key来初始化结构体
	NoUnkeyedLiterals
	// 不允许结构体比较
	DoNotCompare
	// 不允许结构体拷贝、值传递
	DoNotCopy
	Age     int
	Address string
}

func main() {
	_ = &User{Age: 21, Address: "beijing"}

	fmt.Println(unsafe.Sizeof(NoUnkeyedLiterals{}))
	fmt.Printf("%p\n", &DoNotCopy{})
	fmt.Printf("%p\n", &NoUnkeyedLiterals{})
	fmt.Printf("%p\n", &DoNotCompare{})
}
