package slice

func normal(s []int) {
	i := 0
	i += s[0]
	i += s[1]
	i += s[2]
	i += s[3]
	println(i)
}

// 如果能确定访问到的  slice 长度，可以先执行一次让编译器去做优化
// [Bounds Checking Elimination](https://docs.google.com/document/d/1vdAEAjYdzjnPA9WDOQ1e4e05cYVMpqSxJYZT33Cqw2g/edit)
func bce(s []int) {
	_ = s[3]
	i := 0
	i += s[0]
	i += s[1]
	i += s[2]
	i += s[3]
	println(i)
}
