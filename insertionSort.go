package main

import "fmt"

func main() {
	inp := []int{12, 11, 13, 5, 6}
	insertionSort(inp)
}
func insertionSort(nums []int) {

	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		j := i - 1

		for j >= 0 && nums[j] > temp {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = temp
	}
	fmt.Println("the sorted array is ", nums)

}
