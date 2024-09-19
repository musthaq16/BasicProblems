// Segregate 0s and 1s in an array

package main

import (
	"fmt"
)

func main() {

	inp := []int{0, 1, 0, 1, 0, 0, 1, 1, 1, 0}

	var count0, count1 int
	var res []int

	for _, val := range inp {
		if val == 1 {
			count1++
		} else {
			count0++
		}
	}

	for i := 0; i < count0; i++ {
		res = append(res, 0)
	}
	for i := 0; i < count1; i++ {
		res = append(res, 1)
	}
	fmt.Println(res)

}
