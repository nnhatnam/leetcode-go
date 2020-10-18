package _015_3Sum

import "sort"

func threeSum(nums []int) [][]int {
	var result = [][]int{}
	sort.Ints(nums)
	//initialize 3 pointer

	for i := 0; i < len(nums) - 2 ; i++ {
		if i != 0 && nums[i] ==  nums[i-1] { continue }
		j := i + 1
		k := len(nums) -1
		for j < k {

			if nums[i] + nums[j] + nums[k] == 0{
				result = append(result, []int{nums[i], nums[j], nums[k]})


				for j < k && nums[j] == nums[j + 1] {
					j++
				}
				j++
			} else if nums[i] + nums[j] + nums[k] < 0 {
				j++
			} else {
				k--
			}


		}
	}
	return result
}
