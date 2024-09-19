// Program for factorial of a number

package main

import (
	"fmt"
)

func main() {
	inp := 4

	fact := 1
	for i := 1; i <= inp; i++ {
		fact = fact * i
	}

	fmt.Println("The factorial is ", fact)

}