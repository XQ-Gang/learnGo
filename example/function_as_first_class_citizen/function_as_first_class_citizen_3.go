// inspired by https://github.com/go-functional/core
package main

import (
	"fmt"
)

type IntSliceFunctor interface {
	Fmap(fn func(int) int) IntSliceFunctor
}

type intSliceFunctorImpl struct {
	ints []int
}

func (isf intSliceFunctorImpl) Fmap(fn func(int) int) IntSliceFunctor {
	newInts := make([]int, len(isf.ints))
	for i, elt := range isf.ints {
		retInt := fn(elt)
		newInts[i] = retInt
	}
	return intSliceFunctorImpl{ints: newInts}
}

func NewIntSliceFunctor(slice []int) IntSliceFunctor {
	return intSliceFunctorImpl{ints: slice}
}

func main() {
	// 原切片
	intSlice := []int{1, 2, 3, 4}
	fmt.Printf("init a functor from int slice: %#v\n", intSlice) // []int{1, 2, 3, 4}
	f := NewIntSliceFunctor(intSlice)
	fmt.Printf("original functor: %+v\n", f) // {ints:[1 2 3 4]}

	mapperFunc1 := func(i int) int {
		return i + 10
	}

	mapped1 := f.Fmap(mapperFunc1)
	fmt.Printf("mapped functor1: %+v\n", mapped1) // {ints:[11 12 13 14]}

	mapperFunc2 := func(i int) int {
		return i * 3
	}
	mapped2 := mapped1.Fmap(mapperFunc2)
	fmt.Printf("mapped functor2: %+v\n", mapped2)                                 // {ints:[33 36 39 42]}
	fmt.Printf("original functor: %+v\n", f)                                      // 原functor没有改变
	fmt.Printf("composite functor: %+v\n", f.Fmap(mapperFunc1).Fmap(mapperFunc2)) // {ints:[1 2 3 4]}
}
