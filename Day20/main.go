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
	//ans := 0
	_, nodes := parseInput(input)
	m := len(nodes) - 1

	var z *Node
	for i, kd := range nodes {
		k := &nodes[i]
		if nodes[i].value == 0 {
			z = &nodes[i]
			continue
		}
		p := &kd
		if k.value > 0 {
			for j := 0; j < helpers.PosMod(k.value, m); j++ {
				p = p.right
			}
			if k == p {
				continue
			}
			k.right.left = k.left
			k.left.right = k.right
			p.right.left = k
			k.right = p.right
			p.right = k
			k.left = p
		} else {
			for j := 0; j < helpers.PosMod(k.value, m); j++ {
				p = p.left
			}
			if k == p {
				continue
			}
			k.right.left = k.left
			k.left.right = k.right
			p.left.right = k
			k.left = p.left
			p.left = k
			k.right = p

		}
	}
	t := 0
	for i := 0; i < len(nodes); i++ {
		fmt.Println(nodes[i].right.value)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			z = z.right

		}
		t += z.value
	}

	return t
}

func part2(input string) int {
	ans := 0
	return ans
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func parseInput(input string) ([]int, []Node) {
	numlist := []int{}
	nodes := []Node{}
	for _, num := range strings.Split(input, "\n") {
		numlist = append(numlist, helpers.ToInt(num))
		nodes = append(nodes, Node{value: helpers.ToInt(num)})
	}
	fmt.Println(nodes)
	for i, _ := range nodes {
		nodes[i].right = &nodes[helpers.PosMod((i+1), len(nodes))]
		nodes[i].left = &nodes[helpers.PosMod(i-1, len(nodes))]

	}
	fmt.Println(nodes)

	return numlist, nodes
}
