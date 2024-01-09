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
	cards := parseInput(input)
	//fmt.Println(cards)
	for _, card := range cards {
		score := 0
		//fmt.Println("Card", i)
		for winner, _ := range card.winners {
			_, ok := card.cards[winner]

			if ok {
				//fmt.Println("Matches:", winner)
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		ans += score
	}
	return ans
}

func part2(input string) int {
	ans := 0
	cards := parseInput(input)
	queue := make([]Card, len(cards))
	cache := map[int]int{}
	copy(queue, cards)

	for _, card := range cards {
		count := 0
		for winner, _ := range card.winners {
			_, ok := card.cards[winner]
			if ok {
				count++
			}
		}
		cache[card.number] = count
	}
	for true {
		current := queue[0]
		// fmt.Println("Current:", current.number)
		queue = queue[1:]
		// fmt.Print("Queue: ")
		// for _, card := range queue {
		//   fmt.Print(card.number, ", ")
		// }
		// fmt.Println()
		//fmt.Println(queue)
		ans++
		wins := 0
		//fmt.Println(len(queue))
		if len(queue) == 0 {
			break
		}
		// check winners with cards
		wins, _ = cache[current.number]
		if wins == 0 {
			continue
		} else {
			add := make([]Card, wins)
			copy(add, cards[(current.number+1):(current.number+wins+1)])
			queue = append(add, queue...)
		}
	}
	return ans
}

type Card struct {
	winners map[int]bool
	cards   map[int]bool
	number  int
}

func parseInput(input string) []Card {
	cards := []Card{}
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		card := Card{
			winners: map[int]bool{},
			cards:   map[int]bool{},
			number:  i,
		}
		sets := strings.Split(strings.Split(strings.Replace(line, "  ", " ", -1), ": ")[1], " | ")
		for _, winners := range strings.Split(sets[0], " ") {
			//fmt.Println("String:", string(winners))
			//fmt.Println("Int:", helpers.ToInt(winners))
			card.winners[helpers.ToInt(winners)] = true
		}
		for _, cards := range strings.Split(sets[1], " ") {
			card.cards[helpers.ToInt(cards)] = true
		}
		cards = append(cards, card)
	}
	// fmt.Println(cards)
	return cards
}
