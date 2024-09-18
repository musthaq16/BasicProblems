// Program for average of an array 
package main

import "fmt"

func main() {
	inp := []int{1, 2, 3, 4, 5}
	var sum, totalElements int
	for _, v := range inp {
		sum += v
		totalElements++
	}
	fmt.Printf("the sum is %d , the total element present in array is %d, avg is %f", sum, totalElements, float64(sum/totalElements))

}