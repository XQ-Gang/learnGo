package main

import (
	"strings"
)

// 判断两个给定的字符串排序后是否一致
// 给定两个字符串，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
// 给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。

func isRegroup(s1, s2 string) bool {
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}

func main() {
	s1 := "abcdefg"
	s2 := "gfedcba"
	println(isRegroup(s1, s2))
}
