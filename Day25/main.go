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
	//"github.com/jonchen727/2022-AdventofCode/helpers"
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

func part1(input string) string {

	arr := parseInput(input)
	total := convertAndAdd(arr)
	ans := convertToSnafu(total)
	return ans
}

func part2(input string) int {
	ans := 0
	return ans
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func convertAndAdd(input []string) int {
	total := 0
	for _, line := range input {
		coef := 1
		for i := len(line) - 1; i >= 0; i-- {
			x := line[i]
			total += (strings.Index("=-012", string(x)) - 2) * coef
			coef *= 5
		}
	}
	return total
}

func convertToSnafu(num int) string {
	var ans string
	for num != 0 {
		remainder := num % 5
		num /= 5
		if remainder <= 2 {
			ans = fmt.Sprintf("%d%s", remainder, ans)
		} else {
			ans = fmt.Sprintf("%c%s", "   =-"[remainder], ans)
			num++
		}
	}
	return ans
}
