package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

func main() {
	capacity := 10
	//TODO
	observable := rxgo.Range(0, 100, rxgo.WithBufferedChannel(capacity)).
		Map(func(_ context.Context, i interface{}) (interface{}, error) {
			return i, nil
		}, rxgo.WithCPUPool(), rxgo.WithBufferedChannel(capacity)).
		Serialize(0, func(i interface{}) int {
			return i.(int)
		})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
