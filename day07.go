package main

import (
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type folder struct {
	name   string
	subs   []*folder
	files  []file
	parent *folder
	size   int
}

func calcDirectorySize(lines []string) int {
	var size int = 0
	root := readFileSystem(lines)

	sumFolders(root, &size)

	return size
}

func findDeleteFolder(lines []string) int {
	root := readFileSystem(lines)

	spaceNeeded := 30000000 - (70000000 - root.size)
	smallestFolder := root
	findSmallestFolder(root, spaceNeeded, &smallestFolder)

	return smallestFolder.size
}

func findSmallestFolder(current *folder, spaceNeeded int, smallestFolder **folder) {
	if current.size > spaceNeeded && current.size < (*(smallestFolder)).size {
		(*smallestFolder) = current
	}
	for _, f := range current.subs {
		findSmallestFolder(f, spaceNeeded, smallestFolder)
	}
}

func sumFolders(current *folder, sum *int) {
	if current.size <= 100000 {
		*sum += current.size
	}
	for _, f := range current.subs {
		sumFolders(f, sum)
	}
}

func readFileSystem(lines []string) *folder {
	var root *folder = nil
	var current *folder = nil

	for _, v := range lines {
		if strings.Contains(v, "$ cd") {
			parts := strings.Split(v, " ")
			path := parts[2]
			switch path {
			case "..":
				current = current.parent
			default:
				if current == nil {
					new := folder{name: path}
					current = &new
					root = &new
				} else {
					new := folder{name: path, parent: current}
					current.subs = append(current.subs, &new)
					current = &new
				}
			}
		} else {
			parts := strings.Split(v, " ")
			i, err := strconv.Atoi(parts[0])
			if err == nil {
				file := file{name: parts[1], size: i}
				current.files = append(current.files, file)
				publishSize(current, i)
			}
		}
	}
	return root
}

func publishSize(current *folder, i int) {
	current.size += i
	if current.parent != nil {
		publishSize(current.parent, i)
	}
}
