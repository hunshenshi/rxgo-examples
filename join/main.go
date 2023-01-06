package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"time"
)

func main() {

	//ch1 := make(chan rxgo.Item)
	//rand.Seed(time.Now().UnixNano())
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		for i := 0; i < 10; i++ {
	//			n := map[string]int64{"time": rand.Int63n(6), "value": rand.Int63n(6)}
	//			fmt.Println(fmt.Sprintf("src1 n : %v", n))
	//			ch1 <- rxgo.Of(n)
	//		}
	//		time.Sleep(10 * time.Second)
	//	}
	//	close(ch1)
	//}()
	//obs1 := rxgo.FromChannel(ch1)
	//
	//ch2 := make(chan rxgo.Item)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		for i := 0; i < 10; i++ {
	//			n := map[string]int64{"time": rand.Int63n(6), "value": rand.Int63n(6)}
	//			fmt.Println(fmt.Sprintf("src2 n : %v", n))
	//			ch2 <- rxgo.Of(n)
	//		}
	//		time.Sleep(10 * time.Second)
	//	}
	//	close(ch2)
	//}()
	//obs2 := rxgo.FromChannel(ch2)
	//
	//observable := obs1.Join(func(ctx context.Context, l interface{}, r interface{}) (interface{}, error) {
	//	lMap := l.(map[string]int64)
	//	rMap := r.(map[string]int64)
	//	if lMap["time"] == rMap["time"] {
	//		fmt.Println(fmt.Sprintf("%v join %v", lMap, rMap))
	//		lMap["value"] = lMap["value"] + rMap["value"]
	//	}
	//	return lMap, nil
	//}, obs2, func(i interface{}) time.Time {
	//	return time.Unix(0, i.(map[string]int64)["time"]*1000000)
	//}, rxgo.WithDuration(200*time.Millisecond))

	observable := rxgo.Just(
		map[string]int64{"tt": 1, "V": 1},
		map[string]int64{"tt": 4, "V": 2},
		map[string]int64{"tt": 7, "V": 3},
	)().Join(func(ctx context.Context, l interface{}, r interface{}) (interface{}, error) {
		return map[string]interface{}{
			"l": l,
			"r": r,
		}, nil
	}, rxgo.Just(
		map[string]int64{"tt": 2, "V": 5},
		map[string]int64{"tt": 3, "V": 6},
		map[string]int64{"tt": 5, "V": 7},
	)(), func(i interface{}) time.Time { // 时间提取函数，当lTime与rTime的差值小于duration，则在窗口内进行join
		return time.Unix(0, i.(map[string]int64)["tt"]*1000000) // ltt - rtt <= 2 的数据进行join，笛卡尔积
	}, rxgo.WithDuration(2*time.Millisecond))

	for item := range observable.Observe() {
		fmt.Println(fmt.Sprintf("item:%v", item.V))
	}
}
