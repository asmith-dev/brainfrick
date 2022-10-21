/*
	A compiler of BrainF***
	Filename must be entered below (no command line path input).
	This algorithm is not meant to be built and used
	(though it wouldn't be difficult to achieve this with the OS package);
	rather, it is more of a proof of concept.
*/

package main

import (
	"brainfrick/pkg"
	"fmt"
)

// Enter test file name here:
var testFileName = "helloWorld.bf"

func main() {
	// Gets source file
	source := pkg.Read("test/" + testFileName)

	// 30000 is the standard memory size for BrainF***
	memSize := uint16(30000)

	// Initializes memory, pointers, and loop-tracking variables
	mem := make([]byte, memSize)
	memptr := int16(0)
	codeptr := 0
	var bracketPos []int
	scopeLevel := 0

	// Stores the end-of-file location
	eof := len(source)

	// Loops until the file is ended
	for codeptr < eof {
		// Switches current character from the source file
		switch string(source[codeptr]) {
		case "+":
			mem[memptr] += 1
		case "-":
			mem[memptr] -= 1
		case ">":
			// Modding causes overflow to loop back to the beginning of the memory
			memptr = (memptr + 1) % int16(memSize)
		case "<":
			memptr--

			// This statement causes negative pointers to loop to the end of the memory
			if memptr == -1 {
				memptr = int16(memSize - 1)
			}
		case ".":
			fmt.Print(string(mem[memptr]))
		case ",":
			// Users can input whatever, but only the first character is saved in memory
			input := pkg.Input("")
			mem[memptr] = input[0]
		case "[":
			// Appends the bracket's position and increases the scope level
			bracketPos = append(bracketPos, codeptr)
			scopeLevel++

			// If the current memory value is zero, the code pointer skips past that loop
			if mem[memptr] == 0 {
				// Removes the bracket position and stores what the final scope value needs to be
				pkg.Pop(&bracketPos)
				resumeScope := scopeLevel - 1

				// Repeats until the loop is ended or the file ends
				for scopeLevel > resumeScope && codeptr < eof {
					codeptr++

					// Throws an error if the file ends
					if codeptr == eof {
						panic("Unmatched \"[\"")
					}

					// Adjusts the scope level while skipping
					switch string(source[codeptr]) {
					case "[":
						scopeLevel++
					case "]":
						scopeLevel--
					}
				}
			}
		case "]":
			// If the current memory value is zero, the loop ends.
			// Otherwise, the code pointer is reset to the most recent bracket positions
			if mem[memptr] == 0 {
				pkg.Pop(&bracketPos)
				scopeLevel--
			} else {
				codeptr = bracketPos[len(bracketPos)-1]
			}
		}

		// Moves the code pointer to the next character
		codeptr++

		// Throws an error if the file ends without ending all the brackets
		if codeptr == eof && scopeLevel != 0 {
			panic("Unmatched \"[\"")
		}
	}
}
