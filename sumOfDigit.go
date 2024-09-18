// Program for Sum of the digits of a given number

package main

import "fmt"

func main() {
	//row and col should be same
	inp := 687
	sum := 0
	for inp != 0 {
		sum += inp % 10
		inp = inp / 10
	}

	fmt.Println("The sum of digit is ", sum)

}