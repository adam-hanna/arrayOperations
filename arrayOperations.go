package main

import (
	"fmt"
	"math"
	"math/rand"
	// "strconv"
	"sort"
	"time"
)

/* ***************************************************************
*
* THIS SECTION IS FOR STRINGS
*
/* *************************************************************** */

// find the intersection of two arrays.
// e.g. a1 = [1 2 2 4 6]; a2 = [2 4 5]
// Intersect(a1, a2) >> [2 4]
func IntersectString(args ...[]string) []string {
	// create a map to count all the instances of the strings
	arrLength := len(args)
	tempMap := make(map[string]int)
	for _, arg := range args {
		tempArr := DistinctString(arg)
		for idx := range tempArr {
			// how many times have we encountered this elem?
			if _, ok := tempMap[tempArr[idx]]; ok {
				tempMap[tempArr[idx]]++
			} else {
				tempMap[tempArr[idx]] = 1
			}
		}
	}

	// find the keys equal to the length of the input args
	tempArray := make([]string, 0)
	for key, val := range tempMap {
		if val == arrLength {
			tempArray = append(tempArray, key)
		}
	}

	return tempArray
}

// find the intersection of two arrays using a multidimensional array as inputs
func IntersectStringArr(arr [][]string) []string {
	// create a map to count all the instances of the strings
	arrLength := len(arr)
	tempMap := make(map[string]int)
	for idx1 := range arr {
		tempArr := DistinctString(arr[idx1])
		for idx2 := range tempArr {
			// how many times have we encountered this elem?
			if _, ok := tempMap[tempArr[idx2]]; ok {
				tempMap[tempArr[idx2]]++
			} else {
				tempMap[tempArr[idx2]] = 1
			}
		}
	}

	// find the keys equal to the length of the input args
	tempArray := make([]string, 0)
	for key, val := range tempMap {
		if val == arrLength {
			tempArray = append(tempArray, key)
		}
	}

	return tempArray
}

// find the union of two arrays.
// e.g. a1 = [1 2 2 4 6]; a2 = [2 4 5]
// Union(a1, a2) >> [1 2 4 5 6]
func UnionString(args ...[]string) []string {
	// create a temporary map to hold the contents of the arrays
	tempMap := make(map[string]uint8)

	// write the contents of the arrays as keys to the map. The map values don't matter
	for _, arg := range args {
		for idx := range arg {
			tempMap[arg[idx]] = 0
		}
	}

	// the map keys are now unique instances of all of the array contents
	tempArray := make([]string, 0)
	for key := range tempMap {
		tempArray = append(tempArray, key)
	}

	return tempArray
}

// find the union of two arrays using a multidimensional array as inputs
func UnionStringArr(arr [][]string) []string {
	// create a temporary map to hold the contents of the arrays
	tempMap := make(map[string]uint8)

	// write the contents of the arrays as keys to the map. The map values don't matter
	for idx1 := range arr {
		for idx2 := range arr[idx1] {
			tempMap[arr[idx1][idx2]] = 0
		}
	}

	// the map keys are now unique instances of all of the array contents
	tempArray := make([]string, 0)
	for key := range tempMap {
		tempArray = append(tempArray, key)
	}

	return tempArray
}

// find the difference of two arrays.
// e.g. a1 = [1 2 2 4 6]; a2 = [2 4 5]
// Difference(a1, a2) >> [5 6]
func DifferenceString(args ...[]string) []string {
	// create a temporary map to hold the contents of the arrays
	tempMap := make(map[string]int)
	for _, arg := range args {
		tempArr := DistinctString(arg)
		for idx := range tempArr {
			// how many times have we encountered this elem?
			if _, ok := tempMap[tempArr[idx]]; ok {
				tempMap[tempArr[idx]]++
			} else {
				tempMap[tempArr[idx]] = 1
			}
		}
	}

	// write the final val of the diffMap to an array and return
	tempArray := make([]string, 0)
	for key, val := range tempMap {
		if val == 1 {
			tempArray = append(tempArray, key)
		}
	}

	return tempArray
}

// find the difference of two arrays using a multidimensional array as inputs
func DifferenceStringArr(arr [][]string) []string {
	// create a temporary map to hold the contents of the arrays
	tempMap := make(map[string]int)
	for idx1 := range arr {
		tempArr := DistinctString(arr[idx1])
		for idx2 := range tempArr {
			// how many times have we encountered this elem?
			if _, ok := tempMap[tempArr[idx2]]; ok {
				tempMap[tempArr[idx2]]++
			} else {
				tempMap[tempArr[idx2]] = 1
			}
		}
	}

	// write the final val of the diffMap to an array and return
	tempArray := make([]string, 0)
	for key, val := range tempMap {
		if val == 1 {
			tempArray = append(tempArray, key)
		}
	}

	return tempArray
}

// Remove duplicate values from one array.
// e.g. a1 = [1 2 2 4 6]
// Distinct(a1) >> [1 2 4 6]
func DistinctString(arg []string) []string {
	tempMap := make(map[string]uint8)

	for idx := range arg {
		tempMap[arg[idx]] = 0
	}

	tempArray := make([]string, 0)
	for key := range tempMap {
		tempArray = append(tempArray, key)
	}
	return tempArray
}

/* ***************************************************************
*
* THIS SECTION IS FOR uint64's
*
/* *************************************************************** */

// find the intersection of two arrays.
// e.g. a1 = [1 2 2 4 6]; a2 = [2 4 5]
// Intersect(a1, a2) >> [2 4]
func IntersectUint64(args ...[]uint64) []uint64 {
	// create a map to count all the instances of the strings
	arrLength := len(args)
	tempMap := make(map[uint64]int)
	for _, arg := range args {
		tempArr := DistinctUint64(arg)
		for idx := range tempArr {
			// how many times have we encountered this elem?
			if _, ok := tempMap[tempArr[idx]]; ok {
				tempMap[tempArr[idx]]++
			} else {
				tempMap[tempArr[idx]] = 1
			}
		}
	}

	// find the keys equal to the length of the input args
	tempArray := make([]uint64, 0)
	for key, val := range tempMap {
		if val == arrLength {
			tempArray = append(tempArray, key)
		}
	}

	return tempArray
}

// find the intersection of two arrays of distinct vals.
func DistinctIntersectUint64(args ...[]uint64) []uint64 {
	// create a map to count all the instances of the strings
	arrLength := len(args)
	tempMap := make(map[uint64]int)
	for _, arg := range args {
		for idx := range arg {
			// how many times have we encountered this elem?
			if _, ok := tempMap[arg[idx]]; ok {
				tempMap[arg[idx]]++
			} else {
				tempMap[arg[idx]] = 1
			}
		}
	}

	// find the keys equal to the length of the input args
	tempArray := make([]uint64, 0)
	for key, val := range tempMap {
		if val == arrLength {
			tempArray = append(tempArray, key)
		}
	}

	return tempArray
}

func sortedIntersectUintHelper(a1 []uint64, a2 []uint64) []uint64 {
	intersection := make([]uint64, 0)
	n1 := len(a1)
	n2 := len(a2)
	i := 0
	j := 0
	for i < n1 && j < n2 {
		switch {
		case a1[i] > a2[j]:
			j++
		case a2[j] > a1[i]:
			i++
		default:
			intersection = append(intersection, a1[i])
			i++
			j++
		}
	}
	return intersection
}

// find the intersection of two sorted arrays
// assumes no dupes!
// NOTE(@adam-hanna): further improve performance by sorting from smallest to largest
// array length?
func SortedIntersectUint64(args ...[]uint64) []uint64 {
	// create an array to hold the intersection and write the first array to it
	tempIntersection := args[0]
	argsLen := len(args)

	for k := 1; k < argsLen; k++ {
		// do we have any intersections?
		switch len(tempIntersection) {
		case 0:
			// nope! Give them an empty array!
			return tempIntersection

		default:
			// yup, keep chugging
			tempIntersection = sortedIntersectUintHelper(tempIntersection, args[k])
		}
	}

	return tempIntersection
}

// find the intersection of two arrays using a multidimensional array as inputs
func IntersectUint64Arr(arr [][]uint64) []uint64 {
	// create a map to count all the instances of the strings
	arrLength := len(arr)
	tempMap := make(map[uint64]int)
	for idx1 := range arr {
		tempArr := DistinctUint64(arr[idx1])
		for idx2 := range tempArr {
			// how many times have we encountered this elem?
			if _, ok := tempMap[tempArr[idx2]]; ok {
				tempMap[tempArr[idx2]]++
			} else {
				tempMap[tempArr[idx2]] = 1
			}
		}
	}

	// find the keys equal to the length of the input args
	tempArray := make([]uint64, 0)
	for key, val := range tempMap {
		if val == arrLength {
			tempArray = append(tempArray, key)
		}
	}

	return tempArray
}

// find the intersection of two arrays using a multidimensional array as inputs
// assumes no dupes and only works on arrays which are sorted
func SortedIntersectUint64Arr(arr [][]uint64) []uint64 {
	// create an array to hold the intersection and write the first array to it
	tempIntersection := arr[0]
	argsLen := len(arr)

	for k := 1; k < argsLen; k++ {
		// do we have any intersections?
		switch len(tempIntersection) {
		case 0:
			// nope! Give them an empty array!
			return tempIntersection

		default:
			// yup, keep chugging
			tempIntersection = sortedIntersectUintHelper(tempIntersection, arr[k])
		}
	}

	return tempIntersection
}

// find the intersection of two distinct arrays using a multidimensional array as inputs
func DistinctIntersectUint64Arr(arr [][]uint64) []uint64 {
	// create a map to count all the instances of the strings
	arrLength := len(arr)
	tempMap := make(map[uint64]int)
	for idx1 := range arr {
		for idx2 := range arr[idx1] {
			// how many times have we encountered this elem?
			if _, ok := tempMap[arr[idx1][idx2]]; ok {
				tempMap[arr[idx1][idx2]]++
			} else {
				tempMap[arr[idx1][idx2]] = 1
			}
		}
	}

	// find the keys equal to the length of the input args
	tempArray := make([]uint64, 0)
	for key, val := range tempMap {
		if val == arrLength {
			tempArray = append(tempArray, key)
		}
	}

	return tempArray
}

// find the union of two arrays.
// e.g. a1 = [1 2 2 4 6]; a2 = [2 4 5]
// Union(a1, a2) >> [1 2 4 5 6]
func UnionUint64(args ...[]uint64) []uint64 {
	// create a temporary map to hold the contents of the arrays
	tempMap := make(map[uint64]uint8)

	// write the contents of the arrays as keys to the map. The map values don't matter
	for _, arg := range args {
		for idx := range arg {
			tempMap[arg[idx]] = 0
		}
	}

	// the map keys are now unique instances of all of the array contents
	tempArray := make([]uint64, 0)
	for key := range tempMap {
		tempArray = append(tempArray, key)
	}

	return tempArray
}

// find the union of two arrays using a multidimensional array as inputs
func UnionUint64Arr(arr [][]uint64) []uint64 {
	// create a temporary map to hold the contents of the arrays
	tempMap := make(map[uint64]uint8)

	// write the contents of the arrays as keys to the map. The map values don't matter
	for idx1 := range arr {
		for idx2 := range arr[idx1] {
			tempMap[arr[idx1][idx2]] = 0
		}
	}

	// the map keys are now unique instances of all of the array contents
	tempArray := make([]uint64, 0)
	for key := range tempMap {
		tempArray = append(tempArray, key)
	}

	return tempArray
}

// find the difference of two arrays.
// e.g. a1 = [1 2 2 4 6]; a2 = [2 4 5]
// Difference(a1, a2) >> [5 6]
func DifferenceUint64(args ...[]uint64) []uint64 {
	// create a temporary map to hold the contents of the arrays
	tempMap := make(map[uint64]int)
	for _, arg := range args {
		tempArr := DistinctUint64(arg)
		for idx := range tempArr {
			// how many times have we encountered this elem?
			if _, ok := tempMap[tempArr[idx]]; ok {
				tempMap[tempArr[idx]]++
			} else {
				tempMap[tempArr[idx]] = 1
			}
		}
	}

	// write the final val of the diffMap to an array and return
	tempArray := make([]uint64, 0)
	for key, val := range tempMap {
		if val == 1 {
			tempArray = append(tempArray, key)
		}
	}

	return tempArray
}

// find the difference of two arrays using a multidimensional array as inputs
func DifferenceUint64Arr(arr [][]uint64) []uint64 {
	// create a temporary map to hold the contents of the arrays
	tempMap := make(map[uint64]int)
	for idx1 := range arr {
		tempArr := DistinctUint64(arr[idx1])
		for idx2 := range tempArr {
			// how many times have we encountered this elem?
			if _, ok := tempMap[tempArr[idx2]]; ok {
				tempMap[tempArr[idx2]]++
			} else {
				tempMap[tempArr[idx2]] = 1
			}
		}
	}

	// write the final val of the diffMap to an array and return
	tempArray := make([]uint64, 0)
	for key, val := range tempMap {
		if val == 1 {
			tempArray = append(tempArray, key)
		}
	}

	return tempArray
}

// Remove duplicate values from one array.
// e.g. a1 = [1 2 2 4 6]
// Distinct(a1) >> [1 2 4 6]
func DistinctUint64(arg []uint64) []uint64 {
	tempMap := make(map[uint64]uint8)

	for idx := range arg {
		tempMap[arg[idx]] = 0
	}

	tempArray := make([]uint64, 0)
	for key := range tempMap {
		tempArray = append(tempArray, key)
	}
	return tempArray
}

/* ***************************************************************
*
* THIS SECTION IS FOR a little helper benchmark script which can be
* used to test the relative changes in performance when making
* changes to these algorithms.
*
* To use, un-comment out the packages in import.
*
/* *************************************************************** */

func main() {
	aBig := make([][]uint64, 0)
	a1 := make([]uint64, 0)
	a2 := make([]uint64, 0)

	aBig = append(aBig, a1)
	aBig = append(aBig, a2)

	i := 1
	j := 1

	for i = 1; i <= 7; i++ {
		// create the arrays
		arrLength := len(a1)
		for j = 0; j < ((int(math.Pow(10, float64(i))) / 2) - arrLength); j++ {
			a1 = append(a1, uint64(rand.Int()))
			a2 = append(a2, uint64(rand.Int()))
		}

		// get distinct vals
		a1 = DistinctUint64(a1)
		a2 = DistinctUint64(a2)

		// sort the arrs
		sort.Sort(uintArray(a1))
		sort.Sort(uintArray(a2))

		// now run the test
		t0 := time.Now()
		SortedIntersectUint64Arr(aBig)
		t1 := time.Now()
		fmt.Printf("n of %v took %v to run.\n", len(a1)+len(a2), t1.Sub(t0))
	}

}

// Need to create some helpers for our sort.Sort
type uintArray []uint64

func (s uintArray) Len() int           { return len(s) }
func (s uintArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s uintArray) Less(i, j int) bool { return s[i] < s[j] }
