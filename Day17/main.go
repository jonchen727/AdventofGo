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
	jets := parseInput(input)
	ans = simulateBlocks(jets, 2022)
	return ans
}

func part2(input string) int {
	ans := 0
	jets := parseInput(input)
	ans = simulateBlocks(jets, 1000000000000 )
	return ans
}

func parseInput(input string) []complex128 {
	jets := []complex128{}
	for _, val := range strings.Split(input, "") {
		if val == ">" {
			jets = append(jets, complex(float64(1), 0))
			//fmt.Print(val)
		} else {
			jets = append(jets, complex(float64(-1), 0))
			//fmt.Print(val)
		}
	}

	return jets

}

func simulateBlocks(jets []complex128, step int) int {
	rocks := [][]complex128{
		{0, 1, 2, 3},
		{1, 1i, 1 + 1i, 2 + 1i, 1 + 2i},
		{0, 1, 2, 2 + 1i, 2 + 2i},
		{0, 1i, 2i, 3i},
		{0, 1, 1i, 1 + 1i},
	}

	solid := []complex128{}
	for x := 0; x < 7; x++ {
		block := complex(float64(x), -1)
		solid = append(solid, block)
	}

	//fmt.Println(solid)

	// for _, rock := range rocks {
	// 	fmt.Println(rock)
	// }

	rockCount := 0
	rockIndex := 0
	height := 0
	rock := moveRock(rocks[rockIndex], complex(float64(2), float64(height+3)))
	jetidx := 0
	jetcount := 0
	offset := 0
	a := []int{}
	seen := map[string][]int{}

	for rockCount < step {

		moved := []complex128{}
		for true {

			jetidx = (jetcount) % len(jets)
			//fmt.Println(jetidx, rockIndex)
			jet := jets[jetidx]
			moved = moveRock(rock, jet)


			cond1 := true
			for _, val := range moved {
				if real(val) >= 0 && real(val) < 7 {
				} else {
					cond1 = false
				}
			}
			if !Intersection(moved, solid) && cond1 {
				rock = moved
			}
			moved = moveRock(rock, (complex(float64(0), float64(-1))))
			if Intersection(moved, solid) {
				var summary string
				o := height 
				solid, height, summary = Union(solid, rock)
				rockCount++
				a = append(a, height-o)
				if rockCount >= step {
					break
				}
				rockIndex = (rockIndex + 1) % len(rocks)
				rock = moveRock(rocks[rockIndex], complex(float64(2), float64(height+3)))

				key := fmt.Sprintf("%d, %d, %s", jetidx, rockIndex, summary)
				//fmt.Println(key)


				if _, ok := seen[key]; ok {
					fmt.Println("cache hit")
					lastrockCount, lastHeight := seen[key][0], seen[key][1]
					rem := step - rockCount
					rep := rem / (rockCount - lastrockCount)
					offset = rep * (height - lastHeight)
					rockCount += rep * (rockCount - lastrockCount)
					seen = map[string][]int{}
				}
				seen[key] = []int{rockCount, height}

			} else {
				rock = moved
			}
			jetcount++
		}

		//fmt.Println(solid)
	}
	//fmt.Println(solid)
	return height + offset
}


func Union(arr1 []complex128, arr2 []complex128) ([]complex128, int, string) {
	union := map[complex128]bool{}
	maxheight := float64(0)
  topState := []int{}
	for i := 0; i < 7; i++ {
		topState = append(topState, -20)
	}
	for _, val := range arr1 {
		union[val] = true
	}
	for _, val := range arr2 {
		union[val] = true
	}

	unionArr := []complex128{}
	for key, _ := range union {
		real := real(key)
		imag := imag(key)
		topState[int(real)] = helpers.MaxInt(topState[int(real)], int(imag))
		unionArr = append(unionArr, key)
		if imag >= maxheight {
			maxheight = imag
		}
	}
	top := helpers.MaxInt(topState...)

	var summary string
	for i, col := range topState {
		topState[i] = col - top
		summary += fmt.Sprintf("%d,", col-top)
	}

	//fmt.Println(maxheight)
	return unionArr, int(maxheight)+1, summary //, int(maxheight)+1

}

func Intersection(arr1 []complex128, arr2 []complex128) bool {
	intersection := map[complex128]bool{}
	for _, val := range arr1 {
		intersection[val] = true
		//fmt.Println("Intersection", val, intersection)
	}
	for _, val := range arr2 {
		if _, ok := intersection[val]; ok {
			//fmt.Println("Intersection", val, intersection)
			return true
		}
	}

	return false
}

func moveRock(rock []complex128, add complex128) []complex128 {
	moved := []complex128{}
	for _, val := range rock {
		val = val + add
		moved = append(moved, val)
	}

	return moved
}
