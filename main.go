package main

import (
	"fmt"
	"github.com/matinkhosravani/weighted-rand/weightedRand"
	"os"
)

func main() {
	//weights or scores are in float64 format
	items := []interface{}{"one", "two", "four", "eight", "ten"}
	weights := []float64{1.1, 2.2, 4.3, 8.4, 10.5}
	wR := weightedRand.NewWRand(items, weights)
	fmt.Println(wR.GetN(5))

	//weights or scores are in int format
	items = []interface{}{"one", "two", "four", "eight", "ten"}
	weightsInt := []int{1, 2, 4, 8, 5}
	wRInt := weightedRand.NewWRand(items, weightsInt)
	fmt.Println(wRInt.GetN(5))

	//performing weighted random on a map
	m := map[interface{}]int{
		"one":   1,
		"two":   2,
		"four":  4,
		"eight": 8,
		"ten":   10,
	}
	wRMap := weightedRand.NewWRandByMap(m)
	fmt.Println(wRMap.GetN(5))

	//performing weighted random on a slice of object and a specific field
	type object struct {
		id     int
		weight int
	}
	objects := []object{
		{
			id:     1,
			weight: 1,
		},
		{
			id:     2,
			weight: 2,
		},
		{
			id:     3,
			weight: 3,
		},
	}
	var objs []interface{}
	for _, object := range objects {
		objs = append(objs, object)
	}
	wRObj, err := weightedRand.NewWRandByObject[float64]("weight", objs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(wRObj.GetN(5))

	//poping items by their weights
	items = []interface{}{"one", "two", "four", "eight", "ten"}
	weightsPop := []int{1, 2, 4, 8, 5}
	wRPop := weightedRand.NewWRand(items, weightsPop)
	item, err := wRPop.PopN(1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(item)
}
