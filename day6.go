package main

import (
	"fmt"
	"os"
)

const MARKER_SIZE = 14

func main() {
	bytes, _ := os.ReadFile("input.txt")
	buffer := make([]byte, MARKER_SIZE)
	answer := 0
	for i, b := range bytes {
		idx := i % len(buffer)
		buffer[idx] = b
		found := true
	loop:
		for j := 0; j < len(buffer)-1; j++ {
			for k := j + 1; k < len(buffer); k++ {
				if buffer[j] == buffer[k] {
					found = false
					break loop
				}
			}
		}
		if found && i > len(buffer)-1 {
			answer = i + 1
			break
		}
	}
	fmt.Printf("The answer = %d\n", answer)
}
