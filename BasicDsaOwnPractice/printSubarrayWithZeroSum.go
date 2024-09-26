// Print all subarrays with 0 sum

package main

import "fmt"

func main() {
	arr := []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2}
	printAllSubArray(arr)
}

//bruteforce
// func printAllSubArray(nums []int) {

// 	for i := 0; i < len(nums); i++ {
// 		sum := 0

// 		for j := i; j < len(nums); j++ {
// 			sum += nums[j]

// 			if sum == 0 {
// 				fmt.Printf(" the index sub array from %d , %d \n", i, j)
// 			}
// 		}
// 	}

// }

func printAllSubArray(nums []int) {

	hashMap := make(map[int][]int)

	hashMap[0] = []int{-1}
	sum := 0

	for i, num := range nums {
		sum += num

		if val, found := hashMap[sum]; found {
			for _, start := range val {
				fmt.Println(start+1, " ", i)
			}
		}

		hahashMap[sum] = append(hahashMap[sum], i)

	}

}
