// Find a pair with the given sum in an array

package main

import (
	"fmt"
	"sort"
)

func main() {
	inp := []int{8, 7, 2, 5, 3, 1}
	target := 9
	// inp := []int{5, 2, 6, 8, 1, 9}
	// target := 12

	// findAPairBySortingMethod(inp, len(inp), target)
	findAPairByHashMap(inp, target)
}

// BruteForce Approach to find a pair
func findAPair(inp []int, length, target int) {

	for i := 0; i < length; i++ {
		for j := i + 1; j < length-1; j++ {
			if target == inp[i]+inp[j] {
				fmt.Println(" the pair found...", inp[i], inp[j])
				return
			}

		}
	}
	fmt.Println(" the pair not Found ")
}

// Sorting Approach to find a pair
func findAPairBySortingMethod(inp []int, length, target int) {

	sort.Ints(inp)

	low := 0
	high := length - 1

	for low < high {
		sum := inp[low] + inp[high]
		if target == sum {
			fmt.Println(" the pair found...", inp[low], inp[high])
			return
		}

		if sum < target {
			low++
		}
		if sum > target {
			high--
		}
	}
	fmt.Println(" the pair not Found ")

}

// HashMap Approach to find a pair
func findAPairByHashMap(inp []int, target int) {

	posMap := make(map[int]int)

	for i, num := range inp {

		if _, found := posMap[target-num]; found {
			fmt.Println(" the pair found...", target-num, num)
			return
		}

		posMap[num] = i

	}

	fmt.Println(" the pair not Found ")

}
