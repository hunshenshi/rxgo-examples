package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	ch := make(chan rxgo.Item)
	// Data producer
	go producer(ch)

	// Create an Observable
	observable := rxgo.FromChannel(ch).Filter(func(i interface{}) bool {
		res := i.(Customer).Age > 20
		return res
	}).Map(func(_ context.Context, i interface{}) (interface{}, error) {
		customer := i.(Customer)
		if customer.Age > 30 {
			customer.TaxNumber = "19832106687"
		}
		return customer, nil
	}, rxgo.WithContext(context.Background()), rxgo.WithCPUPool(), rxgo.WithPool(5)).Filter(func(i interface{}) bool {
		res := len(i.(Customer).TaxNumber) > 0
		return res
	}).
		//Marshal(json.Marshal)

		GroupByDynamic(func(item rxgo.Item) string {
			return strconv.Itoa(item.V.(Customer).Age)
		}, rxgo.WithBufferedChannel(100)) //.SumInt64()

	// consume the items using ForEach() or Observe()
	c := observable.Observe()
	for item := range c {
		fmt.Println(item.V)
		// for json
		//fmt.Println(string(item.V.([]byte)))

		switch item.V.(type) {
		case rxgo.GroupedObservable: // group operator
			go func() {
				obs := item.V.(rxgo.GroupedObservable)
				fmt.Printf("New observable: %s\n", obs.Key)
				for i := range obs.Observe() {
					//for i := range obs.Count().Observe() {
					fmt.Printf("item: %v\n", i.V)
				}
			}()
		case rxgo.ObservableImpl: // window operator
			obs := item.V.(rxgo.ObservableImpl)
			for i := range obs.Observe() {
				//for i := range obs.Count().Observe() {
				fmt.Printf("item: %v\n", i.V)
			}
		default:
			fmt.Printf("item: %v\n", item.V)
		}
	}

	// ops list
	//(observable, opts...)
}

func producer(ch chan<- rxgo.Item) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		customer := Customer{
			ID:       i,
			Name:     "yi",
			LastName: "zhang",
			Age:      rand.Intn(80),
		}
		ch <- rxgo.Of(customer)
	}
	close(ch)
}

type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	TaxNumber string `json:"taxNumber"`
}
