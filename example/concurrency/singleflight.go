package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/singleflight"
)

var sg singleflight.Group

func main() {
	for i := 0; i < 10; i++ { // 模拟10个并发
		go func() {
			fmt.Println(GetName("username"))
		}()
	}
	// 等待协程执行结束
	time.Sleep(time.Second)
}

// GetName 获取名称
func GetName(cacheKey string) string {
	result, _, _ := sg.Do(cacheKey, func() (ret interface{}, err error) {
		log.Printf("getting %s from database\n", cacheKey)
		log.Printf("setting %s in cache\n", cacheKey)
		return "xiuqiugang", nil
	})
	return result.(string)
}
