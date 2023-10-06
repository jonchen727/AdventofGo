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
		//fmt.Println("Answer:", ans)
	}
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func part1(input string) int {
  grid, start, end := parseInput(input)
	//fmt.Println(len(grid), len(grid[Key{0,0}].canmove))
	fillCanMove(grid)
	
	//fmt.Println(len(input))
	

	//   for i := 0; i < 41; i++ {
	//   	fmt.Println("")
	//   	for j := 0; j < 77; j++ {
	//   		if grid[Key{j,i}].elevation < 10 {
	//   		fmt.Print(grid[Key{j,i}].elevation, "  ")
	//   		} else {
	//   			fmt.Print(grid[Key{j,i}].elevation, " ")
	//   		}
	//  	}
	//  }
	//fmt.Println(grid)
	ans, err := (findShortestPath(grid, start, end))
	if err != nil {
		//fmt.Println(grid[Key{7,37}], grid[Key{8,38}])
		panic(err)

	}
	return ans
}

func part2(input string) int {
	grid, starts, end := parseInput2(input)
	fillCanMove(grid)
	minSteps := math.MaxInt32
	for _, start := range starts {
    ans, err := (findShortestPath(grid, start, end))
		if err != nil {
			continue
		}
		if ans < minSteps {
			minSteps = ans
		}
	}
	return minSteps
}

type Point struct {
	elevation int
	canmove []Key

}

type Key struct {
	X, Y int
}

func parseInput(input string)(grid map[Key]Point, start Key, end Key){
	grid = map[Key]Point{}
	start = Key{}
	end = Key{}
	for r, line := range strings.Split(input, "\n") {
		for c, char := range line {
			key := Key{X: c, Y: r}
			var elevation int
			switch char {
				case 'S':
					elevation = 'a'
					start = key
				case 'E':
					elevation = 'z'
					end = key
				default: 
				  elevation = helpers.ToInt(char)
			}
			elevation = elevation - 'a'
			
			grid[key] = Point{elevation: elevation}
			//fmt.Println(r,c, grid[key])
		}

	}
	return grid, start, end
}

func parseInput2(input string)(grid map[Key]Point, starts []Key, end Key){
	grid = map[Key]Point{}
	starts = []Key{}
	end = Key{}
	for r, line := range strings.Split(input, "\n") {
		for c, char := range line {
			key := Key{X: c, Y: r}
			var elevation int
			switch char {
				case 'S':
					elevation = 'a'
					starts = append(starts,key)
				case 'E':
					elevation = 'z'
					end = key
				case 'a':
					elevation = 'a'
					starts = append(starts,key)
				default: 
				  elevation = helpers.ToInt(char)
			}
			elevation = elevation - 'a'
			
			grid[key] = Point{elevation: elevation}
			//fmt.Println(r,c, grid[key])
		}

	}
	return grid, starts, end
}

func fillCanMove(grid map[Key]Point) {
	for key, point := range grid {
		point.canmove = []Key{}
		for offst := -1; offst <= 1 ; offst += 2 {
			//fmt.Println("Key:", key, "Offset:", offst)
			xdir, ok := grid[Key{X: key.X + offst, Y: key.Y}] 
			if (ok && xdir.elevation <= point.elevation+1) {
				
				//fmt.Println("Key:", key, "Point:", Key{X: key.X + offst, Y: key.Y} )
				//grid[key].canmove = append(grid[key].canmove, Key{X: key.X + offst, Y: key.Y})
				point.canmove = append(point.canmove, Key{X: key.X + offst, Y: key.Y})

			}
			ydir, ok := grid[Key{X: key.X, Y: key.Y + offst}]
			if (ok && ydir.elevation <= point.elevation+1) {
				//fmt.Println("Key:", key, "Point:", Key{X: key.X, Y: key.Y + offst} )
				//grid[key].canmove = append(grid[key].canmove, Key{X: key.X, Y: key.Y + offst})
				point.canmove = append(point.canmove, Key{X: key.X, Y: key.Y + offst})
			}
		}
		//fmt.Println(point)
		grid[key] = point
		//fmt.Println(grid[key].elevation)

		//grid[key] = point
	}
}

// findShortestPath uses a breadth-first search algorithm to find the 
// shortest path between two points in a grid.
func findShortestPath(grid map[Key]Point, start Key, end Key) (int, error) {
	// If the start and end points are the same, return a slice containing only the start point.
	numSteps := 0
	if start == end {
			return numSteps, nil
	}

	// Initialize a queue with a single path containing only the start point and a visited map 
	// to keep track of the points that have been visited.
	queue := []Key{start}
	visited := map[Key]bool{start: true}

	// Loop until the queue is empty.
	for len(queue) > 0 {
		// Create an empty slice to store the current path.
		path := []Key{}

		// Get the size of the queue.
		queueSize := len(queue)

		// Iterate over the queue.
		for i := 0; i < queueSize; i++ {
			// Dequeue the first path from the queue and retrieve the last point in the path.
			current := queue[0]
			queue = queue[1:]

			// Append the current point to the path.
			path = append(path, current)

			// If the last point is the end point, return the length of the path.
			if current == end {
				//fmt.Println(path)
				return numSteps, nil
			}

			// Iterate over the canmove field of the last point in the grid map.
			for _, canmove := range grid[current].canmove {
				// If the point has not been visited, add it to the visited map and create a 
				// new path by appending the canmove point to the end of the current path. Add the new path to the end of the queue.
				if visited[canmove] {
					continue
				}							
				visited[canmove] = true
				queue =  append(queue, canmove)
			}
		}
		// Increment the number of steps taken.
		numSteps++
	}

	// If the end point was not found, return an error.
	return 0, fmt.Errorf("endpoint not found")
}
