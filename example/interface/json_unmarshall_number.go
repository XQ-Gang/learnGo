package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var data = `{"status": 200}`
	// 如果反序列化不指定结构体类型或者变量类型(any)，则JSON中的数字类型默认被反序列化成float64类型
	var res0 map[string]any
	json.Unmarshal([]byte(data), &res0)
	status0, ok := res0["status"].(int)
	fmt.Println(status0)                        // 0
	fmt.Println(ok)                             // false
	fmt.Println(reflect.TypeOf(res0["status"])) // float64

	// 解决方案1：强制类型转化
	var res1 map[string]any
	json.Unmarshal([]byte(data), &res1)
	status1 := int(res1["status"].(float64))
	fmt.Println(int(status1)) // 200

	// 解决方案2：自定义struct
	var res2 struct{ Status int }
	json.Unmarshal([]byte(data), &res2)
	fmt.Println(res2.Status) // 200

	// 解决方案3：UseNumber()
	var res3 map[string]any
	decoder := json.NewDecoder(strings.NewReader(data))
	decoder.UseNumber()
	decoder.Decode(&res3)
	status3, err := res3["status"].(json.Number).Int64()
	fmt.Println(status3)                        // 200
	fmt.Println(err)                            // <nil>
	fmt.Println(reflect.TypeOf(res3["status"])) // json.Number
}
