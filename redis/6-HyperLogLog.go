package redis

import "fmt"

func LearnHyperLogLog() {
	// HyperLogLog 指令都是 PF 开头，这是因为其发明人是 Philippe Flajolet
	key1, key2 := "pf-lang1", "pf-lang2"
	value1, value2, value3, value4 := "Python", "Golang", "Java", "C"

	// PFADD key element [element ...]
	// 添加指定元素到 HyperLogLog 中
	rdb.PFAdd(ctx, key1, value1, value2, value3)
	rdb.PFAdd(ctx, key2, value2, value3, value4)

	// PFCOUNT key [key ...]
	// 返回给定 HyperLogLog 的基数估算值
	fmt.Println(rdb.PFCount(ctx, key1)) // 3
	fmt.Println(rdb.PFCount(ctx, key2)) // 3

	// PFMERGE destkey sourcekey [sourcekey ...]
	// 将多个 HyperLogLog 合并为一个 HyperLogLog
	rdb.PFMerge(ctx, key1, key2)
	fmt.Println(rdb.PFCount(ctx, key1)) // 4
}
