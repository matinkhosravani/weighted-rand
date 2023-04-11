# Weighted Random
**weighted_random** - Go package for selecting items from a slice, map, or objects based on scores or weights.
## Description
weighted_random is a Go package that provides a simple and efficient way to perform weighted random selection on various data structures such as slices, maps, and objects.
The package allows you to assign scores or weights to items, and then select items randomly based on their scores or weights.
The algorithm used in this package is based on the weighted random selection algorithm described in [this](https://trekhleb.medium.com/weighted-random-in-javascript-4748ab3a1500) medium article.
## Installation

You can install weighted_random using the go get command:

```sh
go get github.com/matinkhosravani/weighted-rand
```
## Usage
Importing the Package

```go
import "github.com/matinkhosravani/weighted-rand"
```

### Creating a Weighted Randomizer

To perform weighted random selection, you need to create a weighted_random.WRand object by passing in the items and their corresponding scores or weights. The items can be of any type, and the scores or weights should be of type float64 or int, depending on your use case.

```go

// Example with float64 weights
items := []interface{}{"one", "two", "four", "eight", "ten"}
weights := []float64{1.1, 2.2, 4.3, 8.4, 10.5}
wR := weighted_random.NewWRand(items, weights)

// Example with int weights
items = []interface{}{"one", "two", "four", "eight", "ten"}
weightsInt := []int{1, 2, 4, 8, 5}
wRInt := weighted_random.NewWRand(items, weightsInt)
```
### Performing Weighted Random Selection

Once you have created a weighted_random.WRand object, you can use the GetN() method to perform weighted random selection and get n items from the items list, based on their scores or weights.

```go

// Example of performing weighted random selection and getting 5 items
result := wR.GetN(5)
fmt.Println(result)
```
### Performing Weighted Random Selection on a Map

You can also perform weighted random selection on a map by using the NewWRandByMap() method and passing in a map with keys as items and values as scores or weights.

```go

// Example of performing weighted random selection on a map
m := map[interface{}]int{
"one":   1,
"two":   2,
"four":  4,
"eight": 8,
"ten":   10,
}
wRMap := weighted_random.NewWRandByMap(m)
result := wRMap.GetN(5)
fmt.Println(result)
```
### Performing Weighted Random Selection on a Slice of Objects

If you have a slice of objects and you want to perform weighted random selection based on a specific field of the objects, you can use the NewWRandByObject() method and pass in the field name as a string.

```go
// Example of performing weighted random selection on a slice of objects
type object struct {
		id     int
		weight int
	}
	objects := []object{
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
	for _, object := range objects {
		objs = append(objs, object)
	}
	wRObj, err := NewWRandByObject[float64]("weight", objs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(wRObj.GetN(5))
```

### Performing Weighted Random Pop

Once you have created a weighted_random.WRand object, you can use the PopN() method to perform weighted random Pop and get n items from the items list, based on their scores or weights.

```go

// Example of performing weighted random selection and getting 5 items
result := wR.PopN(5)
fmt.Println(result)
```

## License

Weighted Rand is open-source software released under the [MIT License](https://github.com/matinkhosravani/weighted-rand/blob/main/LICENSE). Please see the [LICENSE](https://github.com/matinkhosravani/weighted-rand/blob/main/LICENSE) file for more details.