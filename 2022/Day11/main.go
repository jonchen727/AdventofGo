package main

import (
	_ "embed"
	"fmt"
	"slices"
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
	monkeys := parseInput(input)
	worry := false
	rounds := 20
	divisor := getdivisor(monkeys)
	ans := monkeybusiness(monkeys, worry, rounds, divisor)
	return ans
}

func part2(input string, links int) int {
	monkeys := parseInput(input)
	worry := true
	rounds := 10000
	divisor := getdivisor(monkeys)
	ans := monkeybusiness(monkeys, worry, rounds, divisor)
	fmt.Println(monkeys)
	return ans
}

func monkeybusiness(monkeys []Monkey, worry bool, rounds int, divisor int) int {
	for i := 0; i < rounds; i++ {
		for i := range monkeys {
			monkeys[i].DoOperation(monkeys, worry, divisor)
		}
	}
	inspections := []int{}
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspected)
	}
	slices.Sort(inspections)
	fmt.Println(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]

}
func parseInput(input string) []Monkey {
	monkeys := []Monkey{}
	for _, v := range strings.Split(input, "\n\n") {
		monkey := Monkey{}
		//var items []int
		for _, line := range strings.Split(v, "\n") {
			split := strings.Split(line, ":")
			prefix := strings.ReplaceAll(split[0], "  ", "")
			switch prefix {
			case "Starting items":
				for _, item := range strings.Split(split[1][1:], ", ") {
					monkey.items = append(monkey.items, helpers.ToInt(item))
				}
			case "Operation":
				_, err := fmt.Sscanf(split[1], " new = old %s %d", &monkey.operation.operation, &monkey.operation.amount)
				if err != nil {
					fmt.Println(err)
				}
			case "Test":
				_, err := fmt.Sscanf(split[1], " divisible by %d", &monkey.test.divisor)
				if err != nil {
					fmt.Println(err)
				}
			case "If true":
				_, err := fmt.Sscanf(split[1], " throw to monkey %d", &monkey.test.truemonkey)
				if err != nil {
					fmt.Println(err)
				}
			case "If false":
				_, err := fmt.Sscanf(split[1], " throw to monkey %d", &monkey.test.falsemonkey)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		//monkey.items = helpers.SliceToInt(strings.Split(items, ","))
		monkey.inspected = 0
		monkeys = append(monkeys, monkey)
	}
	return monkeys

}

func getdivisor(monkeys []Monkey) int {
	divisor := 1
	for _, monkey := range monkeys {
		divisor *= monkey.test.divisor
	}
	return divisor
}

type Monkey struct {
	items     []int
	operation Operation
	test      Test
	inspected int
}

type Test struct {
	divisor     int
	truemonkey  int
	falsemonkey int
}

type Operation struct {
	operation string
	amount    int
}

func (m *Monkey) DoOperation(monkeys []Monkey, worry bool, divisor int) {
	for i := range m.items {
		//fmt.Println("monkey:", i, "items:", m.items)
		switch m.operation.operation {
		case "+":
			m.items[i] += m.operation.amount
		case "-":
			m.items[i] -= m.operation.amount
		case "*":
			if m.operation.amount == 0 {
				m.items[i] *= m.items[i]
			} else {
				m.items[i] *= m.operation.amount
			}
		case "/":
			m.items[i] /= m.operation.amount
		}
		if worry == false {
			m.items[i] = decreaseWorry(m.items[i])
		}
		m.inspected += 1
	}
	m.Test(monkeys, worry, divisor)
}

func (m *Monkey) Test(monkeys []Monkey, worry bool, divisor int) {

	for _, v := range m.items {
		v %= divisor
		if v%m.test.divisor == 0 {
			monkeys[m.test.truemonkey].items = append(monkeys[m.test.truemonkey].items, v)
			if len(m.items) <= 1 {
				m.items = []int{}
			} else {
				m.items = m.items[1:]

			}
			//fmt.Println(m.items)

		} else {
			monkeys[m.test.falsemonkey].items = append(monkeys[m.test.falsemonkey].items, v)
			if len(m.items) <= 1 {
				m.items = []int{}
			} else {
				m.items = m.items[1:]

			}
		}
	}

}

func decreaseWorry(item int) int {
	return item / 3
}
