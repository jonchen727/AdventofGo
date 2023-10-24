package main

import (
	"github.com/jonchen727/2022-AdventofCode/helpers"
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

// part1 is a function that takes a string input and returns an integer.
// It parses the input, generates a distance map, generates a list of all node names,
// and finds the maximum flow between nodes.
func part1(input string) int {
	ans := 0
	valves := parseInput(input)
	distMap, nonempty := generateDistanceMap(valves)
	indicies := generateIndecies(nonempty)
	cache := map[string]int{}
	ans = findMaxFlow(valves, distMap, "AA", 30, indicies, 0, cache)
	//fmt.Println(distMap)

	return ans
}

// part2 is a function that takes a string input and returns an integer.
// It currently does not have any implementation.
func part2(input string) int {
	ans := 0
	valves := parseInput(input)
	distMap, nonempty := generateDistanceMap(valves)
	indicies := generateIndecies(nonempty)
	cache := map[string]int{}
	b := (1<<len(nonempty))-1
	for i := 0; i <= b/2 ; i++ {
		ans = helpers.MaxInt(ans, findMaxFlow(valves, distMap, "AA", 26, indicies, i, cache) +  findMaxFlow(valves, distMap, "AA", 26, indicies, b ^ i, cache))
	}

	return ans
}

// generateIndecies is a function that takes a slice of strings and returns a map of strings to integers.
// It generates a map of node names to their corresponding index in the slice.
func generateIndecies(nonempty []string) map[string]int {
	indicies := map[string]int{}
	for i, valve := range nonempty {
		indicies[valve] = i
	}
	return indicies
}

// generateDistanceMap is a function that takes a map of nodes and returns a map of distances between nodes and a list of all node names.
// It uses breadth-first search to calculate the shortest distance between nodes.
func generateDistanceMap(nodes map[string]*Node) (map[string]map[string]int, []string) {
	valveList := []string{}

	slices.Sort(valveList) // sorts the valveList slice
	//fmt.Println(valveList)
	distMap := make(map[string]map[string]int)

	for _, node := range nodes {
		if node.name != "AA" && node.value == 0 {
			continue
		}
		if node.name != "AA" {
			valveList = append(valveList, node.name)
		}
		distMap[node.name] = make(map[string]int)
		type Valve struct {
			node *Node
			dist int
		}
		queue := []Valve{{nodes[node.name], 0}}
		distMap[node.name] = map[string]int{node.name: 0, "AA": 0}
		visited := map[string]bool{node.name: true}

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			for _, next := range current.node.next {
				if _, ok := visited[next.name]; !ok {
					visited[next.name] = true
					if next.value != 0 {
						distMap[node.name][next.name] = current.dist + 1
					}
					queue = append(queue, Valve{next, current.dist + 1})
				} else {
					continue
				}
			}
		}
		delete(distMap[node.name], node.name)
		if node.name != "AA" {
			delete(distMap[node.name], "AA")
		}
	}
	//fmt.Println(len(distMap))
	return distMap, valveList
}

// findMaxFlow is a function that takes a map of nodes, a map of distances between nodes, a starting node, a time limit, a map of node names to their corresponding index in the slice, a bitmask, and a cache.
// It returns the maximum flow between nodes.
func findMaxFlow(valves map[string]*Node, distmap map[string]map[string]int, start string, time int, indicies map[string]int, bitmask int, cache map[string]int) int {
	
	cachekey := fmt.Sprintf("%d,%s,%d", time, start, bitmask)
	//utilizes a cache using bitmap representation of valves that are on to speed up the process
	if val, ok := cache[cachekey]; ok {
		return val
	}
	maxflow := 0
	for next, dist := range distmap[start] {
		
		bit := 1 << indicies[next]
		//fmt.Println(next, dist)
		if bitmask&bit != 0 {
			continue
		}
		remtime := time - dist - 1
		if remtime <= 0 {
			continue
		}
		//fmt.Println(bitmask, remtime, next, dist, maxflow)
		maxflow = helpers.MaxInt(maxflow, findMaxFlow(valves, distmap, next, remtime, indicies, (bitmask|bit), cache)+(valves[next].value*remtime))
	}
	cache[cachekey] = maxflow
	return maxflow
}

// Node is a struct that represents a node in a graph.
type Node struct {
	name  string
	value int
	next  map[string]*Node
}

// parseInput is a function that takes a string input and returns a map of nodes.
// It parses the input and generates a map of nodes.
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
			nodes[name].next[next] = nodes[next]
		}
	}
	return nodes
}
