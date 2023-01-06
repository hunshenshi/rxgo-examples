package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"strconv"
)

func main() {
	//observable := rxgo.Just(1, errors.New("2"), 3, errors.New("4"), 5)().
	//	OnErrorReturnItem("foo").Filter(func(i interface{}) bool {
	//	switch i.(type) {
	//	case int:
	//		return true
	//	case string:
	//		return i.(string) != "foo"
	//	default:
	//		return true
	//	}
	//})

	//observable := rxgo.Just(1, errors.New("2"), 3, errors.New("4"), 5)().
	//	Map(func(ctx context.Context, i interface{}) (interface{}, error) {
	//		return i, nil
	//	}, rxgo.WithErrorStrategy(rxgo.ContinueOnError))
	//OnErrorReturn(func(err error) interface{} {
	//	return err.Error()
	//})

	observable := rxgo.Just(1, 2, 3, errors.New("4"), 5)().
		Map(func(ctx context.Context, i interface{}) (interface{}, error) {
			switch i.(type) {
			case int:
				return i, errors.New(strconv.Itoa(i.(int)))
			case string:
				return i, errors.New("i.(string)")
			case error:
				return i, errors.New("error")
			default:
				return i, nil
			}
		}).
		Filter(func(i interface{}) bool {
			fmt.Println(i)
			return true
		}, rxgo.WithErrorStrategy(rxgo.ContinueOnError))

	//.GroupByDynamic(func(item rxgo.Item) string {
	//	return strconv.Itoa(item.V.(int))
	//})

	for item := range observable.Observe() {
		fmt.Println(fmt.Sprintf("item: %v", item.V))
	}
}
