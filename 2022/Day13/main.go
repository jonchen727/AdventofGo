package main

import (
	_ "embed"
	"fmt"
	//"slices"
	"strings"
	"reflect"
	//"strconv"
	"flag"
	//"math"
	"time"
	"sort"
	//"github.com/jonchen727/AdventofGo/helpers"
	"encoding/json"
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
	packets := parseInput(input)
	ans = indexSum(packets)
	return ans
}

func part2(input string) int {
	ans := 1
	 packets := []string{"[2]","[6]"}
	 pairs := parseInput(input)
	 
	 for _, pair := range pairs {
		packets = append(packets, pair[0], pair[1])
	 }
	 sort.Slice(packets, func(i, j int) bool {
		left, right := packets[i], packets[j]
		return compare(left, right) == 1
	 })

	 for i, p := range packets {
		if fmt.Sprint(p) == "[2]" || fmt.Sprint(p) == "[6]" {
			ans *= i + 1
		}
	 }
	return ans


}

func parseInput(input string) [][]string {
	ans := [][]string{}
	for _, p := range strings.Split(input, "\n\n") {
		pkt := []string{}
		for _, line := range strings.Split(p, "\n") {
			pkt = append(pkt, line)
		}
		ans = append(ans, pkt)
	}
	return ans
}

func indexSum(packets [][]string) int {
	total := 0
	for i, p := range packets {
		//fmt.Println("Packet", i+1)
		pair1 := p[0]
		pair2 := p[1]
		
		//fmt.Println(compare(pair1, pair2))
		var result bool
		switch compare(pair1, pair2) {
		case 0:
			result = false
		case 1:
			result = true
		}
		if result {
			total += i + 1
		} else {
			continue
		}
	}
	  return total
}



// compare is a function that compares two values, which can be either integers or lists of integers.
// It returns a string indicating whether the left value is less than, greater than, or equal to the right value.
func compare(arr1, arr2 interface{}) int {
	var left, right interface{}
	
	switch arr1.(type) {
	case string:
		left = seperateArrays(arr1.(string))
		right = seperateArrays(arr2.(string))
	default:
		left = arr1
		right = arr2
	}
	
	// Check if the left and right values are lists.
	leftList, leftIsList := asList(left)
	rightList, rightIsList := asList(right)

	//fmt.Println(leftList, rightList)

	// If both values are lists, compare them element by element.
	if leftIsList && rightIsList {
		return compareLists(leftList, rightList)
	}

	// If both values are integers, compare them directly.
	if !leftIsList && !rightIsList {
		return compareIntegers(left, right)
	}

	// If one value is a list and the other is an integer, treat the integer as a list with one element.
	if leftIsList {
		rightList = []interface{}{right}
	} else {
		leftList = []interface{}{left}
	}

	return compareLists(leftList, rightList)
}

// compareLists is a helper function for compare that compares two lists of integers.
// It returns a string indicating whether the left list is less than, greater than, or equal to the right list.
func compareLists(left, right []interface{}) int{
	// Find the minimum length of the two lists.
	minLen := len(left)
	if len(right) < minLen {
		minLen = len(right)
	}

	// Compare the elements of the two lists until a difference is found.
	for i := 0; i < minLen; i++ {
		result := compare(left[i], right[i])
		if result >= 0 {
			return result
		}
	}

	// If the lists are equal up to the minimum length, the longer list is greater.
	if len(left) < len(right) {
		return 1
	} else if len(left) > len(right) {
		return 0 
	}

	return -1
}

// compareIntegers is a helper function for compare that compares two integers.
// It returns a string indicating whether the left integer is less than, greater than, or equal to the right integer.
func compareIntegers(left, right interface{}) int {
	// Convert the left and right values to floats.
	leftFloat, leftIsFloat := left.(float64)
	rightFloat, rightIsFloat := right.(float64)

	// If either value is not a float, return an error message.
	if !leftIsFloat || !rightIsFloat {
		return 0
	}

	// Compare the floats directly.
	if leftFloat < rightFloat {
		return 1
	} else if leftFloat > rightFloat {
		return 0
	}

	return -1
}

// asList is a helper function that converts a value to a list of interface{} values.
// If the value is not a list, it returns false.
func asList(value interface{}) ([]interface{}, bool) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		sliceValue := reflect.ValueOf(value)
		list := make([]interface{}, sliceValue.Len())
		for i := 0; i < sliceValue.Len(); i++ {
			list[i] = sliceValue.Index(i).Interface()
		}
		return list, true
	default:
		return nil, false
	}
}

// Data is a struct that represents a JSON object with a "data" field.
type Data struct {
	Data json.RawMessage `json:"data"`
}

// seperateArrays is a function that takes a string representing a JSON array and returns the array as an interface{} value.
func seperateArrays(input string) interface{} {
	// Create a Data object with the input string as the "data" field.
	data := Data{}
	jsn := `{ "data": `  + input + `}`
	decoder := json.NewDecoder(strings.NewReader(jsn))
	if err := decoder.Decode(&data); err != nil {
		fmt.Println(err)
	}

	// Unmarshal the "data" field into an interface{} value.
	var data1 interface{}
	if err := json.Unmarshal(data.Data, &data1); err != nil {
		fmt.Println(err)
	}

	return data1
}

