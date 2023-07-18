package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print() // one, two, three
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		v.print() // six, six, six
	}

	time.Sleep(3 * time.Second)
}
