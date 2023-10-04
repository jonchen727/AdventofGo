package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"strings"
	//"reflect"
	//"strconv"
	"flag"
	"time"
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
		//fmt.Println("Answer:", ans)
	}
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func part1(input string) int {
	trees := parseInput(input)
	ans := visable(trees)
	return ans
}

func part2(input string) int {
	return 0
}

func parseInput(input string) [][]int {
	ans := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		row := []int{}
		for _, char := range strings.Split(line, "") {
			row = append(row, helpers.ToInt(char))
		}
		ans = append(ans, row)
	}
	//fmt.Println(ans)
	return ans
}

func visable(input [][]int) int {
	ans := (len(input)-2)*2 + (len(input[0])-2)*2 + 4
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			alreadyvisable := false
			height := input[i][j]
			// check top
			for t := 1; t <= i; t++ {
				//fmt.Println("checking top:", height, input[i-t][j])
				if input[i-t][j] >= height || alreadyvisable {
					//fmt.Println("Too Tall:", height, input[i-t][j])
					break
				}
				if i-t == 0 {
					ans++
					alreadyvisable = true
				}
			}
			// check bottom
			for b := 1; b <= len(input)-i-1; b++ {
				//fmt.Println("checking bottom:", height, input[i+b][j])
				if input[i+b][j] >= height || alreadyvisable {
					//fmt.Println("Too Tall:", height, input[i+b][j])
					break
				}
				if i+b == len(input)-1 {
					ans++
					alreadyvisable = true
				}
			}
			// check left
			for l := 1; l <= j; l++ {
				//fmt.Println("checking left:", height, input[i][j-l])
				if input[i][j-l] >= height || alreadyvisable {
					//fmt.Println("Too Tall Left:", height, input[i][j-l])
					break
				}
				if j-l == 0 {
					ans++
					alreadyvisable = true
				}
			}
			// check right
			for r := 1; r <= len(input[i])-j-1; r++ {
				//fmt.Println("checking right:", height, input[i][j+r])
				if input[i][j+r] >= height || alreadyvisable {
					//fmt.Println("Too Tall:", height, input[i][j+r])
					break
				}
				if j+r == len(input[i])-1 {
					ans++
					alreadyvisable = true
				}
			}
		}
	}

	return ans
}
