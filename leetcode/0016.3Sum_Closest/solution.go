package _016_3Sum_Closest

import (
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	diff := math.Inf(+1)

	//closet := math.Inf(+1)
	sort.Ints(nums)

	for i := 0 ; i < len(nums) - 1; i++ {
		//if i != 0 && nums[i] == nums[i - 1] {continue}
		j , k := i + 1, len(nums) - 1
		//j := i + 1
		//if !isAssign {
		//	result = nums[i] + nums[j] + nums[k] - target
		//	isAssign = true
		//}


		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(target - sum)) < math.Abs(diff) {
				diff = float64(target - sum)
			}
			if sum < target {
				j++
			} else {
				k--
			}
			if diff == 0 {
				break
			}

		}
	}

	return target - int(diff)
}

