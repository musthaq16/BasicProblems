// Count number of even and odd elements in an array

package main

import "fmt"

func main() {
	var eve, odd int64
	slic := []int{2, 3, 4, 5, 6}
	for _, v := range slic {
		if v%2 == 0 {
			eve++
		} else {
			odd++
		}
	}
	fmt.Printf("Odd is %d and Even is %d", odd, eve)

}
