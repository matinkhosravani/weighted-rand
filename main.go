package main

import (
	"fmt"
	"weightedRand/weightedRand"
)

func main() {
	weights := []float64{1.1, 2.2, 4.3, 8.4, 10.5}
	items := []interface{}{"one", "two", "four", "eight", "ten"}
	wR := weightedRand.NewWRand(items, weights)
	fmt.Println(wR.GetN(1))

	weightsInt := []int{1, 2, 4, 8, 5}
	items = []interface{}{"one", "two", "four", "eight", "ten"}
	wR2 := weightedRand.NewWRand(items, weightsInt)
	fmt.Println(wR2.GetN(1))
}
