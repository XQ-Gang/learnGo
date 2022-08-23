package main

// 请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
// 给定一个string，请返回一个string，为翻转后的字符串。

func main() {
	s := "我爱你，中国！"
	str := []rune(s)
	l := len(str)
	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}
	println(string(str))
}
