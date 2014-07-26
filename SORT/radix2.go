// Redix Sort written in Go by Austin G Walters for austingwalters.com

package main

import (
	"fmt"
	"strconv"
)

// Finds the largest number in an array
func findLargestNum(array []int) int {
	largestNum := 0

	for i := 0; i < len(array); i++ {
		if array[i] > largestNum {
			largestNum = array[i]
		}
	}
	return largestNum
}

// Radix Sort
func radixSort(array []int) []int {

  fmt.Println("Running Radix Sort on Unsorted List\n")

  // Base 10 is used
	largestNum := findLargestNum(array)
	size := len(array)
	significantDigit := 1
	semiSorted := make([]int, size, size)

	// Loop until we reach the largest significant digit
	for largestNum / significantDigit > 0 {
  
    fmt.Println("\tSorting: " + strconv.Itoa(significantDigit) + "'s place", array)

		bucket := [10]int{0}

