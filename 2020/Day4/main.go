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
	for _, person := range strings.Split(input, "\n\n") {
		if len(strings.Split(person, ":"))-1 >= 7 {
			if len(strings.Split(person, ":"))-1 == 8 {
			ans++
			} else {
				if !strings.Contains(person, "cid") {
					ans++
				}
			}
		}
	}
	return ans
}

func part2(input string) int {
  ans := 0
	passports := parseInput(input)
	for _, person := range passports {
		if (person.numFields == 7 && person.cid == "") || person.numFields == 8 {
			
		}
	}
	return ans
}

func checkPassport (passport Passport) bool {
	var height
	var unit
	if !len(helpers.ToString(passport.byr) == 4) || passport.byr < 1920 && passport.byr > 2002 {
		return false
	}
	if !len(helpers.ToString(passport.iyr) == 4) || passport.iyr < 2010 && passport.iyr > 2020 {
		return false
	}
	if !len(helpers.ToString(passport.eyr) == 4) || passport.eyr < 2020 && passport.eyr > 2030 {
		return false
	}
	_, err := fmt.Sscanf(passport.hgt, "%d%s", &height, &unit)

		

}

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
	cid string
	numFields int
}

func parseInput(input string)	[]Passport {
	passports := []Passport{}
	for _, person := range strings.Split(input, "\n\n") {
		passport := Passport{}
		newLine := strings.Replace(person, "\n", " ", -1)
		for _, field := range strings.Split(newLine, " ") {
			key := strings.Split(field, ":")[0]
			value := strings.Split(field, ":")[1]
			switch key {
			case "byr":
				passport.byr = helpers.ToInt(value)
				passport.numFields++
			case "iyr":
				passport.iyr = helpers.ToInt(value)
				passport.numFields++
			case "eyr":
				passport.eyr = helpers.ToInt(value)
				passport.numFields++
			case "hgt":
				passport.hgt = value
				passport.numFields++
			case "hcl":
				passport.hcl = value
				passport.numFields++
			case "ecl":
				passport.ecl = value
				passport.numFields++
			case "pid":
				passport.pid = value
				passport.numFields++
			case "cid":
				passport.cid = value
				passport.numFields++
			}
		}
		passports = append(passports, passport)
	}
return passports
}
