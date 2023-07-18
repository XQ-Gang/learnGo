package main

func main() {
	m1 := map[int]string{
		1: "hello",
		2: "bye",
	} // map[1:hello 2:bye]

	m1[1] += ", world" // map[1:hello, world 2:bye]

	m2 := map[string][]int{
		"k1": {1, 2},
		"k2": {3, 4},
	} // map[k1:[1 2] k2:[3 4]]
	m2["k1"][0] = 11 // map[k1:[11 2] k2:[3 4]]

	m3 := map[int][10]int{
		1: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}
	m3[1][0] = 11 // 编译错误：cannot assign to m3[1][0] (value of type int)

	type P struct {
		a int
		b float64
	}
	m4 := map[int]P{
		1: {1, 3.14},
		2: {2, 6.28},
	}
	m4[1].a = 11 // 编译错误：cannot assign to struct field m4[1].a in map
}
