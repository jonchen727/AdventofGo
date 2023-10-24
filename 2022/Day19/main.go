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
	ans := 0
	blueprints := parseInput(input)
	for _, blueprint := range blueprints {
		maxSpend := getMaxSpend(blueprint)
		maxGeodes := getMaxGeode(blueprint, 24, maxSpend)
		ans += maxGeodes * blueprint.Number
	}
	return ans
}

func part2(input string) int {
	ans := 1
	blueprints := parseInput(input)
	for i := 0; i <= 2; i++ {
		maxSpend := getMaxSpend(blueprints[i])
		maxGeodes := getMaxGeode(blueprints[i], 32, maxSpend)
		ans *= maxGeodes
	}
	return ans
}

func getMaxGeode(blueprint Blueprint, duration int, maxSpend []int) int {
	maxGeode := 0
	queue := []State{{0, 1, 0, 0, 0, 0, 0, 0, 0}}
	visited := map[string]bool{}
	for len(queue) > 0 {
		currentState := queue[0]
		queue = queue[1:]
		//fmt.Println(currentState)
		if currentState.Time == duration {
			continue
		}

		newStates := generateStates(blueprint, currentState, maxSpend)
		for _, state := range newStates {
			key := fmt.Sprintf("%d,%d,%d,%d,%d,%d,%d,%d,%d", state.Time, state.OreRobot, state.ClayRobot, state.ObsidianRobot, state.GeodeRobot, state.Ore, state.Clay, state.Obsidian, state.Geode)
			if state.Geode < maxGeode-1 {
				continue
			}
			if _, ok := visited[key]; !ok {
				visited[key] = true
				maxGeode = helpers.MaxInt(maxGeode, state.Geode)
				queue = append(queue, state)
			}
		}
	}
	return maxGeode
}

func generateStates(blueprint Blueprint, current State, maxSpend []int) []State {
	states := []State{}
	ore := current.Ore
	clay := current.Clay
	obsidian := current.Obsidian

	rmaxGeode := calculateMaxRobots(blueprint.Geode, ore, clay, obsidian)
	rmaxGeode = helpers.MinInt(rmaxGeode, 1)
	for rGeode := 0; rGeode <= rmaxGeode; rGeode++ {
		ore1, clay1, obsidian1 := updateResources(blueprint.Geode, ore, clay, obsidian, rGeode)
		rmaxObsidian := 0
		if current.ObsidianRobot <= maxSpend[2] {
			rmaxObsidian = calculateMaxRobots(blueprint.Obsidian, ore1, clay1, obsidian1)
			rmaxObsidian = helpers.MinInt(rmaxObsidian, 1)
		}
		for rObsidian := 0; rObsidian <= rmaxObsidian; rObsidian++ {
			ore2, clay2, obsidian2 := updateResources(blueprint.Obsidian, ore1, clay1, obsidian1, rObsidian)
			rmaxClay := 0

			if current.ClayRobot <= maxSpend[1] {
				rmaxClay = calculateMaxRobots(blueprint.Clay, ore2, clay2, obsidian2)
				rmaxClay = helpers.MinInt(rmaxClay, 1)
			}
			for rClay := 0; rClay <= rmaxClay; rClay++ {
				ore3, clay3, obsidian3 := updateResources(blueprint.Clay, ore2, clay2, obsidian2, rClay)
				rmaxOre := 0
				if current.OreRobot <= maxSpend[0] {
					rmaxOre = calculateMaxRobots(blueprint.Ore, ore3, clay3, obsidian3)
					rmaxOre = helpers.MinInt(rmaxOre, 1)
				}
				for rOre := 0; rOre <= rmaxOre; rOre++ {
					ore4, clay4, obsidian4 := updateResources(blueprint.Ore, ore3, clay3, obsidian3, rOre)
					ore4 = helpers.MinInt(ore4, maxSpend[0]+1)
					clay4 = helpers.MinInt(clay4, maxSpend[1]+1)
					obsidian4 = helpers.MinInt(obsidian4, maxSpend[2]+1)
					if rOre+rClay+rObsidian+rGeode > 1 {
						continue
					}
					newState := State{current.Time + 1, current.OreRobot + rOre, current.ClayRobot + rClay, current.ObsidianRobot + rObsidian, current.GeodeRobot + rGeode, current.OreRobot + ore4, current.ClayRobot + clay4, current.ObsidianRobot + obsidian4, current.GeodeRobot + current.Geode}
					// if newState.Geode > 12 {
					// 	fmt.Println(newState)
					// }
					states = append(states, newState)
				}
			}
		}
	}
	return states
}

func updateResources(robot Robot, ore int, clay int, obsidian int, num int) (int, int, int) {
	ore -= (robot.Ore * num)
	clay -= (robot.Clay * num)
	obsidian -= (robot.Obsidian * num)
	return ore, clay, obsidian
}

func calculateMaxRobots(robot Robot, ore int, clay int, obsidian int) int {
	max := math.MaxInt64

	if robot.Ore != 0 {
		max = helpers.MinInt(ore / robot.Ore)
	}
	if robot.Clay != 0 {
		max = helpers.MinInt(max, clay/robot.Clay)
	}
	if robot.Obsidian != 0 {
		max = helpers.MinInt(max, obsidian/robot.Obsidian)
	}
	return max

}

func getMaxSpend(blueprint Blueprint) []int {
	max := []int{}
	max = append(max, helpers.MaxInt(blueprint.Ore.Ore, blueprint.Clay.Ore, blueprint.Obsidian.Ore, blueprint.Geode.Ore))
	max = append(max, helpers.MaxInt(blueprint.Ore.Clay, blueprint.Clay.Clay, blueprint.Obsidian.Clay, blueprint.Geode.Clay))
	max = append(max, helpers.MaxInt(blueprint.Ore.Obsidian, blueprint.Clay.Obsidian, blueprint.Obsidian.Obsidian, blueprint.Geode.Obsidian))
	return max
}

type State struct {
	Time          int
	OreRobot      int
	ClayRobot     int
	ObsidianRobot int
	GeodeRobot    int
	Ore           int
	Clay          int
	Obsidian      int
	Geode         int
}
type Blueprint struct {
	Number   int
	Ore      Robot
	Clay     Robot
	Obsidian Robot
	Geode    Robot
}

type Robot struct {
	Ore      int
	Clay     int
	Obsidian int
}

func parseInput(input string) []Blueprint {
	blueprints := []Blueprint{}
	for _, line := range strings.Split(input, "\n") {
		blueprint := Blueprint{}
		ore := Robot{}
		clay := Robot{}
		obsidian := Robot{}
		geode := Robot{}
		split := strings.Split(line, ": ")
		fmt.Sscanf(split[0], "Blueprint %d", &blueprint.Number)
		fmt.Sscanf(split[1], "Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &ore.Ore, &clay.Ore, &obsidian.Ore, &obsidian.Clay, &geode.Ore, &geode.Obsidian)

		blueprint.Ore = ore
		blueprint.Clay = clay
		blueprint.Obsidian = obsidian
		blueprint.Geode = geode
		blueprints = append(blueprints, blueprint)
	}
	return blueprints
}
