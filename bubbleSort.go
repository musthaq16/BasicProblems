package main

import "fmt"

func main() {
	inp := []int{11, 14, 3, 8, 18, 17, 1}
	// inp := []int{1, 2, 3, 4, 5, 6}

	bubbleSortArray(inp)
}

func bubbleSortArray(nums []int) {
	sorted := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				sorted++
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		if sorted == 0 {
			fmt.Println("the array is already sorted", nums)
			return
		}
	}
	fmt.Println("the sorted array is ", nums)
}
