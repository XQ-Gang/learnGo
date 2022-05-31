package main

import "sync"

const (
	a, b = 1, 11
	c, d
	e, f
)

const g = iota
const h = iota
const i = 4
const j = iota

const (
	_     = iota
	Blue  // 1
	Black // 2
	Red   // 3
	_
	Yellow = iota * 2 // 10
	Green             // 12
)

func main() {
	// learn go-redis/redis/v8
	// redis.LearnRedis()

	// learn standard library
	//standard.LearnStandard()

	const k = iota
	println(a, b, c, d, e, f)
	println(g, h, i, j, k)

	// 零值可用的切片不能通过下标形式操作数据
	var s []int
	s[0] = 12         // 报错！
	s = append(s, 12) // 正确

	// map 没有提供零值可用支持
	var m map[string]int
	m["go"] = 1 // 报错！

	m1 := make(map[string]int)
	m1["go"] = 1 // 正确

	// 尽量避免值复制
	var mu sync.Mutex
	mu1 := mu
	foo(mu)

	// 可以通过指针方式传递类似 Mutex 这样的类型
	var mu sync.Mutex
	foo(&mu)

}
