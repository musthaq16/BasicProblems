// Write a program to reverse digits of a number

package main

import (
	"fmt"
)

func main() {

	inp := 876

	temp := inp
	res := 0
	for temp != 0 {
		res = res*10 + (temp % 10)
		temp = temp / 10
	}
	fmt.Println(res)

}