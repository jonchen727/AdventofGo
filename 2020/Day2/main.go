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
	entries := parseInput(input)
	for _, entry := range entries {
		count := strings.Count(entry.password, entry.letter)
		if count >= entry.min && count <= entry.max {
			ans++
		}
	}
	return ans
}

func part2(input string) int {
	ans := 0
	entries := parseInput(input)
	for _, entry := range entries {
		if (string(entry.password[entry.min-1]) == entry.letter) != (string(entry.password[entry.max-1]) == entry.letter) {
			ans++
		}
	}
	return ans
}

type Entry struct {
	min      int
	max      int
	letter   string
	password string
}

func parseInput(input string) []Entry {
	entries := []Entry{}
	for _, line := range strings.Split(input, "\n") {
		entry := Entry{}
		fmt.Sscanf(line, "%d-%d %1s: %s", &entry.min, &entry.max, &entry.letter, &entry.password)
		entries = append(entries, entry)
	}
	return entries
}
