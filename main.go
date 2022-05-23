package main

import (
	"fmt"
	"github.com/XQ-Gang/learnGo/redis"
	"reflect"
	"runtime"
	"strings"
)

func getFuncName(i interface{}) string {
	names := strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), ".")
	return names[len(names)-1]
}

func wrapFunc(a func()) {
	fmt.Println(strings.Repeat("*", 10) + getFuncName(a) + strings.Repeat("*", 10))
	a()
	fmt.Println(strings.Repeat("*", 10) + getFuncName(a) + strings.Repeat("*", 10))
}

func main() {
	// wrapFunc(redis.LearnString)
	// wrapFunc(redis.LearnHash)
	wrapFunc(redis.LearnList)
}
