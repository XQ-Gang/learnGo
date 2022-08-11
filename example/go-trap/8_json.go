package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	var nilSlice []int
	var emptySlice = make([]int, 0, 5)

	println(nilSlice == nil)   // true
	println(emptySlice == nil) // false

	println(nilSlice, len(nilSlice), cap(nilSlice))       // [0/0]0x0 0 0
	println(emptySlice, len(emptySlice), cap(emptySlice)) // [0/5]0x140000a4030 0 5

	// 1. nil 切片和空切片可能被编码为不同文本
	m1 := map[string][]int{
		"nilSlice":   nilSlice,
		"emptySlice": emptySlice,
	}
	b, _ := json.Marshal(m1)
	println(string(b)) // {"emptySlice":[],"nilSlice":null}

	// 2. 字节切片可能被编码为 base64 编码的文本
	m2 := map[string]any{
		"byteSlice": []byte("hello, go"),
		"string":    "hello, go",
	}
	b, _ = json.Marshal(m2)
	fmt.Println(string(b)) // {"byteSlice":"aGVsbG8sIGdv","string":"hello, go"}

	// 3. 当 JSON 文本中的整型数值被解码为 interface{} 类型时，其底层真实类型为 float64
	s := `{"age": 1}`
	m3 := map[string]any{}
	_ = json.Unmarshal([]byte(s), &m3)
	println(m3["age"].(int)) // panic: interface conversion: interface {} is float64, not int

	m4 := map[string]any{}
	d := json.NewDecoder(strings.NewReader(s))
	d.UseNumber()
	_ = d.Decode(&m4)
	age, _ := m4["age"].(json.Number).Int64()
	println(age) // 1
}
