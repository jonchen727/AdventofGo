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
	ans := 1
	races := parseInput(input, 1)
	for _, race := range races {
		findWins(race)
		ans *= race.wins
	}
	return ans
}

func part2(input string) int {

	race := parseInput(input, 2)[0]
	findWins(race)
	ans := race.wins
	return ans
}

type Race struct {
	time     float64
	distance float64
	wins     int
}

func findWins(race *Race) {
	ways := 0
	for i := 0; i < helpers.ToInt(race.time); i++ {
		elapsed := helpers.ToFloat(i)
		if elapsed*(race.time-elapsed) > race.distance {
			ways++
		}
	}
	race.wins = ways
}

func parseInput(input string, part int) map[int]*Race {
	lines := strings.Split(input, "\n")
	races := map[int]*Race{} // Change to a map of pointers
	if part == 1 {
		for i, line := range lines {
			for strings.Contains(line, "  ") {
				line = strings.Replace(line, "  ", " ", -1)
			}
			line = strings.Split(line, ": ")[1]
			sets := strings.Split(line, " ")
			qty := len(sets)
			//fmt.Println(sets, qty)
			for j := 0; j < qty; j++ {
				if _, ok := races[j]; !ok {
					races[j] = &Race{} // Store a pointer to a new Race
				}
				race := races[j] // race is now a pointer
				//fmt.Println(helpers.ToFloat(sets[j]))
				switch i {
				case 0:
					race.time = helpers.ToFloat(string(sets[j]))
				case 1:
					race.distance = helpers.ToFloat(string(sets[j]))
				}
			}
		}
	}
	if part == 2 {
		for i, line := range lines {
			for strings.Contains(line, " ") {
				line = strings.Replace(line, " ", "", -1)
			}
			line = strings.Split(line, ":")[1]
			if _, ok := races[0]; !ok {
				races[0] = &Race{}
			}
			race := races[0]
			switch i {
			case 0:
				race.time = helpers.ToFloat(string(line))
			case 1:
				race.distance = helpers.ToFloat(string(line))
			}
		}
	}
	return races
}
