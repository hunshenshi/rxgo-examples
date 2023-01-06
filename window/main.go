package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

func main() {
	//TODO window vs buffer
	observe := rxgo.Just(1, 2, 3, 4, 5)().WindowWithCount(2).Observe()

	//fmt.Println("First Observable")
	//for item := range (<-observe).V.(rxgo.Observable).Observe() {
	//	if item.Error() {
	//		//return item.E
	//		fmt.Println(item.E)
	//	}
	//	fmt.Println(item.V)
	//}
	//
	//fmt.Println("Second Observable")
	//for item := range (<-observe).V.(rxgo.Observable).Observe() {
	//	if item.Error() {
	//		//return item.E
	//		fmt.Println(item.E)
	//	}
	//	fmt.Println(item.V)
	//}

	for item := range observe {
		fmt.Println("new obs...")
		//TODO 这种逻辑，代码模版怎么编写
		//for i := range item.V.(rxgo.Observable).FlatMap(func(item rxgo.Item) rxgo.Observable {
		//	fmt.Println(fmt.Sprintf("flatMap:%v", item.V))
		//	return rxgo.Just(item.V)()
		//}).Observe() {
		//for i := range item.V.(rxgo.Observable).Observe() {
		for i := range item.V.(rxgo.Observable).Reduce(func(ctx context.Context, i interface{}, i2 interface{}) (interface{}, error) {
			if i == nil {
				return i2, nil
			}
			return i.(int) + i2.(int), nil
		}).Observe() {
			//for i := range item.V.(rxgo.Observable).TakeLast(1).Observe() {
			fmt.Println(i.V)
		}
	}
}
