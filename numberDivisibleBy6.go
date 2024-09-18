// Check if a large number is divisible by 6 or not

package main

import (
	"fmt"
)

func main() {
	//inp := 2112
	inp := 1124

	if inp%6 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
