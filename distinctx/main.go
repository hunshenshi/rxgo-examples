package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

func main() {
	//TODO Distinct有个KeySet用于存储去重的key
	observable := rxgo.Just(1, 2, 2, 3, 1, 4, 4, 5, 6, 4)().
		Distinct(func(_ context.Context, i interface{}) (interface{}, error) {
			return i, nil
		})

	// 比较当前key与上一个key是否相同
	//observable := rxgo.Just(1, 2, 2, 1, 1, 3)().
	//	DistinctUntilChanged(func(_ context.Context, i interface{}) (interface{}, error) {
	//		return i, nil
	//	})

	//ch := make(chan rxgo.Item)
	//rand.Seed(time.Now().UnixNano())
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		for i := 0; i < 10; i++ {
	//			n := rand.Intn(6)
	//			fmt.Println(n)
	//			ch <- rxgo.Of(n)
	//		}
	//		time.Sleep(10 * time.Second)
	//	}
	//	close(ch)
	//}()
	//
	//observable := rxgo.FromChannel(ch).Distinct(func(ctx context.Context, i interface{}) (interface{}, error) {
	//	return i, nil
	//})

	// 从第一个元素算起，索引为0
	//observable := rxgo.FromChannel(ch).ElementAt(12)

	//observable := rxgo.Just(1, 2, 3)().First()
	for item := range observable.Observe() {
		fmt.Println(fmt.Sprintf("item:%d", item.V))
	}
}
