package main

import (
	_ "embed"
	"fmt"
	"strings"
	"strconv"
	"flag"
	"sort"
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
	var part int
	flag.IntVar(&part, "part", 1, "part of the puzzle to run")
	flag.Parse()
	fmt.Println("Part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Answer:", ans)
		//ans := part1(input)
		//fmt.Println("Answer:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Answer:", ans)
		//fmt.Println("Answer:", ans)
	}

}

func part1(input string) int {
	elves := parseInput(input)
	fmt.Println("Elves:", elves)
	totals := []int{}
		for _, row := range elves {
			totals = append(totals, helpers.SumIntSlice(row))
		}
	fmt.Println("Totals:", totals)
	return helpers.MaxInt(totals...)
}

func part2(input string) int {
	elves := parseInput(input)
	fmt.Println("Elves:", elves)
	totals := []int{}
		for _, row := range elves {
			totals = append(totals, helpers.SumIntSlice(row))
		}
		sort.Ints(totals)
		topThree := 0
		for i :=0; i < 3; i++ {
			topThree += totals[len(totals)-1-i]
		}
		return topThree
	}


func parseInput(input string) (ans [][]int) {
	for _, lines := range strings.Split(input, "\n\n") {
		row := []int{}
		for _, line := range strings.Split(lines, "\n") {
			  line, _ := strconv.Atoi(line)
			  row = append(row, line)
		}
	ans = append(ans, row)
	}
	return ans
}
