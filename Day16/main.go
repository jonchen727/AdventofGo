package main

import (
	//"github.com/jonchen727/2022-AdventofCode/helpers"
	//"container/heap"
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
	"time"
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
	valves := parseInput(input)
	distMap := generateDistanceMap(valves)
	
	return ans
}

func part2(input string) int {
	ans := 0

	return ans
}
func generateDistanceMap(nodes map[string]*Node) map[string]map[string]int {
	// generate a list of all node names
	valveList := []string{}
	for _, node := range nodes {
		valveList = append(valveList, node.name)
	}

	slices.Sort(valveList)
	fmt.Println(valveList)
	distList := make(map[string]map[string]int)
	for _, node := range valveList {
		if nodes[node].value != 0 || node == "AA" {
		distList[node] = make(map[string]int)
		}
	}

	// generate a map of all distances
	for node, _ := range distList {
		for node2, _ := range distList {
			if node != node2 {
					if _, ok := distList[node][node2]; !ok {
					distList[node][node2] = findDistance(nodes, node, node2)
					}
				}
			
		}
	}
	return distList
}
func findDistance(nodes map[string]*Node, start string, end string) int {
	queue := []string{start}
	visited := map[string]bool{start: true}
	dist := map[string]int{start: 0}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == end {
			//fmt.Println("Found", start, end, dist)
			return dist[node]
		}
		for _, next := range nodes[node].next {
			//fmt.Println(start, end, next.name)
			if _, ok := visited[next.name]; !ok {
				visited[next.name] = true
				queue = append(queue, next.name)
				dist[next.name] = dist[node] + 1
			} else {
				continue
			}
		}

	}
	return 0
}

type Node struct {
	name  string
	value int
	next  map[string]*Node
}

func parseInput(input string) map[string]*Node {
	nodes := map[string]*Node{}
	for _, line := range strings.Split(input, "\n") {
		node := Node{}
		fmt.Sscanf(line, "Valve %s has flow rate=%d;", &node.name, &node.value)
		node.next = map[string]*Node{}
		nodes[node.name] = &node

	}

	for _, line := range strings.Split(input, "\n") {
		var name string
		split := strings.Split(line, ";")
		fmt.Sscanf(split[0], "Valve %s", &name)
		replaced := strings.ReplaceAll(split[1], "valves", "valve")
		nextStr := strings.Split(replaced, "valve ")[1]
		// if the next node doesn't exist, create it
		for _, next := range strings.Split(nextStr, ", ") {
			//nxt := nodes[next]
			nodes[name].next[next] = nodes[next]
		}
	}
	return nodes
}
