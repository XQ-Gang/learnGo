package main

import "fmt"

type T2 struct {
	n int
	s string
}

func (T2) M1() {}
func (T2) M2() {}

type NonEmptyInterface interface {
	M1()
	M2()
}

func main() {
	var t = T2{
		n: 17,
		s: "hello, interface",
	}
	var ei interface{}
	ei = t

	var i NonEmptyInterface
	i = t
	fmt.Println(ei)
	fmt.Println(i)
}
