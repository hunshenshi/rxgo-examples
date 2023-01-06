package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"time"
)

func main() {
	// Create the producer
	ch := make(chan rxgo.Item, 1)
	go func() {
		i := 0
		for range time.Tick(time.Second) {
			ch <- rxgo.Of(i)
			i++
		}
	}()

	//TODO buffer vs window
	observable := rxgo.FromChannel(ch).
		//BufferWithTime(rxgo.WithDuration(3 * time.Second))
		//BufferWithTimeOrCount(rxgo.WithDuration(3*time.Second), 2)
		BufferWithCount(3).Filter(func(i interface{}) bool {
		fmt.Println(i)
		for i2 := range i.([]interface{}) {
			fmt.Println(i2)
		}
		return true
	})
	//observable := rxgo.Just(1, 2, 3, 4)().BufferWithCount(3)

	for item := range observable.Observe() {
		fmt.Println(item.V)
		//for _, v := range item.V.([]interface{}) {
		//	fmt.Println(v.(int))
		//}
	}
}
