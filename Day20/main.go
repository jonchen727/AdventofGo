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
	nodes, zeronode := parseInput(input)

	// find the node with value 0

	// move nodes
	for i := 0; i < len(nodes); i++ {
		nodes[i].move(len(nodes))
	}

	// sum values of nodes clockwise from 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			zeronode = zeronode.right
		}
		ans += zeronode.value
	}
	return ans
}

func (node *Node) move(totalLen int) {
	steps := node.value

	// translate steps into absolute steps
	steps %= totalLen - 1

	//if steps = 0 then do nothing
	if steps == 0 {
		return
	}
	//translate backwards into forward steps
	if steps < 0 {
		steps += (totalLen - 1)
	}

	// remove node from neighbors and update references
	// grab current nodes neighbors
	oldleft, oldright := node.left, node.right
	// left neighbor now points to nodes previous right neighbor
	oldleft.right = oldright
	// right neighbor now points to nodes previous left neighbor
	oldright.left = oldleft

	// save starting node
	i := node
	// move right by steps
	for steps > 0 {
		i = i.right
		if i == node {
			panic("repeat")
		}
		steps--
	}

	// save values of where to insert node
	nextleft, nextright := i, i.right
	// update left neighbor
	nextleft.right = node
	// update self to point to left neighbor
	node.left = nextleft
	// update right neighbor
	nextright.left = node
	// update self to point to right neighbor
	node.right = nextright
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

func parseInput(input string) ([]Node, *Node) {

	zeroNode := &Node{}
	nodes := []Node{}
	for _, num := range strings.Split(input, "\n") {
		nodes = append(nodes, Node{value: helpers.ToInt(num)})
	}
	//fmt.Println(nodes)
	for i, _ := range nodes {
		nodes[i].right = &nodes[helpers.PosMod((i+1), len(nodes))]
		nodes[i].left = &nodes[helpers.PosMod((i-1), len(nodes))]
		if nodes[i].value == 0 {
			zeroNode = &nodes[i]
		}
	}
	//fmt.Println(nodes)

	return nodes, zeroNode
}
