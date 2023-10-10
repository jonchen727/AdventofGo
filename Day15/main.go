package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"strings"
	//"reflect"
	//"strconv"
	"flag"
	"math"
	"time"
	//"sort"
	"github.com/jonchen727/2022-AdventofCode/helpers"
)

//go:embed input.txt
var input string

func init() {
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
	sensors, boundary, sensorMap, beaconMap := parseInput(input)
	ans := findNotBeaconLocations(sensors, boundary, sensorMap, beaconMap, 2000000)
	return ans
}

func part2(input string) int {
	sensors, _, _, _ := parseInput(input)
	ans := findTuningFrequency(sensors)
	return ans
}

func findNotBeaconLocations(sensors []Sensor, boundary Boundary, sensorMap map[string]bool, beaconMap map[string]bool, y int) int {
	signalMap := map[string]bool{}
	for _, s := range sensors {
		if (s.sensor.y+s.distance >= y) && (s.sensor.y-s.distance <= y) {

			x1 := helpers.Abs(s.sensor.y-y) - s.distance + s.sensor.x
			x2 := (helpers.Abs(s.sensor.y-y) * -1) + s.distance + s.sensor.x
			dist := helpers.Abs(x1 - x2)
			// fmt.Println(s, x1, x2, dist)
			for i := 0; i <= dist; i++ {
				_, ok := signalMap[fmt.Sprintf("%d,%d", x1+i, y)]
				_, ok2 := sensorMap[fmt.Sprintf("%d,%d", x1+i, y)]
				_, ok3 := beaconMap[fmt.Sprintf("%d,%d", x1+i, y)]
				if ok || ok2 || ok3 {
				} else {
					signalMap[fmt.Sprintf("%d,%d", x1+i, y)] = true
				}
			}
		}
	}
	for i := -4; i < 27; i++ {
		if _, ok := signalMap[fmt.Sprintf("%d,%d", i, y)]; ok {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	return len(signalMap)
}

func findTuningFrequency(sensors []Sensor) int {
	for _, s := range sensors {
		offset := s.distance + 1
		for r := -offset; r <= offset; r++ {
			ty := s.sensor.y + r
			if ty < 0 {
				continue
			}
			if ty > 4000000 {
				break
			}

			xoff := offset - helpers.Abs(r)
			xmin := s.sensor.x - xoff
			xmax := s.sensor.x + xoff

			if xmin >= 0 && xmin <= 4000000 && !isReachable(sensors, xmin, ty) {
				return xmin*4000000 + ty
			}
			if xmax >= 0 && xmax <= 4000000 && !isReachable(sensors, xmax, ty) {
				return xmax*4000000 + ty
			}
		}
	}
	return 0
}

func isReachable(sensors []Sensor, x, y int) bool {
	for _, s := range sensors {
		if s.distance >= helpers.Abs(s.sensor.x-x)+helpers.Abs(s.sensor.y-y) {
			return true
		}
	}
	return false
}

type Point struct {
	x int
	y int
}

type Sensor struct {
	beacon   Point
	sensor   Point
	distance int
}

type Boundary struct {
	xmax int
	xmin int
	ymax int
	ymin int
}

func parseInput(input string) ([]Sensor, Boundary, map[string]bool, map[string]bool) {
	sensorMap := map[string]bool{}
	beaconMap := map[string]bool{}
	xmax := 0
	ymax := 0
	xmin := math.MaxInt64
	ymin := math.MaxInt64
	sensors := []Sensor{}

	for _, line := range strings.Split(input, "\n") {
		sensor := Sensor{}
		var sx, xy, bx, by int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &xy, &bx, &by)
		sensor.beacon = Point{bx, by}
		sensor.sensor = Point{sx, xy}
		sensor.distance = helpers.Abs(sx-bx) + helpers.Abs(xy-by)
		sensors = append(sensors, sensor)
		sensorMap[fmt.Sprintf("%d,%d", sx, xy)] = true
		beaconMap[fmt.Sprintf("%d,%d", bx, by)] = true
		// create bounds
		if sx > xmax {
			xmax = sx + sensor.distance
		}
		if sx < xmin {
			xmin = sx - sensor.distance
		}
		if xy > ymax {
			ymax = xy + sensor.distance
		}
		if xy < ymin {
			ymin = xy - sensor.distance
		}

	}
	// fmt.Println(sensorMap)
	// fmt.Println(beaconMap)
	// fmt.Println(sensors)
	// fmt.Println(xmax, xmin, ymax, ymin)
	return sensors, Boundary{xmax, xmin, ymax, ymin}, sensorMap, beaconMap
}
