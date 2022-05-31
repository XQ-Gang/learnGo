package main

import (
	"context"
	"fmt"
	"github.com/XQ-Gang/learnGo/utils"
)

func LearnPubSub() {
	var rdb = utils.RDB
	var ctx = context.Background()

	pchan := "lang*"
	chan1, chan2, chan3 := "lang1", "lang2", "status"

	// SUBSCRIBE channel [channel ...]
	// 订阅给定的一个或多个频道的信息
	// PSUBSCRIBE pattern [pattern ...]
	// 订阅一个或多个符合给定模式的频道
	sub := rdb.PSubscribe(ctx, pchan)
	err := sub.Subscribe(ctx, chan3)
	if err != nil {
		panic(err)
	}

	// PUBLISH channel message
	// 将信息发送到指定的频道
	rdb.Publish(ctx, chan1, "Python")
	rdb.Publish(ctx, chan1, "Golang")
	rdb.Publish(ctx, chan2, "Java")
	rdb.Publish(ctx, chan2, "C")
	rdb.Publish(ctx, chan3, "quit")

	// PUBSUB subcommand [argument [argument ...]]
	// 查看订阅与发布系统状态
	// 1 - PUBSUB NUMPAT
	// 返回订阅模式的数量
	fmt.Println(rdb.PubSubNumPat(ctx)) // 1
	// 2 - PUBSUB NUMSUB channel [channel ...]
	// 返回给定频道的订阅者数量，订阅模式的客户端不计算在内
	fmt.Println(rdb.PubSubNumSub(ctx, chan1, chan2, chan3)) // map[lang1:0 lang2:0 status:1]
	// 3 - PUBSUB CHANNELS [pattern]
	// 列出当前的活跃频道，订阅模式的客户端不计算在内
	fmt.Println(rdb.PubSubChannels(ctx, "*")) // [status]

	// UNSUBSCRIBE [channel [channel ...]]
	// 指退订给定的频道
	// PUNSUBSCRIBE [pattern [pattern ...]]
	// 退订所有给定模式的频道
	err = sub.PUnsubscribe(ctx, pchan)
	if err != nil {
		panic(err)
	}
	err = sub.Unsubscribe(ctx, chan3)
	if err != nil {
		panic(err)
	}

	// 打印收到的消息
	for msg := range sub.Channel() {
		fmt.Println(msg.Channel, msg.Payload)
		if msg.Channel == chan3 && msg.Payload == "quit" {
			break
		}
	}
}

func main() {
	utils.WrapFunc(LearnPubSub)
}
