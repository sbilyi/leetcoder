/*
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1
Example 2:

Input: [4,1,2,1,2]
Output: 4
*/
package main

import("fmt";)

func main() {
	input := []int{1,2,2,1,4}
	invoke(input, 4, "for %d unique value expected %d and was %d")
	invoke([]int{3,5,3,4,4,7,8,7,8}, 5, "for %d unique value expected %d and was %d")
	// invoke(input, 4, "for %d unique value expected %d and was %d")
}

func invoke(input  []int, expected int, pattern string) {
	result := singleNumber(input)
	fmt.Printf(pattern + "\n", input, expected, result)
}

func singleNumber(nums []int) int {
	sum := 0
	found := make(map[int]int)
	if len(nums) > 0 {
		fmt.Printf("Hello")
		for i := 0; i < len(nums); i++ {
			sum += nums[i]
			if val, ok := found[nums[i]]; ok {
				sum -= val*2
				delete(found, val)
			} else {
				found[nums[i]] = nums[i]
			}
		}
		return sum
	}
	
	return -1
}