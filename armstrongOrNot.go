package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	inp := 1634

	length := len(strconv.Itoa(inp))
	val, _ := strconv.Atoi("9")
	fmt.Println("aooo", val)
	var sum float64
	temp := inp
	for temp != 0 {
		val := temp % 10
		pow := math.Pow(float64(val), float64(length))
		sum += pow
		temp = temp / 10
	}

	if sum == float64(inp) {
		fmt.Println("This is armstrong Number..")
	} else {
		fmt.Println("Its not armstrong..")
	}

}
