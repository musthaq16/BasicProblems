// to count vowels in a string

package main

import "fmt"

func main() {
	inp := "abc de"

	var count int64

	for _, val := range inp {
		if string(val) == "a" || string(val) == "e" || string(val) == "i" || string(val) == "o" || string(val) == "u" {
			count++
		}

	}

	fmt.Println("the vowels count is ", count)

}
