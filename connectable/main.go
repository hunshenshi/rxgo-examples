package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"time"
)

func main() {
	ch := make(chan rxgo.Item)
	go func() {
		ch <- rxgo.Of(1)
		ch <- rxgo.Of(2)
		ch <- rxgo.Of(3)
		close(ch)
	}()
	observable := rxgo.FromChannel(ch, rxgo.WithPublishStrategy())

	observable.Map(func(_ context.Context, i interface{}) (interface{}, error) {
		return i.(int) + 1, nil
	}, rxgo.WithCPUPool())
	//.DoOnNext(func(i interface{}) {
	//	fmt.Printf("First observer: %d\n", i)
	//})
	//o2 := observable.Map(func(_ context.Context, i interface{}) (interface{}, error) {
	//	return i.(int) * 2, nil
	//}, rxgo.WithPublishStrategy())
	//.DoOnNext(func(i interface{}) {
	//	fmt.Printf("Second observer: %d\n", i)
	//})

	//nbConsumers := 3
	//wg := sync.WaitGroup{}
	//wg.Add(nbConsumers)
	//for i := 0; i < nbConsumers; i++ {
	//	go func() {
	//		item1 := observable.Observe()
	//		fmt.Println(i)
	//
	//		for item := range item1 {
	//			fmt.Println(item.V)
	//		}
	//		wg.Done()
	//	}()
	//}
	//
	//wg.Wait()

	// TODO 如何展示多个obs
	_, cancel := observable.Connect(context.Background())
	//go func() {
	defer func() {
		// Do something
		time.Sleep(time.Second)
		// Then cancel the subscription
		cancel()
	}()
	// Wait for the subscription to be disposed
	//<-disposed

	//fmt.Println("o1 end")
	//
	//item2 := o2.Observe()
	//for item := range item2 {
	//	fmt.Println(item.V)
	//}

	go func() {
		for item := range observable.Observe() {
			fmt.Println(fmt.Sprintf("item1:%v", item.V))
		}
	}()

	go func() {
		for item := range observable.Observe() {
			fmt.Println(fmt.Sprintf("item2:%v", item.V))
		}
	}()

	time.Sleep(1 * time.Hour)
}
