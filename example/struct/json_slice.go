package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Age     int      `json:"age,omitempty"`
	Name    string   `json:"name,omitempty"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	jsonStr1 := `{"age": 1,"name": "xiaoming", "hobbies":["羽毛球","乒乓球","篮球"]}`
	a := Student{}
	json.Unmarshal([]byte(jsonStr1), &a)
	aa := a.Hobbies // 发生值拷贝
	fmt.Println(aa) // [羽毛球 乒乓球 篮球]
	ptr1 := fmt.Sprintf("%p", &aa)
	ptr2 := fmt.Sprintf("%p", &aa[0])
	ptr3 := fmt.Sprintf("%p", aa)
	fmt.Println(ptr1 == ptr2) // false
	fmt.Println(ptr1 == ptr3) // false
	fmt.Println(ptr2 == ptr3) // true

	ptr4 := fmt.Sprintf("%p", &a.Hobbies)
	ptr5 := fmt.Sprintf("%p", &a.Hobbies[0])
	ptr6 := fmt.Sprintf("%p", a.Hobbies)
	fmt.Println(ptr4, ptr5, ptr6)
	fmt.Println(ptr1 == ptr4) // false
	fmt.Println(ptr2 == ptr5) // true
	fmt.Println(ptr3 == ptr6) // true

	jsonStr2 := `{"age": 1,"name": "xiaoming", "hobbies":["唱歌","跳舞","乒乓球","篮球","羽毛球"]}`
	json.Unmarshal([]byte(jsonStr2), &a)
	fmt.Println(aa)        // [羽毛球 乒乓球 篮球]
	fmt.Println(a.Hobbies) // [唱歌 跳舞 乒乓球 篮球 羽毛球]
}
