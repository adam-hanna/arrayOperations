[![Build Status](https://travis-ci.org/adam-hanna/arrayOperations.svg?branch=master)](https://travis-ci.org/adam-hanna/arrayOperations) [![Coverage Status](https://coveralls.io/repos/github/adam-hanna/arrayOperations/badge.svg?branch=master)](https://coveralls.io/github/adam-hanna/arrayOperations?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/adam-hanna/arrayOperations)](https://goreportcard.com/report/github.com/adam-hanna/arrayOperations) [![GoDoc](https://godoc.org/github.com/adam-hanna/arrayOperations?status.svg)](https://godoc.org/github.com/adam-hanna/arrayOperations)

# arrayOperations
Small library for performing union, intersect, difference and distinct operations on slices in goLang

I don't promise that these are optimized (not even close!), but they work :)

## Usage

~~~ go
var a = []int{1, 1, 2, 3}
var b = []int{2, 4}

z, ok := Difference(a, b)
if !ok {
	fmt.Println("Cannot find difference")
}

slice, ok := z.Interface().([]int)
if !ok {
	fmt.Println("Cannot convert to slice")
}
fmt.Println(slice, reflect.TypeOf(slice)) // [1, 3] []int
~~~

## License
MIT
