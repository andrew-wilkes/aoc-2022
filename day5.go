package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	data := string(bytes)
	lines := strings.Split(data, "\n")
	stacks := [][]rune{}

	// Parse data
	stacking := true
	for _, line := range lines {
		if stacking {
			for i, ch := range line {
				switch {
				case ch >= 'A' && ch <= 'Z':
					col := i / 4
					stacks_to_add := col - len(stacks) + 1
					if stacks_to_add > 0 {
						for range stacks_to_add {
							stacks = append(stacks, []rune{})
						}
					}
					stacks[col] = append(stacks[col], ch)
				case ch == '1':
					stacking = false
				}
			}
		} else {
			tokens := strings.Split(line, " ")
			if len(tokens) == 6 {
				n, _ := strconv.Atoi(tokens[1])
				from, _ := strconv.Atoi(tokens[3])
				to, _ := strconv.Atoi(tokens[5])
				from--
				to--
				top := make([]rune, n)
				copy(top, stacks[from][:n])
				//slices.Reverse(top) // For part 1
				stacks[to] = append(top, stacks[to]...)
				stacks[from] = stacks[from][n:]
			}
		}
	}
	answer := []string{}
	for _, stack := range stacks {
		answer = append(answer, string(stack[0]))
	}
	astr := strings.Join(answer, "")
	fmt.Printf("The answer = %s\n", astr)
}
