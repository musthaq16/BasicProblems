package main

import "fmt"

func main() {
	inp := []int{12, 11, 13, 5, 6}
	sort := mergeSort(inp)
	fmt.Println("11111111")
	fmt.Println("ssssss", sort)
}

func mergeSort(nums []int) []int {
 
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2

	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)

}

func merge(left, right []int) []int {

	res := []int{}

	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			res = append(res, right[j])
			j++
		} else {
			res = append(res, left[i])
			i++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return res
}
