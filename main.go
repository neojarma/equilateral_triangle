package main

import "fmt"

func main() {
	drawEquilateralTriangle(10)
}

func drawEquilateralTriangle(stack int) {
	// set minimum input stack to 5
	if stack < 5 {
		stack = 5
	}

	// outer loop, equal to the max line of the input
	for i := 0; i < stack; i++ {
		// looping to skip printing character "1"
		for j := 0; j < stack-i-1; j++ {
			fmt.Print(" ")
		}

		// looping for print character "1"
		for k := 0; k < (i*2)+1; k++ {
			fmt.Print("1")
		}

		// start a new line for the next loop
		fmt.Println()
	}
}
