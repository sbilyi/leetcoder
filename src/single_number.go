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

import "fmt"

func singleNumber(nums []int) int {
	sum := 0
	found := make(map[int]int)
	if len(nums) > 0 {
		fmt.Printf("Hello ")
		for i := 0; i < len(nums); i++ {
			sum += nums[i]
			if val, ok := found[nums[i]]; ok {
				sum -= val * 2
				delete(found, val)
			} else {
				found[nums[i]] = nums[i]
			}
		}
		return sum
	}

	return -1
}
