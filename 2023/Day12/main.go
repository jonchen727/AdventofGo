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
	rSprings := parseInput(input)
	for _, r := range rSprings {
		ans += r.findArrangements()
		//fmt.Println(r.springs, r.group, r.arrangements)
	}
	return ans
}

func part2(input string) int {
	ans := 0
	return ans
}

func (r *rSpring) findArrangements() int {
	r.arrangements = count(r.springs, r.group)
	return r.arrangements
}

func count(s string, g []int) int {
	if s == "" {
		if len(g) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if len(g) == 0 {
		if strings.Contains(s, "#") {
			return 0
		} else {
			return 1
		}
	}
	total := 0

	if s[0] == '.' || s[0] == '?' {
		total += count(s[1:], g)
	}
	if s[0] == '#' || s[0] == '?' {
		if (g[0] <= len(s)) && (!strings.Contains(s[:g[0]], ".")) && (g[0] == len(s) || s[g[0]] != '#') {
			if g[0] == len(s) {
				total += count("", g[1:])
			} else {
				total += count(s[g[0]+1:], g[1:])
			}
		}
	}
	return total
}

type rSpring struct {
	springs      string
	group        []int
	arrangements int
}

func parseInput(input string) []rSpring {
	rSprings := []rSpring{}
	split := strings.Split(input, "\n")
	for _, s := range split {
		r := rSpring{}
		split2 := strings.Split(s, " ")
		r.springs = split2[0]
		r.group = []int{}
		for _, num := range strings.Split(split2[1], ",") {
			r.group = append(r.group, helpers.ToInt(num))
		}
		rSprings = append(rSprings, r)
	}
	return rSprings
}
