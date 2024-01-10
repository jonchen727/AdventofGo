package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"reflect"
	"strings"
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
	ans := math.MaxInt64
	seeds, conversions := parseInput(input)
	seedList := buildSeeds(seeds, conversions)
	for _, seed := range seedList {
		ans = helpers.MinInt(ans, seed.Location)
	}
	return ans
}

func part2(input string) int {
	ans := math.MaxInt64
	seeds, conversions := parseInput(input)
	newSeeds := []newSeed{}
	for i := 0; i <= len(seeds)/2; {
		start := helpers.ToInt(seeds[i])
		span := helpers.ToInt(seeds[i+1])
		end := start + span
		seed := newSeed{start, end, span}
		newSeeds = append(newSeeds, seed)
		i = i + 2
	}
	//fmt.Println(conversions["humidity"])

	next := conversions["seed"]
	for true {
		new := []newSeed{}
		for len(newSeeds) > 0 {
			start := newSeeds[0].start
			end := newSeeds[0].end
			//fmt.Println(next.output)
			newSeeds = newSeeds[1:]
			matched := false
			for _, ranges := range next.ranges {
				sdest := ranges.destination
				ssource := ranges.source
				span := ranges.span
				os := helpers.MaxInt(start, ssource)
				oe := helpers.MinInt(end, ssource+span)
				if os < oe {
					new = append(new, newSeed{os - ssource + sdest, oe - ssource + sdest, 0})
					//fmt.Println("added os<oe", newSeed{os - ssource + sdest, oe - ssource + sdest, 0})
					matched = true
					if os > start {
						newSeeds = append(newSeeds, newSeed{start, os, 0})
						//fmt.Println("added os>s", newSeed{start, os, 0})
					}
					if end > oe {
						newSeeds = append(newSeeds, newSeed{oe, end, 0})
						//fmt.Println("added e>oe", newSeed{oe, end, 0})
					}
					break
				}
			}
			if !matched {
				new = append(new, newSeed{start, end, 0})
				//fmt.Println("added no match", newSeed{start, end, 0})
			}
		}
		newSeeds = new
		if next.next != nil {
			next = next.next
		} else {
			break
		}
	}
	for _, seed := range newSeeds {
		ans = helpers.MinInt(ans, seed.start)
	}
	//fmt.Println(newSeeds)
	return ans
}

type newSeed struct {
	start int
	end   int
	span  int
}

type Seed struct {
	Num         int
	Soil        int
	Fertilizer  int
	Water       int
	Light       int
	Temperature int
	Humidity    int
	Location    int
}

type Map struct {
	input    string
	output   string
	ranges   []Ranges
	previous *Map
	next     *Map
}

type Ranges struct {
	source      int
	destination int
	span        int
}

func buildSeeds(seeds []string, conversions map[string]*Map) []Seed {
	seedList := []Seed{}
	for _, seed := range seeds {
		Seed := Seed{}
		Seed.Num = helpers.ToInt(seed)

		//start filling in values
		Seed.Soil = convert(*conversions["seed"], Seed.Num)
		next := conversions["seed"].next
		start_val := Seed.Soil

		for next != nil {
			//fmt.Println(next.output)
			start_val = convert(*next, start_val)
			setField(&Seed, strings.Title(next.output), start_val)
			next = next.next
		}
		seedList = append(seedList, Seed)
	}
	return seedList
}

func convert(item Map, number int) int {
	for _, rng := range item.ranges {
		diff := number - rng.source
		if diff >= 0 && diff <= rng.span {
			return rng.destination + diff
		}
	}
	return number
}

func setField(obj interface{}, name string, value int) {
	reflectValue := reflect.ValueOf(obj)
	if reflectValue.Kind() != reflect.Ptr || reflectValue.IsNil() {
		fmt.Println("Object is not a pointer or is nil")
		return
	}

	reflectValue = reflectValue.Elem()
	fieldVal := reflectValue.FieldByName(name)

	if !fieldVal.IsValid() {
		fmt.Printf("Field %s not found\n", name)
		return
	}

	if !fieldVal.CanSet() {
		fmt.Printf("Field %s cannot be set\n", name)
		return
	}

	if fieldVal.Kind() != reflect.Int {
		fmt.Printf("Field %s is not an int, it is a %s\n", name, fieldVal.Kind())
		return
	}

	fieldVal.SetInt(int64(value))
}

func parseInput(input string) ([]string, map[string]*Map) {
	lines := strings.Split(input, "\n\n")
	seeds := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	//fmt.Println(seeds)
	conversions := make(map[string]*Map)
	for _, types := range lines[1:] {
		Map := Map{}
		split := strings.Split(types, ":\n")

		// fill in input to output
		title := strings.Replace(split[0], "-", " ", -1)
		_, ok := fmt.Sscanf(title, "%s to %s map", &Map.input, &Map.output)
		if ok != nil {
			panic("You suck at coding")
		}

		// fill in conversions
		for _, line := range strings.Split(split[1], "\n") {
			ranges := Ranges{}
			_, ok := fmt.Sscanf(line, "%d %d %d", &ranges.destination, &ranges.source, &ranges.span)
			if ok != nil {
				panic("You suck at coding")
			}
			Map.ranges = append(Map.ranges, ranges)
		}
		//fmt.Println(Map)
		conversions[Map.input] = &Map
	}

	for _, c := range conversions {
		c.next = conversions[c.output]
		if c.input != "seed" {
			c.previous = conversions[c.input]
		}
	}
	return seeds, conversions
}
