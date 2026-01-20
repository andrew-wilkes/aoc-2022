package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type order int

const (
	UNDECIDED = iota
	INORDER
	OUTOFORDER
)

type packetPairStr struct {
	left, right string
}

type listItem struct {
	value int
	list  list
	empty bool
}

type list []listItem

func main() {
	file, _ := os.Open("example.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lineCount := 0
	packetPairStrs := []packetPairStr{}
	var pp packetPairStr
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		switch lineCount {
		case 1:
			pp = packetPairStr{left: line}
		case 2:
			pp.right = line
			packetPairStrs = append(packetPairStrs, pp)
		case 3:
			lineCount = 0
		}
	}

	sumOfIndicesInRightOrder := 0
	for i, pps := range packetPairStrs {
		p1 := parsePacketStr(pps.left)
		p2 := parsePacketStr(pps.right)
		if isInOrder(p1, p2) == INORDER {
			sumOfIndicesInRightOrder += i + 1
		}
	}
	fmt.Printf("Part 1 answer = %d\n", sumOfIndicesInRightOrder)
}

func isInOrder(a, b listItem) order {
	if a.empty && len(a.list) != 0 {
		fmt.Println(a)
		panic(a)
	}
	var l1, l2 list
	switch {
	case len(a.list) > 0:
		if b.empty {
			return OUTOFORDER
		}
		l1 = a.list
		if len(b.list) > 0 {
			l2 = b.list
		} else {
			l2 = list{b}
		}
	case len(b.list) > 0:
		if a.empty {
			return INORDER
		}
		l1 = list{a}
		l2 = b.list
	default:
		if a.value < b.value {
			return INORDER
		}
		if a.value > b.value {
			return OUTOFORDER
		}
		return UNDECIDED
	}
	// Compare lists
	var i int
	var v listItem
	for i, v = range l1 {
		if i == len(l2) {
			return OUTOFORDER
		}
		result := isInOrder(v, l2[i])
		if result != UNDECIDED {
			return result
		}
	}
	if len(l2) > len(l1) {
		return INORDER
	}
	return UNDECIDED
}

func parsePacketStr(s string) listItem {
	thisItem := listItem{}
	digits := []string{}
	evalDidits := func() {
		d, _ := strconv.Atoi(strings.Join(digits, ""))
		thisItem.list = append(thisItem.list, listItem{value: d})
		digits = nil
	}
	i := 0
	for i < len(s) {
		r := s[i]
		i++
		switch {
		case r == '[':
			// Extract inner text for this bracket level
			level := 0
			for j, r2 := range s[i:] {
				if r2 == '[' {
					level++
				} else if r2 == ']' {
					if level == 0 {
						if j == 0 {
							thisItem.list = append(thisItem.list, listItem{empty: true})
						} else {
							thisItem.list = append(thisItem.list, parsePacketStr(s[i:i+j]))
						}
						i = i + j + 1 // Advance the counter
						break
					} else {
						level--
					}
				}
			}
		case r >= '0' && r <= '9':
			digits = append(digits, string(r))

		case r == ',':
			if len(digits) > 0 {
				evalDidits()
			}
		}
	}
	if len(digits) > 0 {
		evalDidits()
	}
	return thisItem
}
