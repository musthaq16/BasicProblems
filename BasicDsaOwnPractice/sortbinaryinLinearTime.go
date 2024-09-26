// Sort binary array in linear time

package main

import "fmt"

func main() {
	inp := []int{1, 0, 1, 0, 1, 0, 0, 1}

	sortBinaryArray(inp)
}

func sortBinaryArray(nums []int) {

	count := 0
	res := []int{}
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		}
	}
	for j = 0; j < count; j++ {
		res = append(res, 0)
	}
	for ; j < len(nums); j++ {
		res = append(res, 1)
	}
	fmt.Println(res)
}
