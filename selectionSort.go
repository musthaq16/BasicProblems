package main

import "fmt"

func main() {
	inp := []int{64, 25, 12, 22, 11}

	selectionSort(inp)
}

func selectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min_idx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min_idx] {
				min_idx = j
			}
		}
		nums[i], nums[min_idx] = nums[min_idx], nums[i]
	}
	fmt.Println("the sorted array is", nums)
}
