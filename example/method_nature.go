package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p field) print1() {
	fmt.Println(p.name)
}

func (p *field) print2() {
	fmt.Println(p.name)
}

func main() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print1() // one two three
		// 等价转换
		go field.print1(*v) // one two three
	}
	for _, v := range data1 {
		go v.print2() // one two three
		// 等价转换
		go (*field).print1(v) // one two three
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go v.print1() // four five six
		// 等价转换
		go field.print1(v) // four five six
	}
	for _, v := range data2 {
		go v.print2() // six six six
		// 等价转换
		go (*field).print1(&v) // six six six
	}

	time.Sleep(1 * time.Second)
}
