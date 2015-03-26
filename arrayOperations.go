// find the intersection of two arrays.
// e.x. a1 = [1 2 2 4 6]; a2 = [2 4 5]
// Intersect(a1, a2) >> [2 4]
func IntersectString(args ...[]string) []string {
	// hold the values of the arrays;
	// the first map of uint64 represents the index of args
	// the second map of strings represents that values in the arg arrays
	tempMap := make(map[uint64]map[string]uint8)
	for n, arg := range args {
		tempMap[uint64(n)] = make(map[string]uint8)
		for idx := range arg {
			tempMap[uint64(n)][arg[idx]] = 0
		}
	}

	// hold the interesctions; incremental keys are successive intersections
	intersectMap := make(map[uint64]map[string]uint8)
	intersectMap[0] = make(map[string]uint8)
	// we need to write the first one
	for key := range tempMap[0] {
		intersectMap[0][key] = 0
	}

	for i := 1; i < len(tempMap); i++ {
		intersectMap[uint64(i)] = make(map[string]uint8)
		for key := range tempMap[uint64(i)] {
			if _, ok := intersectMap[uint64(i-1)][key]; ok {
				intersectMap[uint64(i)][key] = 0
			}
		}
	}

	// write the final val of the intersectMap to an array and return
	tempArray := make([]string, 0)
	for key := range intersectMap[uint64(len(intersectMap)-1)] {
		tempArray = append(tempArray, key)
	}

	return tempArray
}

// find the union of two arrays.
// e.x. a1 = [1 2 2 4 6]; a2 = [2 4 5]
// Union(a1, a2) >> [1 2 4 5 6]
func UnionString(args ...[]string) []string {
	tempMap := make(map[string]uint8)

	for _, arg := range args {
		for idx := range arg {
			tempMap[arg[idx]] = 0
		}
	}

	tempArray := make([]string, 0)
	for key := range tempMap {
		tempArray = append(tempArray, key)
	}
	return tempArray
}

// find the difference of two arrays.
// e.x. a1 = [1 2 2 4 6]; a2 = [2 4 5]
// Difference(a1, a2) >> [5 6]
func DifferenceString(args ...[]string) []string {
	// hold the values of the arrays;
	// the first map of uint64 represents the index of args
	// the second map of strings represents that values in the arg arrays
	tempMap := make(map[uint64]map[string]uint8)
	for n, arg := range args {
		tempMap[uint64(n)] = make(map[string]uint8)
		for idx := range arg {
			tempMap[uint64(n)][arg[idx]] = 0
		}
	}

	// hold the differences
	diffMap := make(map[uint64]map[string]uint8)
	diffMap[0] = make(map[string]uint8)
	// we need to start with one of the arrays
	for key := range tempMap[uint64(0)] {
		diffMap[0][key] = 0
	}

	for i := 1; i < len(tempMap); i++ {
		diffMap[uint64(i)] = make(map[string]uint8)
		for key := range diffMap[uint64(i-1)] {
			if _, ok := tempMap[uint64(i)][key]; !ok {
				diffMap[uint64(i)][key] = 0
			}
		}
		for key := range tempMap[uint64(i)] {
			if _, ok := diffMap[uint64(i-1)][key]; !ok {
				diffMap[uint64(i)][key] = 0
			}
		}
	}

	// write the final val of the diffMap to an array and return
	tempArray := make([]string, 0)
	for key := range diffMap[uint64(len(diffMap)-1)] {
		tempArray = append(tempArray, key)
	}

	return tempArray
}

// Remove duplicate values from one array.
// e.x. a1 = [1 2 2 4 6]
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
