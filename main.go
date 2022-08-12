package main

import (
	"fmt"
	"weightedRand/weightedRand"
)

func main() {
	//items to pick
	items := []interface{}{"one", "two", "four", "eight", "ten"}
	weights := []float64{1.1, 2.2, 4.3, 8.4, 10.5}
	wR := weightedRand.NewWRand(items, weights)
	fmt.Println(wR.GetN(5))

	//items to pick
	items = []interface{}{"one", "two", "four", "eight", "ten"}
	weightsInt := []int{1, 2, 4, 8, 5}
	wRInt := weightedRand.NewWRand(items, weightsInt)
	fmt.Println(wRInt.GetN(5))

	m := map[interface{}]int{
		"one":   1,
		"two":   2,
		"four":  4,
		"eight": 8,
		"ten":   10,
	}
	wRMap := weightedRand.NewWRandByMap(m)
	fmt.Println(wRMap.GetN(5))
}
