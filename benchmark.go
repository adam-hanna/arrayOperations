package main

import (
	"arrayOperations"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	a1 := make([]string, 0)
	a2 := make([]string, 0)

	i := 1
	j := 1

	for i = 1; i <= 7; i++ {
		// create the arrays
		arrLength := len(a1)
		for j = 0; j < ((int(math.Pow(10, float64(i))) / 2) - arrLength); j++ {
			a1 = append(a1, strconv.Itoa(rand.Int()))
			a2 = append(a2, strconv.Itoa(rand.Int()))
		}

		t0 := time.Now()
		IntersectString(a1, a2)
		t1 := time.Now()
		fmt.Printf("n of %v took %v to run.\n", 2*len(a1), t1.Sub(t0))
	}

}
