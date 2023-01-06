package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

func main() {

	// Create the producer
	//ch := make(chan rxgo.Item, 1)
	//go func() {
	//	i := 0
	//	for range time.Tick(time.Second) {
	//		ch <- rxgo.Of(i)
	//		i++
	//	}
	//}()
	//
	//observable := rxgo.FromChannel(ch).
	observable := rxgo.Just(1, 2, 3)().
		//TODO MapReduce中喂给Reduce的数据是排好序的，此处的Reduce如何达到那种效果？
		// Reduce 分组再计算，分组时，如何将相同的key排在一起？ReduceByKey吗？
		Reduce(func(_ context.Context, acc interface{}, elem interface{}) (interface{}, error) {
			if acc == nil {
				return elem, nil
			}
			return acc.(int) + elem.(int), nil
		})

	for item := range observable.Observe() {
		fmt.Println(fmt.Sprintf("item:%v", item.V))
	}
}
