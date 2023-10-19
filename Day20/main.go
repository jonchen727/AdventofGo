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
	"github.com/jonchen727/2022-AdventofCode/helpers"
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
	parseInput(input)
	return ans
}

func part2(input string) int {
  ans := 0
	return ans
}

type Node struct {
	value int
	left *Node
	right *Node

}

func parseInput(input string){
	numlist := []int{}
  nodes := []Node{}
	for _, num := range strings.Split(input, "\n") {
		numlist = append(numlist, helpers.ToInt(num))
		nodes = append(nodes, Node{value: helpers.ToInt(num)})
	}
	fmt.Println(nodes)
	for i, node := range nodes {
		node.right = &nodes[(i+1)%len(nodes)]
		node.left = &nodes[helpers.posMod(i-1,len(nodes))]

	}
	fmt.Println(nodes)

}
