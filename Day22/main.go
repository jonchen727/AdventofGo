package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"strings"
	//"reflect"
	//"strconv"
	"flag"
	"math"
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
	parseInput2(input)
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
	face  string
	cubex int
	cubey int
}

func parseInput2(input string) (map[string]*Node, *Node, []string) {
	nodeMap := make(map[string]*Node)
	rowBoundMap := make(map[int][]int)
	colBoundMap := make(map[int][]int)
	start := &Node{}
	trig := true
	split := strings.Split(input, "\n\n")
	count := (len(strings.Replace(strings.Replace(split[0], " ", "", -1), "\n", "", -1)) + 1) / 6
	maxr := len(strings.Split(split[0], "\n"))
	//maxc := len(strings.Split(strings.Split(split[0], "\n")[0], ""))
	edge := int(math.Sqrt(float64(count)))
	fmt.Println(edge)

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
			if trig == true {
				start = node
				trig = false
			}
		}
	}

	fmt.Println(len(rowBoundMap))
	fmt.Println(len(colBoundMap))

	// Cube Face Planning
	// Top left of map is always going to be the Top Left of the cube
	// Notation is as follows T = Top, B = Bottom, L = Left, R = Right, F = Front, Ba = Back

	// Cube configuration tree
	//     -3    2   -1    0    1    2    3      <- x
	//
	//             			   T -> R -> B -> L      y=0
	//             			   ↓
	//     R <- Ba <- L <- F -> R -> Ba -> L     y=1
	//                     ↓
	//                L <- B -> R					       y=2
	//										 ↓
	//							  L <- Ba -> R               y=3
	//
	//   cmin, cmax, rmin, rmax will represent min and max of each face
	//
	//
	// 1. Everything directly below is continous and crossing and edge will maintain the same direction
	// 2. Everything directly to the right/left is continous and crossing and edge will maintain the same direction
	//
	// Was trying to build a general case but screw it ill just do it based on the input.

	offset := start.col / edge
	// heres some framework for the general case if someone wants to continue it.
	// Shape will be
	//   -1  0  1
	//      [T][R]  0
	//			[F]     1
	//   [L][B]     2
	//	 [Ba]       3

	faceDict := map[string]string{
		"0,0":  "T",
		"1,0":  "R",
		"0,1":  "F",
		"-1,2": "L",
		"0,2":  "B",
		"-1,3": "Ba",
	}
	for i := 0; i < maxr/edge; i++ {
		// breaks it apart into vertial face sections
		cmin := rowBoundMap[((i)*edge)+1][0]
		cmax := rowBoundMap[((i)*edge)+1][1]
		rmin := (i)*edge + 1
		rmax := (i + 1) * edge
		// now go row by row
		for r := rmin; r <= rmax; r++ {
			// break into horizontal face sections
			for c := cmin; c <= cmax; c++ {
				//fmt.Println(r,c)
				node := nodeMap[fmt.Sprintf("%d,%d", r, c)]
				x := (node.col / (edge + 1)) - offset
				y := i
				node.cubex = x
				node.cubey = y
				node.face = faceDict[fmt.Sprintf("%d,%d", x, y)]
			}
		}
	}

	// rowBoundMap has keys of rows and values of [minCol, maxCol]
	// colBoundMap has keys of column and values of [minRow, maxRow]

	//   -1  0  1
	//      [T][R]  0
	//			[F]     1
	//   [L][B]     2
	//	 [Ba]       3
	//
	// Build Cases
	// T Cases x=0 y=0
	// [Left] T -> L y=2 x=-1 rot 90 
	// [Right] T -> R same y nothing 
	// [Up] T -> Ba y=3 x=-1 rot 90  
	// [Down] T -> F same x nothing 

	// R Cases x=1 y=0
	// [Left] R -> T same y nothing 
	// [Right] R -> B x=0 y=2 rot 180
	// [Up] R -> Ba x=-1 y=3 nothing 
	// [Down] R -> F x=0 y=1 rot 270

	// F Cases x=0 y=1
	// [Left] F -> L x=-1 y=2 rot 90
	// [Right] F -> R x=0 y= 1 rot 90
	// [Up] F -> T same y nothing 
	// [Down] F -> B same y nothing 

	// L Cases x=-1 y=2
	// [Left] L -> T x=0 y=0 rot 180
	// [Right] L -> B same y nothing 
	// [Up] L -> F x=0 y=1 rot 270
	// [Down] L -> Ba same x nothing 

	// B Cases x=0 y=2
	// [Left] B -> L same y nothing 
	// [Right] B -> R x=1 y=1 rot 180
	// [Up] B -> F same x nothing 
	// [Down] B -> Ba x=-1 y=3 rot 270

	// Ba Cases x=-1 y=3
	// [Left] Ba -> T x=0 y=0 rot 90
	// [Right] Ba -> B x=0, y=2 rot 90
	// [Up] Ba -> L same x nothing 
	// [Down] Ba -> R x=1 y=0 nothing 



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
