package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"strings"
)

func main() {
	observable := rxgo.Just("Hello, World!", "World", "Golang World")().FlatMap(func(item rxgo.Item) rxgo.Observable {
		obs := rxgo.Just(strings.Split(item.V.(string), " "))()
		return obs
	})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
