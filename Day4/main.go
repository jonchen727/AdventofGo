package main

import (
	_ "embed"
	"fmt"
	"strings"
	//"strconv"
	"flag"
	//"sort"
	"github.com/jonchen727/2022-AdventofCode/helpers"
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
	pairs := parseInput(input)
	//fmt.Println("Pairs:", pairs)
	var overlap int
	for i, pair := range pairs {
		var short string
		var long string
		if len(pair[0]) > len(pair[1]) {
			short = pair[1]
			long = pair[0]
		} else {
			short = pair[0]
			long = pair[1]
		}
		if strings.Contains(long, short) {
			fmt.Println("Overlap",i,":", short)
			fmt.Println("Overlap",i,":", long)
			overlap += 1 
		}
			
	}

	return overlap

}

func part2(input string) int {
	pairs := parseInput(input)
	fmt.Println("Pairs:", pairs)
	var overlap int
	for _, pair := range pairs {
		var short string
		var long string
		if len(pair[0]) > len(pair[1]) {
			short = pair[1]
			long = pair[0]
		} else {
			short = pair[0]
			long = pair[1]
		}

		for _, section := range strings.Split(short, ".")[1:len(strings.Split(short, "."))-1] {
			if strings.Contains(long, "."+section+".") {
				overlap += 1
				break
			}
		}
			
	}

	return overlap
	return 0
}

func parseInput(input string) (ans [][]string) {
	for _, row := range strings.Split(input, "\n") {
		var pair []string
		for _, assignment := range strings.Split(row, ",") {
			sections := "."
			n1 := helpers.ToInt(strings.Split(assignment, "-")[0])
			n2 := helpers.ToInt(strings.Split(assignment, "-")[1])
			for i := n1 ; i <= n2; i++ {
				sections += helpers.ToString(i)
				sections += "."
			}
		  pair = append(pair, sections)
		}
		ans = append(ans, pair)
		
	}
	return ans
}

