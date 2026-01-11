package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Loss = 0
	Draw = 3
	Win  = 6
)

const (
	X = iota + 1
	Y
	Z
)

const (
	A = iota + 1
	B
	C
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	data := string(bytes)
	rounds := strings.Split(data, "\n")
	answer1 := 0
	answer2 := 0
	for _, round := range rounds {
		parts := strings.Split(round, " ")
		answer1 += rpsResult(parts[0], parts[1])
		answer2 += rpsResult2(parts[0], parts[1])
	}
	fmt.Printf("Part 1 answer = %d\n", answer1)
	fmt.Printf("Part 2 answer = %d\n", answer2)
}

func rpsResult(a, b string) int {
	switch a {
	case "A":
		switch b {
		case "X":
			return X + Draw
		case "Y":
			return Y + Win
		case "Z":
			return Z + Loss
		}
	case "B":
		switch b {
		case "X":
			return X + Loss
		case "Y":
			return Y + Draw
		case "Z":
			return Z + Win
		}
	case "C":
		switch b {
		case "X":
			return X + Win
		case "Y":
			return Y + Loss
		case "Z":
			return Z + Draw
		}
	}
	return 0
}

func rpsResult2(a, b string) int {
	switch a {
	case "A":
		switch b {
		case "X":
			return C + Loss
		case "Y":
			return A + Draw
		case "Z":
			return B + Win
		}
	case "B":
		switch b {
		case "X":
			return A + Loss
		case "Y":
			return B + Draw
		case "Z":
			return C + Win
		}
	case "C":
		switch b {
		case "X":
			return B + Loss
		case "Y":
			return C + Draw
		case "Z":
			return A + Win
		}
	}
	return 0
}
