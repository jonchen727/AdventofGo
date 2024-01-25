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
	nodes, start := parseInput(input)
	ans = findMaxDistance(nodes, start)

	return ans
}

func findMaxDistance(nodes map[string]Node, start string) int {
	seen := map[string]bool{}
	loop := map[string]bool{
		start: true,
	}
	queue := []string{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if seen[current] {
			continue
		}
		seen[current] = true

		for _, connection := range nodes[current].connections {
			if _, ok := seen[connection]; !ok {
				loop[connection] = true
				copy := nodes[connection]
				nodes[connection] = copy
				queue = append(queue, connection)
			}
		}
	}
	return (len(loop)) / 2
}

func part2(input string) int {
	ans := 0
	nodes, start := parseInput(input)
	area, nbound := shoeLace(nodes, start)
	ans = int(picksTheorm(area, nbound))
	return ans

}

func picksTheorm(area float64, nbound float64) float64 {
	return area - (nbound / 2) + 1
}

func shoeLace(nodes map[string]Node, start string) (float64, float64) {
	lookup := map[string][][]int{
		"|": [][]int{{-1, 0}, {1, 0}},
		"-": [][]int{{0, 1}, {0, -1}},
		"L": [][]int{{-1, 0}, {0, 1}},
		"J": [][]int{{-1, 0}, {0, -1}},
		"7": [][]int{{1, 0}, {0, -1}},
		"F": [][]int{{0, 1}, {1, 0}},
	}

	var sr, sc int
	_, ok := fmt.Sscanf(start, "%d,%d", &sr, &sc)
	if ok != nil {
		panic("get a new job")
	}

	for piece, connections := range lookup {
		c1 := fmt.Sprintf("%d,%d", sr+connections[0][0], sc+connections[0][1])
		c2 := fmt.Sprintf("%d,%d", sr+connections[1][0], sc+connections[1][1])
		if c1 == nodes[start].connections[0] && c2 == nodes[start].connections[1] {
			startcopy := nodes[start]
			startcopy.pipe = piece
			nodes[start] = startcopy
		}
	}

	seen := map[string]bool{}
	loop := map[string]bool{
		start: true,
	}
	verticies := []Point{}
	if checkVerticies(nodes[start]) {
		point := Point{}
		fmt.Sscanf(start, "%f,%f", &point.y, &point.x)
		verticies = append(verticies, point)
	}
	//force the bfs to go in one direction
	queue := []string{nodes[start].connections[0]}
	loop[queue[0]] = true

	seen[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if seen[current] {
			continue
		}
		seen[current] = true

		if checkVerticies(nodes[current]) {
			point := Point{}
			fmt.Sscanf(current, "%f,%f", &point.y, &point.x)
			verticies = append(verticies, point)
		}

		for _, connection := range nodes[current].connections {
			if _, ok := seen[connection]; !ok {
				loop[connection] = true
				copy := nodes[connection]
				nodes[connection] = copy
				queue = append(queue, connection)
			}
		}
	}

	area := float64(0)
	n := len(verticies)

	for i := 0; i < n-1; i++ {
		area += verticies[i].x * verticies[i+1].y
		area -= verticies[i].y * verticies[i+1].x
	}
	area += verticies[n-1].x * verticies[0].y
	area -= verticies[n-1].y * verticies[0].x
	area = helpers.Abs(area) / float64(2)
	fmt.Println(area, float64(len(loop)))
	return area, float64(len(loop))
}

type Point struct {
	x float64
	y float64
}

func checkVerticies(node Node) bool {
	if node.pipe != "|" && node.pipe != "-" {
		return true
	}
	return false
}

type Node struct {
	connections []string
	pipe        string
	distance    int
}

func parseInput(input string) (map[string]Node, string) {
	rmax := len(strings.Split(input, "\n"))
	cmax := len(strings.Split(input, "\n")[0])
	// direction is always N, E, S, W for foward
	lookup := map[string][][]int{
		"|": [][]int{{-1, 0}, {1, 0}},
		"-": [][]int{{0, 1}, {0, -1}},
		"L": [][]int{{-1, 0}, {0, 1}},
		"J": [][]int{{-1, 0}, {0, -1}},
		"7": [][]int{{1, 0}, {0, -1}},
		"F": [][]int{{0, 1}, {1, 0}},
	}
	var start string
	nodes := map[string]Node{}
	for i, line := range strings.Split(input, "\n") {
		for j, char := range line {
			if char == 'S' {
				start = fmt.Sprintf("%d,%d", i, j)
				nodes[start] = Node{
					connections: []string{},
					pipe:        string(char),
				}
				break
			} else {
				continue
			}
		}
	}
	for i, line := range strings.Split(input, "\n") {
		for j, char := range line {
			if arr, ok := lookup[string(char)]; ok {
				connections := []string{}
				for _, cord := range arr {
					r := i + cord[0]
					c := j + cord[1]
					if r < 0 || c < 0 {
						continue
					}
					if r >= rmax || c >= cmax {
						continue
					}
					if fmt.Sprintf("%d,%d", r, c) == start {
						if nstart, ok := nodes[start]; ok {
							nstart.connections = append(nstart.connections, fmt.Sprintf("%d,%d", i, j))
							nstart.distance = 0
							nodes[start] = nstart
						}
					}
					connections = append(connections, fmt.Sprintf("%d,%d", i+cord[0], j+cord[1]))
				}
				nodes[fmt.Sprintf("%d,%d", i, j)] = Node{
					connections: connections,
					pipe:        string(char),
				}
			} else {
				continue
			}
		}
	}
	return nodes, start
}
