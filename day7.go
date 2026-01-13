package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type fileItem struct {
	name, ext string
	size      int
}

type dir struct {
	name      string
	files     []fileItem
	subDirs   []dir
	parentDir *dir
	size      int
}

func (d *dir) dirIndex(dirName string) int {
	for idx, subDir := range d.subDirs {
		if subDir.name == dirName {
			return idx
		}
	}
	return -1
}

func (d *dir) updateSize(n int) {
	d.size += n
	if d.parentDir != nil {
		d.parentDir.updateSize(n)
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Create file system
	rootDir := dir{name: "/"}
	currentDir := &rootDir
	listing := false
	for scanner.Scan() { // platform agnostic way of splitting lines
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		if tokens[0] == "$" {
			listing = false
			switch tokens[1] {
			case "cd":
				switch tokens[2] {
				case "/":
					currentDir = &rootDir
				case "..":
					currentDir = currentDir.parentDir
				default:
					if idx := currentDir.dirIndex(tokens[2]); idx < 0 {
						panic("Directory doesn't exist")
					} else {
						currentDir = &currentDir.subDirs[idx]
					}
				}
			case "ls":
				listing = true
			}
		} else if listing {
			if tokens[0] == "dir" {
				newDir := dir{name: tokens[1], parentDir: currentDir}
				currentDir.subDirs = append(currentDir.subDirs, newDir)
			} else {
				fileSize, _ := strconv.Atoi(tokens[0])
				fparts := strings.Split(tokens[1], ".")
				var ext string
				if len(fparts) == 2 {
					ext = fparts[1]
				}
				newFile := fileItem{name: fparts[0], ext: ext, size: fileSize}
				currentDir.files = append(currentDir.files, newFile)
				currentDir.updateSize(fileSize)
			}
		}
	}

	var getDirSizes func(d dir) int
	getDirSizes = func(d dir) int {
		n := 0
		if d.size <= 100000 {
			n = d.size
		}
		for _, subd := range d.subDirs {
			n += getDirSizes(subd)
		}
		return n
	}
	answer := getDirSizes(rootDir)
	fmt.Printf("Part 1 answer = %d\n", answer)

	unusedSpace := 70000000 - rootDir.size
	extraSpaceRequired := 30000000 - unusedSpace
	sizeOfSmallestDir := math.MaxInt

	var findDirToDelete func(d dir)
	findDirToDelete = func(d dir) {
		if d.size >= extraSpaceRequired && d.size < sizeOfSmallestDir {
			sizeOfSmallestDir = d.size
		}
		for _, subd := range d.subDirs {
			findDirToDelete(subd)
		}
	}
	findDirToDelete(rootDir)
	fmt.Printf("Part 2 answer = %d\n", sizeOfSmallestDir)
}
