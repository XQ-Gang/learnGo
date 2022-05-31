package main

import "fmt"

func main() {
	u := [...]int{0, 1, 2, 3, 4, 5, 6}
	var s = u[4:5]
	fmt.Println(len(s), cap(s), u) // 1 3 [0 1 2 3 4 5 6]
	s = append(s, 7)
	fmt.Println(len(s), cap(s), u) // 2 3 [0 1 2 3 4 7 6]
	s = append(s, 8)
	fmt.Println(len(s), cap(s), u) // 3 3 [0 1 2 3 4 7 8]

	// 此时切片 s 与 数组 u 共享底层数组，同步改动
	s[0] = 20
	fmt.Println(len(s), cap(s), u) // 3 3 [0 1 2 3 20 7 8]

	// 底层数组剩余空间不满足添加新元素，创建了新的底层数组（长度为原 2 倍）
	s = append(s, 9)
	fmt.Println(len(s), cap(s), u) // 4 6 [0 1 2 3 20 7 8]

	// 此时再修改切片不会影响原来的数组 u
	s[0] = 21
	fmt.Println(len(s), cap(s), u) // 4 6 [0 1 2 3 20 7 8]

	s = append(s, 10)
	fmt.Println(len(s), cap(s), u) // 5 6 [0 1 2 3 20 7 8]
	s = append(s, 11)
	fmt.Println(len(s), cap(s), u) // 6 6 [0 1 2 3 20 7 8]

	// 又创建了新的底层数组
	s = append(s, 12)
	fmt.Println(len(s), cap(s), u) // 7 12 [0 1 2 3 20 7 8]

	// 查看切片
	fmt.Println(s) // [21 7 8 9 10 11 12]
}
