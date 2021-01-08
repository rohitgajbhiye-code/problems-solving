package main

import "fmt"

func main() {
	isPalindrome("mom")
}

/*
* to check palindome we are reversing the string first and then comparing to provided string
 */
func isPalindrome(word string) {
	reverse := []byte{}
	for i := len(word) - 1; i >= 0; i-- {
		reverse = append(reverse, word[i])
	}
	if string(reverse) == word {
		fmt.Println("word is palindrome")
		return
	}
	fmt.Println("word not is palindrome")
}
