package weightedRand

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

type weightTypes interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64
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

func NewWRandByMap[T weightTypes](m map[interface{}]T) *wRand[T] {
	var items []interface{}
	var weights []T

	for item, weight := range m {
		items = append(items, item)
		weights = append(weights, weight)
	}

	return &wRand[T]{
		items:   items,
		weights: weights,
	}
}

func NewWRandByObject[T weightTypes](fieldName string, objects []interface{}) (*wRand[T], error) {
	var weights []T

	for _, object := range objects {
		r := reflect.ValueOf(object)
		f := reflect.Indirect(r).FieldByName(fieldName)
		if !f.IsValid() {
			return nil, fmt.Errorf("field %s does not exist in object %v", fieldName, object)
		}
		switch f.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			weights = append(weights, T(f.Int()))
		case reflect.Float32, reflect.Float64:
			weights = append(weights, T(f.Float()))
		}
	}

	return &wRand[T]{
			items:   objects,
			weights: weights,
		},
		nil
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
func (wR *wRand[T]) PopN(n int) ([]interface{}, error) {
	fmt.Println(wR.weights)
	if n > len(wR.items) {
		return nil, fmt.Errorf("can't pop %d items from an slice with size of %d", n, len(wR.items))
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
