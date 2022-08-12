package weightedRand

import (
	"math/rand"
	"sync"
	"time"
)

type weightTypes interface {
	~float64 | ~float32 | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type wRand[T weightTypes] struct {
	items   []interface{}
	weights []T
}

func NewWRand[T weightTypes](items []interface{}, weights []T) *wRand[T] {
	return &wRand[T]{
		items:   items,
		weights: weights,
	}
}

func (wR *wRand[T]) GetOne() interface{} {
	cumulativeWeights := cumulativeWeights(wR)

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Float64() * float64(cumulativeWeights[len(cumulativeWeights)-1])

	return matchItem(cumulativeWeights, T(randNum), wR)
}

func (wR *wRand[T]) GetN(n int) []interface{} {
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

func matchItem[T weightTypes](cumulativeWeights []T, randNum T, wR *wRand[T]) interface{} {
	for i, v := range cumulativeWeights {
		if v >= randNum {
			return wR.items[i]
		}
	}

	return nil
}

//example : weightes := [1,2,3,4]
// cWeights : [1,3,6,10]
func cumulativeWeights[T weightTypes](wR *wRand[T]) []T {
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
