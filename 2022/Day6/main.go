package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"strings"
	//"reflect"
	//"strconv"
	"flag"
	"time"
	//"sort"
	//"github.com/jonchen727/AdventofGo/helpers"
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
	answer := part1(input)

	fmt.Println("Part 1 Answer:", answer)
	} else {
		ans := part2(input)
		fmt.Println("Part 2 Answer:", ans)
		//fmt.Println("Answer:", ans)
	}
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func part1(input string) int {
// grab 4 chars, then check first char with next 3, then second char with next 2, then third char with last char if dupe is found skip to dupe+1
var ans int
	loop:
	for i := 0; i < len(input)-4; { // loop to grab new set of 4
		
		
		set := input[i:i+4]
		// fmt.Println("chunk:", set)
		out:
		for j := 0; j < len(set)-1; j++ { // loop with in set of 4
			
			// fmt.Println("comparing:", string(set[j]), set[j+1:len(set)])
			for k := j+1 ; k < len(set); k++ { // loop to compare first char with rest of set
				// fmt.Println("i:", i)
				// fmt.Println("j:", j)
				// fmt.Println("k:", k)
				// fmt.Println("comparing:", string(set[j]), string(set[k]))
				if set[j] == set[k] {
					// fmt.Println("found dupe", string(set[j]))
					i = i + 1 //if a match is found grab a new set of 4 shifting 1 over
					break out
				} 
			}
		  if j == len(set)-2 {
				// fmt.Println("no dupes found")
				ans = i+4 // when we hit the end of our j that means we have no dupes and we offset 4 to get the start of packet
				break loop
			}
		}
	}
	return ans
}

func part2(input string) int{
	sop := part1(input)
	var ans int
	loop:
	for i := sop-4; i < len(input)-14; { // loop to grab new set of 14 starting at 4 minus the start of packet
		set := input[i:i+14]
		// fmt.Println("chunk:", set)
		out:
		for j := 0; j < len(set)-1; j++ { // loop with in set of 14
			
			// fmt.Println("comparing:", string(set[j]), set[j+1:len(set)])
			for k := j+1 ; k < len(set); k++ { // loop to compare first char with rest of set
				// fmt.Println("i:", i)
				// fmt.Println("j:", j)
				// fmt.Println("k:", k)
				// fmt.Println("comparing:", string(set[j]), string(set[k]))
				if set[j] == set[k] {
					// fmt.Println("found dupe", string(set[j]))
					i = i + 1 //if a match is found grab a new set of 14 shifting 1 over
					break out
				} 
			}
		  if j == len(set)-2 {
				// fmt.Println("no dupes found")
				ans = i+14 // when we hit the end of our j that means we have no dupes and we offset 14 to get the start message
				break loop
			}
		}
	}
	return ans
  

}


