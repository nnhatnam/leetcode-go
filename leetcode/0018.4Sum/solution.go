package _018_4Sum

import "sort"

func twoSum(nums []int, target int) [][]int {
	results := [][]int{}
	lo , hi := 0 , len(nums) - 1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum < target || (lo > 0 && nums[lo] == nums[lo - 1]) {
			lo++
		} else if sum > target || (hi < len(nums) -1 && nums[hi] == nums[hi-1]){
			hi -= 1
		} else {
			results = append(results, []int{nums[lo], nums[hi]} )
			lo++
			hi--
		}
	}
	return results
}

func kSum(nums []int, target, k int) [][]int {
	sort.Ints(nums)
	results := [][]int{}
	if len(nums) == 0 || nums[0]*k > target || target > nums[len(nums) - 1] * k {
		return results
	}
	if k == 2 {
		return twoSum(nums, target)
	}

	for i := 0; i < len(nums) ; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			for _ , arr := range kSum(nums[i + 1:] , target - nums[i] , k -1) {
				results = append(results, append([]int{ nums[i]} , arr...) )
			}
		}
	}
	return results
}

func fourSum(nums []int, target int) [][]int {
	return kSum(nums, target, 4)
}
