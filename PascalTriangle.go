package main

import "fmt"

func main() {
	var loop int = 5
	multiple := 1
	for i := 1; i <= loop; i++ {
		// loop to print space
		for j := 0; j <= (loop-i)-1; j++ {
			fmt.Print(" ")
		}
		// Print multile using space
		printMultiple(multiple)
		multiple = multiple * 11
		// next line
		fmt.Println()
	}
}

/*
* Recursively print single digit in a Number
 */
func printMultiple(multiple int) {
	if multiple > 0 {
		reminder := multiple % 10
		multiple = multiple / 10
		printMultiple(multiple)
		fmt.Print(" ")
		fmt.Print(reminder)
	}
	return
}
