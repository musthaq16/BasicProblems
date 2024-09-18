// Program to find the maximum element in a Matrix

package main

import (
	"fmt"
	"math"
)

const N = 4
const M = 4

func main() {

	inp := [N][M]int{
		{1, 2, 3, 4},
		{25, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	max := math.MinInt64

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if max < inp[i][j] {
				max = inp[i][j]
			}
		}
	}

	fmt.Println("The max number is ", max)

}
