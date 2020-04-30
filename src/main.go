package main

import (
	"fmt"
)

func main() {
	input := []int{1, 2, 2, 1, 4}
	invoke(input, 4, "for %d unique value expected %d and was %d")
	invoke([]int{3, 5, 3, 4, 4, 7, 8, 7, 8}, 5, "for %d unique value expected %d and was %d")
	fmt.Printf("%d", TestCountCountries())
}

func invoke(input []int, expected int, pattern string) {
	result := singleNumber(input)
	fmt.Printf(pattern+"\n", input, expected, result)
}
