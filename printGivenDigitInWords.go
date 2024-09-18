// print the given digit in words
package main

import (
	"fmt"
)

func main() {
	inp := "1234"
	var val string

	for i, v := range inp {
		fmt.Println("iii", inp[i])
		switch v {
		case 49:
			val += "one "
		case 50:
			val += "two "
		case 51:
			val += "three "
		case 52:
			val += "four "
		case 53:
			val += "five "
		case 54:
			val += "six "
		case 55:
			val += "seven "
		case 56:
			val += "eight "
		case 57:
			val += "nine "
		case 48:
			val += "zero "
		default:
			fmt.Println("Please enter a valid positive number")
		}
	}
	fmt.Println("the numbers in alphabet is ", val)

}
