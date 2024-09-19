// Program to print multiplication table of a number

package main

import "fmt"

func main() {

	inp := 5

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d", inp, i, inp*i)
		fmt.Println()
	}

}