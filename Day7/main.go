package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"strings"
	//"reflect"
	//"strconv"
	"flag"
	"time"
	//"sort"
	//"github.com/jonchen727/2022-AdventofCode/helpers"
)

//go:embed input.txt
var input string
var priorities = map[string]int{}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("input is empty")
	}

}

func main() {
	start := time.Now()
	var part int
	flag.IntVar(&part, "part", 1, "part of the puzzle to run")
	flag.Parse()

	if part == 1 {
	answer := part1(input)

	fmt.Println("Part 1 Answer:", answer)
	} else {
		ans := part2(input)
		fmt.Println("Part 2 Answer:", ans)
		//fmt.Println("Answer:", ans)
	}
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func part1(input string) int {
	root := parseInput(input)
	ans := find100000(root)


	// for i, line := range commands {
	// 	fmt.Println(i, line)
	// }
	return ans
}

func part2(input string) int{
	return 0
}

type dir struct {
	name string
	parent *dir
	children map[string]*dir
	files map[string]*File
	size int
}

type File struct {
	name string
	size int
}



func parseInput(input string) *dir {
	root := &dir{
		name: "root",
		children: map[string]*dir{},
	}
	currentDir := root
	input = input[2:]
	commands := strings.Split(input, "\n$ ")
	//commands = commands[1:]
	//fmt.Println("Commands:", commands)
	for _, command := range commands {
		//fmt.Println(i, command)
		split := strings.Split(command, "\n")
		//fmt.Println("Length:",len(split))
		//fmt.Println("Current Directory:", currentDir)
		if len(split) == 1 {
			//fmt.Println(split)
			var changeDir string
			_, err := fmt.Sscanf(split[0], "cd %s", &changeDir)
			if err != nil {
				fmt.Println("Error:", err)
			}
			if changeDir == ".." {
				currentDir = currentDir.parent
			} else {
				if _, ok := currentDir.children[changeDir]; !ok {
					currentDir.children[changeDir] = &dir{
						name: changeDir,
						parent: currentDir,
						children: map[string]*dir{},
						files: map[string]*File{},
					}
				}
				currentDir = currentDir.children[changeDir]
			}
		} else {
			//fmt.Println("ls command")
			for _, line := range split {
				if strings.Contains(line, "dir") {
					childDirName := line[4:]
					if _, ok := currentDir.children[childDirName]; !ok {
						currentDir.children[childDirName] = &dir{
							name: childDirName,
							parent: currentDir,
							children: map[string]*dir{},
							files: map[string]*File{},
						}
					}
					
				}else if !(strings.HasPrefix(line, "dir") || strings.HasPrefix(line, "ls")) {
					file := File{}
					_, err := fmt.Sscanf(line, "%d %s", &file.size, &file.name)
					if err != nil {
						fmt.Println("Error:", err)
					}
					//fmt.Println("File:", file)
					currentDir.files[file.name] = &file
					//currentDir.size += file.size
				} 
			}
		}


	}
	populateFileSizes(root)
	//fmt.Println("Root:", root)
  return root
}

func populateFileSizes(directory *dir) int {
	totalSize := 0
	for _, childDir := range directory.children {
		//fmt.Println("Child:", childDir.name)
		totalSize += populateFileSizes(childDir)
	}
	for _, file := range directory.files {
		//fmt.Println("File:", file.name, "Size:", file.size)
		totalSize += file.size
	}
	directory.size = totalSize
	return totalSize
}

func find100000(directory *dir) int {
	sum := 0

	if directory.size <= 100000 {
		sum += directory.size
	}

	for _, childDir := range directory.children {
		sum += find100000(childDir)
	}
	return sum
}
