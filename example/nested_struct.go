package main

import (
	"fmt"
)

type Funer interface {
	Name() string
	PrintName()
}

type A struct {
}

func (a *A) Name() string {
	return "a"
}

func (a *A) PrintName() {
	fmt.Println(a.Name())
}

type B struct {
	A
}

func (b *B) Name() string {
	return "b"
}

func getBer() Funer {
	return &B{}
}

func main() {
	b := getBer()
	b.PrintName() // Output: a
}
