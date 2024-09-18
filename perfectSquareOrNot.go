// Check if given number is perfect square

package main

import (
	"fmt"
	"math"
)

func main() {
	var inp float64 = 8

	val := math.Sqrt(inp)
	pow := math.Pow(val, 2)
	if inp == pow {
		fmt.Println("Its Perfect Square")
	} else {
		fmt.Println("Its not perfect Square...")
	}

	//fmt.Println("the vowels count is ", count)

}
