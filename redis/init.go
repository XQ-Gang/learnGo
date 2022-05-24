package redis

import (
	"context"
	"github.com/XQ-Gang/learnGo/utils"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func LearnRedis() {
	utils.WrapFunc(LearnString)
	utils.WrapFunc(LearnHash)
	utils.WrapFunc(LearnList)
	utils.WrapFunc(LearnSet)
	utils.WrapFunc(LearnZSet)
	utils.WrapFunc(LearnHyperLogLog)
	utils.WrapFunc(LearnGEO)
	utils.WrapFunc(LearnPubSub)
}
