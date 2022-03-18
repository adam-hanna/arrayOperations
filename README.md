[![Build Status](https://travis-ci.org/adam-hanna/arrayOperations.svg?branch=master)](https://travis-ci.org/adam-hanna/arrayOperations) [![Coverage Status](https://coveralls.io/repos/github/adam-hanna/arrayOperations/badge.svg?branch=master)](https://coveralls.io/github/adam-hanna/arrayOperations?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/adam-hanna/arrayOperations)](https://goreportcard.com/report/github.com/adam-hanna/arrayOperations) [![GoDoc](https://godoc.org/github.com/adam-hanna/arrayOperations?status.svg)](https://godoc.org/github.com/adam-hanna/arrayOperations)

# arrayOperations
Small utility library for performing common operations on slices in golang.

I don't promise that these are optimized (not even close!), but they work :)

## Quickstart
~~~ go
var a = []int{1, 1, 2, 3}
var b = []int{2, 4}

diff := Difference[int](a, b)

// []int{1, 3}
fmt.Println(diff)
~~~

## API

FindOne[T any](arr []T, guard func(T) bool) (T, bool)

Reduce[T, A any](arr []T, fn func(A, T) A, init A) A

Filter[T any](arr []T, guard func(T) bool) []T

Map[T any](arr []T, transform func(T) T) []T

Distinct[T comparable](arrs ...[]T) []T

Intersect[T comparable](arrs ...[]T) []T

Union[T comparable](arrs ...[]T) []T

Difference[T comparable](arrs ...[]T) []T

## Docs

Docs are available, [here](https://pkg.go.dev/github.com/adam-hanna/arrayOperations)

## License
MIT
