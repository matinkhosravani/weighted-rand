package main

import (
	"fmt"
	"os"
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

	type miniAd struct {
		id     int
		weight int
	}
	miniAds := []miniAd{
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
	for _, miniAd := range miniAds {
		objs = append(objs, miniAd)
	}
	wRObj, err := weightedRand.NewWRandByObject[float64]("weight", objs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(wRObj.GetN(5))
}
