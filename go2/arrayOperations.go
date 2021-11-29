package go2

// Reduce iterates through an array applying the function to each element and the cumulative value and the final state of the cumulative value
//
// arr := []string{"cat","dog","cat","cow"}
// func countAnimals(state map[string]int, animal string) map[string]int {
//   count, ok := state[animal]
//   if !ok {
//     state[animal] = 0
//     return state
//   }
//
//   state[animal] = count + 1
//   return state
// }
// initialState := make([string]int)
// finalState := Reduce(arr, countAnimals, initialState)
// fmt.Println(finalState)
// // output: map["cat":2 "dog":1 "cow":1]
func Reduce[T, A any](arr []T, fn func(A, T) A, init A) A {
  ret := init

  for idx := range arr {
    ret = fn(ret, arr[idx])
  }
  
  return ret
}

// Filter iterates through an array applying the guard function to each element and returns the elements that pass
//
// arr := []int{0,1,2,3,4}
// func isEven(i int) bool {
//   return i % 2 == 0
// }
// newArr := Filter(arr, isEven)
// fmt.Println(newArr)
// // output: [0,2,4]
func Filter[T any](arr []T, guard func(T) bool) []T {
  var ret []T

  for idx := range arr {
    if guard(arr[idx]) {
      ret = append(ret, arr[idx])
    }
  }
  
  return ret
}

// Map iterates through an array applying the transform function to each element and returns the modified array
//
// arr := []int{1,2,3}
// func addTen(i int) int {
//   return i + 10
// }
// newArr := Map(arr, addTen)
// fmt.Println(newArr)
// // output: [11, 12, 13]
func Map[T any](arr []T, transform func(T) T) []T {
  ret := make([]T, len(arr))

  for idx := range arr {
    ret[idx] = transform(arr[idx])
  }
  
  return ret
}

// Distinct returns the unique vals of a slice
//
// [1, 1, 2, 3] >> [1, 2, 3]
func Distinct[T comparable](arrs... []T) []T {
	// put the values of our slice into a map
	// the key's of the map will be the slice's unique values
	m := make(map[T]struct{})
	for idx1 := range arrs {
		for idx2 := range arrs[idx1] {
			m[arrs[idx1][idx2]] = struct{}{}
		}
	}

	// create the output slice and populate it with the map's keys
	res := make([]T, len(m))
	i := 0
	for k := range m {
		res[i] = k
		i++
	}

	return res
}

// Intersect returns a slice of values that are present in all of the input slices
//
// [1, 1, 3, 4, 5, 6] & [2, 3, 6] >> [3, 6]
//
// [1, 1, 3, 4, 5, 6] >> [1, 3, 4, 5, 6]
func Intersect[T comparable](arrs ...[]T) []T {
	m := make(map[T]int)

	var (
		tmpArr []T
		count int
		ok bool
	)
	for idx1 := range arrs {
		tmpArr = Distinct(arrs[idx1])

		for idx2 := range tmpArr {
			count, ok = m[tmpArr[idx2]]
			if !ok {
				m[tmpArr[idx2]] = 1
			} else {
				m[tmpArr[idx2]] = count + 1
			}
		}
	}

	var (
		ret []T
		lenArrs int = len(arrs)
	)
	for k, v := range m {
		if v == lenArrs {
			ret = append(ret, k)
		}
	}

	return ret
}

// Union returns a slice that contains the unique values of all the input slices
//
// [1, 2, 2, 4, 6] & [2, 4, 5] >> [1, 2, 4, 5, 6]
//
// [1, 1, 3, 4, 5, 6] >> [1, 3, 4, 5, 6]
func Union[T comparable](arrs ...[]T) []T {
	m := make(map[T]struct{})

	var tmpArr []T
	for idx := range arrs {
		tmpArr = Distinct(arrs[idx])

		for idx2 := range tmpArr {
			m[tmpArr[idx2]] = struct{}{}
		}
	}

	ret := make([]T, len(m))
	i := 0
	for k := range m {
		ret[i] = k
		i++
	}

	return ret
}

// Difference returns a slice of values that are only present in one of the input slices
//
// [1, 2, 2, 4, 6] & [2, 4, 5] >> [1, 5, 6]
//
// [1, 1, 3, 4, 5, 6] >> [1, 3, 4, 5, 6]
func Difference[T comparable](arrs ...[]T) []T {
	m := make(map[T]int)

	var (
		tmpArr []T
		count int
		ok bool
	)
	for idx1 := range arrs {
		tmpArr = Distinct(arrs[idx1])

		for idx2 := range tmpArr {
			count, ok = m[tmpArr[idx2]]
			if !ok {
				m[tmpArr[idx2]] = 1
			} else {
				m[tmpArr[idx2]] = count + 1
			}
		}
	}

	var (
		ret []T
	)
	for k, v := range m {
		if v == 1 {
			ret = append(ret, k)
		}
	}

	return ret
}
