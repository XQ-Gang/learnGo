package main

var a = 0

func main() {
	println(a, &a) // 0 0x100645e40
	a := 1
	println(a, &a) // 1 0x14000098f60
	a, b := 2, 3
	_ = b
	println(a, &a) // 2 0x14000098f60
	if a == 2 {
		a, c := 4, 5
		_ = c
		println(a, &a) // 4 0x14000098f58
	}
	println(a, &a) // 2 0x14000098f60
	a, d := 6, 7
	_ = d
	println(a, &a) // 6 0x14000098f60
}
