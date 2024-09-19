// Program to check if an array is sorted or not


package main

import "fmt"

func main() {

	inp := []int{20, 20, 45, 89, 89, 80}

	for i := 0; i < len(inp)-1; i++ {
		if inp[i] > inp[i+1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")

}