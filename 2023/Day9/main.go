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
	sets := parseInput(input)
	//fmt.Println(sets)

	triangles := [][][]int{}
	for _, set := range sets {
		triangle := [][]int{}
		triangle = append(triangle, set)
		triangles = append(triangles, triangle)
	}

	for i := range triangles {
		triangles[i] = buildTriangle(triangles[i])
	}


	for i := range triangles {
		triangles[i] = extrapolateRight(triangles[i], 0)
	}

	for _, triangle := range triangles {
		ans += triangle[0][len(triangle[0])-1]
	}


	return ans
}
func extrapolateLeft(triangle [][]int, level int) [][]int {
	max := len(triangle)-1
	if level == 0 {
		triangle[max] = append([]int{0}, triangle[max]...)
	} else {
		diff := triangle[max-level+1][0]
		triangle[max-level] = append([]int{triangle[max-level][0]-diff}, triangle[max-level]...)
	}

	if level == max {
		return triangle
	}
	return extrapolateLeft(triangle, level+1)
}

func extrapolateRight(triangle [][]int, level int) [][]int {
	max := len(triangle)-1
	if level == 0 {
		triangle[max] = append(triangle[max], 0)
	} else {
		diff := triangle[max-level+1][0]
		triangle[max-level] = append(triangle[max-level], triangle[max-level][0]+diff)
	}

	if level == max {
		return triangle
	}

	return extrapolateRight(triangle, level+1)


}

func buildTriangle(triangle [][]int) [][]int {
	max := len(triangle)
	max2 := len(triangle[max-1])
	newTier := []int{}
	state := 0
	for i := 1; i < max2; i++ {
		diff := triangle[max-1][i] - triangle[max-1][i-1]
		if diff == 0 {
			state |= (1 << (i-1))
		}

		newTier = append(newTier, diff)
	}
	triangle = append(triangle, newTier)
	if state == (1<<(max2-1))-1 {
		return triangle
	}

	return buildTriangle(triangle)
}

func part2(input string) int {
	ans := 0
	sets := parseInput(input)
	//fmt.Println(sets)

	triangles := [][][]int{}
	for _, set := range sets {
		triangle := [][]int{}
		triangle = append(triangle, set)
		triangles = append(triangles, triangle)
	}

	for i := range triangles {
		triangles[i] = buildTriangle(triangles[i])
	}


	for i := range triangles {
		triangles[i] = extrapolateLeft(triangles[i], 0)
	}

	for _, triangle := range triangles {
		ans += triangle[0][0]
	}


	return ans
}

func parseInput(input string) [][]int {
	sets := [][]int{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		set := []int{}
		for _, num := range strings.Split(line, " ") {
			set = append(set, helpers.ToInt(num))
		}
		sets = append(sets, set)
	}
	return sets
}
