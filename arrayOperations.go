// find the intersection of two arrays.
// e.g. a1 = [1 2 2 4 6]; a2 = [2 4 5]
// Intersect(a1, a2) >> [2 4]
func IntersectString(args ...[]string) []string {
	// create a map to count all the instances of the strings
	arrLength := len(args)
	tempMap := make(map[string]int)
	for _, arg := range args {
		tempArr := Distinct(arg)
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
	for key := range tempMap {
		if tempMap[key] == arrLength {
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

// find the difference of two arrays.
// e.g. a1 = [1 2 2 4 6]; a2 = [2 4 5]
// Difference(a1, a2) >> [5 6]
func DifferenceString(args ...[]string) []string {
	// create a temporary map to hold the contents of the arrays
	tempMap := make(map[string]int)
	for _, arg := range args {
		tempArr := Distinct(arg)
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
	for key := range tempMap {
		if tempMap[key] == 1 {
			tempArray = append(tempArray, key)
		}
	}

	return tempArray
}

// Remove duplicate values from one array.
// e.g. a1 = [1 2 2 4 6]
// Distinct(a1) >> [1 2 4 6]
func Distinct(arg []string) []string {
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
