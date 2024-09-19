// Check if a word is present in a sentence

package main

import (
	"fmt"
	"strings"
)

func main() {

	inpStr := "Geeks for Geeks"
	wordStr := "Geeks"
	// wordStr := "eeks"

	inpSlice := strings.Fields(inpStr)
	for _, v := range inpSlice {
		if wordStr == v {
			fmt.Println("Word is present in the sentence")
			return
		}
	}
	fmt.Println("Word is not present in the sentence ")

}
