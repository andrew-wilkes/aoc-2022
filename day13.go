package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type order int

const (
	UNDECIDED = iota
	INORDER
	OUTOFORDER
)

// I love Gos easy way to create custom types. No excuse not to do so.
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
	file, _ := os.Open("input.txt")
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

	packets := []listItem{parsePacketStr("[[2]]"), parsePacketStr("[[6]]")} // Add markers for part 2
	sumOfIndicesInRightOrder := 0
	for i, pps := range packetPairStrs {
		p1 := parsePacketStr(pps.left)
		p2 := parsePacketStr(pps.right)
		if isInOrder(p1.list[0], p2.list[0]) == INORDER {
			sumOfIndicesInRightOrder += i + 1
		}
		packets = append(packets, p1, p2) // Needed for part 2
	}
	fmt.Printf("Part 1 answer = %d\n", sumOfIndicesInRightOrder)

  // Part 2
  
	// I'm leveraging the SortFunc rather than rolling my own sort algorithm code.
	slices.SortFunc(packets, func(a, b listItem) int {
		switch isInOrder(a, b) {
		case INORDER:
			return -1
		case OUTOFORDER:
			return 1
		}
		return 0
	})

	// This is a new thing that I tried recenty: function variables.
	// They allow for inlining functions and making use of the closure (access to outer scoped vars) properties of Go functions.
	findMarker := func(n int) int {
		return slices.IndexFunc(packets, func(l listItem) bool { // Another anonymous function ;-)
			// Kind of ugly, but it is specific to my data structure. Saves doing some crazy recursion.
			if len(l.list) == 1 && len(l.list[0].list) == 1 && len(l.list[0].list[0].list) == 1 {
				if l.list[0].list[0].list[0].value == n {
					return true
				}
			}
			return false
		})
	}
	i1 := findMarker(2) + 1
	i2 := findMarker(6) + 1
	fmt.Printf("Part 2 answer = %v\n", i1*i2)
}

// So many conditions to check!
func isInOrder(a, b listItem) order {
	var l1, l2 list
	switch {
	case a.empty && b.empty:
		return UNDECIDED
	case a.empty:
		return INORDER
	case b.empty:
		return OUTOFORDER
	case len(a.list) > 0:
		l1 = a.list
		if len(b.list) > 0 {
			l2 = b.list
		} else {
			l2 = list{b}
		}
	case len(b.list) > 0:
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

// I wrote this function to check the integrity of my parser
func printList(li listItem) {
	if li.empty {
		fmt.Print("[]")
		return
	}
	if len(li.list) == 0 {
		fmt.Print(li.value)
		return
	}
	fmt.Print("[")
	for i, l := range li.list {
		printList(l)
		if i < len(li.list)-1 {
			fmt.Print(",")
		}
	}
	fmt.Print("]")
}

// Just parsing the data was challenging here!
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
