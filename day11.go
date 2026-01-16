package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type opType int

const (
	oldPlus opType = iota
	oldTimes
	oldSquared
)

type monkey struct {
	items                                             []int
	operation                                         opType
	number, divisibleBy, trueTo, falseTo, inspections int
}

func main() {
	monkeys := []monkey{}
	var m monkey
	superMod := 1
	step := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		switch step {
		case 0:
			m = monkey{}
		case 1:
			nums := strings.Split(line[18:], ", ")
			for _, num := range nums {
				n, _ := strconv.Atoi(num)
				m.items = append(m.items, n)
			}
		case 2:
			op := strings.Split(line[23:], " ")
			if op[0] == "+" {
				m.operation = oldPlus
				n, _ := strconv.Atoi(op[1])
				m.number = n
			} else {
				if op[1] == "old" {
					m.operation = oldSquared
				} else {
					m.operation = oldTimes
					n, _ := strconv.Atoi(op[1])
					m.number = n
				}
			}
		case 3:
			n, _ := strconv.Atoi(line[21:])
			m.divisibleBy = n
			superMod *= n
		case 4:
			n, _ := strconv.Atoi(line[29:])
			m.trueTo = n
		case 5:
			n, _ := strconv.Atoi(line[30:])
			m.falseTo = n
			monkeys = append(monkeys, m)
			step = -2
		}
		step++
	}

	for range 10000 {
		for i, m := range monkeys {
			monkeys[i].inspections += len(m.items)
			for _, item := range m.items {
				switch m.operation {
				case oldPlus:
					item += m.number
				case oldTimes:
					item *= m.number
				case oldSquared:
					item *= item
				}
				//item /= 3 // Part 1
				item = item % superMod // Part 2 to restrict the range of numbers

				var dest int
				if item%m.divisibleBy == 0 {
					dest = m.trueTo
				} else {
					dest = m.falseTo
				}
				monkeys[dest].items = append(monkeys[dest].items, item)
			}
			monkeys[i].items = nil
		}
	}
	inspections := make([]int, len(monkeys))
	for i, m := range monkeys {
		inspections[i] = m.inspections
	}
	slices.Sort(inspections)
	slices.Reverse(inspections)
	monkeyBusiness := inspections[0] * inspections[1]

	fmt.Printf("Answer = %d\n", monkeyBusiness)
}
