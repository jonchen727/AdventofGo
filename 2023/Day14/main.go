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
	board, rmax, cmax := parseInput(input)
	//fmt.Println(board, rmax, cmax)

	for r := 0; r < rmax; r++ {
		for c := 0; c < cmax; c++ {
			//fmt.Println(board[fmt.Sprintf("%d,%d", r, c)])
			switch board[fmt.Sprintf("%d,%d", r, c)] {
			case "O":
				if r == 0 {
					break
				}
				for i := r; i >= -1; i-- {
					val, ok := board[fmt.Sprintf("%d,%d", i-1, c)]
					if (val == "#" || val == "O") && ok {
						board[fmt.Sprintf("%d,%d", i, c)] = "O"
						if i != r {
							board[fmt.Sprintf("%d,%d", r, c)] = "."
						}
						//fmt.Println("New Location", i, c)
						break
					}
					if !ok && i != r {
						board[fmt.Sprintf("%d,%d", i, c)] = "O"
						board[fmt.Sprintf("%d,%d", r, c)] = "."
					}

				}
				//fmt.Println("O", r, c)
			case "#":
				//fmt.Println("#")
			case ".":
			}
		}
	}
	for r := 0; r < rmax; r++ {
		for c := 0; c < cmax; c++ {
			if board[fmt.Sprintf("%d,%d", r, c)] == "O" {
				ans += rmax - r
			}
		}
	}
	return ans
}

func part2(input string) int {
	ans := 0
	return ans
}

func parseInput(input string) (map[string]string, int, int) {
	lines := strings.Split(input, "\n")
	board := make(map[string]string)
	rmax := 0
	cmax := 0
	rmax = len(lines)
	for r, line := range lines {
		if cmax == 0 {
			cmax = len(line)
		}
		for c, char := range strings.Split(line, "") {
			board[fmt.Sprintf("%d,%d", r, c)] = char
		}

	}
	return board, rmax, cmax
}
