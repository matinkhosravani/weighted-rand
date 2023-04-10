package weightedRand

import "testing"

func TestNewWRand(t *testing.T) {
	items := []interface{}{"one", "four", "eight"}
	weights := []int{1, 2, 8}
	wr := NewWRand(items, weights)

	if wr.items[0] != items[0] {
		t.Errorf("got %s; want %s", wr.items[0], items[0])
	}

	if wr.weights[0] != weights[0] {
		t.Errorf("got %d; want %d", wr.weights[0], weights[0])
	}
}

func TestNewWRandByMap(t *testing.T) {
	m := map[interface{}]int{
		"one":   1,
		"four":  4,
		"eight": 8,
	}

	wr := NewWRandByMap(m)

	if wr.items[0] != "one" {
		t.Errorf("got %s; want %s", wr.items[0], "one")
	}

	if wr.weights[0] != m["one"] {
		t.Errorf("got %d; want %d", wr.weights[0], m["one"])
	}
}

func TestNewWRandByObject(t *testing.T) {
	type object struct {
		id    int
		name  string
		score int
	}
	objects := []object{
		{
			id:    1,
			name:  "object 1",
			score: 1,
		},
		{
			id:    4,
			name:  "object 2",
			score: 4,
		},
		{
			id:    8,
			name:  "object 3",
			score: 8,
		},
	}
	var objs []interface{}
	for _, object := range objects {
		objs = append(objs, object)
	}

	wr, err := NewWRandByObject[int]("score", objs)

	if err != nil {
		t.Errorf("got %s; want %v", err, nil)
	}
	if wr.items[0] != objects[0] {
		t.Errorf("got %v; want %v", wr.items[0], objects[0])
	}

	if wr.weights[0] != objects[0].score {
		t.Errorf("got %d; want %d", wr.weights[0], objects[0].score)
	}
}
