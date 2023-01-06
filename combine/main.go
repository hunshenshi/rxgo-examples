package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"math/rand"
	"time"
)

func main() {

	ch1 := make(chan rxgo.Item)
	rand.Seed(time.Now().UnixNano())
	go func() {
		for i := 0; i < 5; i++ {
			for i := 0; i < 10; i++ {
				n := rand.Intn(6)
				ch1 <- rxgo.Of(n)
			}
			time.Sleep(10 * time.Second)
		}
		close(ch1)
	}()
	obs1 := rxgo.FromChannel(ch1)

	ch2 := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 5; i++ {
			for i := 0; i < 10; i++ {
				n := rand.Intn(6)
				ch2 <- rxgo.Of(n)
			}
			time.Sleep(10 * time.Second)
		}
		close(ch2)
	}()
	obs2 := rxgo.FromChannel(ch2)

	observable := rxgo.CombineLatest(func(i ...interface{}) interface{} {
		sum := 0
		for _, v := range i {
			if v == nil {
				continue
			}
			fmt.Println(v.(int))
			sum += v.(int)
		}
		return sum
	}, []rxgo.Observable{
		obs1,
		obs2,
		//rxgo.Just(1, 2)(),
		//rxgo.Just(10, 11)(),
	})

	for item := range observable.Observe() {
		fmt.Println(fmt.Sprintf("item:%d", item.V))
	}
}
