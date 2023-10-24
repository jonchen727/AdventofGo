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
var priorities = map[string]int{}

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
		ans := part2(input, 9)
		fmt.Println("Part 2 Answer:", ans)
		//fmt.Println("Answer:", ans)
	}
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func part1(input string) int {
	inst := parseInput(input)
	head := Point{0, 0}
	tail := Point{0, 0}

	visited := map[Point]bool{}

	for _, instruction := range inst {
		//fmt.Println(instruction)
		for i := 0; i < instruction.distance; i++ {

			switch instruction.direction {
			case "R":
				head.X++
			case "L":
				head.X--
			case "U":
				head.Y++
			case "D":
				head.Y--
			}
			if helpers.Abs(head.X-tail.X) > 1 || helpers.Abs(head.Y-tail.Y) > 1 {
				if head.X > tail.X {
					tail.X++
				} else if head.X < tail.X {
					tail.X--
				}

				if head.Y > tail.Y {
					tail.Y++
				} else if head.Y < tail.Y {
					tail.Y--
				}
			}
			fmt.Println(head, tail)
			visited[tail] = true
		}

	}

	fmt.Println(visited)
	ans := len(visited)
	return ans
}

func part2(input string, links int) int {
	links = links+1
	inst := parseInput(input)
	points := []Point{}
	visited := []map[Point]bool{}
	for i := 0; i < links; i++ {
		points = append(points, Point{0, 0})
		visited = append(visited, map[Point]bool{})
	}

	for _, instruction := range inst {
		//fmt.Println(instruction)
		for i := 0; i < instruction.distance; i++ {

			switch instruction.direction {
			case "R":
				points[0].X++
			case "L":
				points[0].X--
			case "U":
				points[0].Y++
			case "D":
				points[0].Y--
			}

			for j := 0; j < links-1; j++ {
				//fmt.Println("Link:", j)
				if helpers.Abs(points[j].X-points[j+1].X) > 1 || helpers.Abs(points[j].Y-points[j+1].Y) > 1 {
					if points[j].X > points[j+1].X {
						points[j+1].X++
					} else if points[j].X < points[j+1].X {
						points[j+1].X--
					}

					if points[j].Y > points[j+1].Y {
						points[j+1].Y++
					} else if points[j].Y < points[j+1].Y {
						points[j+1].Y--
					}
				}
				visited[j+1][points[j+1]] = true
			}
		}

	}
	//for k := 0; k < links; k++ {
	//	fmt.Println("Link:",k,visited[k])
	//}

	ans := len(visited[links-1])
	return ans
}

type instruction struct {
	direction string
	distance  int
}

type Point struct {
	X int
	Y int
}

func parseInput(input string) (ans []instruction) {
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " ")
		ans = append(ans, instruction{
			direction: split[0],
			distance:  helpers.ToInt(split[1]),
		})
	}
	return ans

}
