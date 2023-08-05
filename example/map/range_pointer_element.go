package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	{
		m := make(map[int]*int)
		for i, v := range s {
			m[i] = &v
		}
		fmt.Println(*m[0], *m[1], *m[2]) // 3 3 3
	}
	{
		m := make(map[int]*int)
		for i, v := range s {
			v := v
			m[i] = &v
		}
		fmt.Println(*m[0], *m[1], *m[2]) // 1 2 3
	}
	{
		m := make(map[int]*int)
		for i := range s {
			m[i] = &s[i]
		}
		fmt.Println(*m[0], *m[1], *m[2]) // 1 2 3
	}
}
