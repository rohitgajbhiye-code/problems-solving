package main

import "fmt"

func main() {
	fmt.Println(panagram("some things are meant to be"))
	fmt.Println(panagram("The quick brown fox jumps over the lazy dog"))
	// prints the missing character to consider string as panagram
	panagramMissing("some things are meant to be")
	panagramMissing("The quick brown fox jumps over the lazy dog")
}

/*
 * Efficient solution would be to calculate sum of 97 to 122
 * points -
 * 1. [25]byte would be 0-24
 */
func panagram(source string) bool {
	// Array declaration when you know the size initially
	var count [26]byte = [26]byte{}
	//var index int = 0
	for _, c := range source {
		if c >= 97 && c <= 122 {
			index := c - 97
			count[int(index)] = 1
		}
	}
	for i := 0; i < len(count); i++ {
		if count[i] == 0 {
			return false
		}
	}
	return true
}

func panagramMissing(source string) {
	// Array declaration when you know the size initially
	var count [26]byte = [26]byte{}
	//var index int = 0
	for _, c := range source {
		if c >= 97 && c <= 122 {
			index := c - 97
			count[int(index)] = 1
		}
	}

	for i := 0; i < len(count); i++ {
		if count[i] == 0 {
			fmt.Print(" " + string(i+97))
		}
	}
}
