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
	commands := parseInput(input)
	initXreg := 1
	states := generateState(commands, initXreg)
	ans := sumSigStr([]int{20, 60, 100, 140, 180, 220}, states)
	//fmt.Println(states)
	return ans
}

func part2(input string, links int) int {
	commands := parseInput(input)
	initXreg := 1
	states := generateState(commands, initXreg)
	drawsprites(states)
	ans := 0

	return ans
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans

}

type State struct {
	clock          int
	inst           string
	value          int
	instnum        int
	sXreg          int
	dXreg          int
	aXreg          int
	SignalStrength int
}

func generateState(input []string, initXreg int) []State {
	state := []State{}
	clock := 0
	cycles := map[string]int{
		"noop": 1,
		"addx": 2,
	}
	for i, line := range input {
		if strings.HasPrefix(line, "noop") {
			//fmt.Println(i, "noop command")
			for j := 0; j < cycles["noop"]; j++ {
				clock++
				state = append(state, State{
					clock:          clock,
					inst:           "noop",
					value:          0,
					instnum:        i,
					sXreg:          initXreg,
					dXreg:          initXreg,
					aXreg:          initXreg,
					SignalStrength: initXreg * (clock),
				})
				printlaststate(state)
			}
		}
		if strings.HasPrefix(line, "addx") {
			//fmt.Println(i, "addx command", "value:",strings.Split(line, " ")[1])
			value := helpers.ToInt(strings.Split(line, " ")[1])
			for j := 0; j < cycles["addx"]; j++ {
				clock++
				oldXreg := initXreg
				if j == cycles["addx"]-1 {
					initXreg = initXreg + value
				}
				state = append(state, State{
					clock:          clock,
					inst:           "addx",
					value:          value,
					instnum:        i,
					sXreg:          oldXreg,
					dXreg:          oldXreg,
					aXreg:          initXreg,
					SignalStrength: oldXreg * (clock),
				})
				printlaststate(state)
			}

		}

	}
	//fmt.Println(state)
	return state
}

func sumSigStr(indexs []int, states []State) int {
	ans := 0
	for _, index := range indexs {
		//fmt.Println("index:", index, "SignalStrength:", states[index-1].SignalStrength)
		ans += states[index-1].SignalStrength
	}
	return ans
}

func printlaststate(states []State) {
	//state := states[len(states)-1]
	//fmt.Println("clock:", state.clock, "inst:", state.inst, "value:", state.value, "instnum:", state.instnum, "Start Xreg:", state.sXreg, "During Xreg:", state.dXreg, "After Xreg:", state.aXreg, "SignalStrength:", state.SignalStrength)
}

func drawsprites(states []State) {
	row := 0
	rows := 6
	rowlen := len(states) / rows
	for i, clock := range states {
		rowold := row
		row = i / rowlen
		sposition := clock.dXreg

		if rowold != row {
			fmt.Println("")
		}

		if i%rowlen >= sposition-1 && i%rowlen <= sposition+1 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}

	}

}
