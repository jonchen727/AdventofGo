package main

import (
	_ "embed"
	"fmt"
	"strings"
	//"strconv"
	"flag"
	//"sort"
	"github.com/jonchen727/AdventofGo/helpers"
)

//go:embed input.txt
var input string
var priorities = map[string]int{}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("input is empty")
	}

	//generate priorities
	for i := 0; i < 26; i++ {
		priorities[helpers.ASCIIIntToChar('a'+i)] = i + 1
		priorities[helpers.ASCIIIntToChar('A'+i)] = i + 27

	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part of the puzzle to run")
	flag.Parse()
	fmt.Println("Part", part)

	if part == 1 {
		answer := part1(input)
		fmt.Println("Answer: ", answer)

	} else {
		ans := part2(input)
		fmt.Println("Answer:", ans)
		//fmt.Println("Answer:", ans)
	}

}

func part1(input string) int {
	sacks := parseInput(input)
	prioritiesSum := 0
	fmt.Println("Answer:", sacks)
	for i, sack := range sacks {
		fmt.Println("Sack", i+1, ": ", sack)
		for _, item := range sack[0] {
			if strings.Contains(sack[1], helpers.ToString(item)) {
				fmt.Println("Found:", helpers.ToString(item), "Value:", priorities[helpers.ToString(item)])
				prioritiesSum += priorities[helpers.ToString(item)]
				break //prevent duplicate findings
			}
		}
	}
	return prioritiesSum

}

func part2(input string) int {
	groups := parseInput2(input)
	fmt.Println(groups)
	prioritiesSum := 0
	for group, elve := range groups {
		for _, item := range elve[0] {
			itemstr := helpers.ToString(item)
			if strings.Contains(elve[1], itemstr) && strings.Contains(elve[2], itemstr) {
				fmt.Println("Group:", group, "Item:", itemstr, "Priority:", priorities[itemstr])
				prioritiesSum += priorities[itemstr]
				break
			}
		}
	}

	return prioritiesSum
}

func parseInput(input string) (ans [][]string) {
	for _, lines := range strings.Split(input, "\n") {
		comp1 := lines[0 : len(lines)/2]
		comp2 := lines[len(lines)/2:]
		ans = append(ans, []string{comp1, comp2})
	}
	return ans
}

func parseInput2(input string) (ans [][]string) {
	var line []string
	var group int
	for i, lines := range strings.Split(input, "\n") {
		group = i / 3
		if group == (i / 3) {
			line = append(line, lines)
			if i%3 == 2 {
				ans = append(ans, line)
				line = nil
			}
		}

	}
	return ans
}
