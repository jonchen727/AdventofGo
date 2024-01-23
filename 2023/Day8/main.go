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
	inst, node := parseInput(input)
	start := "AAA"
	for true {
		if start == "ZZZ" {
			break
		}
		
		switch inst[ans%len(inst)] {
		case "R":
			start = node[start][1]
		case "L":
			start = node[start][0]
		}
		ans++
	}
	return ans
}

func part2(input string) int {
  ans := 0
	return ans
}

func parseInput(input string)([]string, map[string][]string){
	split := strings.Split(input, "\n\n")
	inst := []string{}

	for _, val := range strings.Split(split[0],""){
		inst = append(inst, string(val))
	}

	node := make(map[string][]string)
	fmt.Println(split[1])
	for _, val := range strings.Split(split[1], "\n") {
		split2 := strings.Split(val, " = ")
		parts := strings.Split(strings.Trim(split2[1], "()"), ", ")

		node[split2[0]] = []string{parts[0], parts[1]}
	}

	return inst, node
}
