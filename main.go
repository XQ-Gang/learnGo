package main

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
}
