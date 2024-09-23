// Check if a subarray with 0 sum exists or not

package main

import "fmt"

func main() {
	inp := []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2}
	subArrayWithZeroSumExistOrNot(inp)
}

func subArrayWithZeroSumExistOrNot(nums []int) {

	hashmap := make(map[int]struct{})
	sum := 0

	for _, num := range nums {
		sum += num

		if _, found := hashmap[sum]; found {
			fmt.Println("Subarray with zero-sum exists")
			return
		}

		hashmap[sum] = struct{}{}
	}

	fmt.Println("Subarray with zero-sum Not-exists")
}
