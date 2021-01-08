package main

import "fmt"

func main() {
	flag := isPalindromeOptimized("ababas")
	if flag {
		fmt.Println("word is palindrome")
		return
	}
	fmt.Println("word not is palindrome")
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

/*
* Optimized way to check palindrome
 */
func isPalindromeOptimized(word string) bool {
	endIndex := len(word) - 1
	startIndex := 0
	for startIndex <= endIndex {
		if word[startIndex] != word[endIndex] {
			return false
		}
		startIndex++
		endIndex--
	}
	return true

}
