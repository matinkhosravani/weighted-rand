package weighted_rand

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func (wR *WRand[T]) GetOne() interface{} {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := rand.Float64() * float64(wR.cumulativeWeights[len(wR.cumulativeWeights)-1])

	return wR.matchItem(T(randNum))
}

func (wR *WRand[T]) GetN(n int) []interface{} {
	var items []interface{}
	var wg sync.WaitGroup

	itemsChannel := make(chan interface{})

	// seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	last := wR.cumulativeWeights[len(wR.cumulativeWeights)-1]

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			itemsChannel <- func() interface{} {
				randNum := rand.Float64() * float64(last)
				return wR.matchItem(T(randNum))
			}()
		}()
	}

	go func() {
		wg.Wait()
		close(itemsChannel)
	}()

	for item := range itemsChannel {
		items = append(items, item)
	}

	return items
}

func (wR *WRand[T]) PopN(n int) ([]interface{}, error) {
	if n > len(wR.items) {
		return nil, fmt.Errorf("can't pop %d items from a slice with size of %d", n, len(wR.items))
	}

	var items []interface{}
	for i := 0; i < n; i++ {
		item := wR.GetOne()
		for j, other := range wR.items {
			if other == item {
				wR.items = append(wR.items[:j], wR.items[j+1:]...)
				wR.weights = append(wR.weights[:j], wR.weights[j+1:]...)
				wR.cumulativeWeights = cumulativeWeights(wR.items, wR.weights)
			}
		}
		items = append(items, item)
	}

	return items, nil
}

func (wR *WRand[T]) matchItem(randNum T) interface{} {
	for i, v := range wR.cumulativeWeights {
		if v >= randNum {
			return wR.items[i]
		}
	}

	return nil
}
