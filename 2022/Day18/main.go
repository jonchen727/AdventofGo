package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"strings"
	//"reflect"
	"flag"
	"math"
	"strconv"
	"time"
	//"sort"
	//"github.com/jonchen727/AdventofGo/helpers"
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
	points, bounds := parsePoints(input)
	//fmt.Println(points, bounds)
	ans = floodFill(cubes, points, bounds, faceMap)
	return ans
}

func floodFill(cubes [][][]float64, points map[string]bool, bounds Bounds, faceMap map[string]int) int {
	queue := [][]float64{{bounds.minX, bounds.minY, bounds.minZ}}
	air := map[string]bool{}
	air[fmt.Sprintf("%v,%v,%v", queue[0])] = true
	offset := [][]float64{{-0.5, 0, 0}, {0.5, 0, 0}, {0, 0.5, 0}, {0, -0.5, 0}, {0, 0, 0.5}, {0, 0, -0.5}}
	for len(queue) > 0 {
		x, y, z := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]
		for _, val := range offset {
			nx := x + val[0]*2
			ny := y + val[1]*2
			nz := z + val[2]*2
			key := fmt.Sprintf("%v,%v,%v", nx, ny, nz)

			if !((bounds.minX <= nx && nx <= bounds.maxX) && (bounds.minY <= ny && ny <= bounds.maxY) && (bounds.minZ <= nz && nz <= bounds.maxZ)) {
				continue
			}
			_, ok1 := air[key]
			_, ok2 := points[key]

			if ok1 || ok2 {
				continue
			}

			air[key] = true
			queue = append(queue, []float64{nx, ny, nz})

		}
	}
	free := map[string]bool{}
	for key, _ := range air {

		split := strings.Split(key, ",")
		x, _ := strconv.ParseFloat(split[0], 64)
		y, _ := strconv.ParseFloat(split[1], 64)
		z, _ := strconv.ParseFloat(split[2], 64)
		for _, val := range offset {

			nx := x + val[0]
			ny := y + val[1]
			nz := z + val[2]
			key := fmt.Sprintf("%v,%v,%v", nx, ny, nz)
			free[key] = true
		}
	}
	for key, _ := range free {
		if _, ok := faceMap[key]; !ok {
			delete(free, key)
		}
	}
	return len(free)
}

func makeFaceMap(cubes [][][]float64) map[string]int {
	faceMap := map[string]int{}
	for _, cube := range cubes {
		for _, face := range cube {
			faceMap[fmt.Sprintf("%v,%v,%v", face[0], face[1], face[2])]++
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

type Bounds struct {
	minX, maxX, minY, maxY, minZ, maxZ float64
}

func parsePoints(input string) (map[string]bool, Bounds) {
	points := map[string]bool{}
	min := math.Inf(1)
	max := math.Inf(-1)
	bounds := Bounds{min, max, min, max, min, max}
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(split[0], 64)
		y, _ := strconv.ParseFloat(split[1], 64)
		z, _ := strconv.ParseFloat(split[2], 64)
		if x-1 < bounds.minX {
			bounds.minX = x - 1
		}
		if x+1 > bounds.maxX {
			bounds.maxX = x + 1
		}
		if y-1 < bounds.minY {
			bounds.minY = y - 1
		}
		if y+1 > bounds.maxY {
			bounds.maxY = y + 1
		}
		if z-1 < bounds.minZ {
			bounds.minZ = z - 1
		}
		if z+1 > bounds.maxZ {
			bounds.maxZ = z + 1
		}
		points[fmt.Sprintf("%v,%v,%v", x, y, z)] = true
		//fmt.Println(points)
	}
	return points, bounds
}
