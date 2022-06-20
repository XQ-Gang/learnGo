package main

import "fmt"

type T int

func (t T) Error() string {
	return "bad error"
}

func printNilInterface() {
	// nil接口变量
	var i interface{}                 // 空接口类型
	var err error                     // 非空接口类型
	println(i)                        // (0x0,0x0)
	println(err)                      // (0x0,0x0)
	println("i = nil:", i == nil)     // i = nil: true
	println("err = nil:", err == nil) // err = nil: true
	println("i = err:", i == err)     // i = err: true
	println("")
}

func printEmptyInterface() {
	// empty接口变量
	var eif1 interface{} // 空接口类型
	var eif2 interface{} // 空接口类型
	var n, m int = 17, 18

	eif1 = n
	eif2 = m

	println("eif1:", eif1)                // eif1: (0x10726a0,0xc00006ef68)
	println("eif2:", eif2)                // eif2: (0x10726a0,0xc00006ef60)
	println("eif1 = eif2:", eif1 == eif2) // eif1 = eif2: false

	eif2 = 17
	println("eif1:", eif1)                // eif1: (0x10726a0,0xc00006ef68)
	println("eif2:", eif2)                // eif2: (0x10726a0,0x1099828)
	println("eif1 = eif2:", eif1 == eif2) // eif1 = eif2: true

	eif2 = int64(17)
	println("eif1:", eif1)                // eif1: (0x10726a0,0xc00006ef68)
	println("eif2:", eif2)                // eif2: (0x1072760,0x1099828)
	println("eif1 = eif2:", eif1 == eif2) // eif1 = eif2: false

	println("")
}

func printNonEmptyInterface() {
	var err1 error // 非空接口类型
	var err2 error // 非空接口类型
	err1 = (*T)(nil)
	println("err1:", err1)              // err1: (0x10c0708,0x0)
	println("err1 = nil:", err1 == nil) // err1 = nil: false

	err1 = T(5)
	err2 = T(6)
	println("err1:", err1)                // err1: (0x10c0768,0x10c0210)
	println("err2:", err2)                // err2: (0x10c0768,0x10c0218)
	println("err1 = err2:", err1 == err2) // err1 = err2: false

	err2 = fmt.Errorf("%d\n", 5)
	println("err1:", err1)                // err1: (0x10c0768,0x10c0210)
	println("err2:", err2)                // err2: (0x10c0688,0xc000010250)
	println("err1 = err2:", err1 == err2) // err1 = err2: false

	println("")
}

func printEmptyInterfaceAndNonEmptyInterface() {
	var eif interface{} = T(5)
	var err error = T(5)
	println("eif:", eif)              // eif: (0x1007ff7c0,0x1007f4f78)
	println("err:", err)              // err: (0x10080b3e8,0x1007f4f78)
	println("eif = err:", eif == err) // eif = err: true

	err = T(6)
	println("eif:", eif)              // eif: (0x1007ff7c0,0x1007f4f78)
	println("err:", err)              // err: (0x10080b3e8,0x1007f4f80)
	println("eif = err:", eif == err) // eif = err: false
}

func main() {
	printNilInterface()
	printEmptyInterface()
	printNonEmptyInterface()
	printEmptyInterfaceAndNonEmptyInterface()
}
