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
	elfMap, grid := parseInput(input)
	grid = moveElvs(elfMap, grid, 10)
	ans = (((grid.xmax - grid.xmin) + 1) * ((grid.ymax - grid.ymin) + 1)) - len(elfMap)
	return ans
}

func part2(input string) int {
	ans := 0
	elfMap, grid := parseInput(input)
	ans = findNoMovement(elfMap, grid, 999)
	return ans
}
func findNoMovement(elfMap map[string]bool, grid Grid, rounds int) int {
	round := 0
	dirQueue := []string{"North", "South", "West", "East"}

	for i := 0; i < rounds; i++ {
		nextMap := make(map[string][]string)
		// fmt.Println("Round:", i+1, "First Dir:", dirQueue)
		for key, _ := range elfMap {
			var x, y int
			moves := []string{}
			fmt.Sscanf(key, "%d,%d", &x, &y)
			// North New Positions
			N := fmt.Sprintf("%d,%d", x, y-1)
			NE := fmt.Sprintf("%d,%d", x+1, y-1)
			NW := fmt.Sprintf("%d,%d", x-1, y-1)
			// South New Positions
			S := fmt.Sprintf("%d,%d", x, y+1)
			SE := fmt.Sprintf("%d,%d", x+1, y+1)
			SW := fmt.Sprintf("%d,%d", x-1, y+1)
			// West New Positions
			W := fmt.Sprintf("%d,%d", x-1, y)
			// East New Positions
			E := fmt.Sprintf("%d,%d", x+1, y)
			// Check North Positions

			checkMap := map[string][]string{
				"North": {N, NE, NW},
				"South": {S, SE, SW},
				"West":  {W, NW, SW},
				"East":  {E, NE, SE},
			}

			for _, dir := range dirQueue {
				if directionCheck(checkMap[dir], elfMap) {
					moves = append(moves, checkMap[dir][0])
				}
			}

			if len(moves) == 0 || len(moves) == 4 {
				continue
			} else {
				if _, ok := nextMap[moves[0]]; ok {
					nextMap[moves[0]] = append(nextMap[moves[0]], key)
				} else {
					nextMap[moves[0]] = []string{key}
				}
			}
		}
		if len(nextMap) == 0 {
			round = i + 1
			break
		}
		for key, val := range nextMap {
			if len(val) > 1 {
				continue
			}
			elfMap[key] = true
			delete(elfMap, val[0])
			var x, y int
			fmt.Sscanf(key, "%d,%d", &x, &y)
			grid.xmin = helpers.MinInt(grid.xmin, x)
			grid.xmax = helpers.MaxInt(grid.xmax, x)
			grid.ymin = helpers.MinInt(grid.ymin, y)
			grid.ymax = helpers.MaxInt(grid.ymax, y)
		}
		//fmt.Println(grid, elfMap)
		apnd := dirQueue[0]
		dirQueue = append(dirQueue[1:], apnd)
	}
	// fmt.Println("Final Grid", grid)
	// for y := grid.ymin; y <= grid.ymax; y++ {
	// 	for x := grid.xmin; x <= grid.xmax; x++ {
	// 		if _, ok := elfMap[fmt.Sprintf("%d,%d", x, y)]; ok {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	return round
}
func moveElvs(elfMap map[string]bool, grid Grid, rounds int) Grid {

	dirQueue := []string{"North", "South", "West", "East"}

	for i := 0; i < rounds; i++ {
		nextMap := make(map[string][]string)
		fmt.Println("Round:", i+1, "First Dir:", dirQueue)
		for key, _ := range elfMap {
			var x, y int
			moves := []string{}
			fmt.Sscanf(key, "%d,%d", &x, &y)
			// North New Positions
			N := fmt.Sprintf("%d,%d", x, y-1)
			NE := fmt.Sprintf("%d,%d", x+1, y-1)
			NW := fmt.Sprintf("%d,%d", x-1, y-1)
			// South New Positions
			S := fmt.Sprintf("%d,%d", x, y+1)
			SE := fmt.Sprintf("%d,%d", x+1, y+1)
			SW := fmt.Sprintf("%d,%d", x-1, y+1)
			// West New Positions
			W := fmt.Sprintf("%d,%d", x-1, y)
			// East New Positions
			E := fmt.Sprintf("%d,%d", x+1, y)
			// Check North Positions

			checkMap := map[string][]string{
				"North": {N, NE, NW},
				"South": {S, SE, SW},
				"West":  {W, NW, SW},
				"East":  {E, NE, SE},
			}

			for _, dir := range dirQueue {
				if directionCheck(checkMap[dir], elfMap) {
					moves = append(moves, checkMap[dir][0])
				}
			}

			if len(moves) == 0 || len(moves) == 4 {
				continue
			} else {
				if _, ok := nextMap[moves[0]]; ok {
					nextMap[moves[0]] = append(nextMap[moves[0]], key)
				} else {
					nextMap[moves[0]] = []string{key}
				}
			}
		}
		for key, val := range nextMap {
			if len(val) > 1 {
				continue
			}
			elfMap[key] = true
			delete(elfMap, val[0])
			var x, y int
			fmt.Sscanf(key, "%d,%d", &x, &y)
			grid.xmin = helpers.MinInt(grid.xmin, x)
			grid.xmax = helpers.MaxInt(grid.xmax, x)
			grid.ymin = helpers.MinInt(grid.ymin, y)
			grid.ymax = helpers.MaxInt(grid.ymax, y)
		}
		fmt.Println(grid, elfMap)
		apnd := dirQueue[0]
		dirQueue = append(dirQueue[1:], apnd)
	}
	fmt.Println("Final Grid", grid)
	for y := grid.ymin; y <= grid.ymax; y++ {
		for x := grid.xmin; x <= grid.xmax; x++ {
			if _, ok := elfMap[fmt.Sprintf("%d,%d", x, y)]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return grid
}

func directionCheck(dir []string, elfMap map[string]bool) bool {
	move := true
	for _, key := range dir {
		if _, ok := elfMap[key]; ok {
			move = false
			break
		} else {
			continue
		}
	}
	return move
}

type Grid struct {
	xmin int
	xmax int
	ymin int
	ymax int
}

func parseInput(input string) (map[string]bool, Grid) {
	elfMap := make(map[string]bool)
	grid := Grid{math.MaxInt32, math.MinInt32, math.MaxInt32, math.MinInt32}
	for y, line := range strings.Split(input, "\n") {
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				elfMap[fmt.Sprintf("%d,%d", x, y)] = true
			}
			grid.xmin = helpers.MinInt(grid.xmin, x)
			grid.xmax = helpers.MaxInt(grid.xmax, x)
			grid.ymin = helpers.MinInt(grid.ymin, y)
			grid.ymax = helpers.MaxInt(grid.ymax, y)

		}
	}
	return elfMap, grid
}
