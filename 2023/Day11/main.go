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
	galaxies, rmax, cmax := parseInput(input)
	expandSpace(galaxies, rmax, cmax, 1)

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			ans += ManhattanDistance(galaxies[i], galaxies[j])
		}
	}
	return ans
}

func part2(input string) int64 {
	ans := int64(0)
	galaxies, rmax, cmax := parseInput(input)
	expandSpace(galaxies, rmax, cmax, 0)

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			ans += ManhattanDistanceExp(galaxies[i], galaxies[j], 1000000)
		}
	}
	return ans
}

type Galaxy struct {
	r      int
	c      int
	symbol string
	rExp   int64
	cExp   int64
}

func ManhattanDistance(g1, g2 Galaxy) int {
	return helpers.Abs(g1.r-g2.r) + helpers.Abs(g1.c-g2.c)
}

func ManhattanDistanceExp(g1, g2 Galaxy, factor int64) int64 {
	c := int64(0)
	r := int64(0)
	dexpc := g1.cExp - g2.cExp
	dexpr := g1.rExp - g2.rExp
	dc := int64((g1.c - g2.c))-dexpc
	dr := int64((g1.r - g2.r))-dexpr

	c = helpers.Abs((dexpc*factor) + dc)
	r = helpers.Abs((dexpr*factor) + dr)

	// c = helpers.Abs(g1.c4)
	// if g1.rExp == g2.rExp {
	// 	if g1.r == g2.r {
	// 		r = 0
	// 	} else {
	// 		r = 1
	// 	}
	// } else {
	// 	r = factor*helpers.Abs(g1.rExp-g2.rExp)
	// }
	// if g1.cExp == g2.cExp {
	// 	if g1.c == g2.c {
	// 		c = 0
	// 	} else {
	// 		c = 1
	// 	}
	// } else {
	// 	c = factor*helpers.Abs(g1.cExp-g2.cExp)
	// }
	return r + c

}

func expandSpace(galaxies []Galaxy, rmax int, cmax int, factor int) {
	rspace := map[int][]*Galaxy{}
	//rkey := []int{}
	cspace := map[int][]*Galaxy{}
	//ckey := []int{}
	for g, galaxy := range galaxies {
		// make a map of all row values that are occupied
		for i := 0; i < rmax; i++ {
			if i == galaxy.r {
				if _, okr := rspace[galaxy.r]; !okr {
					rspace[i] = []*Galaxy{&galaxies[g]}
				} else {
					rspace[i] = append(rspace[i], &galaxies[g])
				}
			}
		}
		// make a map of all column values that are occupied
		for i := 0; i < cmax; i++ {
			if i == galaxy.c {
				if _, okc := cspace[galaxy.c]; !okc {
					cspace[i] = []*Galaxy{&galaxies[g]}
				} else {
					cspace[i] = append(cspace[i], &galaxies[g])
				}
			}
		}
	}
	// expand rows
	for i := 0; i < rmax; i++ {
		if _, ok := rspace[i]; !ok {
			for j := i + 1; j < rmax; j++ {
				for k, _ := range rspace[j] {
					rspace[j][k].r += factor
					rspace[j][k].rExp += 1
				}
			}
		}
	}

	for i := 0; i < cmax; i++ {
		if _, ok := cspace[i]; !ok {
			for j := i + 1; j < cmax; j++ {
				for k, _ := range cspace[j] {
					cspace[j][k].c += factor
					cspace[j][k].cExp += 1
				}
			}
		}
	}
}

func parseInput(input string) ([]Galaxy, int, int) {
	galaxies := []Galaxy{}
	rmax := len(strings.Split(input, "\n"))
	cmax := len(strings.Split(strings.Split(input, "\n")[0], ""))
	for r, line := range strings.Split(input, "\n") {
		for c, char := range strings.Split(line, "") {
			if string(char) != "." {
				galaxy := Galaxy{
					r:      r,
					c:      c,
					symbol: string(char),
				}
				galaxies = append(galaxies, galaxy)
			}
		}
	}
	return galaxies, rmax, cmax
}
