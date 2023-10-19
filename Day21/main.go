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

type Monkey struct {
	name string
	value int 
	monkey1 *Monkey
	monkey2 *Monkey
	op string
}

func (m Monkey) getValue(monkeys map[string]Monkey) int {
	if m.value != 0 {
		return m.value
	}
	switch m.op {
	case "+":
		return m.monkey1.getValue(monkeys) + m.monkey2.getValue(monkeys)
	case "*":
		return m.monkey1.getValue(monkeys) * m.monkey2.getValue(monkeys)
	case "/":
		return m.monkey1.getValue(monkeys) / m.monkey2.getValue(monkeys)
	case "-":
		return m.monkey1.getValue(monkeys) - m.monkey2.getValue(monkeys)
	default:
		return m.value
	}
}

func part1(input string) int {
	monkeys := parseInput(input)
	root, rootExists := monkeys["root"]
	if !rootExists {
		fmt.Println("Root node not found.")
		panic("Root node not found.")
	}
	ans := root.getValue(monkeys)
	return ans
}

func part2(input string) int {
  ans := 0
	return ans
}

func parseInput(input string) map[string]Monkey {
	monkeys := make(map[string]Monkey)
	tmp := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, ": ")
		name := split[0]
		data := strings.Split(split[1], " ")
		if len(data) == 1 {
			m := Monkey{name: name, value: helpers.ToInt(data[0])}
			monkeys[name] = m
		} else {
			tmp[name] = data
			monkeys[name] = Monkey{name: name}
		}
	}

	for k, v := range tmp {
		monkey := monkeys[k]
		monkey1 := monkeys[v[0]]
		monkey2 := monkeys[v[2]]
		op := v[1]
		monkey.monkey1 = &monkey1
		monkey.monkey2 = &monkey2
		monkey.op = op
		monkeys[k] = monkey
	}

	return monkeys
}
