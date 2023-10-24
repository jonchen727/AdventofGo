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
	//ans := 0
	cords, boundary := parseInput(input)
	cave := GenerateCave(cords)
	ans := fillSand(cave, boundary, Point{500, 0})

	return ans
}

func part2(input string) int {

	cords, boundary := parseInput(input)
	cave := GenerateCave(cords)
	ans := fillSand2(cave, boundary, Point{500, 0})

	return ans
}

type Point struct {
	x int
	y int
}

type Bounds struct {
	xmin int
	xmax int
	ymin int
	ymax int
}

func parseInput(input string) ([][]Point, Bounds) {
	cords := [][]Point{}
	xmin := math.MaxInt32
	xmax := 0
	ymin := math.MaxInt32
	ymax := 0
	for _, line := range strings.Split(input, "\n") {
		row := []Point{}
		for _, cord := range strings.Split(line, " -> ") {
			split := strings.Split(cord, ",")
			x := helpers.ToInt(split[0])
			y := helpers.ToInt(split[1])
			if x < xmin {
				xmin = x
			}
			if x > xmax {
				xmax = x
			}
			if y < ymin {
				ymin = y
			}
			if y > ymax {
				ymax = y
			}
			row = append(row, Point{x, y})
		}
		cords = append(cords, row)
	}
	return cords, Bounds{xmin, xmax, ymin, ymax}
}

func GenerateCave(cave [][]Point) map[string]bool {
	caveMap := map[string]bool{}
	for _, row := range cave {
		for c := 1; c < len(row); c++ {
			xdif := helpers.Abs(row[c].x - row[c-1].x)
			ydif := helpers.Abs(row[c].y - row[c-1].y)
			if xdif > 0 {
				xdir := xdif / (row[c].x - row[c-1].x)
				for x := 0; x <= xdif; x++ {
					if _, ok := caveMap[fmt.Sprintf("%d,%d", row[c-1].x+(x*xdir), row[c-1].y)]; !ok {
						caveMap[fmt.Sprintf("%d,%d", row[c-1].x+(x*xdir), row[c-1].y)] = true
					}

				}
			}
			if ydif > 0 {
				ydir := ydif / (row[c].y - row[c-1].y)
				for y := 0; y <= ydif; y++ {
					if _, ok := caveMap[fmt.Sprintf("%d,%d", row[c-1].x, row[c-1].y+(y*ydir))]; !ok {
						caveMap[fmt.Sprintf("%d,%d", row[c-1].x, row[c-1].y+(y*ydir))] = true
					}
				}
			}
		}
	}

	return caveMap
}

func fillSand(caveMap map[string]bool, boundary Bounds, start Point) int {
	counter := 0
	sandMap := map[string]bool{}
out:
	for true {
		sand := start
		//fmt.Println("sand:", sand)
		itr := 0
		// loop till sand stops
		for true {
			//fmt.Println(counter, itr, sand)
			_, caveok := caveMap[fmt.Sprintf("%d,%d", sand.x, sand.y+1)]
			_, sandok := sandMap[fmt.Sprintf("%d,%d", sand.x, sand.y+1)]
			//fmt.Println(caveok, sandok)
			if caveok || sandok { // if sand hits something below it
				// try to move down left
				_, caveok = caveMap[fmt.Sprintf("%d,%d", sand.x-1, sand.y+1)]
				_, sandok = sandMap[fmt.Sprintf("%d,%d", sand.x-1, sand.y+1)]
				if !caveok && !sandok {
					sand = Point{sand.x - 1, sand.y + 1} // if there is no sand or cave below the left down position move there
				} else { // if there is sand in the left down position try to move right down
					_, caveok = caveMap[fmt.Sprintf("%d,%d", sand.x+1, sand.y+1)]
					_, sandok = sandMap[fmt.Sprintf("%d,%d", sand.x+1, sand.y+1)]
					if !caveok && !sandok { // if there is no sand or cave below the right down position move there
						sand = Point{sand.x + 1, sand.y + 1}
					} else { //if there is sand to the right and left of the current sand then stop and add the current sand position to the sandMap
						sandMap[fmt.Sprintf("%d,%d", sand.x, sand.y)] = true
						break
					}
				}
			} else { // if there is nothing below the sand move it down and continue
				sand = Point{sand.x, sand.y + 1}
				if sand.x < boundary.xmin || sand.x > boundary.xmax || sand.y > boundary.ymax {
					break out
				}
			}
			itr++ //increment the itr
		}
		counter++ //increment the counter and drop the next grain of sand
		//fmt.Println(sandMap)
	}
	return counter
}

func fillSand2(caveMap map[string]bool, boundary Bounds, start Point) int {
	counter := 0
	sandMap := map[string]bool{}
out:
	for true {
		sand := start
		//fmt.Println("sand:", sand)
		itr := 0
		// loop till sand stops
		for true {
			// if sand hits top then break out
			if _, ok := sandMap[fmt.Sprintf("%d,%d", start.x, start.y)]; ok {
				break out
			}
			//fmt.Println(counter, itr, sand)
			_, caveok := caveMap[fmt.Sprintf("%d,%d", sand.x, sand.y+1)]
			_, sandok := sandMap[fmt.Sprintf("%d,%d", sand.x, sand.y+1)]
			//fmt.Println(caveok, sandok)
			if caveok || sandok { // if sand hits something below it
				// try to move down left
				_, caveok = caveMap[fmt.Sprintf("%d,%d", sand.x-1, sand.y+1)]
				_, sandok = sandMap[fmt.Sprintf("%d,%d", sand.x-1, sand.y+1)]
				if !caveok && !sandok {
					sand = Point{sand.x - 1, sand.y + 1} // if there is no sand or cave below the left down position move there
				} else { // if there is sand in the left down position try to move right down
					_, caveok = caveMap[fmt.Sprintf("%d,%d", sand.x+1, sand.y+1)]
					_, sandok = sandMap[fmt.Sprintf("%d,%d", sand.x+1, sand.y+1)]
					if !caveok && !sandok { // if there is no sand or cave below the right down position move there
						sand = Point{sand.x + 1, sand.y + 1}
					} else { //if there is sand to the right and left of the current sand then stop and add the current sand position to the sandMap
						sandMap[fmt.Sprintf("%d,%d", sand.x, sand.y)] = true
						break
					}
				}
			} else { // if there is nothing below the sand move it down and continue
				//stop stand at bottom floor
				if sand.y == boundary.ymax+1 {
					sandMap[fmt.Sprintf("%d,%d", sand.x, sand.y)] = true
					break
				}
				sand = Point{sand.x, sand.y + 1}
			}
			itr++ //increment the itr
		}
		counter++ //increment the counter and drop the next grain of sand
		//fmt.Println(sandMap)
	}
	return counter
}
