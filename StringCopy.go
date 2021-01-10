package main

import "fmt"

func main() {
	var source string = "source string"
	var destination string = "destination string"
	// destination = copy(source, destination)
	destination = copyRecusively(source, destination)
	fmt.Println(destination)
}

/*
* Below function copies source string to destination
* Points to Note -
*
* 1. Strings in go are just a pointer to a (read-only) slice and a length. Thus,
*    when you pass them to a function the pointers are passed as value instead of the whole string.
*
* 2. Constants cannot be evaluated at run time example - const length int = len(source)
 */
func copy(source string, destination string) string {
	var length int = len(source)
	dest := make([]rune, length)
	for i, c := range source {
		dest[i] = c
	}
	destination = string(dest)
	return destination
}
