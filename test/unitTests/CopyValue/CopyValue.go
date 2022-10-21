/*
	Checking to see if variable copying is stupid, like it is in Python
*/

package main

import "fmt"

func main() {
	// Declaring a variable and then copying it
	var1 := 10
	var2 := var1

	// Dissimilar changes to each
	var1++
	var2--

	// Print results
	fmt.Println(var1)
	fmt.Println(var2)

	// The variables should be different, and they are.
	// Conclusion: Go is not stupid like Python
}
