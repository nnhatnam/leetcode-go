package _365_How_Many_Numbers_Are_Smaller_Than_the_Current_Number

import (
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {

	result := make([]int, len(nums))
	copy(result, nums)
	sort.Ints(result)
	hashmap := map[int]int{}
	for i := len(result) - 1 ; i > 0 ; i-- {
		j := i - 1
		for ; j >= 0 ; j-- {
			if result[i] == result[j] {
				continue
			}
			hashmap[result[i]] = j + 1
			break
		}
	}
	for i := 0; i < len(nums); i++ {
		result[i] = hashmap[nums[i]]
	}
	return result
}
