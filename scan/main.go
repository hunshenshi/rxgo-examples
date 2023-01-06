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
	//		ch <- rxgo.Of(strconv.Itoa(i))
	//		i++
	//	}
	//}()

	//observable := rxgo.FromChannel(ch).
	observable := rxgo.Just("a", "b", "c", "d")().
		Scan(func(ctx context.Context, i interface{}, i2 interface{}) (interface{}, error) {
			//if i == nil {
			//	return i2, nil
			//}
			return i2, nil
		})

	//observable := rxgo.Just(1, 2, 3, 4, 5)().
	//	Scan(func(_ context.Context, acc interface{}, elem interface{}) (interface{}, error) {
	//		if acc == nil {
	//			return elem, nil
	//		}
	//		return acc.(int) + elem.(int), nil
	//	})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
