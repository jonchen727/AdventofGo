package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	//"reflect"
	//"strconv"
	"flag"
	//"sort"
	//"github.com/jonchen727/AdventofGo/helpers"
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

	fmt.Println("Part 1 Answer:", answer)
	} else {
		ans := part2(input)
		fmt.Println("Part 2 Answer:", ans)
		//fmt.Println("Answer:", ans)
	}

}

func part1(input string) string {
	blocks, steps := parseInput(input)
	var ans string
	for _, step := range steps {
		move := blocks[step.from][len(blocks[step.from])-step.amt:]
    fmt.Println(move)
		slices.Reverse(move)
		blocks[step.from] = blocks[step.from][:len(blocks[step.from])-step.amt]
		blocks[step.to] = append(blocks[step.to], move...)
	}

	fmt.Println("Blocks:", blocks)
	for _, block := range blocks {
		ans += block[len(block)-1]
	}

	return ans
}

func part2(input string) string {
	blocks, steps := parseInput(input)
	var ans string
	for _, step := range steps {
		move := blocks[step.from][len(blocks[step.from])-step.amt:]
    fmt.Println(move)
		blocks[step.from] = blocks[step.from][:len(blocks[step.from])-step.amt]
		blocks[step.to] = append(blocks[step.to], move...)
	}

	fmt.Println("Blocks:", blocks)
	for _, block := range blocks {
		ans += block[len(block)-1]
	}

	return ans

}

type command struct {
	amt  int
	from int
	to   int
}

func parseInput(input string) ([][]string, []command) {
	split := strings.Split(input, "\n\n")
	blocks := split[0]
	rawcommands := split[1]

	fmt.Println(blocks)
	fmt.Println(rawcommands)

	blocks = strings.ReplaceAll(blocks, "    ", "[0]")
	blocks = strings.ReplaceAll(blocks, " ", "")
	blocks = strings.ReplaceAll(blocks, "][", "")
	blocks = strings.ReplaceAll(blocks, "[", "")
	blocks = strings.ReplaceAll(blocks, "]", "")
	splitblocks := [][]string{}
	sortedblocks := [][]string{}

	for _, row := range strings.Split(blocks, "\n") {
		splitblocks = append(splitblocks, strings.Split(row, ""))
	}

	fmt.Println(splitblocks)

	tcol, trow := len(splitblocks[0]), len(splitblocks)

	for c := 0; c < tcol; c++ {
		stack := []string{}
		for r := trow - 2; r >= 0; r-- {
			fmt.Println(splitblocks[r][c])
			char := splitblocks[r][c]
			if char != "0" {
				stack = append(stack, char)
			}

		}
		sortedblocks = append(sortedblocks, stack)
	}
	fmt.Println(sortedblocks)

	commands := []command{}

	for _, row := range strings.Split(rawcommands, "\n") {
		org := command{}
		_, err := fmt.Sscanf(row, "move %d from %d to %d", &org.amt, &org.from, &org.to)
		if err != nil {
			panic(err)
		}
		//shift from and to to be 0 indexed
		org.from--
		org.to--
		commands = append(commands, org)
	}
	fmt.Println(commands)

	return sortedblocks, commands
}
