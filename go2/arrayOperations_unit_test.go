package go2

import (
	"fmt"
	"reflect"
	"testing"
)

func isEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	m1 := make(map[T]int)
	var (
		count int
		ok bool
	)
	for idx := range a {
		count, ok = m1[a[idx]]	
		if !ok {
			m1[a[idx]] = 1
		} else {
			m1[a[idx]] = count + 1
		}
	}

	m2 := make(map[T]int)
	for idx := range b {
		count, ok = m2[b[idx]]	
		if !ok {
			m2[b[idx]] = 1
		} else {
			m2[b[idx]] = count + 1
		}
	}

	for k, v := range m1 {
		count, ok = m2[k]
		if ! ok {
			return false
		}
		if count != v {
			return false
		}
	}
	for k, v := range m2 {
		count, ok = m1[k]
		if ! ok {
			return false
		}
		if count != v {
			return false
		}
	}

	return true
}

func countAnimals(state map[string]int, animal string) map[string]int {
  count, ok := state[animal]
  if !ok {
    state[animal] = 1
    return state
  }

  state[animal] = count + 1
  return state
}

func TestReduce(t *testing.T) {
  arr := []string{"cat","dog","cat","cow"}
  initialState := make(map[string]int)
  finalState := Reduce(arr, countAnimals, initialState)
  expected := map[string]int{
    "cat": 2,
    "dog": 1,
    "cow": 1,
  }
  
  if !reflect.DeepEqual(finalState, expected) {
    t.Errorf("expected %v, received %v", expected, finalState)
  }
}

func guard(i int) bool {
  return i % 2 == 0
}

func TestFilter(t *testing.T) {
  arr := []int{0,1,2,3,4}
  newArr := Filter(arr, guard)
  expected := []int{0,2,4}
  
  if !reflect.DeepEqual(newArr, expected) {
    t.Errorf("expected %v, received %v", expected, newArr)
  }
}

func addTen(i int) int {
  return i + 10
}

func TestMap(t *testing.T) {
  arr := []int{1,2,3}
  newArr := Map(arr, addTen)
  expected := []int{11,12,13}
  
  if !reflect.DeepEqual(newArr, expected) {
    t.Errorf("expected %v, received %v", expected, newArr)
  }
}

type testDistinctInput[T comparable] [][]T
type testDistinctOutput[T comparable] []T
type testDistinct[T comparable] struct {
	input testDistinctInput[T]
	output testDistinctOutput[T]
}

func TestDistinct(t *testing.T) {
		for i, tt := range []testDistinct[int] {
			testDistinct[int]{
				input: testDistinctInput[int]{
					[]int{1, 2, 3, 3, 4},
				},
				output: []int{1, 2, 3, 4},
			},
		} {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				actual := Distinct(tt.input...)

				if !isEqual(actual, tt.output) {
					t.Errorf("expected: %v %T, received: %v %T", tt.output, tt.output, actual, actual)
				}
		})
	}
}

type testIntersectInput[T comparable] [][]T
type testIntersectOutput[T comparable] []T
type testIntersect[T comparable] struct {
	input testIntersectInput[T]
	output testIntersectOutput[T]
}

func TestIntersect(t *testing.T) {
		for i, tt := range []testIntersect[int] {
			testIntersect[int]{
				input: testIntersectInput[int]{
					[]int{1, 2, 3, 3, 4},
					[]int{1, 2, 4},
					[]int{1, 3, 4},
				},
				output: []int{1, 4},
			},
		} {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				actual := Intersect(tt.input...)

				if !isEqual(actual, tt.output) {
					t.Errorf("expected: %v %T, received: %v %T", tt.output, tt.output, actual, actual)
				}
		})
	}
}

type testUnionInput[T comparable] [][]T
type testUnionOutput[T comparable] []T
type testUnion[T comparable] struct {
	input testUnionInput[T]
	output testUnionOutput[T]
}

func TestUnion(t *testing.T) {
		for i, tt := range []testUnion[int] {
			testUnion[int]{
				input: testUnionInput[int]{
					[]int{1, 2, 3, 3, 4},
					[]int{1, 2, 4, 5},
					[]int{1, 3, 4, 6},
				},
				output: []int{1, 2, 3, 4, 5, 6},
			},
		} {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				actual := Union(tt.input...)

				if !isEqual(actual, tt.output) {
					t.Errorf("expected: %v %T, received: %v %T", tt.output, tt.output, actual, actual)
				}
		})
	}
}

type testDifferenceInput[T comparable] [][]T
type testDifferenceOutput[T comparable] []T
type testDifference[T comparable] struct {
	input testDifferenceInput[T]
	output testDifferenceOutput[T]
}

func TestDifference(t *testing.T) {
		for i, tt := range []testDifference[int] {
			testDifference[int]{
				input: testDifferenceInput[int]{
					[]int{1, 2, 3, 3, 4},
					[]int{1, 2, 5},
					[]int{1, 3, 6},
				},
				output: []int{4, 5, 6},
			},
			//testDifference[interface{}]{
				//input: testDifferenceInput[string]{
					//[]string{"1", "2", "3", "3", "4"},
					//[]string{"1", "2", "5"},
					//[]string{"1", "3", "6"},
				//},
				//output: []string{"4", "5", "6"},
			//},
		} {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				actual := Difference(tt.input...)

				if !isEqual(actual, tt.output) {
					t.Errorf("expected: %v %T, received: %v %T", tt.output, tt.output, actual, actual)
				}
		})
	}
}
