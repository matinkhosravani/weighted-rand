package weightedRand

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func (wR *WRand[T]) GetOne() interface{} {
	cumulativeWeights := cumulativeWeights(wR)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := rand.Float64() * float64(cumulativeWeights[len(cumulativeWeights)-1])

	return wR.matchItem(cumulativeWeights, T(randNum))
}

func (wR *WRand[T]) GetN(n int) []interface{} {
	var items []interface{}
	var wg sync.WaitGroup
	var mutex sync.RWMutex

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			items = append(items, wR.GetOne())
			mutex.Unlock()
		}()
	}
	wg.Wait()

	return items
}

func (wR *WRand[T]) PopN(n int) ([]interface{}, error) {
	if n > len(wR.items) {
		return nil, fmt.Errorf("can't pop %d items from a slice with size of %d", n, len(wR.items))
	}
	var items []interface{}
	for i := 0; i < n; i++ {
		item := wR.GetOne()
		for i, other := range wR.items {
			if other == item {
				wR.items = append(wR.items[:i], wR.items[i+1:]...)
				wR.weights = append(wR.weights[:i], wR.weights[i+1:]...)
			}
		}
		items = append(items, item)
	}

	return items, nil
}

func (wR *WRand[T]) matchItem(cumulativeWeights []T, randNum T) interface{} {
	for i, v := range cumulativeWeights {
		if v >= randNum {
			return wR.items[i]
		}
	}

	return nil
}

// example : weightes := [1,2,3,4]
// cWeights : [1,3,6,10]
func cumulativeWeights[T weightTypes](wR *WRand[T]) []T {
	cWeights := make([]T, len(wR.items))
	for i, v := range wR.weights {
		if i == 0 {
			cWeights[i] = v
		} else {
			cWeights[i] = v + cWeights[i-1]
		}
	}
	return cWeights
}
