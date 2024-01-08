package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"strings"
	//"reflect"
	"flag"
	"strconv"
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
	nums := parseInput(input)
	for _, num := range nums {
		ans += num
	}
	return ans
}

func part2(input string) int {
	ans := 0
	nums := parseInput2(input)
	for idx, num := range nums {
		if idx == 4 {
			fmt.Println(num)
		}
		ans += num
	}
	return ans
}

func replaceDigits(input string) string {
	digits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	newline := input

	low := 9999
	high := -1
	var low_value, high_value string
	for key, _ := range digits {
		low_idx := strings.Index(newline, key)
		high_idx := strings.LastIndex(newline, key)

		if low_idx != -1 {
			if low_idx < low {
				low = low_idx
				low_value = key
			}
		}
		if high_idx != -1 {
			if high_idx > high {
				high = high_idx
				high_value = key
			}
		}
	}
	if low == 9999 && high == -1 {
		return newline
	}
	if low < 9999 && high > -1 {
		tmp := newline[:low] + digits[low_value] + newline[low:high] + digits[high_value] + newline[high:]
		newline = tmp
	} else if low < 9999 {
		tmp := newline[:low] + digits[low_value] + newline[low:]
		newline = tmp
	} else if high > -1 {
		tmp := newline[:high] + digits[high_value] + newline[high:]
		newline = tmp
	}

	return newline
}

func parseInput2(input string) []int {

	lines := strings.Split(input, "\n")
	nums := []int{}
	for _, line := range lines {
		var sNum string
		newline := replaceDigits(line)
		for _, char := range newline {
			if _, err := strconv.Atoi(string(char)); err == nil {
				sNum = sNum + string(char)
			}
		}
		num := helpers.ToInt(string(sNum[0]) + string(sNum[len(sNum)-1]))
		nums = append(nums, num)
	}
	return nums
}

func parseInput(input string) []int {
	lines := strings.Split(input, "\n")
	nums := []int{}
	for _, line := range lines {
		var sNum string
		for _, char := range line {
			if _, err := strconv.Atoi(string(char)); err == nil {
				sNum = sNum + string(char)
			}
		}
		num := helpers.ToInt(string(sNum[0]) + string(sNum[len(sNum)-1]))
		nums = append(nums, num)
	}
	return nums
}
