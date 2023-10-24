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
	"github.com/jonchen727/AdventofGo/helpers"
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
	stuff := parseInput(input)
	for i := 0; i < len(stuff); i++ {
		for j := i + 1; j < len(stuff); j++ {
			if stuff[i]+stuff[j] == 2020 {
				ans = stuff[i] * stuff[j]
				break
			}
		}
	}

	return ans
}

func part2(input string) int {
	ans := 0
	stuff := parseInput(input)
	for i := 0; i < len(stuff); i++ {
		for j := i + 1; j < len(stuff); j++ {
			for k := j + 1; k < len(stuff); k++ {
				if stuff[i]+stuff[j]+stuff[k] == 2020 {
					ans = stuff[i] * stuff[j] * stuff[k]
					break
				}
			}
		}
	}
	return ans
}

func parseInput(input string) []int {
	stuff := []int{}
	for _, line := range strings.Split(input, "\n") {
		stuff = append(stuff, helpers.ToInt(line))
	}
	return stuff
}
