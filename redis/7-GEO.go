package main

import (
	"context"
	"fmt"
	"github.com/XQ-Gang/learnGo/utils"
	"github.com/go-redis/redis/v8"
)

func LearnGEO() {
	var rdb = utils.RDB
	var ctx = context.Background()

	key := "China"
	member1 := redis.GeoLocation{Name: "Beijing", Longitude: 116.40, Latitude: 39.90}
	member2 := redis.GeoLocation{Name: "Tianjing", Longitude: 117.20, Latitude: 39.12}
	member3 := redis.GeoLocation{Name: "Taiyuan", Longitude: 112.55, Latitude: 37.87}

	// GEOADD key longitude latitude member [longitude latitude member ...]
	// 存储指定的地理空间位置
	rdb.GeoAdd(ctx, key, &member1)
	rdb.GeoAdd(ctx, key, &member2)
	rdb.GeoAdd(ctx, key, &member3)

	// GEOPOS key member [member ...]
	// 用于从给定的 key 里返回所有指定名称的位置（经度和纬度）
	pos, _ := rdb.GeoPos(ctx, key, member1.Name, member2.Name, member3.Name).Result()
	fmt.Println(member1.Name, pos[0].Longitude, pos[0].Latitude) // Beijing 116.39999896287918 39.900000091670925
	fmt.Println(member2.Name, pos[1].Longitude, pos[1].Latitude) // Tianjing 117.19999998807907 39.12000048819218
	fmt.Println(member3.Name, pos[2].Longitude, pos[2].Latitude) // Taiyuan 112.54999905824661 37.87000021374302

	// GEODIST key member1 member2 [m|km|ft|mi]
	// 用于返回两个给定位置之间的距离
	fmt.Println(rdb.GeoDist(ctx, key, member1.Name, member2.Name, "m"))  // 110631.271
	fmt.Println(rdb.GeoDist(ctx, key, member1.Name, member2.Name, "km")) // 110.6313

	// GEORADIUS key longitude latitude radius m|km|ft|mi [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count] [ASC|DESC] [STORE key] [STOREDIST key]
	// 以给定的经纬度为中心，返回键包含的位置元素当中，与中心的距离不超过给定最大距离的所有位置元素
	// GEORADIUSBYMEMBER key member radius m|km|ft|mi [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count] [ASC|DESC] [STORE key] [STOREDIST key]
	// 找出位于指定范围内的元素，中心点是由给定的位置元素决定的，而不是使用经度和纬度来决定中心点
	members, _ := rdb.GeoRadius(ctx, key, 117, 39.5, &redis.GeoRadiusQuery{
		Radius:      50,
		Unit:        "km",
		WithCoord:   true, // 将位置元素的经度和纬度也一并返回
		WithDist:    true, // 在返回位置元素的同时， 将位置元素与中心之间的距离也一并返回
		WithGeoHash: true, // 以 52 位有符号整数的形式，返回位置元素经过原始 geohash 编码的有序集合分值
		Count:       0,    // 限定返回的记录数
		Sort:        "",   // 从近到远 ASC 从远到近 DESC，默认不排序
		Store:       "",
		StoreDist:   "",
	}).Result()
	fmt.Println(members) // [{Tianjing 117.19999998807907 39.12000048819218 45.6361 4069186099572142}]

	members, _ = rdb.GeoRadiusByMember(ctx, key, member1.Name, &redis.GeoRadiusQuery{
		Radius: 100,
		Unit:   "km",
	}).Result()
	fmt.Println(members) // [{Beijing 0 0 0 0}]

	// GEOSEARCH key FROMMEMBER member | FROMLONLAT longitude latitude BYRADIUS radius M | KM | FT | MI | BYBOX width height M | KM | FT | MI [ ASC | DESC] [ COUNT count [ANY]] [WITHCOORD] [WITHDIST] [WITHHASH]
	names, _ := rdb.GeoSearch(ctx, key, &redis.GeoSearchQuery{
		Member:     member1.Name, // frommember
		Longitude:  0,            // fromlonlat
		Latitude:   0,            // fromlonlat
		Radius:     0,            // byradius
		RadiusUnit: "",           // byradius
		BoxWidth:   5000,         // bybox
		BoxHeight:  1000,         // bybox
		BoxUnit:    "km",
		Sort:       "",
		Count:      0,
		CountAny:   false,
	}).Result()
	fmt.Println(names) // [Taiyuan Tianjing Beijing]

	// GEOHASH key member [member ...]
	// 用于获取一个或多个位置元素的 geohash 值
	hashs, _ := rdb.GeoHash(ctx, key, member1.Name, member2.Name).Result()
	fmt.Println(hashs) // [wx4fbxxfke0 wwgqdcw6tb0]

	// 删除测试数据
	// rdb.Del(ctx, key)
}

func main() {
	utils.WrapFunc(LearnGEO)
}
