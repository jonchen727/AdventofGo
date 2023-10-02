package main

import (
	_ "embed"
	"fmt"
	"strings"
	//"strconv"
	"flag"
	//"sort"
	//"github.com/jonchen727/2022-AdventofCode/helpers"


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
	var part int
	flag.IntVar(&part, "part", 1, "part of the puzzle to run")
	flag.Parse()
	fmt.Println("Part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Answer:", ans)
		//ans := part1(input)
		//fmt.Println("Answer:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Answer:", ans)
		//fmt.Println("Answer:", ans)
	}

}
const (
	Win = 6
	Draw = 3 
	Loss = 0

	Rock = 1
	Paper = 2
	Scissors = 3
 )
// A = Rock = X
// B = Paper = Y
// C = Scissors = Z


func part1(input string) int {
	rounds := parseInput(input)
	
	choices := map[string]int{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}
	fmt.Println("Rounds:" , rounds)
	totalScore := 0

	for _, round := range rounds {
		choiceScore := choices[round[1]]
		switch round[0] {
		case "A": // They play rock
			switch round[1] {
			case "X": // We play rock
				totalScore += choiceScore
				totalScore += Draw
			case "Y": // We play paper
				totalScore += choiceScore
				totalScore += Win
			case "Z": // We play scissors
				totalScore += choiceScore
				totalScore += Loss
			}
			case "B": // They play paper
				switch round[1] {
				case "X": // We play rock
					totalScore += choiceScore
					totalScore += Loss
				case "Y": // We play paper
					totalScore += choiceScore
					totalScore += Draw
				case "Z": // We play scissors
					totalScore += choiceScore
					totalScore += Win
				}
			case "C": // They play scissors
				switch round[1] {
				case "X": // We play rock
					totalScore += choiceScore
					totalScore += Win
				case "Y": // We play paper
					totalScore += choiceScore
					totalScore += Loss
				case "Z": // We play scissors
					totalScore += choiceScore
					totalScore += Draw
				}
		}
	}
	return totalScore

}

func part2(input string) int {
	rounds := parseInput(input)
	
	choices := map[string]int{
		"X": Loss,
		"Y": Draw,
		"Z": Win,
	}
	fmt.Println("Rounds:" , rounds)
	totalScore := 0
	
	for _, round := range rounds {
		choiceScore := choices[round[1]]
		switch round[0] {
			case "A": // They play rock
				switch round[1] {
				case "X": // We need to lose
					totalScore += choiceScore
					totalScore += Scissors
				case "Y": // We need to draw
					totalScore += choiceScore
					totalScore += Rock
				case "Z": // We need to win
					totalScore += choiceScore
					totalScore += Paper
				}
			case "B": // They play paper
				switch round[1] {
				case "X": // We need to lose
					totalScore += choiceScore
					totalScore += Rock
				case "Y": // We need to draw
					totalScore += choiceScore
					totalScore += Paper
				case "Z": // We need to win
					totalScore += choiceScore
					totalScore += Scissors
				}
			case "C": // They play scissors
				switch round[1] {
				case "X": // We need to lose
					totalScore += choiceScore
					totalScore += Paper
				case "Y": // We need to draw
					totalScore += choiceScore
					totalScore += Scissors
				case "Z": // We need to win
					totalScore += choiceScore
					totalScore += Rock
				}
			}
		}
		return totalScore
	}


func parseInput(input string) (ans [][]string) {
	for _, lines := range strings.Split(input, "\n") {
			ans = append(ans, strings.Split(lines, " "))
	}
	return ans
}
