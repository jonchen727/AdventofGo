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
	nodeMap, start, ins := parseInput(input)
	end, finalDir := FollowPath(nodeMap, start, ins)
	// fmt.Println(start)

	var add int
	switch finalDir {
	case "R":
		add = 0
	case "D":
		add = 1
	case "L":
		add = 2
	case "U":
		add = 3
	}
	ans = (1000 * end.row) + (4 * end.col) + add
	return ans
}

func part2(input string) int {
	ans := 0
	return ans
}

func FollowPath(nodeMap map[string]*Node, start *Node, instructions []string) (*Node, string) {
	// fmt.Println(start)
	// fmt.Println(nodeMap)
	var FinalDir string

	for idx, ins := range instructions {
		var dir string
		var dist int
		fmt.Sscanf(ins, "%1s%d", &dir, &dist)
		if idx == len(instructions)-1 {
			FinalDir = dir
		}
		for i := 0; i < dist; i++ {
			next := &Node{}
			switch dir {
			case "R":
				next = start.right
			case "L":
				next = start.left
			case "U":
				next = start.up
			case "D":
				next = start.down
			}
			if next.block == "wall" {
				break
			}
			start = next
		}
	}
	return start, FinalDir
}

type Node struct {
	row   int
	col   int
	block string
	left  *Node
	right *Node
	up    *Node
	down  *Node
}

func parseInput(input string) (map[string]*Node, *Node, []string) {
	nodeMap := make(map[string]*Node)
	rowBoundMap := make(map[int][]int)
	colBoundMap := make(map[int][]int)
	start := &Node{}
	split := strings.Split(input, "\n\n")
	ctr := 0
	for r, row := range strings.Split(split[0], "\n") {
		rowBoundMap[r+1] = []int{999, 0}
		for c, col := range strings.Split(row, "") {
			if _, ok := colBoundMap[c+1]; !ok {
				colBoundMap[c+1] = []int{999, 0}
			}
			var block string
			switch col {
			case " ":
				continue
			case "#":
				block = "wall"
			case ".":
				block = "path"
			}
			node := &Node{row: r + 1, col: c + 1, block: block}
			nodeMap[fmt.Sprintf("%d,%d", r+1, c+1)] = node
			rowBoundMap[r+1][0] = helpers.MinInt(c+1, rowBoundMap[r+1][0])
			rowBoundMap[r+1][1] = helpers.MaxInt(c+1, rowBoundMap[r+1][1])
			colBoundMap[c+1][0] = helpers.MinInt(r+1, colBoundMap[c+1][0])
			colBoundMap[c+1][1] = helpers.MaxInt(r+1, colBoundMap[c+1][1])
			if ctr == 0 {
				start = node
			}
			ctr++
		}
	}
	// rowBoundMap has keys of rows and values of [minCol, maxCol]
	// colBoundMap has keys of column and values of [minRow, maxRow]

	for _, node := range nodeMap {

		// check up
		if nu, ok := nodeMap[fmt.Sprintf("%d,%d", node.row-1, node.col)]; ok {
			node.up = nu
		} else {
			wrap := colBoundMap[node.col][1]
			node.up = nodeMap[fmt.Sprintf("%d,%d", wrap, node.col)]
		}

		// check down
		if nd, ok := nodeMap[fmt.Sprintf("%d,%d", node.row+1, node.col)]; ok {
			node.down = nd
		} else {
			wrap := colBoundMap[node.col][0]
			node.down = nodeMap[fmt.Sprintf("%d,%d", wrap, node.col)]
		}

		// check left
		if nl, ok := nodeMap[fmt.Sprintf("%d,%d", node.row, node.col-1)]; ok {
			node.left = nl
		} else {
			wrap := rowBoundMap[node.row][1]
			node.left = nodeMap[fmt.Sprintf("%d,%d", node.row, wrap)]
		}

		// check right
		if nr, ok := nodeMap[fmt.Sprintf("%d,%d", node.row, node.col+1)]; ok {
			node.right = nr
		} else {
			wrap := rowBoundMap[node.row][0]
			node.right = nodeMap[fmt.Sprintf("%d,%d", node.row, wrap)]
		}
	}

	split[1] = strings.Replace(split[1], "R", ",R,", -1)
	split[1] = strings.Replace(split[1], "L", ",L,", -1)

	// Simplify instructions into directions and distance
	facing := 0
	newIns := []string{}
	for _, ins := range strings.Split(split[1], ",") {
		var dir string
		switch ins {
		case "R":
			facing++
			continue
		case "L":
			facing--
			continue
		}
		switch helpers.PosMod(facing, 4) {
		case 0:
			dir = "R"
		case 1:
			dir = "D"
		case 2:
			dir = "L"
		case 3:
			dir = "U"
		}
		newIns = append(newIns, fmt.Sprintf("%s%s", dir, ins))
	}
	return nodeMap, start, newIns
}
