package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"strings"
)

func main() {
	observable := rxgo.Just("Hello, World!", "World")().Filter(func(i interface{}) bool {
		res := strings.Contains(i.(string), "World")
		return res
	})
	ch := observable.Observe()
	for i := range ch {
		if i.Error() {
			fmt.Println("error")
			fmt.Println(i.E)
		}
		fmt.Println(i.V)
	}

}
