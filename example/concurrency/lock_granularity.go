package main

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type SafeRander struct {
	pos     uint32
	randers [128]*rand.Rand
	locks   [128]*sync.Mutex
}

// NewSafeRander 创建独立的多个随机数生成器，通过原子操作获取对应的随机数生成器的下标，
// 确保两次相邻的调用总是引用不同的随机数生成器，锁的粒度就被大大降低了，
// 其锁的竞争次数也降低了。本质上也是一种空间换取时间的策略。
func NewSafeRander() *SafeRander {
	var randers [128]*rand.Rand
	var locks [128]*sync.Mutex
	for i := 0; i < 128; i++ {
		randers[i] = rand.New(rand.NewSource(time.Now().UnixNano()))
		locks[i] = new(sync.Mutex)
	}
	return &SafeRander{
		randers: randers,
		locks:   locks,
	}
}

func (sr *SafeRander) Intn(n int) int {
	// 原子操作，确保两次调用获取的是不同的随机数生成器
	x := atomic.AddUint32(&sr.pos, 1)
	x %= 128
	sr.locks[x].Lock()
	n = sr.randers[x].Intn(n)
	sr.locks[x].Unlock()
	return n
}
