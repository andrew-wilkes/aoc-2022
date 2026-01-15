package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sigstr := 0
	cycle := 0
	x := 1
	op_cycles := 0
	var instr string
	var operand int
	for scanner.Scan() {
		// Load instruction
		line := scanner.Text()
		data := strings.Split(line, " ")
		instr = data[0]
		if len(data) > 1 {
			operand, _ = strconv.Atoi(data[1])
		}
		switch instr {
		case "noop":
			op_cycles = 1
		case "addx":
			op_cycles = 2
		}
		for range op_cycles {
			cycle++
			if (cycle-20)%40 == 0 {
				ss := x * cycle
				sigstr += ss
			}
			// CRT output
			px := (cycle-1)%40 - x
			if px >= -1 && px <= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if cycle%40 == 0 {
				fmt.Println("")
			}
		}
		if instr == "addx" {
			x += operand
		}
	}
	fmt.Printf("Part 1 answer = %d\n", sigstr)
}
