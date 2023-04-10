package weightedRand

import (
	"reflect"
	"testing"
)

func TestGetN(t *testing.T) {
	items := []interface{}{"a", "b", "c"}
	weights := []float64{0.01, 0.1, 0.2}

	wR := NewWRand(items, weights)

	numIterations := 1000
	results := wR.GetN(numIterations)

	if len(results) != numIterations {
		t.Errorf("GetN did not return %d items, returned %d", numIterations, len(results))
	}

	// check that each item is one of the possible options
	for _, item := range results {
		if !contains(items, item) {
			t.Errorf("GetN returned unexpected item %v", item)
		}
	}
}

func BenchmarkGetN(b *testing.B) {
	items := []interface{}{"a", "b", "c", "d", "e"}
	weights := []float64{0.1, 0.2, 0.3, 0.2, 0.2}

	wR := NewWRand(items, weights)

	for i := 0; i < b.N; i++ {
		wR.GetN(5)
	}
}

func contains(items []interface{}, item interface{}) bool {
	for _, i := range items {
		if reflect.DeepEqual(i, item) {
			return true
		}
	}
	return false
}
