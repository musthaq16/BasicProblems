// Check if a subarray with 0 sum exists or not

package main

import "fmt"

func main() {
	inp := []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2}
	subArrayWithZeroSumExistOrNot(inp)
}

func subArrayWithZeroSumExistOrNot(nums []int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == 0 {
			fmt.Println("Subarray with zero-sum exists")
			return
		}
	}
	fmt.Println("Subarray with zero-sum Not-exists")
}
