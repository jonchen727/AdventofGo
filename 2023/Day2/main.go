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
	games := parseInput(input)
	red := 12
	green := 13
	blue := 14
	for idx, game := range games {
		if game.maxRed <= red && game.maxGreen <= green && game.maxBlue <= blue {
			ans += idx + 1
		}
	}
	return ans
}

func part2(input string) int {
	ans := 0
	return ans
}

type Set struct {
	red   int
	green int
	blue  int
}

type Game struct {
	maxRed   int
	maxGreen int
	maxBlue  int
	set      []Set
}

func parseInput(input string) []Game {
	lines := strings.Split(input, "\n")
	Games := []Game{}
	for _, line := range lines {
		sets := strings.Split(line, ": ")[1]
		Game := Game{0, 0, 0, []Set{}}
		Sets := []Set{}
		for _, set := range strings.Split(sets, "; ") {
			Set := Set{0, 0, 0}
			for _, roll := range strings.Split(set, ", ") {
				var color string
				var num int
				_, err := fmt.Sscanf(roll, "%d %s", &num, &color)
				if err != nil {
					panic(err)
				}
				switch color {
				case "red":
					Set.red = num
					Game.maxRed = helpers.MaxInt(Game.maxRed, num)
				case "green":
					Set.green = num
					Game.maxGreen = helpers.MaxInt(Game.maxGreen, num)
				case "blue":
					Set.blue = num
					Game.maxBlue = helpers.MaxInt(Game.maxBlue, num)
				}
			}
			Sets = append(Sets, Set)
		}
		Game.set = Sets
		Games = append(Games, Game)
	}
	return Games
}
