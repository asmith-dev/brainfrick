/*
	Trying to devise a method to pop a slice.

	Eventually need to implement with an interface parameter,
	and then switch the type and pop it accordingly.

	For now, this will only pop a []int
*/

package main

import "fmt"

func pop(arr *[]int) {
	// Cannot index a pointer to a slice, so need temp
	temp := *arr

	// Handling empty slices
	if len(temp) == 0 {
		panic("Cannot pop an empty array.")
	}

	// Pops the array
	*arr = temp[:len(temp)-1]
}

func main() {
	// Example slice, with printing before and after popping
	example := []int{1, 2, 3, 4}
	fmt.Println(example)
	pop(&example)
	fmt.Println(example)
}
