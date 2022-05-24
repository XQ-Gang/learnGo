package utils

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func getFuncName(i interface{}) string {
	names := strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), ".")
	return names[len(names)-1]
}

func center(str string, width int, fillchar string) string {
	length := len(str)
	if width-length < 0 {
		return str
	}
	left, right := (width-length)/2, (width-length+1)/2
	res := strings.Repeat(fillchar, left) + str + strings.Repeat(fillchar, right)
	return res
}

func WrapFunc(f func()) {
	funcName := getFuncName(f)
	fmt.Println(center(funcName, 50, "*"))
	f()
	fmt.Println(center(funcName, 50, "*"))
}
