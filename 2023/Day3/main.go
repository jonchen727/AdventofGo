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
	arr, parts, _ := parseInput(input)
	for _, part := range parts {
		coordinates := [][]int{}
		left := []int{part.minCol - 1, part.row}
		right := []int{part.maxCol + 1, part.row}
		top := [][]int{}
		bottom := [][]int{}
		for i := part.minCol - 1; i <= part.maxCol+1; i++ {
			top = append(top, []int{i, part.row - 1})
			bottom = append(bottom, []int{i, part.row + 1})
		}
		coordinates = append(coordinates, []int{left[0], left[1]})
		coordinates = append(coordinates, []int{right[0], right[1]})
		coordinates = append(coordinates, top...)
		coordinates = append(coordinates, bottom...)

		for _, coordinate := range coordinates {
			if coordinate[0] < 0 || coordinate[1] < 0 {
				continue
			}
			if coordinate[0] >= len(arr[0]) || coordinate[1] >= len(arr) {
				continue
			}
			_, err := strconv.Atoi(string(arr[coordinate[1]][coordinate[0]]))
			if err == nil || string(arr[coordinate[1]][coordinate[0]]) == "." {
				continue
			}
			if err != nil {
				ans += part.number
				//fmt.Println(string(arr[coordinate[1]][coordinate[0]]))
				//fmt.Println(part.number, "[", coordinate[0], ",", coordinate[1], "]", coordinates)
				break
			}
		}
	}
	return ans
}

func part2(input string) int {
	ans := 0
	_, parts, gearMap := parseInput(input)
	partMap := map[string]int{}
	for idx, part := range parts {
		for i := part.minCol; i <= part.maxCol; i++ {
			partMap[strconv.Itoa(i)+","+strconv.Itoa(part.row)] = idx
		}
	}
	Gears := []Gear{}
	for key, _ := range gearMap {
		Gear := Gear{}
		counter := 0
		seen := map[int]bool{}
		//build 3x3 grid around gear
		coordinates := [][]int{}
		left := []int{helpers.ToInt(strings.Split(key, ",")[0]) - 1, helpers.ToInt(strings.Split(key, ",")[1])}
		right := []int{helpers.ToInt(strings.Split(key, ",")[0]) + 1, helpers.ToInt(strings.Split(key, ",")[1])}
		top := [][]int{}
		bottom := [][]int{}
		for i := helpers.ToInt(strings.Split(key, ",")[0]) - 1; i <= helpers.ToInt(strings.Split(key, ",")[0])+1; i++ {
			top = append(top, []int{i, helpers.ToInt(strings.Split(key, ",")[1]) - 1})
			bottom = append(bottom, []int{i, helpers.ToInt(strings.Split(key, ",")[1]) + 1})
		}
		coordinates = append(coordinates, []int{left[0], left[1]})
		coordinates = append(coordinates, []int{right[0], right[1]})
		coordinates = append(coordinates, top...)
		coordinates = append(coordinates, bottom...)
		//fmt.Println("Gear:", key, "Coordinates:", coordinates)
		for j, coordinate := range coordinates {
			_, ok := partMap[strconv.Itoa(coordinate[0])+","+strconv.Itoa(coordinate[1])]
			if ok {
				//fmt.Println("found part")
				if _, ok := seen[partMap[strconv.Itoa(coordinate[0])+","+strconv.Itoa(coordinate[1])]]; ok {
					//fmt.Println("already seen")
					if j == len(coordinates)-1 && counter == 2 {
						Gears = append(Gears, Gear)
					}
					continue
				}
				Gear.parts = append(Gear.parts, parts[partMap[strconv.Itoa(coordinate[0])+","+strconv.Itoa(coordinate[1])]])
				seen[partMap[strconv.Itoa(coordinate[0])+","+strconv.Itoa(coordinate[1])]] = true
				counter++
			}
			//fmt.Println("counter:", counter)
			if j == len(coordinates)-1 && counter == 2 {
				Gears = append(Gears, Gear)
				//fmt.Println(Gears)
			}
		}
	}
	for _, gear := range Gears {
		ratio := 1
		for _, part := range gear.parts {
			ratio *= part.number
		}
		ans += ratio
	}
	//fmt.Println(partMap)
	return ans
}

type Part struct {
	number int
	row    int
	minCol int
	maxCol int
}

type Gear struct {
	parts []Part
}

func parseInput(input string) ([]string, []Part, map[string]bool) {
	lines := strings.Split(input, "\n")
	arr := []string{}
	parts := []Part{}
	gearMap := map[string]bool{}
	for i, line := range lines {
		arr = append(arr, line)
		var ram string
		start := 9999
		end := 0
		toggle := 0
		for j, char := range line {
			if _, err := strconv.Atoi(string(char)); err != nil {
				if string(char) == "*" {
					gearMap[strconv.Itoa(j)+","+strconv.Itoa(i)] = true
				}

				if toggle == 1 {
					end = j - 1
					toggle = 0
					part := Part{
						number: helpers.ToInt(ram),
						row:    i,
						minCol: start,
						maxCol: end,
					}
					parts = append(parts, part)
					ram = ""
					continue
				} else {
					continue
				}
			} else {
				if toggle == 1 && j == len(line)-1 {
					ram += string(char)
					end = j
					toggle = 0
					part := Part{
						number: helpers.ToInt(ram),
						row:    i,
						minCol: start,
						maxCol: end,
					}
					parts = append(parts, part)
				}

				if toggle == 0 {
					start = j
					ram += string(char)
					toggle = 1
				} else {
					ram += string(char)
				}
			}
		}
	}
	return arr, parts, gearMap
}
