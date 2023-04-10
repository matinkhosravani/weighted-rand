package weightedRand

import (
	"fmt"
	"reflect"
)

type weightTypes interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type WRand[T weightTypes] struct {
	items   []interface{}
	weights []T
}

func NewWRand[T weightTypes](items []interface{}, weights []T) *WRand[T] {
	return &WRand[T]{
		items:   items,
		weights: weights,
	}
}

func NewWRandByMap[T weightTypes](m map[interface{}]T) *WRand[T] {
	var items []interface{}
	var weights []T

	for item, weight := range m {
		items = append(items, item)
		weights = append(weights, weight)
	}

	return &WRand[T]{
		items:   items,
		weights: weights,
	}
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

	return &WRand[T]{
			items:   objects,
			weights: weights,
		},
		nil
}
