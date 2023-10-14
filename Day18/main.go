package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"strings"
	//"reflect"
	"flag"
	"strconv"
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

func part1(input string) int {
	ans := 0
	cubes := parseInput(input)
	faceMap := makeFaceMap(cubes)
	for _, val := range faceMap {
		if val == 1 {
			ans++
		}
	}

	return ans
}

func part2(input string) int {
	ans := 0
	cubes := parseInput(input)
	faceMap := makeFaceMap(cubes)
  for _, val := range faceMap {
		if val == 2 {
			ans++
		}

	}
	return ans
}


func makeFaceMap(cubes [][][]float64) map[string]int {
	faceMap := map[string]int{}
	for _, cube := range cubes {
		for _, face := range cube {
			faceMap[fmt.Sprintf("%v", face)]++
		}
	}
	return faceMap
}

func parseInput(input string) [][][]float64 {
	offset := [][]float64{{-0.5, 0, 0}, {0.5, 0, 0}, {0, 0.5, 0}, {0, -0.5, 0}, {0, 0, 0.5}, {0, 0, -0.5}}
	//fmt.Println(offset)

	cubes := [][][]float64{}
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(split[0], 64)
		y, _ := strconv.ParseFloat(split[1], 64)
		z, _ := strconv.ParseFloat(split[2], 64)

		faces := [][]float64{}
		for _, val := range offset {
			face := []float64{x + val[0], y + val[1], z + val[2]}
			faces = append(faces, face)
		}
		cubes = append(cubes, faces)
	}
	return cubes
}
