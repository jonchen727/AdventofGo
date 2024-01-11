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
	hands := parseInput(input, 1)
	handsMap := make(map[int][]Hand)
	for _, hand := range hands {
		handsMap[kindToInt(hand.Kind)] = append(handsMap[kindToInt(hand.Kind)], hand)
	}
	for _, hands := range handsMap {
		sort.Slice(hands, func(i, j int) bool {
			return compareHands(hands[i], hands[j], 1)
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
	hands := parseInput(input, 2)
	handsMap := make(map[int][]Hand)
	for _, hand := range hands {
		handsMap[kindToInt(hand.Kind)] = append(handsMap[kindToInt(hand.Kind)], hand)
	}
	for _, hands := range handsMap {
		sort.Slice(hands, func(i, j int) bool {
			return compareHands(hands[i], hands[j], 2)
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

var rankValues = map[string]int{
	"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10,
	"9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2,
}

var rankValues2 = map[string]int{
	"A": 14, "K": 13, "Q": 12, "J": 1, "T": 10,
	"9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2,
}

func compareHands(h1 Hand, h2 Hand, part int) bool {
	rankMap := map[string]int{}
	switch part {
	case 1:
		rankMap = rankValues
	case 2:
		rankMap = rankValues2
	}
	for i := 0; i < len(h1.Cards) && i < len(h2.Cards); i++ {
		if rankMap[h1.Cards[i]] != rankMap[h2.Cards[i]] {
			return rankMap[h1.Cards[i]] < rankMap[h2.Cards[i]]
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

func (h *Hand) classify(part int) {
	frequencyMap := make(map[string]int)
	jokerCount := 0

	for _, card := range h.Cards {
		if card == "J" && part == 2 {
			jokerCount++
		} else {
			frequencyMap[card]++
		}
	}

	h.Kind = make([]bool, 5) // 5-bit register for hand type
	// Five of a Kind
	if countInMap(frequencyMap, 5-jokerCount) || jokerCount == 5 {
		h.Kind[4] = true
		return
	}

	// Four of a Kind
	if countInMap(frequencyMap, 4-jokerCount) {
		h.Kind[3] = true
		return
	}

	// Full House
	if (countInMap(frequencyMap, 3) && countInMap(frequencyMap, 2)) || (countPairs(frequencyMap) == 2 && jokerCount == 1) {
		h.Kind[2] = true
		h.Kind[0] = true
		return
	}

	// Three of a Kind
	if countInMap(frequencyMap, 3-jokerCount) {
		h.Kind[2] = true
		return
	}

	// Two Pair
	if countPairs(frequencyMap) == 2 && jokerCount == 0 {
		h.Kind[1] = true
		return
	}

	// One Pair
	if countPairs(frequencyMap) == 1 || jokerCount == 1 {
		h.Kind[0] = true
		return
	}

	// High Card is represented by all bits being false
}

func (h *Hand) sortCards(part int) {
	h.Sorted = make([]string, len(h.Cards))
	copy(h.Sorted, h.Cards)
	rankMap := map[string]int{}
	switch part {
	case 1:
		rankMap = rankValues
	case 2:
		rankMap = rankValues2
	}

	sort.Slice(h.Sorted, func(i, j int) bool {
		return rankMap[h.Sorted[i]] > rankMap[h.Sorted[j]]
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

func parseInput(input string, part int) []Hand {
	lines := strings.Split(input, "\n")
	hands := []Hand{}
	for _, line := range lines {
		hand := Hand{}
		split := strings.Split(line, " ")
		for _, card := range split[0] {
			hand.Cards = append(hand.Cards, string(card))
		}
		hand.classify(part)
		hand.Bet = helpers.ToInt(split[1])
		hands = append(hands, hand)
	}
	return hands
}
