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
	"github.com/jonchen727/AdventofGo/helpers"
	"sort"
	"time"
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
	hands := parseInput(input)
	handsMap := make(map[int][]Hand)
	for _, hand := range hands {
		handsMap[kindToInt(hand.Kind)] = append(handsMap[kindToInt(hand.Kind)], hand)
	}
	for _, hands := range handsMap {
		sort.Slice(hands, func(i, j int) bool {
			return compareHands(hands[i], hands[j])
		})
	}
	keys := make([]int, 0, len(handsMap))
	for k := range handsMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.IntSlice(keys))

	var ranked []Hand
	for _, k := range keys {
		ranked = append(ranked, handsMap[k]...)
	}

	for i, hand := range ranked {
		ans += (i + 1) * hand.Bet
		//fmt.Println(hand)
	}
	return ans
}

func part2(input string) int {
	ans := 0
	return ans
}

var rankValues = map[string]int{
	"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10,
	"9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2,
}

func compareHands(h1, h2 Hand) bool {
	for i := 0; i < len(h1.Cards) && i < len(h2.Cards); i++ {
		if rankValues[h1.Cards[i]] != rankValues[h2.Cards[i]] {
			return rankValues[h1.Cards[i]] < rankValues[h2.Cards[i]]
		}
	}
	return false // If all cards are equal, return false (or handle as needed)
}

type Hand struct {
	Cards  []string
	Bet    int
	Sorted []string
	Kind   []bool
}

func kindToInt(kind []bool) int {
	result := 0
	for i, isSet := range kind {
		if isSet {
			result |= 1 << i
		}
	}
	return result
}

func (h *Hand) classify() {
	frequencyMap := make(map[string]int)

	h.sortCards()

	for _, card := range h.Cards {
		frequencyMap[card]++
	}

	h.Kind = make([]bool, 5) // 5-bit register for hand type

	// Check for Five of a Kind
	if countInMap(frequencyMap, 5) {
		h.Kind[4] = true
		return
	}
	// Check for Four of a Kind
	if countInMap(frequencyMap, 4) {
		h.Kind[3] = true
		return
	}
	// Check for Full House (Three of a Kind and a Pair)
	if countInMap(frequencyMap, 3) && countInMap(frequencyMap, 2) {
		h.Kind[2] = true
		h.Kind[0] = true
		return
	}
	// Check for Three of a Kind
	if countInMap(frequencyMap, 3) {
		h.Kind[2] = true
		return
	}
	// Check for Two Pair
	if countPairs(frequencyMap) == 2 {
		h.Kind[1] = true
		return
	}
	// Check for One Pair
	if countPairs(frequencyMap) == 1 {
		h.Kind[0] = true
		return
	}
	// High Card is represented by all bits being false
}

func (h *Hand) sortCards() {
	h.Sorted = make([]string, len(h.Cards))
	copy(h.Sorted, h.Cards)
	sort.Slice(h.Sorted, func(i, j int) bool {
		return rankValues[h.Sorted[i]] > rankValues[h.Sorted[j]]
	})
}

// countInMap checks if a count exists in the frequency map
func countInMap(frequencyMap map[string]int, count int) bool {
	for _, c := range frequencyMap {
		if c == count {
			return true
		}
	}
	return false
}

// countPairs counts the number of pairs in the frequency map
func countPairs(frequencyMap map[string]int) int {
	pairs := 0
	for _, c := range frequencyMap {
		if c == 2 {
			pairs++
		}
	}
	return pairs
}

func parseInput(input string) []Hand {
	lines := strings.Split(input, "\n")
	hands := []Hand{}
	for _, line := range lines {
		hand := Hand{}
		split := strings.Split(line, " ")
		for _, card := range split[0] {
			hand.Cards = append(hand.Cards, string(card))
		}
		hand.classify()
		hand.Bet = helpers.ToInt(split[1])
		hands = append(hands, hand)
	}
	return hands
}
