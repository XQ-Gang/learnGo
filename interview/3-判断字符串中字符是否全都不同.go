package main

// 判断字符串中字符是否全都不同
// 请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
// 给定一个string，请返回一个bool值，true代表所有字符全都不同，false代表存在相同的字符。
// 保证字符串中的字符为【ASCII字符】。

// 思路：使用位运算来判断是否重复

func isUniqString(s string) bool { // 256 个字符 256 = 64 + 64 + 64 + 64
	var mark1, mark2, mark3, mark4 uint64
	var mark *uint64
	for _, r := range s {
		n := uint64(r)
		if n < 64 {
			mark = &mark1
		} else if n < 128 {
			mark = &mark2
			n -= 64
		} else if n < 192 {
			mark = &mark3
			n -= 128
		} else {
			mark = &mark4
			n -= 192
		}

		if (*mark)&(1<<n) != 0 {
			return false
		}
		*mark = (*mark) | uint64(1<<n)
	}
	return true
}

func main() {
	s := "aba"
	println(isUniqString(s))
}
