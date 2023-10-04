package helpers

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

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}