package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"time"
)

func main() {
	//count := 3
	// group 配合 buffer 使用
	//TODO buffer的大小有什么影响
	//observable := rxgo.Range(0, 10).GroupBy(count, func(item rxgo.Item) int {
	//	return item.V.(int) % count
	//}, rxgo.WithBufferedChannel(9))
	//
	//n := 0
	//for i := range observable.Observe() {
	//	//groupedObservable := i.V.(rxgo.GroupedObservable) //  *rxgo.ObservableImpl, not rxgo.GroupedObservable
	//	fmt.Println("New observable: " + strconv.Itoa(n))
	//
	//	for i := range i.V.(rxgo.Observable).Observe() {
	//		fmt.Printf("item: %v\n", i.V)
	//	}
	//	n++
	//}

	//obs := observable.Observe()
	//
	//fmt.Println("first obs...")
	//for item := range (<-obs).V.(rxgo.Observable).Observe() {
	//	fmt.Println(item.V)
	//}
	//
	//fmt.Println("second obs...")
	//for item := range (<-obs).V.(rxgo.Observable).Observe() {
	//	fmt.Println(item.V)
	//}
	//
	//fmt.Println("third obs...")
	//for item := range (<-obs).V.(rxgo.Observable).Observe() {
	//	fmt.Println(item.V)
	//}

	observable := rxgo.Range(0, 24).GroupByDynamic(func(item rxgo.Item) string {
		//time.Sleep(1 * time.Second)
		//return strconv.Itoa(item.V.(int) % 2) // 分组的key，其值存在一个map中
		if item.V.(int) <= 5 {
			return "1"
		} else if item.V.(int) <= 11 {
			return "2"
		} else if item.V.(int) <= 17 {
			return "1"
		} else if item.V.(int) <= 23 {
			return "2"
		} else {
			return "3"
		}
	}, rxgo.WithBufferedChannel(2)) //TODO buffer设置多少合适？？？  buffer是用来缓存不同分组下的元素的
	// 如果数据

	for i := range observable.Observe() {
		//fmt.Println("&")
		go func() {
			groupedObservable := i.V.(rxgo.GroupedObservable)
			key := groupedObservable.Key
			fmt.Printf("New observable: %s\n", key)

			for i := range groupedObservable.AverageInt().Observe() {
				fmt.Printf("key: %s, item: %v\n", key, i.V)
			}
		}()

	}

	time.Sleep(1 * time.Hour)
	//for item := range observable.Observe() {
	//	fmt.Println(item.V)
	//}
}
