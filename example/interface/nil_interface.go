package main

import (
	"fmt"
)

func main() {
	var p map[string]string
	var i interface{} = p
	fmt.Println(i == nil) // false
}
