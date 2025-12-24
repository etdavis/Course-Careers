/* Question 2
In Go, write a function named oddSumMaxPair that accepts a slice of integers and returns a
slice containing the pair of integers that have the highest sum among all pairs with odd sums.
The input slice will always contain at least two integers. The pairs should be considered as
(numbers[i], numbers[j]) where j>i. If multiple pairs have the same highest odd sum, return the
pair that occurs first.
The returned slice should have the first element of the pair as the first element and the second
element of the pair as the second element. The sum of the pair should be an odd number.
Remember to handle the edge case where there are no pairs with an odd sum. In this case, the
function should return a nil or empty slice. */

//go:build practice2
//run with $ go run -tags practice2 practice2.go
package main

import (
	"fmt"
	"math"
)

func oddSumMaxPair(nums []int) []int {
	largestSum := int(math.Inf(-1))
	pair := []int{}
	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum > largestSum && sum % 2 == 1 {
				largestSum = sum
				pair = []int{nums[i], nums[j]}
			}
		}
	}
	return pair
}

func main() {
	fmt.Println(oddSumMaxPair([]int{7, 3, 9, 4, 2, 1, 5})) // {9, 4}
	fmt.Println(oddSumMaxPair([]int{10, 15, 3, 7, 8, 2, 12, 17})) // {12, 17}
}