package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"strings"
	//"reflect"
	//"strconv"
	"flag"
	//"math"
	"time"
	//"sort"
	//"github.com/jonchen727/AdventofGo/helpers"
)

//go:embed input.txt
var input string

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
		ans := part1(input)

		fmt.Println("Part 1 Answer:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Part 2 Answer:", ans)
	}
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func part1(input string) int {
	ans := 0
	treeMap := parseInput(input)
	line := strings.Split(input, "\n")
	for i := 1; i < len(line); i++ {
		if _, ok := treeMap[fmt.Sprintf("%d,%d", i, (i*3)%len(strings.Split(line[0], "")))]; ok {
			ans++
		}
	}
	return ans
}

func part2(input string) int {
	ans := 1
	treeMap := parseInput(input)
	line := strings.Split(input, "\n")
	set := [][]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}
	for _, v := range set {
		trees := 0
		for i := 1; i < len(line)/v[0]; i++ {
			if _, ok := treeMap[fmt.Sprintf("%d,%d", i*v[0], (i*v[1])%len(strings.Split(line[0], "")))]; ok {
				trees++
			}
		}
		ans *= trees
	}
	return ans
}

func parseInput(input string) map[string]bool {
	treeMap := map[string]bool{}
	for i, line := range strings.Split(input, "\n") {
		for j, line := range strings.Split(line, "") {
			if line == "#" {
				treeMap[fmt.Sprintf("%d,%d", i, j)] = true
			}
		}
	}
	return treeMap
}
