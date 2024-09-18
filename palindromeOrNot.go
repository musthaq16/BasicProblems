// Check if a number is Palindrome

package main

import "fmt"

func main() {
	inp := 12321

	reverseNum := 0
	temp := inp

	for temp != 0 {
		reverseNum = reverseNum*10 + (temp % 10)
		temp = temp / 10
	}

	if reverseNum == inp {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	fmt.Println()

}
