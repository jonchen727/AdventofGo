package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"strings"
	//"reflect"
	"flag"
	//"strconv"
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

type Monkey struct {
	name    string
	value   int
	monkey1 *Monkey
	monkey2 *Monkey
	op      string
}

func (m *Monkey) getValue(itr int, part int) int {

	if m.name == "root" && part == 2 {

		m.monkey1.getValue(itr, part)
		m.monkey2.getValue(itr, part)
		return 1

	}

	if m.value > 0 {
		return m.value
	}

	if itr > 0 && (m.monkey1.name == "humn" || m.monkey2.name == "humn" ) && part == 2 {
		m.value = -1
		return m.value
	}
	switch m.op {
	case "+":
		itr++

		m1 := m.monkey1.getValue(itr, part)
		m2 := m.monkey2.getValue(itr, part)
		if checkNeg(m1, m2) {
			m.value = helpers.MinInt(m1, m2) - 1
		} else {
			m.value = m.monkey1.getValue(itr, part) + m.monkey2.getValue(itr, part)
		}
		return m.value
	case "*":
		itr++
		m1 := m.monkey1.getValue(itr, part)
		m2 := m.monkey2.getValue(itr, part)
		if checkNeg(m1, m2) {
			m.value = helpers.MinInt(m1, m2) - 1
		} else {
			m.value = m.monkey1.getValue(itr, part) * m.monkey2.getValue(itr, part)
		}
		return m.value
	case "/":
		itr++
		m1 := m.monkey1.getValue(itr, part)
		m2 := m.monkey2.getValue(itr, part)
		if checkNeg(m1, m2) {
			m.value = helpers.MinInt(m1, m2) - 1
		} else {
			m.value = m.monkey1.getValue(itr, part) / m.monkey2.getValue(itr, part)
		}
		return m.value
	case "-":
		itr++
		m1 := m.monkey1.getValue(itr, part)
		m2 := m.monkey2.getValue(itr, part)
		if checkNeg(m1, m2) {
			m.value = helpers.MinInt(m1, m2) - 1
		} else {
			m.value = m.monkey1.getValue(itr, part) - m.monkey2.getValue(itr, part)
		}
		return m.value

	default:
		//fmt.Println("default")
		return m.value
	}
}

func checkNeg(n1, n2 int) bool {
	if n1 < 0 || n2 < 0 {
		return true
	}
	return false
}

func part1(input string) int {
	monkeys := parseInput(input)
	root, rootExists := monkeys["root"]
	if !rootExists {
		fmt.Println("Root node not found.")
		panic("Root node not found.")
	}

	itr := 0
	ans := root.getValue(itr, 1)
	fmt.Println(monkeys["root"].monkey2)
	fmt.Println(monkeys["root"].monkey1)
	return ans
}

func part2(input string) int {

	monkeys := parseInput(input)
	monkeys["root"].op = "/"
	monkeys["root"].value = 1
	monkeys["humn"].value = 0
	monkeys["humn"].op = "?"

	root, rootExists := monkeys["root"]
	if !rootExists {
		fmt.Println("Root node not found.")
		panic("Root node not found.")
	}
	itr := 0
	root.getValue(itr, 2)
	

	//fmt.Println(monkeys["root"].monkey2)
	//fmt.Println(monkeys["root"].monkey1)

	maxct := 0
	nodemap := map[int][]*Monkey{}
	for _, v := range monkeys {
		if v.value < 1 {
			nodemap[v.value] = append(nodemap[v.value], v)
			maxct = helpers.MinInt(maxct, v.value)
		}
	}
	root.fillMissingVals()
	for _, val := range nodemap[0] {
		val.getValue(itr, 2)
	}
	for i := maxct; i < 0; i++ {
		nodes := nodemap[i]
		for j := 0; j < len(nodes); j++ {
			nodes[j].fillMissingVals()
			//fmt.Println(nodes[j])
		}
	}
	// for _, v := range nodemap {
	// 	for j := 0; j < len(v); j++ {
	// 		fmt.Println(v[j], v[j].monkey1, v[j].monkey2)
	// 	}
	// }

	for _, value := range monkeys {
		if value.value < 0 {
			fmt.Println("you done fucked up")
		}
	}
	return monkeys["humn"].value
}
func (m *Monkey) fillMissingVals() {
	// fmt.Println(m)
	// fmt.Println(m.monkey1, m.monkey2)
	var missing1 bool
	if m.monkey1.value <= 0 {
		missing1 = true
	} else { 
		missing1 = false
		}
		
	

	switch m.op {
	case "+":
		// num = n1 + n2
		// n1 = num - n2
		// n2 = num - n1

		if missing1 {
			// fmt.Println("x=", m.value, "-", m.monkey2.value)
			m.monkey1.value = m.value - m.monkey2.value
		} else {
			// fmt.Println("x=", m.value, "-", m.monkey1.value)
			m.monkey2.value = m.value - m.monkey1.value
		}
	case "*":
		// num = n1 * n2
		// n1 = num / n2
		// n2 = num / n1
		if missing1 {
			// fmt.Println("x=", m.value, "/", m.monkey2.value)
			m.monkey1.value = m.value / m.monkey2.value
		} else {
			// fmt.Println("x=", m.value, "/", m.monkey1.value)
			m.monkey2.value = m.value / m.monkey1.value
		}
	case "/":
		// num = n1 / n2
		// n1 = num * n2
		// n2 = n1/num
		if missing1 {
			// fmt.Println("x=", m.value, "*", m.monkey2.value)
			m.monkey1.value = m.value * m.monkey2.value
		} else {

			// fmt.Println("x=", m.value, "*", m.monkey1.value)
			m.monkey2.value = m.monkey1.value / m.value
		}

	case "-":
		// num = n1 - n2
		// n1 = num + n2
		// n2 = n1 - num
		if missing1 {
			// fmt.Println("x=", m.value, "+", m.monkey2.value)
			m.monkey1.value = m.value + m.monkey2.value
		} else {
			// fmt.Println("x=", m.value, "+", m.monkey1.value)
			m.monkey2.value = m.monkey1.value - m.value
		}
	}
}

func parseInput(input string) map[string]*Monkey {
	monkeys := make(map[string]*Monkey)
	tmp := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, ": ")
		name := split[0]
		data := strings.Split(split[1], " ")
		if len(data) == 1 {
			val := helpers.ToInt(data[0])
			m := Monkey{name: name, value: val}
			monkeys[name] = &m
		} else {
			tmp[name] = data
			monkeys[name] = &Monkey{name: name}
		}
	}

	for k, v := range tmp {
		monkey := monkeys[k]
		monkey1 := monkeys[v[0]]
		monkey2 := monkeys[v[2]]
		op := v[1]
		monkey.monkey1 = monkey1
		monkey.monkey2 = monkey2
		monkey.op = op
		monkeys[k] = monkey
	}

	return monkeys
}
