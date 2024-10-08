package main

import "fmt"

func Sum(nums []int) int {
	sum := 0

	// for i := 0; i < len(nums); i++ {
	// 	sum += nums[i]
	// }

	// using 'range'
	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumAll(numsToSum ...[]int) []int { // here 'numsToSum' is a slice of slices: [][]int{}

	// empty slice
	var tempSum []int

	// LOGIC: 2 - create a new slice using make
	// lengthOfNums := len(numsToSum)
	// sums := make([]int, lengthOfNums)

	// range through the big slice
	for _, num := range numsToSum {
		sumOfSingleSlice := Sum(num)
		tempSum = append(tempSum, sumOfSingleSlice) // append to a slice
	}

	// LOGIC: 2
	// for i, nums := range numsToSum {
	// 	sums[i] = Sum(nums)
	// }

	return tempSum
}

func SumAllTails(numsToSum ...[]int) []int {

	// empty slice
	var sumSlice []int

	// loop through the big slice
	for _, num := range numsToSum {
		if len(num) == 0 {
			sumSlice = append(sumSlice, 0)
		} else {
			trimTail := num[1:] // trim the head of each slice - 1st element
			sumSlice = append(sumSlice, Sum(trimTail))
		}
	}

	return sumSlice
}

// to test the logic
func main() {
	x := []int{1, 2, 3}
	y := []int{0, 9, 2}

	fmt.Println(SumAllTails(x, y))
}
