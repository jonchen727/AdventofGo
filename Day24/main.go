package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	//"reflect"
	//"strconv"
	"flag"
	//"math"
	"time"
	//"sort"
	//"github.com/jonchen727/2022-AdventofCode/helpers"
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

type State struct {
	y, x, turn int
}

func part1(input string) int {
	ans := 0
	bounds, start, end, blizzardMap := parseInput(input)
	queue := []State{State{start[0], start[1], 0}}
	fmt.Println(end, queue)
	mod := 0
	cache := map[int]map[string][]string{}
	cache[0] = blizzardMap
	turn := 0
	for true {
		turn++
		newBlizzardMap := moveBlizzard(blizzardMap, bounds)
		cache[turn] = newBlizzardMap
		//fmt.Println("%v", newBlizzardMap)
		if fmt.Sprintf("%v", cache[1]) == fmt.Sprintf("%v", newBlizzardMap) && turn != 1 {
			fmt.Println("dupe found")
			mod = turn - 1
			break
		}
		blizzardMap = newBlizzardMap

	}

	seen := map[string]int{}
out:
	for true {
		//fmt.Println(queue)
		current := queue[0]
		fmt.Println(current.turn)
		blizzardMap = cache[current.turn%mod]
		queue = queue[1:]
		if current.y == end[0] && current.x == end[1] {
			fmt.Println("found")
			break out
		}
		// for y := 0; y < bounds.bottom; y++ {
		// 	for x := 0; x < bounds.right; x++ {

		// 		if b, ok := blizzardMap[fmt.Sprintf("%d,%d", y, x)]; ok {
		// 			if len(b) > 1 {
		// 				fmt.Print(helpers.ToString(len(b)))
		// 				if current.x == x && current.y == y {
		// 					fmt.Print("&")
		// 				}
		// 			} else {
		// 				fmt.Print(b[0])
		// 				if current.x == x && current.y == y {
		// 					fmt.Print("&")
		// 				}
		// 			}

		// 		} else {
		// 			if current.x == x && current.y == y {
		// 				fmt.Print("E")
		// 			}else{
		// 					fmt.Print(".")
		// 				}

		// 		}

		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println()

		//fmt.Println(blizzardMap)

		nextBlizzardMap := cache[(current.turn+1)%mod]
		//fmt.Println(seen)
		for x := -1; x <= 1; x++ {
			if current.x+x > 0 && current.x+x < bounds.right && current.y > 0 {
				//_, ok1 := blizzardMap[fmt.Sprintf("%d,%d", current.y, current.x)]
				_, ok2 := nextBlizzardMap[fmt.Sprintf("%d,%d", current.y, current.x+x)]
				if !ok2 {

					if s, ok := seen[fmt.Sprintf("%d,%d,%d", current.y, current.x+x, current.turn)]; !ok {
						queue = append(queue, State{current.y, current.x + x, current.turn + 1})
						seen[fmt.Sprintf("%d,%d,%d", current.y, current.x+x, current.turn)] = 1
					} else {
						if s > 5 {
							fmt.Println("seen too manytimes")
							continue
						} else {
							queue = append(queue, State{current.y, current.x + x, current.turn + 1})
							seen[fmt.Sprintf("%d,%d,%d", current.y, current.x+x, current.turn)] = s + 1
						}
					}
				}
			}
		}
		for y := -1; y <= 1; y++ {
			// fmt.Println(current.y+y,current.x)
			if current.y+y == end[0] && current.x == end[1] {
				fmt.Println("found")
				ans = current.turn
				break out
			}
			if current.y+y >= bounds.top && current.y+y < bounds.bottom {
				//_, ok1 := blizzardMap[fmt.Sprintf("%d,%d", current.y, current.x)]
				_, ok2 := nextBlizzardMap[fmt.Sprintf("%d,%d", current.y+y, current.x)]
				if !ok2 {

					if s, ok := seen[fmt.Sprintf("%d,%d,%d", current.y+y, current.x, current.turn)]; !ok {
						queue = append(queue, State{current.y + y, current.x, current.turn + 1})
						seen[fmt.Sprintf("%d,%d,%d", current.y+y, current.x, current.turn)] = 1
					} else {
						if s > 5 {
							fmt.Println("seen too manytimes")
							continue

						} else {
							queue = append(queue, State{current.y + y, current.x, current.turn + 1})
							seen[fmt.Sprintf("%d,%d,%d", current.y+y, current.x, current.turn)] = s + 1
						}
					}
				}
			}
		}

	}

	//fmt.Println(blizzardMap)
	return ans
}

func moveBlizzard(blizzardMap map[string][]string, bounds Bounds) map[string][]string {
	newBlizzardMap := map[string][]string{}
	// fmt.Println(turns)
	for pos, blizzards := range blizzardMap {
		var x, y int
		fmt.Sscanf(pos, "%d,%d", &y, &x)
		for _, blizzard := range blizzards {
			// fmt.Println(blizzard)
			var newPos string
			switch blizzard {
			case "^":
				yNew := y - 1
				if yNew == bounds.top {
					yNew = bounds.bottom - 1
				}
				newPos = fmt.Sprintf("%d,%d", yNew, x)
			case "v":
				yNew := y + 1
				if yNew == bounds.bottom {
					yNew = bounds.top + 1
				}
				newPos = fmt.Sprintf("%d,%d", yNew, x)
			case "<":
				xNew := x - 1
				if xNew == bounds.left {
					xNew = bounds.right - 1
				}
				newPos = fmt.Sprintf("%d,%d", y, xNew)
			case ">":
				xNew := x + 1
				if xNew == bounds.right {
					xNew = bounds.left + 1
				}
				newPos = fmt.Sprintf("%d,%d", y, xNew)
			}
			if _, ok := newBlizzardMap[newPos]; ok {
				newBlizzardMap[newPos] = append(newBlizzardMap[newPos], blizzard)
				slices.Sort(newBlizzardMap[newPos])
			} else {
				newBlizzardMap[newPos] = []string{blizzard}
			}
		}
	}
	// fmt.Println("turns0")
	return newBlizzardMap
}

func part2(input string) int {
	ans := 0
	return ans
}

type Bounds struct {
	top, bottom, left, right int
}

func parseInput(input string) (Bounds, []int, []int, map[string][]string) {
	split := strings.Split(input, "\n")
	var start []int
	var end []int
	bounds := Bounds{0, len(split) - 1, 0, (len(strings.Split(split[0], "")) - 1)}
	blizzardMap := map[string][]string{}

	for i, line := range split {
		for j, char := range line {
			c := string(char)
			//fmt.Println(c)
			switch c {
			case "#":
				continue
			case ".":
				if i == 0 {
					start = []int{i, j}
				}
				if i == bounds.bottom {
					end = []int{i, j}
				}
			default:
				key := fmt.Sprintf("%d,%d", i, j)
				blizzardMap[key] = []string{c}
			}
		}
	}
	return bounds, start, end, blizzardMap

}
