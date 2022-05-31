package main

import (
	"context"
	"fmt"
	"github.com/XQ-Gang/learnGo/utils"
	"time"
)

// Context context.Context 接口
type Context interface {
	// Deadline 返回 context.Context 被取消的时间，也就是完成工作的截止日期
	Deadline() (deadline time.Time, ok bool)
	// Done 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭
	Done() <-chan struct{}
	// Err 返回 context.Context 结束的原因，它只会在 Done 方法对应的 Channel 关闭时返回非空的值
	// 1. 如果 context.Context 被取消，会返回 Canceled 错误；
	// 2. 如果 context.Context 超时，会返回 DeadlineExceeded 错误；
	Err() error
	// Value 从 context.Context 中获取键对应的值，对于同一个上下文来说，
	// 多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据
	Value(key interface{}) interface{}
}

func LearnContext() {
	root := context.Background()
	ctx1 := context.WithValue(root, "k1", "v1")
	ctx2, cancel1 := context.WithCancel(ctx1)
	defer cancel1()
	ctx3, cancel2 := context.WithTimeout(ctx2, 2*time.Second)
	defer cancel2()
	ctx4 := context.WithValue(ctx2, "k2", "v2")
	ctx5 := context.WithValue(ctx4, "k3", "v3")
	ctx6 := context.WithValue(ctx3, "k4", "v4")
	ctx7, cancel3 := context.WithDeadline(ctx4, time.Now().Add(4*time.Second))
	defer cancel3()
	ctx8, cancel4 := context.WithCancel(ctx6)
	defer cancel4()

	fmt.Println(ctx7.Value("k1")) // v1
	fmt.Println(ctx6.Value("k4")) // v4
	fmt.Println(ctx5.Value("k2")) // v2

	fmt.Println(ctx7) // context.Background.WithValue(type string, val v1).WithCancel.WithValue(type string, val v2).WithDeadline(2022-05-24 23:55:36.196879 +0800 CST m=+4.000265802 [3.999736898s])

	go func() {
		select {
		case <-ctx7.Done():
			fmt.Println("ctx7 canceled")
		}
	}()

	go func() {
		select {
		// ctx8 canceled before ctx7, because ctx3 canceled
		case <-ctx8.Done():
			fmt.Println("ctx8 canceled")
		}
	}()

	time.Sleep(5 * time.Second)
}

func main() {
	utils.WrapFunc(LearnContext)
}
