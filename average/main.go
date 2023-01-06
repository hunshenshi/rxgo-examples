package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"math/rand"
	"time"
)

func main() {
	//observable := rxgo.Just(1, 2, 3, 4)().AverageInt()

	ch1 := make(chan rxgo.Item)
	rand.Seed(time.Now().UnixNano())
	go func() {
		for i := 0; i < 5; i++ {
			for i := 0; i < 10; i++ {
				n := rand.Intn(6)
				fmt.Println(fmt.Sprintf("src1 n : %v", n))
				ch1 <- rxgo.Of(n)
			}
			time.Sleep(10 * time.Second)
		}
		close(ch1)
	}()
	//TODO stream的avg是如何计算的，结合窗口算子使用
	//observable := rxgo.FromChannel(ch1).BufferWithTimeOrCount(rxgo.WithDuration(3*time.Second), 2).AverageInt()
	observable := rxgo.FromChannel(ch1).WindowWithCount(4)

	for item := range observable.Observe() {
		for i := range item.V.(rxgo.Observable).AverageInt().Observe() {
			//for i := range item.V.(rxgo.Observable).SumInt64().Observe() {
			fmt.Println(fmt.Sprintf("item:%v", i.V))
		}
	}
}
