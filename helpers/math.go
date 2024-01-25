package helpers

import (
	"reflect"
	"golang.org/x/exp/constraints"
)

func MaxInt(nums ...int) int {
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func SumIntSlice(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func MinInt(nums ...int) int {
	minNum := nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}

func MinFloat(nums ...float64) float64 {
	minNum := nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}

type Numeric interface {
	constraints.Integer | constraints.Float
}

func MaxFloat(nums ...float64) float64 {
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func Abs[T Numeric](num T) T {
	var zero T
	if num < zero {
			return -num
	}
	return num
}



func DeepEqual(a1, a2 interface{}) bool {
	return reflect.DeepEqual(a1, a2)
}

func PosMod(d, m int) int {
	var res int = d % m
	if res < 0 && m > 0 {
		return res + m
	}
	return res
}
