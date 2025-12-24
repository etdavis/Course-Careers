/* Question 1
In Go, write a function named twoLargest that accepts a slice of integers and returns a slice
containing the two largest integers from the input slice. The returned slice should have the
larger integer as the first element and the second largest integer as the second element.
The input slice will always contain at least two integers.
If the largest and the second largest number in the input slice are the same, the output slice
should contain two copies of this number. */

//go:build practice1
//run with $ go run -tags practice1 practice1.go
package main

import (
	"fmt"
	"math"
)

func twoLargest(nums[]int) []int {
	largest := int(math.Inf(-1))
	secondLargest := int(math.Inf(-1))

	for _, num := range nums {
		if num > largest {
			secondLargest = largest
			largest = num
		} else if num > secondLargest {
			secondLargest = num
		}
	}
	return []int{largest, secondLargest}
}

func main() {
	fmt.Println(twoLargest([]int{7, 3, 7, 4, 2, 1, 5})) // {7, 7}
	fmt.Println(twoLargest([]int{7, 3, 9, 4, 2, 1, 5})) // {9, 7}
}