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
	blocks := parseInput(input)
	for _, block := range blocks {
		r := findMirror(block, false)
		ans += r * 100
	}
	for _, block := range zipSlices(blocks) {
		c := findMirror(block, false)
		ans += c
	}
	return ans
}

func findMirror(block []string, mismatch bool) int {
	for i := 1; i < len(block); i++ {
		top := reverseSlice(block[:i])
		bottom := block[i:]
		mismatches := 0

		switch mismatch {
		case true:
<<<<<<< HEAD
			if len(top) > len(bottom) {
				top = top[:len(bottom)]
			} else {
				bottom = bottom[:len(top)]
=======
			for i := 0; i < len(top) && i < len(bottom); i++ {
				if top[i] != bottom[i] {
					mismatches++
				}
>>>>>>> 4bc8ccc (add day 14 part 1)
			}
			mismatches = countMismatches(top, bottom)

			if mismatches == 1 {
				return i
			}

		case false:
			if len(top) > len(bottom) {
				top = top[:len(bottom)]
			} else {
				bottom = bottom[:len(top)]
			}
			if slicesEqual(top, bottom) {
				return i
			}
		}
	}
	return 0
}

func reverseSlice(s []string) []string {
	r := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	fmt.Println(s,r)
	return r
}

func countMismatches(a, b []string) int {
	mismatches := 0
	if len(a) != len(b) {
		return mismatches
	}
	for i, x := range a {
		for j := range x {
			if a[i][j] != b[i][j] {
				mismatches++
			}
		}
	}
	return mismatches
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func zipSlices(blocks [][]string) [][]string {
	newArr := [][]string{}
	for _, block := range blocks {
		newBlock := []string{}
		for i, line := range block {
			for j, char := range line {
				if i == 0 {
					newBlock = append(newBlock, string(char))
				} else {
					newBlock[j] += string(char)
				}
			}
		}
		newArr = append(newArr, newBlock)
	}
	return newArr
}

func part2(input string) int {
	ans := 0
	blocks := parseInput(input)
	for _, block := range blocks {
		r := findMirror(block, true)
		ans += r * 100
	}
	for _, block := range zipSlices(blocks) {
		c := findMirror(block, true)
		ans += c
	}
	return ans
}

func parseInput(input string) [][]string {
	arr := [][]string{}
	split := strings.Split(input, "\n\n")

	for _, block := range split {
		aBlock := []string{}
		for _, line := range strings.Split(block, "\n") {
			aBlock = append(aBlock, line)
		}
		arr = append(arr, aBlock)
	}
	return arr
}
