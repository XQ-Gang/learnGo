package main

import "fmt"

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

	numbers := [256]int{'a': 8, 'b': 7, 'c': 4, 'd': 3, 'e': 2, 'y': 1, 'x': 5}
	fnumbers := [...]float64{-1, 4: 2, 3, 7: 4, 9: 5}
	fmt.Println(numbers)
	fmt.Println('a' == 97)
	fmt.Println(fnumbers)

	type Point struct {
		x float64
		y float64
	}
	m := map[string]*Point{
		"Persepolis": {29.9, 52.9},
		"Uluru":      {-25.4, 131.0},
		"Googleplex": {37.4, -122.1},
	}
}
