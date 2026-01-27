package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	flowRate int
	tunnels  []string // IDs of connected nodes
}

func main() {
	file, _ := os.Open("example.txt")
	defer file.Close()

	nodes := map[string]node{}
	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("Valve ([A-Z]+) has flow rate=([0-9]+); tunnels? leads? to valves? (.+)")
	for scanner.Scan() {
		line := scanner.Text()
		m := r.FindStringSubmatch(line)
		node := node{num(m[2]), strings.Split(m[3], ", ")}
		nodes[m[1]] = node
	}

	// Part 1
	// Evaluate all paths through the graph up to 30 steps

	paths := [][]string{}
	var tracer func(nodeId string, time int, path []string)
	tracer = func(nodeId string, time int, path []string) {
		path = append(path, nodeId)
		time++
		if time > 30 {
			p := make([]string, len(path))
			copy(p, path)
			paths = append(paths, p)
			return
		}
		end := true
	outer:
		for _, id := range nodes[nodeId].tunnels {
			i := len(path) - 2
			for i > 0 {
				if path[i] == id && path[i-1] == nodeId {
					continue outer
				}
				i--
			}
			tracer(id, time, path)
			end = false
		}
		if end {
			p := make([]string, len(path))
			copy(p, path)
			paths = append(paths, p)
		}
	}
	tracer("AA", 0, []string{})

	// Apply valve on/off through each path to find the highest value of pressure release within 30 mins.
	maxPressure := 0
	var valves func(idx int, currentPressure, totalPressure, time int, valveClosed bool, path []string)
	valves = func(idx int, currentPressure, totalPressure, time int, valveClosed bool, path []string) {
		totalPressure += currentPressure
		if time == 30 {
			if totalPressure > maxPressure {
				maxPressure = totalPressure
				return
			}
		} else {
			if valveClosed {
				// Open valve
				flowRate := nodes[path[idx]].flowRate
				if flowRate > 0 {
					valves(idx, flowRate+currentPressure, totalPressure, time+1, false, path)
				}
			}
			idx++
			if idx < len(path) {
				// Move
				valves(idx, currentPressure, totalPressure, time+1, true, path)
			}
		}
	}
	for _, path := range paths {
		valves(0, 0, 0, 1, false, path)
	}
	fmt.Printf("Part 1 answer = %d\n", maxPressure)
}

// A simple utility function to simplify code above.
func num(s string) int {
	n, _ := strconv.Atoi(s) // Assume that there will be no error or else we should get a panic.
	return n
}
