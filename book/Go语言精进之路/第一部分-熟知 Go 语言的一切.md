# 熟知 Go 语言的一切

## 1 了解 Go 语言的诞生与演进

诞生历史：C++ 复杂性高，编译构建速度慢，编写服务端程序时不便支持并发。

正式发布：Go 语言项目在 **2009 年 11 月 10 日**正式开源，这一天被 Go 官方确定为 Go 语言诞生日。

2009 年 TIOBE 年度最佳编程语言。

相关名词：

- 吉祥物：mascot。
- Go 程序员：Gopher。
- Go 语言官方技术大会：GopherCon。
- 国内最负盛名的 Go 技术大会：GopherChina。

## 2 选择适当的 Go 语言版本

重要版本发布历史：

- 2012 年 3 月 28 日：Go1.0 正式发布。同时发布了 “[Go1 兼容性](https://tip.golang.org/doc/go1compat)” 承诺。
- 2015 年 8 月 19 日：Go1.5 版本发布。Go 实现了**自举**。
- 2018 年 8 月 25 日：Go1.11 版本发布。引入新的包管理机制 **Go Module**。

Go 团队已经将版本发布节奏稳定在每年发布两次大版本上。

Go 团队承诺对最新的两个稳定大版本提供支持。

Go 开发团队一直**建议大家使用最新的发布版**。

## 3 理解 Go 语言的设计哲学

设计哲学：

- **追求简单，少即是多**。拒绝走语言特性融合的道路。推崇 “最小方式” 思维。
- **偏好组合，正交解耦**。Go 采用了组合的方式，也是唯一的方式。
  - 垂直组合：**类型嵌入**，将已经实现的功能嵌入新类型中，以快速满足新类型的功能需求。通过类型嵌入，快速让一个新类型复用其他类型已经实现的能力，实现功能的垂直扩展。
  - 水平组合：通过 **interface** 将程序各个部分组合在一起。接口只是方法集合，且与实现者之间的关系是隐式的，它让程序各个部分之间的耦合降至最低，同时是连接程序各个部分的“纽带”。隐式的 interface 实现会不经意间自然而然满足依赖抽象、里式替换、接口隔离等设计原则。
- **原生并发，轻量高效**。
  - 采用轻量级协程 goroutine，使得 Go 应用在面向多核硬件时更具可扩展性。goroutine 占用的资源非常少，Go 运行时默认为每个 goroutine 分配的**栈空间**仅 **2KB**。goroutine 调度的切换也不用陷入操作系统内核层完成，代价很低。
  - 提供了支持并发的语法元素和机制。通过内置的 **channel** 传递消息或实现同步，并通过 **select** 实现多路 channel 的并发控制。
  - 并发原则对 Go 开发者在程序设计层面的影响。“并发不是并行”。
- **面向工程，“自带电池”**。Go设计者将所有工程问题浓缩为一个词：**scale**。**gofmt** 统一了 Go 语言的编码风格。

## 4 使用 Go 语言原生编程思维来写 Go 代码

> 不能影响到你的编程思维方式的编程语言不值得学习和使用。——首届图灵奖得主 艾伦·佩利

引例：如何找到小于或等于给定整数 n 的素数？埃拉托斯特尼素数筛算法。

```Go
package main

// Generate Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Filter Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
```

我们学习和使用一门编程语言，目标是用这门语言的原生思维方式编写高质量代码。掌握 Go 原生编程思维就是我们通往高质量 Go 编程的学习方向和必经之路。

## 参考

《Go 语言精进之路：从新手到高手的编程思想、方法和技巧》——白明