// Difference between sum of the squares of first n natural numbers and square of sum

package main

import "fmt"

func main() {
	inp := 3

	var sumOfNumbers, sumOfNumbersSquare int

	for i := 1; i <= inp; i++ {
		sumOfNumbers += i
	}

	squareOfSum := sumOfNumbers * sumOfNumbers

	for i := 1; i <= inp; i++ {
		sumOfNumbersSquare += i * i
	}

	fmt.Println("The Absolute Difference is :", squareOfSum-sumOfNumbersSquare)

}
