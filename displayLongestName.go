// Display the Longest Name

package main

import (
	"fmt"
)

func main() {

	inp := []string{"GeeksforGeeks", "FreeCodeCamp", "StackOverFlow", "MyCodeSchool"}

	var names []string
	maxine := make(map[string]int)
	var maxLength int

	for _, val := range inp {
		length := len(val)
		if maxLength < length {
			maxLength = length
		}
		maxine[val] = length
	}
	for i, val := range maxine {
		if maxLength == val {
			names = append(names, i)
		}
	}
	fmt.Println(names)

}
