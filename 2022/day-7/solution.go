package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	totalDiskSpace = 70000000
	required       = 30000000
)

type File struct {
	name  string
	size  int
	dir   bool
	files map[string]*File
}

func (f *File) String() string {
	return fmt.Sprintf("File{name: %v, size: %v, dir: %t, \n\tfiles: %s}", f.name, f.size, f.dir, f.files)
}

func (f *File) AddFile(name string, size int) {
	f.files[name] = &File{name: name, size: size, files: make(map[string]*File)}
	f.size += size
}

func (f *File) AddDir(name string) {
	f.files[name] = &File{name: name, dir: true, files: make(map[string]*File)}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	terminalOutput := bufio.NewScanner(file)

	disk := &File{name: "disk", files: make(map[string]*File)}
	disk.files["/"] = &File{name: "/", dir: true, files: make(map[string]*File)}
	buildDisk(disk, terminalOutput)

	// part 1
	fmt.Printf("Ans part 1: %d\n", sumDirsWithSize(disk.files, 100000))

	// part 2
	root := disk.files["/"]
	remainingSpace := totalDiskSpace - disk.size
	spaceNeeded := required - remainingSpace
	remove := smallestDirToRemove(root.files, spaceNeeded, nil)
	fmt.Printf("Ans part 2: %s size: %d\n", remove.name, remove.size)
}

func smallestDirToRemove(files map[string]*File, limit int, smallest *File) *File {
	for _, f := range files {
		if f.dir && f.size >= limit {
			smallest = smallestDirToRemove(f.files, limit, smallest)
			if smallest == nil || f.size < smallest.size {
				smallest = f
			}
		}
	}
	return smallest
}

func sumDirsWithSize(files map[string]*File, limit int) int {
	totalSize := 0
	for _, f := range files {
		if f.dir {
			totalSize += sumDirsWithSize(f.files, limit)
			if f.size <= limit {
				totalSize += f.size
			}
		}
	}
	return totalSize
}

func buildDisk(f *File, scanner *bufio.Scanner) int {
	for scanner.Scan() {
		totalSize := f.size
		cmd := strings.Split(scanner.Text(), " ")
		if "$" == cmd[0] {
			if "cd" == cmd[1] {
				dir := cmd[2]
				if ".." == dir {
					return totalSize
				} else {
					f.size += buildDisk(f.files[dir], scanner)
				}
			}
		} else {
			//listing of files
			name := cmd[1]
			if cmd[0] == "dir" {
				f.AddDir(name)
			} else {
				s := size(cmd[0])
				f.AddFile(name, s)
				totalSize += s
			}
		}
	}
	return f.size
}

func size(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		os.Exit(3)
	}
	return i
}
