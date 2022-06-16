package main

import (
	"fmt"
	"reflect"
)

// DumpMethodSet To dump method set of type T, you should pass a pointer to T
// to DumpMethodSet，include interface type.
//
// e.g.
// for interface type I:
//   utils.DumpMethodSet((*I)(nil))
//
// for non-interface type T:
//   var t T
//   utils.DumpMethodSet(&t)
//
// for non-interface type *T:
//   var pt = &T{}
//   utils.DumpMethodSet(&pt)
//
func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elemTyp := v.Elem()

	n := elemTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", elemTyp)
		return
	}

	fmt.Printf("%s's method set:\n", elemTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", elemTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}

type T1 struct{}

func (T1) T1M1()   { println("T1's M1") }
func (T1) T1M2()   { println("T1's M2") }
func (*T1) PT1M3() { println("PT1's M3") }

type T2 struct{}

func (T2) T2M1()   { println("T2's M1") }
func (T2) T2M2()   { println("T2's M2") }
func (*T2) PT2M3() { println("PT2's M3") }

type T struct {
	T1
	*T2
}

func main() {
	t := T{
		T1: T1{},
		T2: &T2{},
	}

	println("call method through t:")
	t.T1M1()
	t.T1M2()
	t.PT1M3()
	t.T2M1()
	t.T2M2()
	t.PT2M3()

	println("\ncall method through pt:")
	pt := &t
	pt.T1M1()
	pt.T1M2()
	pt.PT1M3()
	pt.T2M1()
	pt.T2M2()
	pt.PT2M3()
	println("")

	var t1 T1
	var pt1 *T1
	DumpMethodSet(&t1)
	DumpMethodSet(&pt1)
	t1.PT1M3()

	var t2 T2
	var pt2 *T2
	DumpMethodSet(&t2)
	DumpMethodSet(&pt2)

	DumpMethodSet(&t)
	DumpMethodSet(&pt)
}
