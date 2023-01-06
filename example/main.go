package main

import (
	"context"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	events := rxgo.Create([]rxgo.Producer{func(_ context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(expensiveReadFromDisk(0))
		next <- rxgo.Of(expensiveReadFromDisk(1))
		next <- rxgo.Of(expensiveReadFromDisk(2))
	}}, rxgo.WithPublishStrategy())

	total := events.Count()
	filtered := events.Filter(func(i interface{}) bool {
		return i.(int) > 0
	}).Count()

	events.Connect(context.Background())

	t, _ := total.Get()
	fmt.Printf("   Total: %d\n", t.V)

	ch := filtered.Observe()
	fmt.Println((<-ch).V)

	f, _ := filtered.Get()
	fmt.Printf("Filtered: %d\n", f.V)
}

func expensiveReadFromDisk(e int) int {
	fmt.Printf("Reading event: %d\n", e)
	return e
}
