// Find Second largest element in an array

package main

import (
	"fmt"
	"math"
)

func main() {

	inp := []int{12, 35, 1, 10, 34, 1}

	var secondMax, firstMax = math.MinInt64, math.MinInt64

	for i := 0; i < len(inp); i++ {
		if inp[i] > firstMax {
			secondMax = firstMax
			firstMax = inp[i]
		} else if (inp[i] != firstMax) && (inp[i] > secondMax) {
			secondMax = inp[i]
		}
	}
	fmt.Println("the second largest Number is ", secondMax)

}
