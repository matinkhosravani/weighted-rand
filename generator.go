package weighted_rand

import (
	"fmt"
	"reflect"
)

type weightTypes interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type WRand[T weightTypes] struct {
	items             []interface{}
	weights           []T
	cumulativeWeights []T
}

func NewWRand[T weightTypes](items []interface{}, weights []T) *WRand[T] {
	return &WRand[T]{
		items:             items,
		weights:           weights,
		cumulativeWeights: cumulativeWeights(items, weights),
	}
}

func NewWRandByMap[T weightTypes](m map[interface{}]T) *WRand[T] {
	var items []interface{}
	var weights []T

	for item, weight := range m {
		items = append(items, item)
		weights = append(weights, weight)
	}

	return NewWRand(items, weights)
}

func NewWRandByObject[T weightTypes](fieldName string, objects []interface{}) (*WRand[T], error) {
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

	return NewWRand(objects, weights), nil
}

// example : weights := [1,2,3,4]
// cWeights : [1,3,6,10]
func cumulativeWeights[T weightTypes](items []interface{}, weights []T) []T {
	cWeights := make([]T, len(items))
	for i, v := range weights {
		if i == 0 {
			cWeights[i] = v
		} else {
			cWeights[i] = v + cWeights[i-1]
		}
	}
	return cWeights
}
