/*
	This script includes a variety of useful functions to condense workflow.
*/

package pkg

import (
	"fmt"
	"os"
)

// HandleError simplifies general error handling.
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Input is a Python-like implementation for getting user input in a condensed format.
func Input(str string) string {
	var response string
	fmt.Print(str)

	_, err := fmt.Scanln(&response)
	HandleError(err)

	return response
}

// Read reads an entire file into a string and returns it
// Note: path is relative to go.mod
func Read(path string) string {
	source, err := os.ReadFile(path)
	HandleError(err)
	return string(source)
}

// Pop removes the last entry of a slice of integers
// Can be reimplemented with different types
func Pop(arr *[]int) {
	// Cannot index a pointer to a slice, so need temp
	temp := *arr

	// Handling empty slices
	if len(temp) == 0 {
		panic("Cannot pop an empty array.")
	}

	// Pops the array
	*arr = temp[:len(temp)-1]
}
